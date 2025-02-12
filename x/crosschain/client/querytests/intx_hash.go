package querytests

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"google.golang.org/grpc/codes"

	tmcli "github.com/tendermint/tendermint/libs/cli"
	"github.com/zeta-chain/zetacore/testutil/nullify"
	"github.com/zeta-chain/zetacore/x/crosschain/client/cli"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
	"google.golang.org/grpc/status"
	"strconv"
)

func (s *CliTestSuite) TestShowInTxHashToCctx() {
	ctx := s.network.Validators[0].ClientCtx
	objs := s.state.InTxHashToCctxList
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc       string
		idInTxHash string

		args []string
		err  error
		obj  types.InTxHashToCctx
	}{
		{
			desc:       "found",
			idInTxHash: objs[0].InTxHash,

			args: common,
			obj:  objs[0],
		},
		{
			desc:       "not found",
			idInTxHash: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		s.Run(tc.desc, func() {
			args := []string{
				tc.idInTxHash,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowInTxHashToCctx(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				s.Require().True(ok)
				s.Require().ErrorIs(stat.Err(), tc.err)
			} else {
				s.Require().NoError(err)
				var resp types.QueryGetInTxHashToCctxResponse
				s.Require().NoError(s.network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				s.Require().NotNil(resp.InTxHashToCctx)
				s.Require().Equal(nullify.Fill(&tc.obj),
					nullify.Fill(&resp.InTxHashToCctx),
				)
			}
		})
	}
}

func (s *CliTestSuite) TestListInTxHashToCctx() {
	ctx := s.network.Validators[0].ClientCtx
	objs := s.state.InTxHashToCctxList
	cctxCount := len(s.state.CrossChainTxs)
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	s.Run("ByOffset", func() {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListInTxHashToCctx(), args)
			s.Require().NoError(err)
			var resp types.QueryAllInTxHashToCctxResponse
			s.Require().NoError(s.network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			s.Require().LessOrEqual(len(resp.InTxHashToCctx), step)
			s.Require().Subset(nullify.Fill(objs),
				nullify.Fill(resp.InTxHashToCctx),
			)
		}
	})
	s.Run("ByKey", func() {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListInTxHashToCctx(), args)
			s.Require().NoError(err)
			var resp types.QueryAllInTxHashToCctxResponse
			s.Require().NoError(s.network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			s.Require().LessOrEqual(len(resp.InTxHashToCctx), step)
			s.Require().Subset(nullify.Fill(objs),
				nullify.Fill(resp.InTxHashToCctx),
			)
			next = resp.Pagination.NextKey
		}
	})
	s.Run("Total", func() {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListInTxHashToCctx(), args)
		s.Require().NoError(err)
		var resp types.QueryAllInTxHashToCctxResponse
		s.Require().NoError(s.network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		s.Require().NoError(err)
		// saving CCTX also adds a new mapping
		s.Require().Equal(len(objs)+cctxCount, int(resp.Pagination.Total))
		s.Require().ElementsMatch(nullify.Fill(objs),
			nullify.Fill(resp.InTxHashToCctx),
		)
	})
}

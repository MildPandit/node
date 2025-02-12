// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crosschain/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the metacore module's genesis state.
type GenesisState struct {
	Params              Params             `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	OutTxTrackerList    []OutTxTracker     `protobuf:"bytes,2,rep,name=outTxTrackerList,proto3" json:"outTxTrackerList"`
	Keygen              *Keygen            `protobuf:"bytes,3,opt,name=keygen,proto3" json:"keygen,omitempty"`
	Tss                 *TSS               `protobuf:"bytes,4,opt,name=tss,proto3" json:"tss,omitempty"`
	GasPriceList        []*GasPrice        `protobuf:"bytes,5,rep,name=gasPriceList,proto3" json:"gasPriceList,omitempty"`
	ChainNoncesList     []*ChainNonces     `protobuf:"bytes,6,rep,name=chainNoncesList,proto3" json:"chainNoncesList,omitempty"`
	CrossChainTxs       []*CrossChainTx    `protobuf:"bytes,7,rep,name=CrossChainTxs,proto3" json:"CrossChainTxs,omitempty"`
	NodeAccountList     []*NodeAccount     `protobuf:"bytes,8,rep,name=nodeAccountList,proto3" json:"nodeAccountList,omitempty"`
	LastBlockHeightList []*LastBlockHeight `protobuf:"bytes,9,rep,name=lastBlockHeightList,proto3" json:"lastBlockHeightList,omitempty"`
	InTxHashToCctxList  []InTxHashToCctx   `protobuf:"bytes,10,rep,name=inTxHashToCctxList,proto3" json:"inTxHashToCctxList"`
	PermissionFlags     *PermissionFlags   `protobuf:"bytes,11,opt,name=permissionFlags,proto3" json:"permissionFlags,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd51403692d571f4, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetOutTxTrackerList() []OutTxTracker {
	if m != nil {
		return m.OutTxTrackerList
	}
	return nil
}

func (m *GenesisState) GetKeygen() *Keygen {
	if m != nil {
		return m.Keygen
	}
	return nil
}

func (m *GenesisState) GetTss() *TSS {
	if m != nil {
		return m.Tss
	}
	return nil
}

func (m *GenesisState) GetGasPriceList() []*GasPrice {
	if m != nil {
		return m.GasPriceList
	}
	return nil
}

func (m *GenesisState) GetChainNoncesList() []*ChainNonces {
	if m != nil {
		return m.ChainNoncesList
	}
	return nil
}

func (m *GenesisState) GetCrossChainTxs() []*CrossChainTx {
	if m != nil {
		return m.CrossChainTxs
	}
	return nil
}

func (m *GenesisState) GetNodeAccountList() []*NodeAccount {
	if m != nil {
		return m.NodeAccountList
	}
	return nil
}

func (m *GenesisState) GetLastBlockHeightList() []*LastBlockHeight {
	if m != nil {
		return m.LastBlockHeightList
	}
	return nil
}

func (m *GenesisState) GetInTxHashToCctxList() []InTxHashToCctx {
	if m != nil {
		return m.InTxHashToCctxList
	}
	return nil
}

func (m *GenesisState) GetPermissionFlags() *PermissionFlags {
	if m != nil {
		return m.PermissionFlags
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "zetachain.zetacore.crosschain.GenesisState")
}

func init() { proto.RegisterFile("crosschain/genesis.proto", fileDescriptor_dd51403692d571f4) }

var fileDescriptor_dd51403692d571f4 = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x13, 0x5b, 0x53, 0x9d, 0x54, 0x2a, 0xa3, 0xe0, 0x12, 0xe8, 0xb6, 0x46, 0xc4, 0xa2,
	0x74, 0x83, 0xd5, 0xab, 0x07, 0x13, 0xb0, 0x95, 0x94, 0x1a, 0x37, 0x7b, 0x10, 0x41, 0xc6, 0xc9,
	0x38, 0xee, 0x2e, 0x49, 0x76, 0xc2, 0xce, 0x04, 0xb6, 0x7e, 0x0a, 0x0f, 0x7e, 0xa8, 0x1e, 0x7b,
	0xf4, 0x24, 0x92, 0x7c, 0x11, 0x99, 0xb7, 0x13, 0x33, 0x1b, 0x83, 0x9b, 0x4b, 0x18, 0xf2, 0xfe,
	0xff, 0xdf, 0x7b, 0x6f, 0xdf, 0x9b, 0x41, 0x0e, 0x4b, 0x85, 0x94, 0x2c, 0xa2, 0x71, 0xd2, 0x0a,
	0x79, 0xc2, 0x65, 0x2c, 0xbd, 0x49, 0x2a, 0x94, 0xc0, 0xfb, 0xdf, 0xb8, 0xa2, 0x10, 0xf0, 0xe0,
	0x24, 0x52, 0xee, 0x2d, 0xc5, 0x8d, 0x03, 0xcb, 0x28, 0xa6, 0x8a, 0xa8, 0x8c, 0xa8, 0x94, 0xb2,
	0x21, 0x4f, 0x73, 0x7f, 0xe3, 0x91, 0x25, 0x88, 0x13, 0x1d, 0x8f, 0xa8, 0x8c, 0x88, 0x12, 0x84,
	0x31, 0x95, 0x19, 0xd1, 0x43, 0x4b, 0x34, 0xe1, 0xe9, 0x38, 0x96, 0x32, 0x16, 0x09, 0xf9, 0x3a,
	0xa2, 0xa1, 0xa9, 0xa3, 0x71, 0x3f, 0x14, 0xa1, 0x80, 0x63, 0x4b, 0x9f, 0xcc, 0xbf, 0x0f, 0x2c,
	0xe3, 0x90, 0x5f, 0x86, 0x3c, 0x59, 0xc8, 0xad, 0x80, 0x92, 0x0b, 0x48, 0xc3, 0x6e, 0x93, 0x4a,
	0x32, 0x49, 0x63, 0xc6, 0x4d, 0x6c, 0xdf, 0x8a, 0xc1, 0x2f, 0x49, 0x44, 0xc2, 0xf8, 0xc2, 0xda,
	0xb4, 0xc2, 0x23, 0x2a, 0x15, 0x19, 0x8c, 0x04, 0x1b, 0x92, 0x88, 0xc7, 0x61, 0xa4, 0x8c, 0xc6,
	0xfe, 0x18, 0x70, 0x24, 0x39, 0xe8, 0x6f, 0x9f, 0x76, 0x8e, 0x44, 0x7c, 0xe1, 0x84, 0x32, 0x26,
	0xa6, 0x89, 0x5a, 0xd3, 0xcd, 0x84, 0xa6, 0x74, 0x6c, 0x92, 0x37, 0x7f, 0xec, 0xa0, 0xdd, 0xd3,
	0x7c, 0x2c, 0x7d, 0x45, 0x15, 0xc7, 0x1d, 0x54, 0xcb, 0x05, 0x4e, 0xf5, 0xb0, 0x7a, 0x54, 0x3f,
	0x79, 0xec, 0xfd, 0x77, 0x4c, 0x5e, 0x0f, 0xc4, 0xed, 0xed, 0xab, 0x5f, 0x07, 0x15, 0xdf, 0x58,
	0xf1, 0x27, 0x74, 0x57, 0x4c, 0x55, 0x90, 0x05, 0xf9, 0xc0, 0xce, 0x63, 0xa9, 0x9c, 0x1b, 0x87,
	0x5b, 0x47, 0xf5, 0x93, 0x67, 0x25, 0xb8, 0x77, 0x96, 0xcd, 0x40, 0xff, 0x41, 0xe1, 0x57, 0xa8,
	0x96, 0x8f, 0xc4, 0xd9, 0xda, 0xa8, 0xc6, 0x2e, 0x88, 0x7d, 0x63, 0xc2, 0x2f, 0xd1, 0x96, 0x92,
	0xd2, 0xd9, 0x06, 0x6f, 0xb3, 0xc4, 0x1b, 0xf4, 0xfb, 0xbe, 0x96, 0xe3, 0x2e, 0xda, 0x0d, 0xa9,
	0xec, 0xe9, 0xb9, 0x42, 0x3f, 0x37, 0xa1, 0x9f, 0x27, 0x25, 0xf6, 0x53, 0x63, 0xf1, 0x0b, 0x66,
	0x1c, 0xa0, 0x3d, 0x88, 0x5f, 0xc0, 0x22, 0x00, 0xaf, 0x06, 0xbc, 0xa7, 0x25, 0xbc, 0xce, 0xd2,
	0xe5, 0xaf, 0x22, 0xf0, 0x7b, 0x74, 0xa7, 0xa3, 0xa5, 0x20, 0x0a, 0x32, 0xe9, 0xec, 0x6c, 0xf4,
	0xcd, 0x6d, 0x8f, 0x5f, 0x24, 0xe8, 0x42, 0xf5, 0x3a, 0xbd, 0xce, 0xb7, 0x09, 0x0a, 0xbd, 0xb5,
	0x51, 0xa1, 0x17, 0x4b, 0x97, 0xbf, 0x8a, 0xc0, 0x9f, 0xd1, 0x3d, 0xbd, 0xe9, 0x6d, 0xbd, 0xe8,
	0x67, 0xb0, 0xe7, 0x40, 0xbe, 0x0d, 0x64, 0xaf, 0x84, 0x7c, 0x5e, 0x74, 0xfa, 0xeb, 0x50, 0x98,
	0x21, 0xac, 0x1b, 0x38, 0xa3, 0x32, 0x0a, 0x44, 0x87, 0xa9, 0x0c, 0x12, 0x20, 0x48, 0x70, 0x5c,
	0x92, 0xe0, 0x6d, 0xc1, 0x68, 0xb6, 0x70, 0x0d, 0x0e, 0x7f, 0x40, 0x7b, 0xcb, 0x37, 0xe5, 0x8d,
	0x7e, 0x52, 0x9c, 0x3a, 0x2c, 0x55, 0x59, 0x0b, 0xbd, 0xa2, 0xcb, 0x5f, 0xc5, 0xb4, 0xbb, 0x57,
	0x33, 0xb7, 0x7a, 0x3d, 0x73, 0xab, 0xbf, 0x67, 0x6e, 0xf5, 0xfb, 0xdc, 0xad, 0x5c, 0xcf, 0xdd,
	0xca, 0xcf, 0xb9, 0x5b, 0xf9, 0xf8, 0x3c, 0x8c, 0x55, 0x34, 0x1d, 0x78, 0x4c, 0x8c, 0x5b, 0x1a,
	0x7d, 0x9c, 0x5f, 0xea, 0x45, 0x96, 0x56, 0xd6, 0xb2, 0xdf, 0xa7, 0xcb, 0x09, 0x97, 0x83, 0x1a,
	0x5c, 0xf5, 0x17, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x3d, 0x8a, 0xdf, 0x8b, 0x05, 0x00,
	0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PermissionFlags != nil {
		{
			size, err := m.PermissionFlags.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	if len(m.InTxHashToCctxList) > 0 {
		for iNdEx := len(m.InTxHashToCctxList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InTxHashToCctxList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.LastBlockHeightList) > 0 {
		for iNdEx := len(m.LastBlockHeightList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LastBlockHeightList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.NodeAccountList) > 0 {
		for iNdEx := len(m.NodeAccountList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NodeAccountList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.CrossChainTxs) > 0 {
		for iNdEx := len(m.CrossChainTxs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CrossChainTxs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.ChainNoncesList) > 0 {
		for iNdEx := len(m.ChainNoncesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChainNoncesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.GasPriceList) > 0 {
		for iNdEx := len(m.GasPriceList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GasPriceList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Tss != nil {
		{
			size, err := m.Tss.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Keygen != nil {
		{
			size, err := m.Keygen.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OutTxTrackerList) > 0 {
		for iNdEx := len(m.OutTxTrackerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OutTxTrackerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.OutTxTrackerList) > 0 {
		for _, e := range m.OutTxTrackerList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.Keygen != nil {
		l = m.Keygen.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Tss != nil {
		l = m.Tss.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.GasPriceList) > 0 {
		for _, e := range m.GasPriceList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ChainNoncesList) > 0 {
		for _, e := range m.ChainNoncesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CrossChainTxs) > 0 {
		for _, e := range m.CrossChainTxs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.NodeAccountList) > 0 {
		for _, e := range m.NodeAccountList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LastBlockHeightList) > 0 {
		for _, e := range m.LastBlockHeightList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.InTxHashToCctxList) > 0 {
		for _, e := range m.InTxHashToCctxList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.PermissionFlags != nil {
		l = m.PermissionFlags.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutTxTrackerList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutTxTrackerList = append(m.OutTxTrackerList, OutTxTracker{})
			if err := m.OutTxTrackerList[len(m.OutTxTrackerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keygen", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Keygen == nil {
				m.Keygen = &Keygen{}
			}
			if err := m.Keygen.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tss", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tss == nil {
				m.Tss = &TSS{}
			}
			if err := m.Tss.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPriceList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasPriceList = append(m.GasPriceList, &GasPrice{})
			if err := m.GasPriceList[len(m.GasPriceList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainNoncesList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainNoncesList = append(m.ChainNoncesList, &ChainNonces{})
			if err := m.ChainNoncesList[len(m.ChainNoncesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CrossChainTxs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CrossChainTxs = append(m.CrossChainTxs, &CrossChainTx{})
			if err := m.CrossChainTxs[len(m.CrossChainTxs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAccountList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeAccountList = append(m.NodeAccountList, &NodeAccount{})
			if err := m.NodeAccountList[len(m.NodeAccountList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockHeightList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastBlockHeightList = append(m.LastBlockHeightList, &LastBlockHeight{})
			if err := m.LastBlockHeightList[len(m.LastBlockHeightList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InTxHashToCctxList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InTxHashToCctxList = append(m.InTxHashToCctxList, InTxHashToCctx{})
			if err := m.InTxHashToCctxList[len(m.InTxHashToCctxList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PermissionFlags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PermissionFlags == nil {
				m.PermissionFlags = &PermissionFlags{}
			}
			if err := m.PermissionFlags.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)

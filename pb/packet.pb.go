// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/packet.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Packet struct {
	SequenceId           uint32   `protobuf:"varint,1,opt,name=sequence_id,json=sequenceId,proto3" json:"sequence_id,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	AckStartSeq          []uint32 `protobuf:"varint,3,rep,packed,name=ack_start_seq,json=ackStartSeq,proto3" json:"ack_start_seq,omitempty"`
	AckSeqCount          []uint32 `protobuf:"varint,4,rep,packed,name=ack_seq_count,json=ackSeqCount,proto3" json:"ack_seq_count,omitempty"`
	BytesRead            uint64   `protobuf:"varint,5,opt,name=bytes_read,json=bytesRead,proto3" json:"bytes_read,omitempty"`
	ClientIds            []string `protobuf:"bytes,6,rep,name=client_ids,json=clientIds,proto3" json:"client_ids,omitempty"`
	WindowSize           uint32   `protobuf:"varint,7,opt,name=window_size,json=windowSize,proto3" json:"window_size,omitempty"`
	Mtu                  uint32   `protobuf:"varint,8,opt,name=mtu,proto3" json:"mtu,omitempty"`
	Close                bool     `protobuf:"varint,9,opt,name=close,proto3" json:"close,omitempty"`
	Handshake            bool     `protobuf:"varint,10,opt,name=handshake,proto3" json:"handshake,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_packet_dc5a9d1a21ec3142, []int{0}
}
func (m *Packet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packet.Unmarshal(m, b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
}
func (dst *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(dst, src)
}
func (m *Packet) XXX_Size() int {
	return xxx_messageInfo_Packet.Size(m)
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

func (m *Packet) GetSequenceId() uint32 {
	if m != nil {
		return m.SequenceId
	}
	return 0
}

func (m *Packet) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Packet) GetAckStartSeq() []uint32 {
	if m != nil {
		return m.AckStartSeq
	}
	return nil
}

func (m *Packet) GetAckSeqCount() []uint32 {
	if m != nil {
		return m.AckSeqCount
	}
	return nil
}

func (m *Packet) GetBytesRead() uint64 {
	if m != nil {
		return m.BytesRead
	}
	return 0
}

func (m *Packet) GetClientIds() []string {
	if m != nil {
		return m.ClientIds
	}
	return nil
}

func (m *Packet) GetWindowSize() uint32 {
	if m != nil {
		return m.WindowSize
	}
	return 0
}

func (m *Packet) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *Packet) GetClose() bool {
	if m != nil {
		return m.Close
	}
	return false
}

func (m *Packet) GetHandshake() bool {
	if m != nil {
		return m.Handshake
	}
	return false
}

func init() {
	proto.RegisterType((*Packet)(nil), "pb.Packet")
}

func init() { proto.RegisterFile("pb/packet.proto", fileDescriptor_packet_dc5a9d1a21ec3142) }

var fileDescriptor_packet_dc5a9d1a21ec3142 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xd0, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x05, 0x60, 0xe5, 0xa7, 0xa1, 0xb9, 0x25, 0x02, 0x59, 0x0c, 0x1e, 0x40, 0x58, 0x9d, 0x3c,
	0xc1, 0xc0, 0x23, 0x30, 0x75, 0x43, 0xce, 0x03, 0x58, 0x37, 0xf6, 0x95, 0x6a, 0xa5, 0xe4, 0xcf,
	0x8e, 0x2a, 0xfa, 0x48, 0x3c, 0x25, 0x8a, 0x23, 0xda, 0xed, 0xfa, 0x3b, 0x67, 0x38, 0x32, 0x3c,
	0x0c, 0xcd, 0xfb, 0x80, 0xa6, 0xa5, 0xf0, 0x36, 0x4c, 0x7d, 0xe8, 0x59, 0x3a, 0x34, 0xfb, 0xdf,
	0x14, 0x8a, 0xaf, 0x88, 0xec, 0x15, 0x76, 0x9e, 0xc6, 0x99, 0x3a, 0x43, 0xda, 0x59, 0x9e, 0x88,
	0x44, 0x56, 0x0a, 0xfe, 0xe9, 0x60, 0x19, 0x83, 0xdc, 0x62, 0x40, 0x9e, 0x8a, 0x44, 0xde, 0xab,
	0x78, 0xb3, 0x3d, 0x54, 0x68, 0x5a, 0xed, 0x03, 0x4e, 0x41, 0x7b, 0x1a, 0x79, 0x26, 0x32, 0x59,
	0xa9, 0x1d, 0x9a, 0xb6, 0x5e, 0xac, 0xa6, 0xf1, 0xda, 0xa1, 0x51, 0x9b, 0x7e, 0xee, 0x02, 0xcf,
	0x6f, 0x1d, 0x1a, 0x3f, 0x17, 0x62, 0x2f, 0x00, 0xcd, 0x4f, 0x20, 0xaf, 0x27, 0x42, 0xcb, 0x37,
	0x22, 0x91, 0xb9, 0x2a, 0xa3, 0x28, 0x42, 0xbb, 0xc4, 0xe6, 0xe4, 0xa8, 0x0b, 0xda, 0x59, 0xcf,
	0x0b, 0x91, 0xc9, 0x52, 0x95, 0xab, 0x1c, 0xac, 0x5f, 0xa6, 0x9f, 0x5d, 0x67, 0xfb, 0xb3, 0xf6,
	0xee, 0x42, 0xfc, 0x6e, 0x9d, 0xbe, 0x52, 0xed, 0x2e, 0xc4, 0x1e, 0x21, 0xfb, 0x0e, 0x33, 0xdf,
	0xc6, 0x60, 0x39, 0xd9, 0x13, 0x6c, 0xcc, 0xa9, 0xf7, 0xc4, 0x4b, 0x91, 0xc8, 0xad, 0x5a, 0x1f,
	0xec, 0x19, 0xca, 0x23, 0x76, 0xd6, 0x1f, 0xb1, 0x25, 0x0e, 0x31, 0xb9, 0x41, 0x53, 0xc4, 0x7f,
	0xfb, 0xf8, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x94, 0xb5, 0xa3, 0x8f, 0x4a, 0x01, 0x00, 0x00,
}

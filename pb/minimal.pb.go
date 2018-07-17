// Code generated by protoc-gen-go. DO NOT EDIT.
// source: minimal.proto

package protocols_p2p

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// a protocol define a set of reuqest and responses
type AddPeerRequest struct {
	// method specific data
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *AddPeerRequest) Reset()                    { *m = AddPeerRequest{} }
func (m *AddPeerRequest) String() string            { return proto.CompactTextString(m) }
func (*AddPeerRequest) ProtoMessage()               {}
func (*AddPeerRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *AddPeerRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type AddPeerResponse struct {
	// response specific data
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *AddPeerResponse) Reset()                    { *m = AddPeerResponse{} }
func (m *AddPeerResponse) String() string            { return proto.CompactTextString(m) }
func (*AddPeerResponse) ProtoMessage()               {}
func (*AddPeerResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *AddPeerResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type Collation struct {
	ShardID int64  `protobuf:"varint,1,opt,name=shardID" json:"shardID,omitempty"`
	Period  int64  `protobuf:"varint,2,opt,name=period" json:"period,omitempty"`
	Blobs   string `protobuf:"bytes,3,opt,name=blobs" json:"blobs,omitempty"`
}

func (m *Collation) Reset()                    { *m = Collation{} }
func (m *Collation) String() string            { return proto.CompactTextString(m) }
func (*Collation) ProtoMessage()               {}
func (*Collation) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Collation) GetShardID() int64 {
	if m != nil {
		return m.ShardID
	}
	return 0
}

func (m *Collation) GetPeriod() int64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *Collation) GetBlobs() string {
	if m != nil {
		return m.Blobs
	}
	return ""
}

type CollationRequest struct {
	ShardID int64  `protobuf:"varint,1,opt,name=shardID" json:"shardID,omitempty"`
	Period  int64  `protobuf:"varint,2,opt,name=period" json:"period,omitempty"`
	Hash    string `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
}

func (m *CollationRequest) Reset()                    { *m = CollationRequest{} }
func (m *CollationRequest) String() string            { return proto.CompactTextString(m) }
func (*CollationRequest) ProtoMessage()               {}
func (*CollationRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *CollationRequest) GetShardID() int64 {
	if m != nil {
		return m.ShardID
	}
	return 0
}

func (m *CollationRequest) GetPeriod() int64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *CollationRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type CollationResponse struct {
	Success   bool       `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Collation *Collation `protobuf:"bytes,2,opt,name=collation" json:"collation,omitempty"`
}

func (m *CollationResponse) Reset()                    { *m = CollationResponse{} }
func (m *CollationResponse) String() string            { return proto.CompactTextString(m) }
func (*CollationResponse) ProtoMessage()               {}
func (*CollationResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *CollationResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *CollationResponse) GetCollation() *Collation {
	if m != nil {
		return m.Collation
	}
	return nil
}

type NotifyShardsRequest struct {
	ShardIDs []int64 `protobuf:"varint,1,rep,packed,name=shardIDs" json:"shardIDs,omitempty"`
}

func (m *NotifyShardsRequest) Reset()                    { *m = NotifyShardsRequest{} }
func (m *NotifyShardsRequest) String() string            { return proto.CompactTextString(m) }
func (*NotifyShardsRequest) ProtoMessage()               {}
func (*NotifyShardsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *NotifyShardsRequest) GetShardIDs() []int64 {
	if m != nil {
		return m.ShardIDs
	}
	return nil
}

func init() {
	proto.RegisterType((*AddPeerRequest)(nil), "protocols.p2p.AddPeerRequest")
	proto.RegisterType((*AddPeerResponse)(nil), "protocols.p2p.AddPeerResponse")
	proto.RegisterType((*Collation)(nil), "protocols.p2p.Collation")
	proto.RegisterType((*CollationRequest)(nil), "protocols.p2p.CollationRequest")
	proto.RegisterType((*CollationResponse)(nil), "protocols.p2p.CollationResponse")
	proto.RegisterType((*NotifyShardsRequest)(nil), "protocols.p2p.NotifyShardsRequest")
}

func init() { proto.RegisterFile("minimal.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xcf, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x59, 0xab, 0xeb, 0x76, 0x64, 0xfd, 0x11, 0x45, 0x82, 0x27, 0xc9, 0x69, 0x51, 0x28,
	0xb8, 0x82, 0x77, 0xd1, 0x8b, 0x17, 0x91, 0xec, 0xc5, 0x6b, 0xda, 0x8e, 0x36, 0x90, 0x36, 0xb1,
	0xd3, 0x3d, 0xf8, 0xdf, 0x4b, 0xc7, 0xa6, 0xbb, 0x9e, 0x64, 0x4f, 0xc9, 0x23, 0x1f, 0x5f, 0xde,
	0x0c, 0xcc, 0x6b, 0xdb, 0xd8, 0xda, 0xb8, 0x2c, 0xb4, 0xbe, 0xf3, 0x62, 0xce, 0x47, 0xe1, 0x1d,
	0x65, 0x61, 0x19, 0xd4, 0x0d, 0x1c, 0x3f, 0x96, 0xe5, 0x1b, 0x62, 0xab, 0xf1, 0x6b, 0x8d, 0xd4,
	0x09, 0x09, 0x87, 0x35, 0x12, 0x99, 0x4f, 0x94, 0x93, 0xeb, 0xc9, 0x22, 0xd5, 0x31, 0xaa, 0x5b,
	0x38, 0x19, 0x59, 0x0a, 0xbe, 0x21, 0xec, 0x61, 0x5a, 0x17, 0x05, 0x12, 0x31, 0x3c, 0xd3, 0x31,
	0xaa, 0x15, 0xa4, 0x4f, 0xde, 0x39, 0xd3, 0x59, 0xdf, 0x30, 0x56, 0x99, 0xb6, 0x7c, 0x79, 0x66,
	0x2c, 0xd1, 0x31, 0x8a, 0x4b, 0x98, 0x06, 0x6c, 0xad, 0x2f, 0xe5, 0x1e, 0x3f, 0x0c, 0x49, 0x5c,
	0xc0, 0x41, 0xee, 0x7c, 0x4e, 0x32, 0xe1, 0x0e, 0xbf, 0x41, 0xbd, 0xc3, 0xe9, 0x28, 0xdd, 0xea,
	0xbb, 0xa3, 0x5b, 0xc0, 0x7e, 0x65, 0xa8, 0x1a, 0xd4, 0x7c, 0x57, 0x08, 0x67, 0x5b, 0xe6, 0xff,
	0xa6, 0x13, 0x0f, 0x90, 0x16, 0x11, 0x67, 0xfb, 0xd1, 0x52, 0x66, 0x7f, 0x36, 0x9b, 0x6d, 0x74,
	0x1b, 0x54, 0xdd, 0xc1, 0xf9, 0xab, 0xef, 0xec, 0xc7, 0xf7, 0xaa, 0xef, 0x48, 0x71, 0x86, 0x2b,
	0x98, 0x0d, 0xa5, 0xfb, 0x9f, 0x92, 0x45, 0xa2, 0xc7, 0x9c, 0x4f, 0x59, 0x7b, 0xff, 0x13, 0x00,
	0x00, 0xff, 0xff, 0xcc, 0x49, 0xd2, 0x37, 0xc8, 0x01, 0x00, 0x00,
}

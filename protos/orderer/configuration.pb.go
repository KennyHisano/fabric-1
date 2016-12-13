// Code generated by protoc-gen-go.
// source: orderer/configuration.proto
// DO NOT EDIT!

package orderer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/hyperledger/fabric/protos/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ConsensusType struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *ConsensusType) Reset()                    { *m = ConsensusType{} }
func (m *ConsensusType) String() string            { return proto.CompactTextString(m) }
func (*ConsensusType) ProtoMessage()               {}
func (*ConsensusType) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type BatchSize struct {
	// Simply specified as messages for now, in the future we may want to allow this to be specified by size in bytes
	Messages uint32 `protobuf:"varint,1,opt,name=messages" json:"messages,omitempty"`
}

func (m *BatchSize) Reset()                    { *m = BatchSize{} }
func (m *BatchSize) String() string            { return proto.CompactTextString(m) }
func (*BatchSize) ProtoMessage()               {}
func (*BatchSize) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

// When submitting a new chain configuration transaction to create a new chain, the first configuration item must of of type
// Orderer with Key CreationPolicy and contents of a Marshaled CreationPolicy.  The policy should be set to the policy which
// was supplied by the ordering service for the client's chain creation.  The digest should be the hash of the concatenation
// of the remaining ConfigurationItem bytes.  The signatures of the configuration item should satisfy the policy for chain creation
type CreationPolicy struct {
	// The name of the policy which should be used to validate the creation of this chain
	Policy string `protobuf:"bytes,1,opt,name=policy" json:"policy,omitempty"`
	// The hash of the concatenation of remaining configuration item bytes
	Digest []byte `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (m *CreationPolicy) Reset()                    { *m = CreationPolicy{} }
func (m *CreationPolicy) String() string            { return proto.CompactTextString(m) }
func (*CreationPolicy) ProtoMessage()               {}
func (*CreationPolicy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type ChainCreators struct {
	// A list of policies, any of which may be specified as the chain creation policy in a chain creation request
	Policies []string `protobuf:"bytes,1,rep,name=policies" json:"policies,omitempty"`
}

func (m *ChainCreators) Reset()                    { *m = ChainCreators{} }
func (m *ChainCreators) String() string            { return proto.CompactTextString(m) }
func (*ChainCreators) ProtoMessage()               {}
func (*ChainCreators) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func init() {
	proto.RegisterType((*ConsensusType)(nil), "orderer.ConsensusType")
	proto.RegisterType((*BatchSize)(nil), "orderer.BatchSize")
	proto.RegisterType((*CreationPolicy)(nil), "orderer.CreationPolicy")
	proto.RegisterType((*ChainCreators)(nil), "orderer.ChainCreators")
}

func init() { proto.RegisterFile("orderer/configuration.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x3c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x59, 0x95, 0xd5, 0x06, 0xeb, 0x21, 0x82, 0x2c, 0xeb, 0x65, 0xa9, 0x07, 0x0b, 0x4a,
	0x73, 0xf0, 0x0f, 0xc8, 0xf6, 0x0f, 0x48, 0xf5, 0xe4, 0x2d, 0x4d, 0x67, 0xd3, 0xc0, 0x36, 0x13,
	0x26, 0xe9, 0xa1, 0xfe, 0x7a, 0xd9, 0xd9, 0xb0, 0xa7, 0xbc, 0x2f, 0x2f, 0x79, 0x8f, 0x19, 0xf1,
	0x8c, 0x34, 0x00, 0x01, 0x29, 0x83, 0xfe, 0xe0, 0xec, 0x4c, 0x3a, 0x39, 0xf4, 0x4d, 0x20, 0x4c,
	0x28, 0x6f, 0xb3, 0xb9, 0x7d, 0x34, 0x38, 0x4d, 0xe8, 0xd5, 0xf9, 0x38, 0xbb, 0xd5, 0x8b, 0x28,
	0x5b, 0xf4, 0x11, 0x7c, 0x9c, 0xe3, 0xcf, 0x12, 0x40, 0x4a, 0x71, 0x93, 0x96, 0x00, 0x9b, 0xd5,
	0x6e, 0x55, 0x17, 0x1d, 0xeb, 0xea, 0x55, 0x14, 0x7b, 0x9d, 0xcc, 0xf8, 0xed, 0xfe, 0x40, 0x6e,
	0xc5, 0xdd, 0x04, 0x31, 0x6a, 0x0b, 0x91, 0x1f, 0x95, 0xdd, 0x85, 0xab, 0x4f, 0xf1, 0xd0, 0x12,
	0x70, 0xfb, 0x17, 0x1e, 0x9d, 0x59, 0xe4, 0x93, 0x58, 0x07, 0x56, 0x39, 0x30, 0xd3, 0xe9, 0x7e,
	0x70, 0x16, 0x62, 0xda, 0x5c, 0xed, 0x56, 0xf5, 0x7d, 0x97, 0xa9, 0x7a, 0x13, 0x65, 0x3b, 0x6a,
	0xe7, 0x39, 0x06, 0x29, 0x9e, 0xea, 0xf8, 0x8b, 0xe3, 0xba, 0xeb, 0xba, 0xe8, 0x2e, 0xbc, 0x6f,
	0x7e, 0xdf, 0xad, 0x4b, 0xe3, 0xdc, 0x37, 0x06, 0x27, 0x35, 0x2e, 0x01, 0xe8, 0x08, 0x83, 0x05,
	0x52, 0x07, 0xdd, 0x93, 0x33, 0x8a, 0x67, 0x8c, 0x2a, 0x6f, 0xa0, 0x5f, 0x33, 0x7f, 0xfc, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x8b, 0x76, 0xe7, 0x02, 0x30, 0x01, 0x00, 0x00,
}
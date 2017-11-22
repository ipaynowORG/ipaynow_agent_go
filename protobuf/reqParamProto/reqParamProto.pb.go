// Code generated by protoc-gen-go. DO NOT EDIT.
// source: reqParamProto.proto

/*
Package reqParamProto is a generated protocol buffer package.

It is generated from these files:
	reqParamProto.proto

It has these top-level messages:
	Param
*/
package reqParamProto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import transactionProto "protobuf_test/protobuf/transactionProto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type METHOD int32

const (
	METHOD_PAY                          METHOD = 0
	METHOD_RECEIVE                      METHOD = 1
	METHOD_QUERY_TRANS                  METHOD = 2
	METHOD_QUERY_BALANCE                METHOD = 3
	METHOD_TRANS                        METHOD = 4
	METHOD_QUERY_TRANS_DETAIL           METHOD = 5
	METHOD_QUERY_BALANCE_V1             METHOD = 6
	METHOD_AGENT_PAY_REFUND_QUERY       METHOD = 7
	METHOD_AGENT_PAY_REFUND_BATCH_QUERY METHOD = 8
)

var METHOD_name = map[int32]string{
	0: "PAY",
	1: "RECEIVE",
	2: "QUERY_TRANS",
	3: "QUERY_BALANCE",
	4: "TRANS",
	5: "QUERY_TRANS_DETAIL",
	6: "QUERY_BALANCE_V1",
	7: "AGENT_PAY_REFUND_QUERY",
	8: "AGENT_PAY_REFUND_BATCH_QUERY",
}
var METHOD_value = map[string]int32{
	"PAY":                          0,
	"RECEIVE":                      1,
	"QUERY_TRANS":                  2,
	"QUERY_BALANCE":                3,
	"TRANS":                        4,
	"QUERY_TRANS_DETAIL":           5,
	"QUERY_BALANCE_V1":             6,
	"AGENT_PAY_REFUND_QUERY":       7,
	"AGENT_PAY_REFUND_BATCH_QUERY": 8,
}

func (x METHOD) String() string {
	return proto.EnumName(METHOD_name, int32(x))
}
func (METHOD) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Param struct {
	Method METHOD                              `protobuf:"varint,1,opt,name=method,enum=METHOD" json:"method,omitempty"`
	Sign   *transactionProto.SignedTransaction `protobuf:"bytes,2,opt,name=sign" json:"sign,omitempty"`
}

func (m *Param) Reset()                    { *m = Param{} }
func (m *Param) String() string            { return proto.CompactTextString(m) }
func (*Param) ProtoMessage()               {}
func (*Param) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Param) GetMethod() METHOD {
	if m != nil {
		return m.Method
	}
	return METHOD_PAY
}

func (m *Param) GetSign() *transactionProto.SignedTransaction {
	if m != nil {
		return m.Sign
	}
	return nil
}

func init() {
	proto.RegisterType((*Param)(nil), "Param")
	proto.RegisterEnum("METHOD", METHOD_name, METHOD_value)
}

func init() { proto.RegisterFile("reqParamProto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xdd, 0xb6, 0x49, 0x74, 0x82, 0xba, 0x8e, 0x12, 0x4a, 0x11, 0x0c, 0x1e, 0xa4, 0x78,
	0x08, 0x58, 0x9f, 0x60, 0x9b, 0x8c, 0xb6, 0x50, 0xe3, 0xba, 0xdd, 0x16, 0x72, 0x5a, 0xa2, 0x0d,
	0x35, 0x87, 0x26, 0x9a, 0xe6, 0xe5, 0x7c, 0x3b, 0x31, 0x29, 0x58, 0xf1, 0xb2, 0x87, 0xef, 0xfb,
	0xfe, 0x3d, 0x0c, 0x9c, 0x57, 0xd9, 0xa7, 0x4c, 0xab, 0x74, 0x23, 0xab, 0xb2, 0x2e, 0x83, 0x8f,
	0x9f, 0x77, 0xe0, 0xd5, 0x55, 0x5a, 0x6c, 0xd3, 0xb7, 0x3a, 0x2f, 0x8b, 0x3d, 0x7e, 0x2d, 0xc1,
	0x6a, 0x5a, 0xbc, 0x02, 0x7b, 0x93, 0xd5, 0xef, 0xe5, 0xaa, 0xcf, 0x7c, 0x36, 0x3c, 0x19, 0x39,
	0xc1, 0x13, 0xe9, 0xc9, 0x73, 0xa4, 0x76, 0x18, 0x6f, 0xa0, 0xb7, 0xcd, 0xd7, 0x45, 0xbf, 0xe3,
	0xb3, 0xa1, 0x3b, 0xc2, 0x60, 0x9e, 0xaf, 0x8b, 0x6c, 0xa5, 0x7f, 0xbf, 0x55, 0x8d, 0xbf, 0xfd,
	0x62, 0x60, 0xb7, 0x53, 0x74, 0xa0, 0x2b, 0x45, 0xc2, 0x0f, 0xd0, 0x05, 0x47, 0x51, 0x48, 0xd3,
	0x25, 0x71, 0x86, 0xa7, 0xe0, 0xbe, 0x2c, 0x48, 0x25, 0x46, 0x2b, 0x11, 0xcf, 0x79, 0x07, 0xcf,
	0xe0, 0xb8, 0x05, 0x63, 0x31, 0x13, 0x71, 0x48, 0xbc, 0x8b, 0x47, 0x60, 0xb5, 0xb6, 0x87, 0x1e,
	0xe0, 0x5e, 0x6e, 0x22, 0xd2, 0x62, 0x3a, 0xe3, 0x16, 0x5e, 0x00, 0xff, 0xb3, 0x32, 0xcb, 0x3b,
	0x6e, 0xe3, 0x00, 0x3c, 0xf1, 0x48, 0xb1, 0x36, 0x52, 0x24, 0x46, 0xd1, 0xc3, 0x22, 0x8e, 0x4c,
	0x93, 0x71, 0x07, 0x7d, 0xb8, 0xfc, 0xe7, 0xc6, 0x42, 0x87, 0x93, 0x5d, 0x71, 0xf8, 0x6a, 0x37,
	0x47, 0xb9, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x76, 0x21, 0x1e, 0x23, 0x43, 0x01, 0x00, 0x00,
}
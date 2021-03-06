// Code generated by protoc-gen-go.
// source: proto/protolion.proto
// DO NOT EDIT!

package protolion

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Level is a logging level.
type Level int32

const (
	Level_LEVEL_DEBUG Level = 0
	Level_LEVEL_INFO  Level = 1
	Level_LEVEL_WARN  Level = 2
	Level_LEVEL_ERROR Level = 3
	Level_LEVEL_FATAL Level = 4
	Level_LEVEL_PANIC Level = 5
	Level_LEVEL_NONE  Level = 6
)

var Level_name = map[int32]string{
	0: "LEVEL_DEBUG",
	1: "LEVEL_INFO",
	2: "LEVEL_WARN",
	3: "LEVEL_ERROR",
	4: "LEVEL_FATAL",
	5: "LEVEL_PANIC",
	6: "LEVEL_NONE",
}
var Level_value = map[string]int32{
	"LEVEL_DEBUG": 0,
	"LEVEL_INFO":  1,
	"LEVEL_WARN":  2,
	"LEVEL_ERROR": 3,
	"LEVEL_FATAL": 4,
	"LEVEL_PANIC": 5,
	"LEVEL_NONE":  6,
}

func (x Level) String() string {
	return proto.EnumName(Level_name, int32(x))
}
func (Level) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Entry is the object serialized for logging.
type Entry struct {
	// id may not be set depending on logger options
	// it is up to the user to determine if id is required
	Id        string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Level     Level                      `protobuf:"varint,2,opt,name=level,enum=lion.Level" json:"level,omitempty"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	// both context and fields may be set
	Context []*Entry_Message  `protobuf:"bytes,4,rep,name=context" json:"context,omitempty"`
	Fields  map[string]string `protobuf:"bytes,5,rep,name=fields" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// one of event, message, writer_output will be set
	Event        *Entry_Message `protobuf:"bytes,6,opt,name=event" json:"event,omitempty"`
	Message      string         `protobuf:"bytes,7,opt,name=message" json:"message,omitempty"`
	WriterOutput []byte         `protobuf:"bytes,8,opt,name=writer_output,json=writerOutput,proto3" json:"writer_output,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Entry) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Entry) GetContext() []*Entry_Message {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Entry) GetFields() map[string]string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Entry) GetEvent() *Entry_Message {
	if m != nil {
		return m.Event
	}
	return nil
}

// Message represents a serialized protobuf message.
// The name is the name registered with lion.
type Entry_Message struct {
	Encoding string `protobuf:"bytes,1,opt,name=encoding" json:"encoding,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Value    []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Entry_Message) Reset()                    { *m = Entry_Message{} }
func (m *Entry_Message) String() string            { return proto.CompactTextString(m) }
func (*Entry_Message) ProtoMessage()               {}
func (*Entry_Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func init() {
	proto.RegisterType((*Entry)(nil), "lion.Entry")
	proto.RegisterType((*Entry_Message)(nil), "lion.Entry.Message")
	proto.RegisterEnum("lion.Level", Level_name, Level_value)
}

func init() { proto.RegisterFile("proto/protolion.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0x41, 0x8f, 0x93, 0x40,
	0x14, 0x16, 0x28, 0xed, 0xf6, 0x51, 0xeb, 0xe4, 0xa9, 0x71, 0xc2, 0x45, 0xd4, 0x0b, 0x9a, 0x48,
	0x93, 0x7a, 0x59, 0xbd, 0x75, 0x95, 0x9a, 0x4d, 0x10, 0xcc, 0x64, 0xd5, 0xc4, 0xcb, 0x86, 0x5d,
	0x66, 0x09, 0x11, 0x98, 0x06, 0x86, 0xea, 0xc6, 0x7f, 0xe8, 0xaf, 0x32, 0x9d, 0x81, 0xb2, 0x87,
	0xbd, 0x90, 0xef, 0xfb, 0xde, 0x37, 0xdf, 0xbc, 0xf7, 0x18, 0x78, 0xba, 0x6b, 0x84, 0x14, 0x2b,
	0xf5, 0x2d, 0x0b, 0x51, 0x07, 0x0a, 0xe1, 0xe4, 0x80, 0xdd, 0xe7, 0xb9, 0x10, 0x79, 0xc9, 0x75,
	0xf5, 0xaa, 0xbb, 0x59, 0xc9, 0xa2, 0xe2, 0xad, 0x4c, 0xab, 0x9d, 0xb6, 0xbd, 0xfc, 0x67, 0x81,
	0x1d, 0xd6, 0xb2, 0xb9, 0xc5, 0x25, 0x98, 0x45, 0x46, 0x0d, 0xcf, 0xf0, 0xe7, 0xcc, 0x2c, 0x32,
	0x7c, 0x01, 0x76, 0xc9, 0xf7, 0xbc, 0xa4, 0xa6, 0x67, 0xf8, 0xcb, 0xb5, 0x13, 0xa8, 0xf0, 0xe8,
	0x20, 0x31, 0x5d, 0xc1, 0x53, 0x98, 0x1f, 0xf3, 0xa8, 0xe5, 0x19, 0xbe, 0xb3, 0x76, 0x03, 0x7d,
	0x63, 0x30, 0xdc, 0x18, 0x5c, 0x0c, 0x0e, 0x36, 0x9a, 0xf1, 0x2d, 0xcc, 0xae, 0x45, 0x2d, 0xf9,
	0x1f, 0x49, 0x27, 0x9e, 0xe5, 0x3b, 0xeb, 0xc7, 0x3a, 0x5e, 0xb5, 0x12, 0x7c, 0xe1, 0x6d, 0x9b,
	0xe6, 0x9c, 0x0d, 0x1e, 0x5c, 0xc1, 0xf4, 0xa6, 0xe0, 0x65, 0xd6, 0x52, 0x5b, 0xb9, 0x9f, 0xdd,
	0x75, 0x6f, 0x55, 0x45, 0x61, 0xd6, 0xdb, 0xf0, 0x35, 0xd8, 0x7c, 0xcf, 0x6b, 0x49, 0xa7, 0xaa,
	0xab, 0x7b, 0xd3, 0xb5, 0x03, 0x29, 0xcc, 0x2a, 0xad, 0xd0, 0x99, 0x1a, 0x7e, 0xa0, 0xf8, 0x0a,
	0x1e, 0xfe, 0x6e, 0x0a, 0xc9, 0x9b, 0x4b, 0xd1, 0xc9, 0x5d, 0x27, 0xe9, 0x89, 0x67, 0xf8, 0x0b,
	0xb6, 0xd0, 0x62, 0xa2, 0x34, 0x37, 0x81, 0x59, 0x1f, 0x88, 0x2e, 0x9c, 0xf0, 0xfa, 0x5a, 0x64,
	0x45, 0x9d, 0xf7, 0x7b, 0x3c, 0x72, 0x44, 0x98, 0xd4, 0x69, 0xc5, 0xd5, 0x32, 0xe7, 0x4c, 0x61,
	0x7c, 0x02, 0xf6, 0x3e, 0x2d, 0x3b, 0xae, 0x56, 0xb7, 0x60, 0x9a, 0xb8, 0xef, 0xc1, 0xb9, 0x33,
	0x11, 0x12, 0xb0, 0x7e, 0xf1, 0xdb, 0x3e, 0xef, 0x00, 0xc7, 0x63, 0x3a, 0x4b, 0x93, 0x0f, 0xe6,
	0xa9, 0xf1, 0xe6, 0x2f, 0xd8, 0xea, 0xff, 0xe0, 0x23, 0x70, 0xa2, 0xf0, 0x7b, 0x18, 0x5d, 0x7e,
	0x0a, 0xcf, 0xbe, 0x7d, 0x26, 0x0f, 0x70, 0x09, 0xa0, 0x85, 0xf3, 0x78, 0x9b, 0x10, 0x63, 0xe4,
	0x3f, 0x36, 0x2c, 0x26, 0xe6, 0x78, 0x20, 0x64, 0x2c, 0x61, 0xc4, 0x1a, 0x85, 0xed, 0xe6, 0x62,
	0x13, 0x91, 0xc9, 0x28, 0x7c, 0xdd, 0xc4, 0xe7, 0x1f, 0x89, 0x3d, 0x46, 0xc4, 0x49, 0x1c, 0x92,
	0xe9, 0x99, 0xf3, 0x73, 0x7e, 0x7c, 0x83, 0x57, 0x53, 0x05, 0xdf, 0xfd, 0x0f, 0x00, 0x00, 0xff,
	0xff, 0xcc, 0xd0, 0x4d, 0xc7, 0x9d, 0x02, 0x00, 0x00,
}

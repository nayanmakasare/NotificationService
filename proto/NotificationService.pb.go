// Code generated by protoc-gen-go. DO NOT EDIT.
// source: NotificationService.proto

package NotificationService

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MessageToPublish struct {
	ExchangeName         string   `protobuf:"bytes,1,opt,name=exchangeName,proto3" json:"exchangeName,omitempty"`
	RoutingKeyName       string   `protobuf:"bytes,2,opt,name=routingKeyName,proto3" json:"routingKeyName,omitempty"`
	MessageTosend        []byte   `protobuf:"bytes,3,opt,name=messageTosend,proto3" json:"messageTosend,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageToPublish) Reset()         { *m = MessageToPublish{} }
func (m *MessageToPublish) String() string { return proto.CompactTextString(m) }
func (*MessageToPublish) ProtoMessage()    {}
func (*MessageToPublish) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a9b3db0602ab621, []int{0}
}

func (m *MessageToPublish) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageToPublish.Unmarshal(m, b)
}
func (m *MessageToPublish) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageToPublish.Marshal(b, m, deterministic)
}
func (m *MessageToPublish) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageToPublish.Merge(m, src)
}
func (m *MessageToPublish) XXX_Size() int {
	return xxx_messageInfo_MessageToPublish.Size(m)
}
func (m *MessageToPublish) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageToPublish.DiscardUnknown(m)
}

var xxx_messageInfo_MessageToPublish proto.InternalMessageInfo

func (m *MessageToPublish) GetExchangeName() string {
	if m != nil {
		return m.ExchangeName
	}
	return ""
}

func (m *MessageToPublish) GetRoutingKeyName() string {
	if m != nil {
		return m.RoutingKeyName
	}
	return ""
}

func (m *MessageToPublish) GetMessageTosend() []byte {
	if m != nil {
		return m.MessageTosend
	}
	return nil
}

func init() {
	proto.RegisterType((*MessageToPublish)(nil), "NotificationService.MessageToPublish")
}

func init() { proto.RegisterFile("NotificationService.proto", fileDescriptor_8a9b3db0602ab621) }

var fileDescriptor_8a9b3db0602ab621 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xf4, 0xcb, 0x2f, 0xc9,
	0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x0b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc6, 0x22, 0xa5, 0xd4, 0xc2, 0xc8, 0x25, 0xe0, 0x9b,
	0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0x1a, 0x92, 0x1f, 0x50, 0x9a, 0x94, 0x93, 0x59, 0x9c, 0x21, 0xa4,
	0xc4, 0xc5, 0x93, 0x5a, 0x91, 0x9c, 0x91, 0x98, 0x97, 0x9e, 0xea, 0x97, 0x98, 0x9b, 0x2a, 0xc1,
	0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x22, 0x26, 0xa4, 0xc6, 0xc5, 0x57, 0x94, 0x5f, 0x5a, 0x92,
	0x99, 0x97, 0xee, 0x9d, 0x5a, 0x09, 0x56, 0xc5, 0x04, 0x56, 0x85, 0x26, 0x2a, 0xa4, 0xc2, 0xc5,
	0x9b, 0x0b, 0x33, 0xbf, 0x38, 0x35, 0x2f, 0x45, 0x82, 0x59, 0x81, 0x51, 0x83, 0x27, 0x08, 0x55,
	0x30, 0x89, 0x0d, 0xec, 0x44, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xae, 0x8e, 0x04, 0x81,
	0xbf, 0x00, 0x00, 0x00,
}

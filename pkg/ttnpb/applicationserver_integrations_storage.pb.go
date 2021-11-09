// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/applicationserver_integrations_storage.proto

package ttnpb

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GetStoredApplicationUpRequest struct {
	// Query upstream messages from all end devices of an application. Cannot be used in conjunction with end_device_ids.
	ApplicationIds *ApplicationIdentifiers `protobuf:"bytes,1,opt,name=application_ids,json=applicationIds,proto3" json:"application_ids,omitempty"`
	// Query upstream messages from a single end device. Cannot be used in conjunction with application_ids.
	EndDeviceIds *EndDeviceIdentifiers `protobuf:"bytes,2,opt,name=end_device_ids,json=endDeviceIds,proto3" json:"end_device_ids,omitempty"`
	// Query upstream messages of a specific type. If not set, then all upstream messages are returned.
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// Limit number of results.
	Limit *types.UInt32Value `protobuf:"bytes,4,opt,name=limit,proto3" json:"limit,omitempty"`
	// Query upstream messages after this timestamp only. Cannot be used in conjunction with last.
	After *types.Timestamp `protobuf:"bytes,5,opt,name=after,proto3" json:"after,omitempty"`
	// Query upstream messages before this timestamp only. Cannot be used in conjunction with last.
	Before *types.Timestamp `protobuf:"bytes,6,opt,name=before,proto3" json:"before,omitempty"`
	// Query uplinks on a specific FPort only.
	FPort *types.UInt32Value `protobuf:"bytes,7,opt,name=f_port,json=fPort,proto3" json:"f_port,omitempty"`
	// Order results.
	Order string `protobuf:"bytes,8,opt,name=order,proto3" json:"order,omitempty"`
	// The names of the upstream message fields that should be returned. See the API reference
	// for allowed field names for each type of upstream message.
	FieldMask *types.FieldMask `protobuf:"bytes,9,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	// Query upstream messages that have arrived in the last minutes or hours. Cannot be used in conjunction with after and before.
	Last                 *types.Duration `protobuf:"bytes,10,opt,name=last,proto3" json:"last,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetStoredApplicationUpRequest) Reset()      { *m = GetStoredApplicationUpRequest{} }
func (*GetStoredApplicationUpRequest) ProtoMessage() {}
func (*GetStoredApplicationUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff0e9f52f73d254, []int{0}
}
func (m *GetStoredApplicationUpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStoredApplicationUpRequest.Unmarshal(m, b)
}
func (m *GetStoredApplicationUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStoredApplicationUpRequest.Marshal(b, m, deterministic)
}
func (m *GetStoredApplicationUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoredApplicationUpRequest.Merge(m, src)
}
func (m *GetStoredApplicationUpRequest) XXX_Size() int {
	return xxx_messageInfo_GetStoredApplicationUpRequest.Size(m)
}
func (m *GetStoredApplicationUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoredApplicationUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoredApplicationUpRequest proto.InternalMessageInfo

func (m *GetStoredApplicationUpRequest) GetApplicationIds() *ApplicationIdentifiers {
	if m != nil {
		return m.ApplicationIds
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetEndDeviceIds() *EndDeviceIdentifiers {
	if m != nil {
		return m.EndDeviceIds
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetStoredApplicationUpRequest) GetLimit() *types.UInt32Value {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetAfter() *types.Timestamp {
	if m != nil {
		return m.After
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetBefore() *types.Timestamp {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetFPort() *types.UInt32Value {
	if m != nil {
		return m.FPort
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

func (m *GetStoredApplicationUpRequest) GetFieldMask() *types.FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetLast() *types.Duration {
	if m != nil {
		return m.Last
	}
	return nil
}

type GetStoredApplicationUpCountRequest struct {
	// Count upstream messages from all end devices of an application. Cannot be used in conjunction with end_device_ids.
	ApplicationIds *ApplicationIdentifiers `protobuf:"bytes,1,opt,name=application_ids,json=applicationIds,proto3" json:"application_ids,omitempty"`
	// Count upstream messages from a single end device. Cannot be used in conjunction with application_ids.
	EndDeviceIds *EndDeviceIdentifiers `protobuf:"bytes,2,opt,name=end_device_ids,json=endDeviceIds,proto3" json:"end_device_ids,omitempty"`
	// Count upstream messages of a specific type. If not set, then all upstream messages are returned.
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// Count upstream messages after this timestamp only. Cannot be used in conjunction with last.
	After *types.Timestamp `protobuf:"bytes,4,opt,name=after,proto3" json:"after,omitempty"`
	// Count upstream messages before this timestamp only. Cannot be used in conjunction with last.
	Before *types.Timestamp `protobuf:"bytes,5,opt,name=before,proto3" json:"before,omitempty"`
	// Count uplinks on a specific FPort only.
	FPort *types.UInt32Value `protobuf:"bytes,6,opt,name=f_port,json=fPort,proto3" json:"f_port,omitempty"`
	// Count upstream messages that have arrived in the last minutes or hours. Cannot be used in conjunction with after and before.
	Last                 *types.Duration `protobuf:"bytes,7,opt,name=last,proto3" json:"last,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetStoredApplicationUpCountRequest) Reset()      { *m = GetStoredApplicationUpCountRequest{} }
func (*GetStoredApplicationUpCountRequest) ProtoMessage() {}
func (*GetStoredApplicationUpCountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff0e9f52f73d254, []int{1}
}
func (m *GetStoredApplicationUpCountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStoredApplicationUpCountRequest.Unmarshal(m, b)
}
func (m *GetStoredApplicationUpCountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStoredApplicationUpCountRequest.Marshal(b, m, deterministic)
}
func (m *GetStoredApplicationUpCountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoredApplicationUpCountRequest.Merge(m, src)
}
func (m *GetStoredApplicationUpCountRequest) XXX_Size() int {
	return xxx_messageInfo_GetStoredApplicationUpCountRequest.Size(m)
}
func (m *GetStoredApplicationUpCountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoredApplicationUpCountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoredApplicationUpCountRequest proto.InternalMessageInfo

func (m *GetStoredApplicationUpCountRequest) GetApplicationIds() *ApplicationIdentifiers {
	if m != nil {
		return m.ApplicationIds
	}
	return nil
}

func (m *GetStoredApplicationUpCountRequest) GetEndDeviceIds() *EndDeviceIdentifiers {
	if m != nil {
		return m.EndDeviceIds
	}
	return nil
}

func (m *GetStoredApplicationUpCountRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetStoredApplicationUpCountRequest) GetAfter() *types.Timestamp {
	if m != nil {
		return m.After
	}
	return nil
}

func (m *GetStoredApplicationUpCountRequest) GetBefore() *types.Timestamp {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *GetStoredApplicationUpCountRequest) GetFPort() *types.UInt32Value {
	if m != nil {
		return m.FPort
	}
	return nil
}

func (m *GetStoredApplicationUpCountRequest) GetLast() *types.Duration {
	if m != nil {
		return m.Last
	}
	return nil
}

type GetStoredApplicationUpCountResponse struct {
	// Number of stored messages by end device ID.
	Count                map[string]uint32 `protobuf:"bytes,1,rep,name=count,proto3" json:"count,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetStoredApplicationUpCountResponse) Reset()      { *m = GetStoredApplicationUpCountResponse{} }
func (*GetStoredApplicationUpCountResponse) ProtoMessage() {}
func (*GetStoredApplicationUpCountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff0e9f52f73d254, []int{2}
}
func (m *GetStoredApplicationUpCountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStoredApplicationUpCountResponse.Unmarshal(m, b)
}
func (m *GetStoredApplicationUpCountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStoredApplicationUpCountResponse.Marshal(b, m, deterministic)
}
func (m *GetStoredApplicationUpCountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoredApplicationUpCountResponse.Merge(m, src)
}
func (m *GetStoredApplicationUpCountResponse) XXX_Size() int {
	return xxx_messageInfo_GetStoredApplicationUpCountResponse.Size(m)
}
func (m *GetStoredApplicationUpCountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoredApplicationUpCountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoredApplicationUpCountResponse proto.InternalMessageInfo

func (m *GetStoredApplicationUpCountResponse) GetCount() map[string]uint32 {
	if m != nil {
		return m.Count
	}
	return nil
}

func init() {
	proto.RegisterType((*GetStoredApplicationUpRequest)(nil), "ttn.lorawan.v3.GetStoredApplicationUpRequest")
	golang_proto.RegisterType((*GetStoredApplicationUpRequest)(nil), "ttn.lorawan.v3.GetStoredApplicationUpRequest")
	proto.RegisterType((*GetStoredApplicationUpCountRequest)(nil), "ttn.lorawan.v3.GetStoredApplicationUpCountRequest")
	golang_proto.RegisterType((*GetStoredApplicationUpCountRequest)(nil), "ttn.lorawan.v3.GetStoredApplicationUpCountRequest")
	proto.RegisterType((*GetStoredApplicationUpCountResponse)(nil), "ttn.lorawan.v3.GetStoredApplicationUpCountResponse")
	golang_proto.RegisterType((*GetStoredApplicationUpCountResponse)(nil), "ttn.lorawan.v3.GetStoredApplicationUpCountResponse")
	proto.RegisterMapType((map[string]uint32)(nil), "ttn.lorawan.v3.GetStoredApplicationUpCountResponse.CountEntry")
	golang_proto.RegisterMapType((map[string]uint32)(nil), "ttn.lorawan.v3.GetStoredApplicationUpCountResponse.CountEntry")
}

func init() {
	proto.RegisterFile("lorawan-stack/api/applicationserver_integrations_storage.proto", fileDescriptor_6ff0e9f52f73d254)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/applicationserver_integrations_storage.proto", fileDescriptor_6ff0e9f52f73d254)
}

var fileDescriptor_6ff0e9f52f73d254 = []byte{
	// 921 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x56, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xf6, 0x38, 0xb6, 0x4b, 0x26, 0x69, 0x82, 0x46, 0x15, 0x5a, 0x4c, 0xe3, 0x5a, 0x2e, 0x42,
	0xb9, 0x78, 0xb7, 0xb2, 0x2f, 0x81, 0x43, 0xa5, 0x86, 0x16, 0x14, 0x50, 0x55, 0x98, 0x36, 0x1c,
	0x72, 0x59, 0x8d, 0x77, 0x9f, 0x37, 0x83, 0xd7, 0x33, 0xdb, 0x99, 0xb1, 0x83, 0x15, 0x45, 0x02,
	0xfe, 0x02, 0x24, 0xfe, 0x89, 0x4a, 0x1c, 0x38, 0x71, 0x43, 0x82, 0x03, 0x27, 0x6e, 0xc0, 0x01,
	0x4e, 0x48, 0xa4, 0x48, 0x70, 0xe4, 0xc4, 0xa1, 0x27, 0xe4, 0xd9, 0xf5, 0xef, 0x26, 0x75, 0x2b,
	0x04, 0x97, 0xde, 0xe6, 0xbd, 0xf9, 0xde, 0x37, 0x6f, 0xde, 0xfb, 0x9e, 0x66, 0xf0, 0xf5, 0x58,
	0x2a, 0x76, 0xc4, 0x44, 0x5d, 0x1b, 0x16, 0x74, 0x3c, 0x96, 0x70, 0x8f, 0x25, 0x49, 0xcc, 0x03,
	0x66, 0xb8, 0x14, 0x1a, 0x54, 0x1f, 0x94, 0xcf, 0x85, 0x81, 0x48, 0xa5, 0x1e, 0x5f, 0x1b, 0xa9,
	0x58, 0x04, 0x6e, 0xa2, 0xa4, 0x91, 0x64, 0xc3, 0x18, 0xe1, 0x66, 0x1c, 0x6e, 0xbf, 0x59, 0xbe,
	0x11, 0x71, 0x73, 0xd8, 0x6b, 0xb9, 0x81, 0xec, 0x7a, 0x20, 0xfa, 0x72, 0x90, 0x28, 0xf9, 0xd1,
	0xc0, 0xb3, 0xe0, 0xa0, 0x1e, 0x81, 0xa8, 0xf7, 0x59, 0xcc, 0x43, 0x66, 0xc0, 0x5b, 0x58, 0xa4,
	0x94, 0xe5, 0xfa, 0x14, 0x45, 0x24, 0x23, 0x99, 0x06, 0xb7, 0x7a, 0x6d, 0x6b, 0x59, 0xc3, 0xae,
	0x32, 0xf8, 0xe5, 0x48, 0xca, 0x28, 0x86, 0x34, 0x75, 0x21, 0xa4, 0x49, 0xf3, 0xcc, 0x76, 0xab,
	0xd9, 0xee, 0x98, 0xa3, 0xcd, 0x21, 0x0e, 0xfd, 0x2e, 0xd3, 0x9d, 0x0c, 0x71, 0x65, 0x1e, 0x61,
	0x78, 0x17, 0xb4, 0x61, 0xdd, 0x24, 0x03, 0x54, 0xe6, 0x01, 0x61, 0x2f, 0xad, 0xc5, 0x59, 0xfb,
	0x47, 0x8a, 0x25, 0x09, 0xa8, 0x51, 0x0a, 0x57, 0x17, 0x4b, 0xcc, 0x43, 0x10, 0x86, 0xb7, 0xf9,
	0x04, 0x54, 0x5d, 0x04, 0x75, 0x41, 0x6b, 0x16, 0x41, 0x86, 0xa8, 0xfd, 0x5d, 0xc4, 0x5b, 0x6f,
	0x83, 0xb9, 0x6b, 0xa4, 0x82, 0xf0, 0xc6, 0xa4, 0x47, 0xfb, 0x09, 0x85, 0xfb, 0x3d, 0xd0, 0x86,
	0xdc, 0xc1, 0x9b, 0x53, 0xbd, 0xf3, 0x79, 0xa8, 0x1d, 0x54, 0x45, 0xdb, 0x6b, 0x8d, 0xd7, 0xdc,
	0xd9, 0x2e, 0xb9, 0x53, 0xe1, 0x7b, 0x93, 0x54, 0xe8, 0x06, 0x9b, 0xf6, 0x6b, 0xf2, 0x0e, 0xde,
	0x00, 0x11, 0xfa, 0x21, 0xf4, 0x79, 0x00, 0x96, 0x2f, 0x6f, 0xf9, 0x5e, 0x9d, 0xe7, 0xbb, 0x25,
	0xc2, 0x9b, 0x16, 0x34, 0xcd, 0xb6, 0x0e, 0x13, 0xaf, 0x26, 0xdf, 0x21, 0x5c, 0x30, 0x83, 0x04,
	0x9c, 0x95, 0x2a, 0xda, 0x5e, 0xdd, 0xfd, 0x12, 0x3d, 0xda, 0xfd, 0x02, 0xa9, 0x07, 0x88, 0xe6,
	0xe8, 0x46, 0x2f, 0x89, 0xb9, 0xe8, 0xf8, 0xd9, 0x85, 0xe9, 0xda, 0x87, 0x92, 0x0b, 0x9f, 0x05,
	0x01, 0x24, 0x86, 0xae, 0x87, 0xf2, 0x48, 0xd8, 0x6d, 0x16, 0x74, 0xe8, 0xc5, 0xb1, 0x25, 0x66,
	0x4d, 0x0d, 0xc2, 0xd0, 0xcd, 0xb1, 0xd9, 0x66, 0x3c, 0x86, 0x70, 0xca, 0x71, 0xbf, 0x07, 0x3d,
	0x08, 0x69, 0x79, 0xd6, 0xe1, 0x73, 0x31, 0x12, 0x5f, 0x48, 0x37, 0x63, 0x99, 0x55, 0x4e, 0xcb,
	0xb8, 0x0f, 0x21, 0x5d, 0x1f, 0xca, 0x7f, 0x78, 0xf3, 0x90, 0x19, 0x46, 0x6d, 0xf6, 0xa4, 0x81,
	0x8b, 0x31, 0xef, 0x72, 0xe3, 0x14, 0x6c, 0x25, 0x2e, 0xbb, 0x69, 0xf3, 0xdd, 0x51, 0xf3, 0xdd,
	0xfd, 0x3d, 0x61, 0x9a, 0x8d, 0x0f, 0x58, 0xdc, 0x03, 0x9a, 0x42, 0xc9, 0x35, 0x5c, 0x64, 0x6d,
	0x03, 0xca, 0x29, 0xda, 0x98, 0xf2, 0x42, 0xcc, 0xbd, 0x91, 0xe2, 0x68, 0x0a, 0x24, 0x0d, 0x5c,
	0x6a, 0x41, 0x5b, 0x2a, 0x70, 0x4a, 0x4f, 0x0c, 0xc9, 0x90, 0xa4, 0x89, 0x4b, 0x6d, 0x3f, 0x91,
	0xca, 0x38, 0x17, 0x96, 0x49, 0xad, 0xfd, 0x9e, 0x54, 0x86, 0xec, 0xe0, 0xa2, 0x54, 0x21, 0x28,
	0xe7, 0x05, 0xdb, 0x95, 0xda, 0xa3, 0xdd, 0x2b, 0x6a, 0x8b, 0xe6, 0xe8, 0x7a, 0x5d, 0x41, 0x00,
	0xbc, 0x0f, 0xa1, 0xcf, 0x0c, 0x5d, 0x9b, 0x36, 0xd2, 0x00, 0xf2, 0x3a, 0xc6, 0x93, 0x51, 0x72,
	0x56, 0xcf, 0x48, 0xf3, 0xad, 0x21, 0xe4, 0x36, 0xd3, 0x1d, 0xba, 0xda, 0x1e, 0x2d, 0x49, 0x1d,
	0x17, 0x62, 0xa6, 0x8d, 0x83, 0x6d, 0xd0, 0xcb, 0x0b, 0x41, 0x37, 0xb3, 0xf9, 0xa2, 0x16, 0x56,
	0xfb, 0xb1, 0x80, 0x6b, 0x8f, 0x17, 0xfe, 0x9b, 0xb2, 0x27, 0xcc, 0x73, 0xf5, 0xff, 0x77, 0xea,
	0x1f, 0x2b, 0xb9, 0xf0, 0xf4, 0x4a, 0x2e, 0x3e, 0x83, 0x92, 0x4b, 0xcb, 0x2b, 0x79, 0x24, 0xaa,
	0x0b, 0xcb, 0x89, 0xea, 0x2b, 0x84, 0xaf, 0x9e, 0x2b, 0x2a, 0x9d, 0x0c, 0x9f, 0x3f, 0x72, 0x0f,
	0x17, 0x83, 0xa1, 0xc3, 0x41, 0xd5, 0x95, 0xed, 0xb5, 0xc6, 0xf5, 0xf9, 0xde, 0x2f, 0xc1, 0xe1,
	0x5a, 0xeb, 0x96, 0x30, 0x6a, 0x40, 0x53, 0xb2, 0xf2, 0x0e, 0xc6, 0x13, 0x27, 0x79, 0x11, 0xaf,
	0x74, 0x60, 0x60, 0xd5, 0xba, 0x4a, 0x87, 0x4b, 0x72, 0x09, 0x17, 0xfb, 0xc3, 0xcb, 0x59, 0xc5,
	0x5d, 0xa4, 0xa9, 0xf1, 0x46, 0x7e, 0x07, 0x35, 0xbe, 0x29, 0xe2, 0x4b, 0x33, 0x47, 0xdd, 0x4d,
	0x9f, 0x63, 0xf2, 0x75, 0x1e, 0xbf, 0xf4, 0xf8, 0x64, 0x48, 0x7d, 0xb9, 0xa4, 0xb3, 0x41, 0x2a,
	0x6f, 0x9d, 0x33, 0x2f, 0xfb, 0x49, 0xed, 0x07, 0xf4, 0xe9, 0x4f, 0xbf, 0x7f, 0x9e, 0xff, 0x1e,
	0x91, 0x63, 0x8f, 0xe9, 0x99, 0xdf, 0x82, 0x77, 0x3c, 0x3b, 0x30, 0xee, 0xdc, 0x40, 0xce, 0xd9,
	0x27, 0x5e, 0x0a, 0x5d, 0x8c, 0x1b, 0x2f, 0x4f, 0xbc, 0x84, 0x05, 0x9d, 0xe1, 0x2b, 0xe8, 0x65,
	0xff, 0x0e, 0xef, 0x78, 0xa8, 0xc4, 0x93, 0x83, 0x77, 0xc9, 0xde, 0xe2, 0xf1, 0x4f, 0x3a, 0xef,
	0x0c, 0xb2, 0x6b, 0x88, 0xfc, 0x91, 0xc7, 0xaf, 0x9c, 0xd3, 0x4b, 0xd2, 0x78, 0xaa, 0xc6, 0xa7,
	0x85, 0x6c, 0x3e, 0x83, 0x58, 0x6a, 0xbf, 0xa6, 0xe5, 0xfd, 0x19, 0x91, 0x4f, 0xd0, 0xff, 0x58,
	0x5f, 0xcf, 0x0a, 0xf5, 0xe0, 0x7d, 0x72, 0xe7, 0x5f, 0xab, 0x72, 0x4a, 0xb9, 0x7b, 0xfb, 0x97,
	0xdf, 0x2a, 0xb9, 0x8f, 0x4f, 0x2b, 0xe8, 0xc1, 0x69, 0x05, 0xfd, 0x79, 0x5a, 0xc9, 0xfd, 0x75,
	0x5a, 0x41, 0x9f, 0x3d, 0xac, 0xe4, 0xbe, 0x7d, 0x58, 0x41, 0x07, 0x5e, 0x24, 0x5d, 0x73, 0x08,
	0xe6, 0x90, 0x8b, 0x48, 0xbb, 0x02, 0xcc, 0x91, 0x54, 0x1d, 0x6f, 0xf6, 0x73, 0xd4, 0x6f, 0x7a,
	0x49, 0x27, 0xf2, 0x8c, 0x11, 0x49, 0xab, 0x55, 0xb2, 0x13, 0xde, 0xfc, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x92, 0xe8, 0xc8, 0x6b, 0xc9, 0x0a, 0x00, 0x00,
}

func (this *GetStoredApplicationUpRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetStoredApplicationUpRequest)
	if !ok {
		that2, ok := that.(GetStoredApplicationUpRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ApplicationIds.Equal(that1.ApplicationIds) {
		return false
	}
	if !this.EndDeviceIds.Equal(that1.EndDeviceIds) {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if !this.Limit.Equal(that1.Limit) {
		return false
	}
	if !this.After.Equal(that1.After) {
		return false
	}
	if !this.Before.Equal(that1.Before) {
		return false
	}
	if !this.FPort.Equal(that1.FPort) {
		return false
	}
	if this.Order != that1.Order {
		return false
	}
	if !this.FieldMask.Equal(that1.FieldMask) {
		return false
	}
	if !this.Last.Equal(that1.Last) {
		return false
	}
	return true
}
func (this *GetStoredApplicationUpCountRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetStoredApplicationUpCountRequest)
	if !ok {
		that2, ok := that.(GetStoredApplicationUpCountRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ApplicationIds.Equal(that1.ApplicationIds) {
		return false
	}
	if !this.EndDeviceIds.Equal(that1.EndDeviceIds) {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if !this.After.Equal(that1.After) {
		return false
	}
	if !this.Before.Equal(that1.Before) {
		return false
	}
	if !this.FPort.Equal(that1.FPort) {
		return false
	}
	if !this.Last.Equal(that1.Last) {
		return false
	}
	return true
}
func (this *GetStoredApplicationUpCountResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetStoredApplicationUpCountResponse)
	if !ok {
		that2, ok := that.(GetStoredApplicationUpCountResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Count) != len(that1.Count) {
		return false
	}
	for i := range this.Count {
		if this.Count[i] != that1.Count[i] {
			return false
		}
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApplicationUpStorageClient is the client API for ApplicationUpStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApplicationUpStorageClient interface {
	// Returns a stream of application messages that have been stored in the database.
	GetStoredApplicationUp(ctx context.Context, in *GetStoredApplicationUpRequest, opts ...grpc.CallOption) (ApplicationUpStorage_GetStoredApplicationUpClient, error)
	// Returns how many application messages have been stored in the database for an application or end device.
	GetStoredApplicationUpCount(ctx context.Context, in *GetStoredApplicationUpCountRequest, opts ...grpc.CallOption) (*GetStoredApplicationUpCountResponse, error)
}

type applicationUpStorageClient struct {
	cc *grpc.ClientConn
}

func NewApplicationUpStorageClient(cc *grpc.ClientConn) ApplicationUpStorageClient {
	return &applicationUpStorageClient{cc}
}

func (c *applicationUpStorageClient) GetStoredApplicationUp(ctx context.Context, in *GetStoredApplicationUpRequest, opts ...grpc.CallOption) (ApplicationUpStorage_GetStoredApplicationUpClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ApplicationUpStorage_serviceDesc.Streams[0], "/ttn.lorawan.v3.ApplicationUpStorage/GetStoredApplicationUp", opts...)
	if err != nil {
		return nil, err
	}
	x := &applicationUpStorageGetStoredApplicationUpClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ApplicationUpStorage_GetStoredApplicationUpClient interface {
	Recv() (*ApplicationUp, error)
	grpc.ClientStream
}

type applicationUpStorageGetStoredApplicationUpClient struct {
	grpc.ClientStream
}

func (x *applicationUpStorageGetStoredApplicationUpClient) Recv() (*ApplicationUp, error) {
	m := new(ApplicationUp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *applicationUpStorageClient) GetStoredApplicationUpCount(ctx context.Context, in *GetStoredApplicationUpCountRequest, opts ...grpc.CallOption) (*GetStoredApplicationUpCountResponse, error) {
	out := new(GetStoredApplicationUpCountResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationUpStorage/GetStoredApplicationUpCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationUpStorageServer is the server API for ApplicationUpStorage service.
type ApplicationUpStorageServer interface {
	// Returns a stream of application messages that have been stored in the database.
	GetStoredApplicationUp(*GetStoredApplicationUpRequest, ApplicationUpStorage_GetStoredApplicationUpServer) error
	// Returns how many application messages have been stored in the database for an application or end device.
	GetStoredApplicationUpCount(context.Context, *GetStoredApplicationUpCountRequest) (*GetStoredApplicationUpCountResponse, error)
}

// UnimplementedApplicationUpStorageServer can be embedded to have forward compatible implementations.
type UnimplementedApplicationUpStorageServer struct {
}

func (*UnimplementedApplicationUpStorageServer) GetStoredApplicationUp(req *GetStoredApplicationUpRequest, srv ApplicationUpStorage_GetStoredApplicationUpServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStoredApplicationUp not implemented")
}
func (*UnimplementedApplicationUpStorageServer) GetStoredApplicationUpCount(ctx context.Context, req *GetStoredApplicationUpCountRequest) (*GetStoredApplicationUpCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStoredApplicationUpCount not implemented")
}

func RegisterApplicationUpStorageServer(s *grpc.Server, srv ApplicationUpStorageServer) {
	s.RegisterService(&_ApplicationUpStorage_serviceDesc, srv)
}

func _ApplicationUpStorage_GetStoredApplicationUp_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetStoredApplicationUpRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ApplicationUpStorageServer).GetStoredApplicationUp(m, &applicationUpStorageGetStoredApplicationUpServer{stream})
}

type ApplicationUpStorage_GetStoredApplicationUpServer interface {
	Send(*ApplicationUp) error
	grpc.ServerStream
}

type applicationUpStorageGetStoredApplicationUpServer struct {
	grpc.ServerStream
}

func (x *applicationUpStorageGetStoredApplicationUpServer) Send(m *ApplicationUp) error {
	return x.ServerStream.SendMsg(m)
}

func _ApplicationUpStorage_GetStoredApplicationUpCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoredApplicationUpCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationUpStorageServer).GetStoredApplicationUpCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationUpStorage/GetStoredApplicationUpCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationUpStorageServer).GetStoredApplicationUpCount(ctx, req.(*GetStoredApplicationUpCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApplicationUpStorage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.ApplicationUpStorage",
	HandlerType: (*ApplicationUpStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStoredApplicationUpCount",
			Handler:    _ApplicationUpStorage_GetStoredApplicationUpCount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStoredApplicationUp",
			Handler:       _ApplicationUpStorage_GetStoredApplicationUp_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "lorawan-stack/api/applicationserver_integrations_storage.proto",
}

func (this *GetStoredApplicationUpRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetStoredApplicationUpRequest{`,
		`ApplicationIds:` + strings.Replace(fmt.Sprintf("%v", this.ApplicationIds), "ApplicationIdentifiers", "ApplicationIdentifiers", 1) + `,`,
		`EndDeviceIds:` + strings.Replace(fmt.Sprintf("%v", this.EndDeviceIds), "EndDeviceIdentifiers", "EndDeviceIdentifiers", 1) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Limit:` + strings.Replace(fmt.Sprintf("%v", this.Limit), "UInt32Value", "types.UInt32Value", 1) + `,`,
		`After:` + strings.Replace(fmt.Sprintf("%v", this.After), "Timestamp", "types.Timestamp", 1) + `,`,
		`Before:` + strings.Replace(fmt.Sprintf("%v", this.Before), "Timestamp", "types.Timestamp", 1) + `,`,
		`FPort:` + strings.Replace(fmt.Sprintf("%v", this.FPort), "UInt32Value", "types.UInt32Value", 1) + `,`,
		`Order:` + fmt.Sprintf("%v", this.Order) + `,`,
		`FieldMask:` + strings.Replace(fmt.Sprintf("%v", this.FieldMask), "FieldMask", "types.FieldMask", 1) + `,`,
		`Last:` + strings.Replace(fmt.Sprintf("%v", this.Last), "Duration", "types.Duration", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetStoredApplicationUpCountRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetStoredApplicationUpCountRequest{`,
		`ApplicationIds:` + strings.Replace(fmt.Sprintf("%v", this.ApplicationIds), "ApplicationIdentifiers", "ApplicationIdentifiers", 1) + `,`,
		`EndDeviceIds:` + strings.Replace(fmt.Sprintf("%v", this.EndDeviceIds), "EndDeviceIdentifiers", "EndDeviceIdentifiers", 1) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`After:` + strings.Replace(fmt.Sprintf("%v", this.After), "Timestamp", "types.Timestamp", 1) + `,`,
		`Before:` + strings.Replace(fmt.Sprintf("%v", this.Before), "Timestamp", "types.Timestamp", 1) + `,`,
		`FPort:` + strings.Replace(fmt.Sprintf("%v", this.FPort), "UInt32Value", "types.UInt32Value", 1) + `,`,
		`Last:` + strings.Replace(fmt.Sprintf("%v", this.Last), "Duration", "types.Duration", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetStoredApplicationUpCountResponse) String() string {
	if this == nil {
		return "nil"
	}
	keysForCount := make([]string, 0, len(this.Count))
	for k := range this.Count {
		keysForCount = append(keysForCount, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForCount)
	mapStringForCount := "map[string]uint32{"
	for _, k := range keysForCount {
		mapStringForCount += fmt.Sprintf("%v: %v,", k, this.Count[k])
	}
	mapStringForCount += "}"
	s := strings.Join([]string{`&GetStoredApplicationUpCountResponse{`,
		`Count:` + mapStringForCount + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringApplicationserverIntegrationsStorage(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}

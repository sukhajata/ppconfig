// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config-service.proto

package config

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SetDesiredRequest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	FieldName            string   `protobuf:"bytes,2,opt,name=fieldName,proto3" json:"fieldName,omitempty"`
	FieldValue           string   `protobuf:"bytes,3,opt,name=fieldValue,proto3" json:"fieldValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetDesiredRequest) Reset()         { *m = SetDesiredRequest{} }
func (m *SetDesiredRequest) String() string { return proto.CompactTextString(m) }
func (*SetDesiredRequest) ProtoMessage()    {}
func (*SetDesiredRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f48c1c3d9797070f, []int{0}
}

func (m *SetDesiredRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDesiredRequest.Unmarshal(m, b)
}
func (m *SetDesiredRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDesiredRequest.Marshal(b, m, deterministic)
}
func (m *SetDesiredRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDesiredRequest.Merge(m, src)
}
func (m *SetDesiredRequest) XXX_Size() int {
	return xxx_messageInfo_SetDesiredRequest.Size(m)
}
func (m *SetDesiredRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDesiredRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetDesiredRequest proto.InternalMessageInfo

func (m *SetDesiredRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *SetDesiredRequest) GetFieldName() string {
	if m != nil {
		return m.FieldName
	}
	return ""
}

func (m *SetDesiredRequest) GetFieldValue() string {
	if m != nil {
		return m.FieldValue
	}
	return ""
}

type CheckConsistencyRequest struct {
	DeviceEUI            string   `protobuf:"bytes,1,opt,name=deviceEUI,proto3" json:"deviceEUI,omitempty"`
	Slot                 int32    `protobuf:"varint,2,opt,name=slot,proto3" json:"slot,omitempty"`
	FieldIndex           int32    `protobuf:"varint,3,opt,name=fieldIndex,proto3" json:"fieldIndex,omitempty"`
	NumRetries           int32    `protobuf:"varint,4,opt,name=numRetries,proto3" json:"numRetries,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckConsistencyRequest) Reset()         { *m = CheckConsistencyRequest{} }
func (m *CheckConsistencyRequest) String() string { return proto.CompactTextString(m) }
func (*CheckConsistencyRequest) ProtoMessage()    {}
func (*CheckConsistencyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f48c1c3d9797070f, []int{1}
}

func (m *CheckConsistencyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckConsistencyRequest.Unmarshal(m, b)
}
func (m *CheckConsistencyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckConsistencyRequest.Marshal(b, m, deterministic)
}
func (m *CheckConsistencyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckConsistencyRequest.Merge(m, src)
}
func (m *CheckConsistencyRequest) XXX_Size() int {
	return xxx_messageInfo_CheckConsistencyRequest.Size(m)
}
func (m *CheckConsistencyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckConsistencyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckConsistencyRequest proto.InternalMessageInfo

func (m *CheckConsistencyRequest) GetDeviceEUI() string {
	if m != nil {
		return m.DeviceEUI
	}
	return ""
}

func (m *CheckConsistencyRequest) GetSlot() int32 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *CheckConsistencyRequest) GetFieldIndex() int32 {
	if m != nil {
		return m.FieldIndex
	}
	return 0
}

func (m *CheckConsistencyRequest) GetNumRetries() int32 {
	if m != nil {
		return m.NumRetries
	}
	return 0
}

type UpdateReportedRequest struct {
	DeviceEUI            string   `protobuf:"bytes,1,opt,name=deviceEUI,proto3" json:"deviceEUI,omitempty"`
	FieldIndex           int32    `protobuf:"varint,2,opt,name=fieldIndex,proto3" json:"fieldIndex,omitempty"`
	FieldValue           string   `protobuf:"bytes,3,opt,name=fieldValue,proto3" json:"fieldValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateReportedRequest) Reset()         { *m = UpdateReportedRequest{} }
func (m *UpdateReportedRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateReportedRequest) ProtoMessage()    {}
func (*UpdateReportedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f48c1c3d9797070f, []int{2}
}

func (m *UpdateReportedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateReportedRequest.Unmarshal(m, b)
}
func (m *UpdateReportedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateReportedRequest.Marshal(b, m, deterministic)
}
func (m *UpdateReportedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateReportedRequest.Merge(m, src)
}
func (m *UpdateReportedRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateReportedRequest.Size(m)
}
func (m *UpdateReportedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateReportedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateReportedRequest proto.InternalMessageInfo

func (m *UpdateReportedRequest) GetDeviceEUI() string {
	if m != nil {
		return m.DeviceEUI
	}
	return ""
}

func (m *UpdateReportedRequest) GetFieldIndex() int32 {
	if m != nil {
		return m.FieldIndex
	}
	return 0
}

func (m *UpdateReportedRequest) GetFieldValue() string {
	if m != nil {
		return m.FieldValue
	}
	return ""
}

type Response struct {
	Reply                string   `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_f48c1c3d9797070f, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

type ConfigFields struct {
	Fields               []*ConfigField `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ConfigFields) Reset()         { *m = ConfigFields{} }
func (m *ConfigFields) String() string { return proto.CompactTextString(m) }
func (*ConfigFields) ProtoMessage()    {}
func (*ConfigFields) Descriptor() ([]byte, []int) {
	return fileDescriptor_f48c1c3d9797070f, []int{4}
}

func (m *ConfigFields) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFields.Unmarshal(m, b)
}
func (m *ConfigFields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFields.Marshal(b, m, deterministic)
}
func (m *ConfigFields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFields.Merge(m, src)
}
func (m *ConfigFields) XXX_Size() int {
	return xxx_messageInfo_ConfigFields.Size(m)
}
func (m *ConfigFields) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFields.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFields proto.InternalMessageInfo

func (m *ConfigFields) GetFields() []*ConfigField {
	if m != nil {
		return m.Fields
	}
	return nil
}

type ConfigField struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Index                int32    `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Desired              string   `protobuf:"bytes,3,opt,name=desired,proto3" json:"desired,omitempty"`
	Reported             string   `protobuf:"bytes,4,opt,name=reported,proto3" json:"reported,omitempty"`
	FieldType            string   `protobuf:"bytes,5,opt,name=fieldType,proto3" json:"fieldType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigField) Reset()         { *m = ConfigField{} }
func (m *ConfigField) String() string { return proto.CompactTextString(m) }
func (*ConfigField) ProtoMessage()    {}
func (*ConfigField) Descriptor() ([]byte, []int) {
	return fileDescriptor_f48c1c3d9797070f, []int{5}
}

func (m *ConfigField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigField.Unmarshal(m, b)
}
func (m *ConfigField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigField.Marshal(b, m, deterministic)
}
func (m *ConfigField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigField.Merge(m, src)
}
func (m *ConfigField) XXX_Size() int {
	return xxx_messageInfo_ConfigField.Size(m)
}
func (m *ConfigField) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigField.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigField proto.InternalMessageInfo

func (m *ConfigField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ConfigField) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ConfigField) GetDesired() string {
	if m != nil {
		return m.Desired
	}
	return ""
}

func (m *ConfigField) GetReported() string {
	if m != nil {
		return m.Reported
	}
	return ""
}

func (m *ConfigField) GetFieldType() string {
	if m != nil {
		return m.FieldType
	}
	return ""
}

type NewConfigRequest struct {
	DeviceEUI            string   `protobuf:"bytes,1,opt,name=deviceEUI,proto3" json:"deviceEUI,omitempty"`
	Identifier           string   `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewConfigRequest) Reset()         { *m = NewConfigRequest{} }
func (m *NewConfigRequest) String() string { return proto.CompactTextString(m) }
func (*NewConfigRequest) ProtoMessage()    {}
func (*NewConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f48c1c3d9797070f, []int{6}
}

func (m *NewConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewConfigRequest.Unmarshal(m, b)
}
func (m *NewConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewConfigRequest.Marshal(b, m, deterministic)
}
func (m *NewConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewConfigRequest.Merge(m, src)
}
func (m *NewConfigRequest) XXX_Size() int {
	return xxx_messageInfo_NewConfigRequest.Size(m)
}
func (m *NewConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewConfigRequest proto.InternalMessageInfo

func (m *NewConfigRequest) GetDeviceEUI() string {
	if m != nil {
		return m.DeviceEUI
	}
	return ""
}

func (m *NewConfigRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type Identifier struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identifier) Reset()         { *m = Identifier{} }
func (m *Identifier) String() string { return proto.CompactTextString(m) }
func (*Identifier) ProtoMessage()    {}
func (*Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_f48c1c3d9797070f, []int{7}
}

func (m *Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identifier.Unmarshal(m, b)
}
func (m *Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identifier.Marshal(b, m, deterministic)
}
func (m *Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identifier.Merge(m, src)
}
func (m *Identifier) XXX_Size() int {
	return xxx_messageInfo_Identifier.Size(m)
}
func (m *Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_Identifier proto.InternalMessageInfo

func (m *Identifier) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type GetConfigByNameRequest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	FieldName            string   `protobuf:"bytes,2,opt,name=fieldName,proto3" json:"fieldName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetConfigByNameRequest) Reset()         { *m = GetConfigByNameRequest{} }
func (m *GetConfigByNameRequest) String() string { return proto.CompactTextString(m) }
func (*GetConfigByNameRequest) ProtoMessage()    {}
func (*GetConfigByNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f48c1c3d9797070f, []int{8}
}

func (m *GetConfigByNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConfigByNameRequest.Unmarshal(m, b)
}
func (m *GetConfigByNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConfigByNameRequest.Marshal(b, m, deterministic)
}
func (m *GetConfigByNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConfigByNameRequest.Merge(m, src)
}
func (m *GetConfigByNameRequest) XXX_Size() int {
	return xxx_messageInfo_GetConfigByNameRequest.Size(m)
}
func (m *GetConfigByNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConfigByNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetConfigByNameRequest proto.InternalMessageInfo

func (m *GetConfigByNameRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *GetConfigByNameRequest) GetFieldName() string {
	if m != nil {
		return m.FieldName
	}
	return ""
}

func init() {
	proto.RegisterType((*SetDesiredRequest)(nil), "config.SetDesiredRequest")
	proto.RegisterType((*CheckConsistencyRequest)(nil), "config.CheckConsistencyRequest")
	proto.RegisterType((*UpdateReportedRequest)(nil), "config.UpdateReportedRequest")
	proto.RegisterType((*Response)(nil), "config.Response")
	proto.RegisterType((*ConfigFields)(nil), "config.ConfigFields")
	proto.RegisterType((*ConfigField)(nil), "config.ConfigField")
	proto.RegisterType((*NewConfigRequest)(nil), "config.NewConfigRequest")
	proto.RegisterType((*Identifier)(nil), "config.Identifier")
	proto.RegisterType((*GetConfigByNameRequest)(nil), "config.GetConfigByNameRequest")
}

func init() { proto.RegisterFile("config-service.proto", fileDescriptor_f48c1c3d9797070f) }

var fileDescriptor_f48c1c3d9797070f = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0xcd, 0x9f, 0x26, 0xbf, 0x66, 0xda, 0x1f, 0x0d, 0xd3, 0x00, 0x26, 0x2a, 0x25, 0xda, 0x53,
	0x25, 0x20, 0x87, 0x72, 0xa3, 0x42, 0x08, 0x0c, 0x84, 0x5c, 0xaa, 0xca, 0xa5, 0xbd, 0x87, 0x78,
	0x52, 0x56, 0xb8, 0x6b, 0xd7, 0xbb, 0x01, 0xf2, 0x19, 0xf8, 0x92, 0x7c, 0x14, 0xb4, 0xbb, 0x5e,
	0xdb, 0x0d, 0xab, 0xc0, 0x81, 0x9b, 0xe7, 0xed, 0xcc, 0xbc, 0xc9, 0x9b, 0x37, 0x81, 0xc1, 0x3c,
	0x15, 0x0b, 0x7e, 0xf5, 0x4c, 0x52, 0xfe, 0x95, 0xcf, 0x69, 0x9c, 0xe5, 0xa9, 0x4a, 0xb1, 0x6b,
	0x51, 0x76, 0x03, 0x77, 0xcf, 0x49, 0xbd, 0x25, 0xc9, 0x73, 0x8a, 0x23, 0xba, 0x59, 0x92, 0x54,
	0x78, 0x08, 0xc0, 0x63, 0x12, 0x8a, 0x2f, 0x38, 0xe5, 0x41, 0x73, 0xd4, 0x3c, 0xea, 0x45, 0x35,
	0x04, 0x0f, 0xa0, 0xb7, 0xe0, 0x94, 0xc4, 0xa7, 0xb3, 0x6b, 0x0a, 0x5a, 0xe6, 0xb9, 0x02, 0x74,
	0xb5, 0x09, 0x2e, 0x67, 0xc9, 0x92, 0x82, 0xb6, 0xad, 0xae, 0x10, 0xf6, 0xa3, 0x09, 0x0f, 0xc2,
	0xcf, 0x34, 0xff, 0x12, 0xa6, 0x42, 0x72, 0xa9, 0x48, 0xcc, 0x57, 0x8e, 0xf9, 0x00, 0x7a, 0x31,
	0xe9, 0x31, 0xdf, 0x5d, 0x4c, 0x0b, 0xe2, 0x0a, 0x40, 0x84, 0x2d, 0x99, 0xa4, 0xca, 0x50, 0x76,
	0x22, 0xf3, 0x5d, 0xb2, 0x4d, 0x45, 0x4c, 0xdf, 0x0d, 0x5b, 0x27, 0xaa, 0x21, 0xfa, 0x5d, 0x2c,
	0xaf, 0x23, 0x52, 0x39, 0x27, 0x19, 0x6c, 0xd9, 0xf7, 0x0a, 0x61, 0x4b, 0xb8, 0x77, 0x91, 0xc5,
	0x33, 0x45, 0x11, 0x65, 0x69, 0xae, 0x2a, 0x11, 0x36, 0x8f, 0x72, 0x9b, 0xb6, 0xe5, 0xa3, 0xdd,
	0x28, 0xc2, 0x08, 0xb6, 0x23, 0x92, 0x59, 0x2a, 0x24, 0xe1, 0x00, 0x3a, 0x39, 0x65, 0xc9, 0xaa,
	0x60, 0xb1, 0x01, 0x3b, 0x81, 0xdd, 0xd0, 0xec, 0xe8, 0xbd, 0xae, 0x92, 0xf8, 0x04, 0xba, 0xa6,
	0x5e, 0x06, 0xcd, 0x51, 0xfb, 0x68, 0xe7, 0x78, 0x7f, 0x6c, 0x57, 0x38, 0xae, 0x65, 0x45, 0x45,
	0x8a, 0xd6, 0x78, 0xa7, 0x86, 0x6b, 0xe5, 0x84, 0x5e, 0x96, 0x65, 0x30, 0xdf, 0x9a, 0x96, 0xd7,
	0xa6, 0xb7, 0x01, 0x06, 0xf0, 0x5f, 0x6c, 0xdd, 0x50, 0x4c, 0xed, 0x42, 0x1c, 0xc2, 0x76, 0x5e,
	0x68, 0x64, 0x74, 0xec, 0x45, 0x65, 0x5c, 0x3a, 0xe2, 0xe3, 0x2a, 0xa3, 0xa0, 0x53, 0x73, 0x84,
	0x06, 0xd8, 0x19, 0xf4, 0x4f, 0xe9, 0x9b, 0x9d, 0xe7, 0xaf, 0xe5, 0xad, 0x39, 0xb0, 0xb5, 0xee,
	0x40, 0xf6, 0x14, 0x60, 0x5a, 0xf9, 0xf1, 0x0f, 0x7e, 0x65, 0x97, 0x70, 0x7f, 0x42, 0xca, 0xf2,
	0xbf, 0x59, 0x69, 0x93, 0xfe, 0x13, 0xa7, 0x1f, 0xff, 0x6c, 0xc3, 0xff, 0xb6, 0xeb, 0xb9, 0x3d,
	0x2e, 0x3c, 0x01, 0xa8, 0xce, 0x09, 0x1f, 0xba, 0x15, 0xfd, 0x76, 0x62, 0xc3, 0xbe, 0x7b, 0x72,
	0x2e, 0x60, 0x0d, 0x9c, 0x40, 0x7f, 0xfd, 0x2e, 0xf0, 0x71, 0xb9, 0x65, 0xff, 0xc5, 0x78, 0x1b,
	0x85, 0x70, 0xe7, 0xb6, 0xa7, 0xf1, 0x91, 0xcb, 0xf2, 0x7a, 0xdd, 0xdb, 0xe4, 0x03, 0xec, 0xad,
	0x89, 0x86, 0x87, 0x2e, 0xcd, 0xaf, 0xe6, 0xd0, 0x67, 0x49, 0xd6, 0xc0, 0x17, 0xb0, 0x3b, 0x21,
	0xf5, 0x3a, 0x49, 0x2c, 0x8c, 0xe8, 0xd2, 0xaa, 0x15, 0x0e, 0x07, 0x9e, 0x52, 0xc9, 0x1a, 0xf8,
	0x12, 0xf6, 0xc3, 0x9c, 0x66, 0x8a, 0xce, 0x48, 0xc4, 0x5c, 0x5c, 0x6d, 0x68, 0xe1, 0xfb, 0x11,
	0xaf, 0x60, 0xcf, 0x96, 0x97, 0xfe, 0xc3, 0xc0, 0xa5, 0xad, 0x5b, 0xd2, 0xd7, 0xe0, 0x53, 0xd7,
	0xfc, 0x5d, 0x3e, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xcb, 0x18, 0xfe, 0x46, 0x05, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfigServiceClient is the client API for ConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigServiceClient interface {
	SetDesired(ctx context.Context, in *SetDesiredRequest, opts ...grpc.CallOption) (*Response, error)
	CheckConsistency(ctx context.Context, in *CheckConsistencyRequest, opts ...grpc.CallOption) (*Response, error)
	UpdateReported(ctx context.Context, in *UpdateReportedRequest, opts ...grpc.CallOption) (*Response, error)
	GetConfigByName(ctx context.Context, in *GetConfigByNameRequest, opts ...grpc.CallOption) (*ConfigField, error)
	GetAllConfig(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*ConfigFields, error)
	CreatePendingConfig(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Response, error)
	CreateNewConfig(ctx context.Context, in *NewConfigRequest, opts ...grpc.CallOption) (*Response, error)
}

type configServiceClient struct {
	cc *grpc.ClientConn
}

func NewConfigServiceClient(cc *grpc.ClientConn) ConfigServiceClient {
	return &configServiceClient{cc}
}

func (c *configServiceClient) SetDesired(ctx context.Context, in *SetDesiredRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/config.ConfigService/SetDesired", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) CheckConsistency(ctx context.Context, in *CheckConsistencyRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/config.ConfigService/CheckConsistency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) UpdateReported(ctx context.Context, in *UpdateReportedRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/config.ConfigService/UpdateReported", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) GetConfigByName(ctx context.Context, in *GetConfigByNameRequest, opts ...grpc.CallOption) (*ConfigField, error) {
	out := new(ConfigField)
	err := c.cc.Invoke(ctx, "/config.ConfigService/GetConfigByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) GetAllConfig(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*ConfigFields, error) {
	out := new(ConfigFields)
	err := c.cc.Invoke(ctx, "/config.ConfigService/GetAllConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) CreatePendingConfig(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/config.ConfigService/CreatePendingConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) CreateNewConfig(ctx context.Context, in *NewConfigRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/config.ConfigService/CreateNewConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigServiceServer is the server API for ConfigService service.
type ConfigServiceServer interface {
	SetDesired(context.Context, *SetDesiredRequest) (*Response, error)
	CheckConsistency(context.Context, *CheckConsistencyRequest) (*Response, error)
	UpdateReported(context.Context, *UpdateReportedRequest) (*Response, error)
	GetConfigByName(context.Context, *GetConfigByNameRequest) (*ConfigField, error)
	GetAllConfig(context.Context, *Identifier) (*ConfigFields, error)
	CreatePendingConfig(context.Context, *Identifier) (*Response, error)
	CreateNewConfig(context.Context, *NewConfigRequest) (*Response, error)
}

// UnimplementedConfigServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConfigServiceServer struct {
}

func (*UnimplementedConfigServiceServer) SetDesired(ctx context.Context, req *SetDesiredRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDesired not implemented")
}
func (*UnimplementedConfigServiceServer) CheckConsistency(ctx context.Context, req *CheckConsistencyRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckConsistency not implemented")
}
func (*UnimplementedConfigServiceServer) UpdateReported(ctx context.Context, req *UpdateReportedRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReported not implemented")
}
func (*UnimplementedConfigServiceServer) GetConfigByName(ctx context.Context, req *GetConfigByNameRequest) (*ConfigField, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfigByName not implemented")
}
func (*UnimplementedConfigServiceServer) GetAllConfig(ctx context.Context, req *Identifier) (*ConfigFields, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllConfig not implemented")
}
func (*UnimplementedConfigServiceServer) CreatePendingConfig(ctx context.Context, req *Identifier) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePendingConfig not implemented")
}
func (*UnimplementedConfigServiceServer) CreateNewConfig(ctx context.Context, req *NewConfigRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewConfig not implemented")
}

func RegisterConfigServiceServer(s *grpc.Server, srv ConfigServiceServer) {
	s.RegisterService(&_ConfigService_serviceDesc, srv)
}

func _ConfigService_SetDesired_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDesiredRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).SetDesired(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/SetDesired",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).SetDesired(ctx, req.(*SetDesiredRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_CheckConsistency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckConsistencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).CheckConsistency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/CheckConsistency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).CheckConsistency(ctx, req.(*CheckConsistencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_UpdateReported_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReportedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).UpdateReported(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/UpdateReported",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).UpdateReported(ctx, req.(*UpdateReportedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_GetConfigByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).GetConfigByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/GetConfigByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).GetConfigByName(ctx, req.(*GetConfigByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_GetAllConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).GetAllConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/GetAllConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).GetAllConfig(ctx, req.(*Identifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_CreatePendingConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).CreatePendingConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/CreatePendingConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).CreatePendingConfig(ctx, req.(*Identifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_CreateNewConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).CreateNewConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/CreateNewConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).CreateNewConfig(ctx, req.(*NewConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConfigService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "config.ConfigService",
	HandlerType: (*ConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetDesired",
			Handler:    _ConfigService_SetDesired_Handler,
		},
		{
			MethodName: "CheckConsistency",
			Handler:    _ConfigService_CheckConsistency_Handler,
		},
		{
			MethodName: "UpdateReported",
			Handler:    _ConfigService_UpdateReported_Handler,
		},
		{
			MethodName: "GetConfigByName",
			Handler:    _ConfigService_GetConfigByName_Handler,
		},
		{
			MethodName: "GetAllConfig",
			Handler:    _ConfigService_GetAllConfig_Handler,
		},
		{
			MethodName: "CreatePendingConfig",
			Handler:    _ConfigService_CreatePendingConfig_Handler,
		},
		{
			MethodName: "CreateNewConfig",
			Handler:    _ConfigService_CreateNewConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config-service.proto",
}
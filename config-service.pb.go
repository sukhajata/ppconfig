// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config-service.proto

package config

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
	FieldValue           []byte   `protobuf:"bytes,3,opt,name=fieldValue,proto3" json:"fieldValue,omitempty"`
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

func (m *UpdateReportedRequest) GetFieldValue() []byte {
	if m != nil {
		return m.FieldValue
	}
	return nil
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
	Description          string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
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

func (m *ConfigField) GetDescription() string {
	if m != nil {
		return m.Description
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

type UpdateFirmwareRequest struct {
	Firmware             string   `protobuf:"bytes,1,opt,name=firmware,proto3" json:"firmware,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateFirmwareRequest) Reset()         { *m = UpdateFirmwareRequest{} }
func (m *UpdateFirmwareRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateFirmwareRequest) ProtoMessage()    {}
func (*UpdateFirmwareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f48c1c3d9797070f, []int{9}
}

func (m *UpdateFirmwareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateFirmwareRequest.Unmarshal(m, b)
}
func (m *UpdateFirmwareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateFirmwareRequest.Marshal(b, m, deterministic)
}
func (m *UpdateFirmwareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFirmwareRequest.Merge(m, src)
}
func (m *UpdateFirmwareRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateFirmwareRequest.Size(m)
}
func (m *UpdateFirmwareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFirmwareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFirmwareRequest proto.InternalMessageInfo

func (m *UpdateFirmwareRequest) GetFirmware() string {
	if m != nil {
		return m.Firmware
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
	proto.RegisterType((*UpdateFirmwareRequest)(nil), "config.UpdateFirmwareRequest")
}

func init() { proto.RegisterFile("config-service.proto", fileDescriptor_f48c1c3d9797070f) }

var fileDescriptor_f48c1c3d9797070f = []byte{
	// 562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x51, 0x73, 0xd2, 0x40,
	0x10, 0x06, 0x0a, 0x08, 0x0b, 0x5a, 0xdc, 0xa2, 0x46, 0xa6, 0x56, 0xe6, 0x9e, 0x3a, 0xa3, 0xf2,
	0xd0, 0xbe, 0xd9, 0x71, 0x1c, 0x45, 0x8b, 0xbc, 0x74, 0x3a, 0xa9, 0xed, 0x3b, 0x92, 0xa5, 0xde,
	0x18, 0x92, 0xf4, 0xee, 0xb0, 0xf2, 0x1b, 0xfc, 0x1d, 0x3e, 0xfa, 0x1f, 0x9d, 0xdc, 0xe5, 0x92,
	0x10, 0x33, 0xe8, 0x43, 0xdf, 0xb2, 0xdf, 0xed, 0xee, 0xb7, 0x7c, 0xfb, 0x2d, 0xd0, 0x9f, 0x87,
	0xc1, 0x82, 0x5f, 0xbf, 0x92, 0x24, 0xbe, 0xf3, 0x39, 0x8d, 0x22, 0x11, 0xaa, 0x10, 0x9b, 0x06,
	0x65, 0x37, 0xf0, 0xf0, 0x82, 0xd4, 0x07, 0x92, 0x5c, 0x90, 0xe7, 0xd2, 0xcd, 0x8a, 0xa4, 0xc2,
	0x03, 0x00, 0xee, 0x51, 0xa0, 0xf8, 0x82, 0x93, 0x70, 0xaa, 0xc3, 0xea, 0x61, 0xdb, 0xcd, 0x21,
	0xb8, 0x0f, 0xed, 0x05, 0x27, 0xdf, 0x3b, 0x9b, 0x2d, 0xc9, 0xa9, 0xe9, 0xe7, 0x0c, 0x88, 0xab,
	0x75, 0x70, 0x35, 0xf3, 0x57, 0xe4, 0xec, 0x98, 0xea, 0x0c, 0x61, 0x3f, 0xab, 0xf0, 0x64, 0xfc,
	0x95, 0xe6, 0xdf, 0xc6, 0x61, 0x20, 0xb9, 0x54, 0x14, 0xcc, 0xd7, 0x96, 0x79, 0x1f, 0xda, 0x1e,
	0xc5, 0x63, 0x7e, 0xbc, 0x9c, 0x26, 0xc4, 0x19, 0x80, 0x08, 0x75, 0xe9, 0x87, 0x4a, 0x53, 0x36,
	0x5c, 0xfd, 0x9d, 0xb2, 0x4d, 0x03, 0x8f, 0x7e, 0x68, 0xb6, 0x86, 0x9b, 0x43, 0xe2, 0xf7, 0x60,
	0xb5, 0x74, 0x49, 0x09, 0x4e, 0xd2, 0xa9, 0x9b, 0xf7, 0x0c, 0x61, 0x2b, 0x78, 0x74, 0x19, 0x79,
	0x33, 0x45, 0x2e, 0x45, 0xa1, 0x50, 0x99, 0x08, 0xdb, 0x47, 0xd9, 0xa4, 0xad, 0x95, 0xd1, 0x16,
	0x44, 0xe8, 0x6e, 0x88, 0x30, 0x84, 0x96, 0x4b, 0x32, 0x0a, 0x03, 0x49, 0xd8, 0x87, 0x86, 0xa0,
	0xc8, 0x5f, 0x27, 0x2c, 0x26, 0x60, 0x27, 0xd0, 0x1d, 0xeb, 0x1d, 0x9d, 0xc6, 0x55, 0x12, 0x5f,
	0x40, 0x53, 0xd7, 0x4b, 0xa7, 0x3a, 0xdc, 0x39, 0xec, 0x1c, 0xed, 0x8d, 0xcc, 0x0a, 0x47, 0xb9,
	0x2c, 0x37, 0x49, 0x61, 0xbf, 0xab, 0xd0, 0xc9, 0xe1, 0xb1, 0x72, 0x41, 0xbc, 0x2c, 0xc3, 0xa0,
	0xbf, 0x63, 0x5a, 0x9e, 0x9b, 0xde, 0x04, 0xe8, 0xc0, 0x3d, 0xcf, 0xb8, 0x21, 0x59, 0x9d, 0x0d,
	0x71, 0x00, 0x2d, 0x91, 0x68, 0xa4, 0x75, 0x6c, 0xbb, 0x69, 0x9c, 0x3a, 0xe2, 0xf3, 0x3a, 0x22,
	0xa7, 0x91, 0x73, 0x44, 0x0c, 0xe0, 0x10, 0x3a, 0x1e, 0xc9, 0xb9, 0xe0, 0x91, 0xe2, 0x61, 0xe0,
	0x34, 0xf5, 0x7b, 0x1e, 0x62, 0xe7, 0xd0, 0x3b, 0xa3, 0x5b, 0x33, 0xf1, 0x7f, 0x2f, 0x20, 0xe7,
	0xd1, 0x5a, 0xd1, 0xa3, 0xec, 0x25, 0xc0, 0x34, 0x73, 0xec, 0x3f, 0x1c, 0xcd, 0xae, 0xe0, 0xf1,
	0x84, 0x94, 0xe1, 0x7f, 0xbf, 0x8e, 0x6d, 0x7c, 0x27, 0xb7, 0xc0, 0x8e, 0xad, 0xbb, 0x4e, 0xb9,
	0x58, 0xde, 0xce, 0x44, 0xda, 0x76, 0x00, 0xad, 0x45, 0x02, 0x25, 0x4d, 0xd3, 0xf8, 0xe8, 0x57,
	0x1d, 0xee, 0x9b, 0x51, 0x2e, 0xcc, 0xcd, 0xe2, 0x09, 0x40, 0x76, 0xa5, 0xf8, 0xd4, 0x6e, 0xfe,
	0xaf, 0xcb, 0x1d, 0xf4, 0xec, 0x93, 0x35, 0x17, 0xab, 0xe0, 0x04, 0x7a, 0xc5, 0x73, 0xc3, 0xe7,
	0xa9, 0x79, 0xca, 0x0f, 0xb1, 0xb4, 0xd1, 0x18, 0x1e, 0x6c, 0x9e, 0x0a, 0x3e, 0xb3, 0x59, 0xa5,
	0x27, 0x54, 0xda, 0xe4, 0x13, 0xec, 0x16, 0x94, 0xc6, 0x03, 0x9b, 0x56, 0xbe, 0x82, 0x41, 0x99,
	0xd3, 0x59, 0x05, 0x5f, 0x43, 0x77, 0x42, 0xea, 0x9d, 0xef, 0x1b, 0x18, 0xd1, 0xa6, 0x65, 0x7b,
	0x1f, 0xf4, 0x4b, 0x4a, 0x25, 0xab, 0xe0, 0x1b, 0xd8, 0x1b, 0x0b, 0x9a, 0x29, 0x3a, 0xa7, 0xc0,
	0xe3, 0xc1, 0xf5, 0x96, 0x16, 0x65, 0x3f, 0xe2, 0x2d, 0xec, 0x9a, 0xf2, 0xd4, 0xb4, 0xe8, 0xd8,
	0xb4, 0xa2, 0x8f, 0xb7, 0x4b, 0x69, 0x7d, 0x51, 0x94, 0xb2, 0xe0, 0x97, 0xb2, 0x26, 0x5f, 0x9a,
	0xfa, 0xaf, 0xfc, 0xf8, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x32, 0x4b, 0x87, 0xe2, 0x05,
	0x00, 0x00,
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
	UpdateFirmware(ctx context.Context, in *UpdateFirmwareRequest, opts ...grpc.CallOption) (*Response, error)
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

func (c *configServiceClient) UpdateFirmware(ctx context.Context, in *UpdateFirmwareRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/config.ConfigService/UpdateFirmware", in, out, opts...)
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
	UpdateFirmware(context.Context, *UpdateFirmwareRequest) (*Response, error)
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
func (*UnimplementedConfigServiceServer) UpdateFirmware(ctx context.Context, req *UpdateFirmwareRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFirmware not implemented")
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

func _ConfigService_UpdateFirmware_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFirmwareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).UpdateFirmware(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/UpdateFirmware",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).UpdateFirmware(ctx, req.(*UpdateFirmwareRequest))
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
		{
			MethodName: "UpdateFirmware",
			Handler:    _ConfigService_UpdateFirmware_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config-service.proto",
}

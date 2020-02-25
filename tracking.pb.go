// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tracking.proto

package proto

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

type AddCaptureRequest struct {
	// x-position on the captured image
	CaptureX float32 `protobuf:"fixed32,1,opt,name=captureX,proto3" json:"captureX,omitempty"`
	// y-position on the captured image
	CaptureY             float32  `protobuf:"fixed32,2,opt,name=captureY,proto3" json:"captureY,omitempty"`
	Time                 int32    `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	ObjectUuid           string   `protobuf:"bytes,4,opt,name=objectUuid,proto3" json:"objectUuid,omitempty"`
	CameraUuid           string   `protobuf:"bytes,5,opt,name=cameraUuid,proto3" json:"cameraUuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddCaptureRequest) Reset()         { *m = AddCaptureRequest{} }
func (m *AddCaptureRequest) String() string { return proto.CompactTextString(m) }
func (*AddCaptureRequest) ProtoMessage()    {}
func (*AddCaptureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62a1d65b4bf956e8, []int{0}
}

func (m *AddCaptureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddCaptureRequest.Unmarshal(m, b)
}
func (m *AddCaptureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddCaptureRequest.Marshal(b, m, deterministic)
}
func (m *AddCaptureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddCaptureRequest.Merge(m, src)
}
func (m *AddCaptureRequest) XXX_Size() int {
	return xxx_messageInfo_AddCaptureRequest.Size(m)
}
func (m *AddCaptureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddCaptureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddCaptureRequest proto.InternalMessageInfo

func (m *AddCaptureRequest) GetCaptureX() float32 {
	if m != nil {
		return m.CaptureX
	}
	return 0
}

func (m *AddCaptureRequest) GetCaptureY() float32 {
	if m != nil {
		return m.CaptureY
	}
	return 0
}

func (m *AddCaptureRequest) GetTime() int32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *AddCaptureRequest) GetObjectUuid() string {
	if m != nil {
		return m.ObjectUuid
	}
	return ""
}

func (m *AddCaptureRequest) GetCameraUuid() string {
	if m != nil {
		return m.CameraUuid
	}
	return ""
}

type AddCaptureReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddCaptureReply) Reset()         { *m = AddCaptureReply{} }
func (m *AddCaptureReply) String() string { return proto.CompactTextString(m) }
func (*AddCaptureReply) ProtoMessage()    {}
func (*AddCaptureReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_62a1d65b4bf956e8, []int{1}
}

func (m *AddCaptureReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddCaptureReply.Unmarshal(m, b)
}
func (m *AddCaptureReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddCaptureReply.Marshal(b, m, deterministic)
}
func (m *AddCaptureReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddCaptureReply.Merge(m, src)
}
func (m *AddCaptureReply) XXX_Size() int {
	return xxx_messageInfo_AddCaptureReply.Size(m)
}
func (m *AddCaptureReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AddCaptureReply.DiscardUnknown(m)
}

var xxx_messageInfo_AddCaptureReply proto.InternalMessageInfo

type GetAllObjectsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllObjectsRequest) Reset()         { *m = GetAllObjectsRequest{} }
func (m *GetAllObjectsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllObjectsRequest) ProtoMessage()    {}
func (*GetAllObjectsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62a1d65b4bf956e8, []int{2}
}

func (m *GetAllObjectsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllObjectsRequest.Unmarshal(m, b)
}
func (m *GetAllObjectsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllObjectsRequest.Marshal(b, m, deterministic)
}
func (m *GetAllObjectsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllObjectsRequest.Merge(m, src)
}
func (m *GetAllObjectsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllObjectsRequest.Size(m)
}
func (m *GetAllObjectsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllObjectsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllObjectsRequest proto.InternalMessageInfo

type GetAllObjectsReply struct {
	Objects              []*ObjectInfo `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetAllObjectsReply) Reset()         { *m = GetAllObjectsReply{} }
func (m *GetAllObjectsReply) String() string { return proto.CompactTextString(m) }
func (*GetAllObjectsReply) ProtoMessage()    {}
func (*GetAllObjectsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_62a1d65b4bf956e8, []int{3}
}

func (m *GetAllObjectsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllObjectsReply.Unmarshal(m, b)
}
func (m *GetAllObjectsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllObjectsReply.Marshal(b, m, deterministic)
}
func (m *GetAllObjectsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllObjectsReply.Merge(m, src)
}
func (m *GetAllObjectsReply) XXX_Size() int {
	return xxx_messageInfo_GetAllObjectsReply.Size(m)
}
func (m *GetAllObjectsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllObjectsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllObjectsReply proto.InternalMessageInfo

func (m *GetAllObjectsReply) GetObjects() []*ObjectInfo {
	if m != nil {
		return m.Objects
	}
	return nil
}

type GetObjectRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetObjectRequest) Reset()         { *m = GetObjectRequest{} }
func (m *GetObjectRequest) String() string { return proto.CompactTextString(m) }
func (*GetObjectRequest) ProtoMessage()    {}
func (*GetObjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62a1d65b4bf956e8, []int{4}
}

func (m *GetObjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetObjectRequest.Unmarshal(m, b)
}
func (m *GetObjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetObjectRequest.Marshal(b, m, deterministic)
}
func (m *GetObjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetObjectRequest.Merge(m, src)
}
func (m *GetObjectRequest) XXX_Size() int {
	return xxx_messageInfo_GetObjectRequest.Size(m)
}
func (m *GetObjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetObjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetObjectRequest proto.InternalMessageInfo

func (m *GetObjectRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type GetObjectReply struct {
	Locations            []*ObjectLocation `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetObjectReply) Reset()         { *m = GetObjectReply{} }
func (m *GetObjectReply) String() string { return proto.CompactTextString(m) }
func (*GetObjectReply) ProtoMessage()    {}
func (*GetObjectReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_62a1d65b4bf956e8, []int{5}
}

func (m *GetObjectReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetObjectReply.Unmarshal(m, b)
}
func (m *GetObjectReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetObjectReply.Marshal(b, m, deterministic)
}
func (m *GetObjectReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetObjectReply.Merge(m, src)
}
func (m *GetObjectReply) XXX_Size() int {
	return xxx_messageInfo_GetObjectReply.Size(m)
}
func (m *GetObjectReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetObjectReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetObjectReply proto.InternalMessageInfo

func (m *GetObjectReply) GetLocations() []*ObjectLocation {
	if m != nil {
		return m.Locations
	}
	return nil
}

type ObjectInfo struct {
	Uuid                 string          `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name                 string          `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Note                 string          `protobuf:"bytes,2,opt,name=note,proto3" json:"note,omitempty"`
	LastLocation         *ObjectLocation `protobuf:"bytes,3,opt,name=lastLocation,proto3" json:"lastLocation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ObjectInfo) Reset()         { *m = ObjectInfo{} }
func (m *ObjectInfo) String() string { return proto.CompactTextString(m) }
func (*ObjectInfo) ProtoMessage()    {}
func (*ObjectInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_62a1d65b4bf956e8, []int{6}
}

func (m *ObjectInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectInfo.Unmarshal(m, b)
}
func (m *ObjectInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectInfo.Marshal(b, m, deterministic)
}
func (m *ObjectInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectInfo.Merge(m, src)
}
func (m *ObjectInfo) XXX_Size() int {
	return xxx_messageInfo_ObjectInfo.Size(m)
}
func (m *ObjectInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectInfo proto.InternalMessageInfo

func (m *ObjectInfo) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ObjectInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ObjectInfo) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *ObjectInfo) GetLastLocation() *ObjectLocation {
	if m != nil {
		return m.LastLocation
	}
	return nil
}

type ObjectLocation struct {
	X                    float32  `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float32  `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z                    float32  `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
	Time                 int32    `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectLocation) Reset()         { *m = ObjectLocation{} }
func (m *ObjectLocation) String() string { return proto.CompactTextString(m) }
func (*ObjectLocation) ProtoMessage()    {}
func (*ObjectLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_62a1d65b4bf956e8, []int{7}
}

func (m *ObjectLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectLocation.Unmarshal(m, b)
}
func (m *ObjectLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectLocation.Marshal(b, m, deterministic)
}
func (m *ObjectLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectLocation.Merge(m, src)
}
func (m *ObjectLocation) XXX_Size() int {
	return xxx_messageInfo_ObjectLocation.Size(m)
}
func (m *ObjectLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectLocation.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectLocation proto.InternalMessageInfo

func (m *ObjectLocation) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *ObjectLocation) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *ObjectLocation) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

func (m *ObjectLocation) GetTime() int32 {
	if m != nil {
		return m.Time
	}
	return 0
}

type UpdatePositionsRequest struct {
	// optional tracking object uuid. if this is empty, will update position for every object
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePositionsRequest) Reset()         { *m = UpdatePositionsRequest{} }
func (m *UpdatePositionsRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePositionsRequest) ProtoMessage()    {}
func (*UpdatePositionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62a1d65b4bf956e8, []int{8}
}

func (m *UpdatePositionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePositionsRequest.Unmarshal(m, b)
}
func (m *UpdatePositionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePositionsRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePositionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePositionsRequest.Merge(m, src)
}
func (m *UpdatePositionsRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePositionsRequest.Size(m)
}
func (m *UpdatePositionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePositionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePositionsRequest proto.InternalMessageInfo

func (m *UpdatePositionsRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type UpdatePositionsReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePositionsReply) Reset()         { *m = UpdatePositionsReply{} }
func (m *UpdatePositionsReply) String() string { return proto.CompactTextString(m) }
func (*UpdatePositionsReply) ProtoMessage()    {}
func (*UpdatePositionsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_62a1d65b4bf956e8, []int{9}
}

func (m *UpdatePositionsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePositionsReply.Unmarshal(m, b)
}
func (m *UpdatePositionsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePositionsReply.Marshal(b, m, deterministic)
}
func (m *UpdatePositionsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePositionsReply.Merge(m, src)
}
func (m *UpdatePositionsReply) XXX_Size() int {
	return xxx_messageInfo_UpdatePositionsReply.Size(m)
}
func (m *UpdatePositionsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePositionsReply.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePositionsReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AddCaptureRequest)(nil), "proto.AddCaptureRequest")
	proto.RegisterType((*AddCaptureReply)(nil), "proto.AddCaptureReply")
	proto.RegisterType((*GetAllObjectsRequest)(nil), "proto.GetAllObjectsRequest")
	proto.RegisterType((*GetAllObjectsReply)(nil), "proto.GetAllObjectsReply")
	proto.RegisterType((*GetObjectRequest)(nil), "proto.GetObjectRequest")
	proto.RegisterType((*GetObjectReply)(nil), "proto.GetObjectReply")
	proto.RegisterType((*ObjectInfo)(nil), "proto.ObjectInfo")
	proto.RegisterType((*ObjectLocation)(nil), "proto.ObjectLocation")
	proto.RegisterType((*UpdatePositionsRequest)(nil), "proto.UpdatePositionsRequest")
	proto.RegisterType((*UpdatePositionsReply)(nil), "proto.UpdatePositionsReply")
}

func init() { proto.RegisterFile("tracking.proto", fileDescriptor_62a1d65b4bf956e8) }

var fileDescriptor_62a1d65b4bf956e8 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xd9, 0xd4, 0x01, 0x32, 0x94, 0x84, 0x8c, 0x5a, 0x63, 0x5c, 0x81, 0xac, 0x3d, 0x20,
	0x4b, 0xa0, 0x1e, 0xda, 0x13, 0x07, 0x24, 0x22, 0x84, 0xaa, 0x4a, 0x88, 0x22, 0x43, 0x25, 0x7a,
	0xdc, 0xda, 0x0b, 0x32, 0x38, 0x5e, 0x63, 0xaf, 0x51, 0xdd, 0x2b, 0x0f, 0xc2, 0x1b, 0xf0, 0x8c,
	0x68, 0x3f, 0xe2, 0x8f, 0xc6, 0xe5, 0xe4, 0xdd, 0xf9, 0xfd, 0x3d, 0x33, 0x3b, 0xf3, 0x87, 0xb9,
	0x2c, 0x59, 0xfc, 0x23, 0xcd, 0xbf, 0x1d, 0x16, 0xa5, 0x90, 0x02, 0xa7, 0xfa, 0x43, 0xff, 0x10,
	0x58, 0xae, 0x92, 0xe4, 0x2d, 0x2b, 0x64, 0x5d, 0xf2, 0x88, 0xff, 0xac, 0x79, 0x25, 0xd1, 0x87,
	0xfb, 0xb1, 0x89, 0x7c, 0xf1, 0x48, 0x40, 0xc2, 0x49, 0xd4, 0xde, 0x7b, 0xec, 0xc2, 0x9b, 0x0c,
	0xd8, 0x05, 0x22, 0x38, 0x32, 0x5d, 0x73, 0x6f, 0x27, 0x20, 0xe1, 0x34, 0xd2, 0x67, 0x7c, 0x06,
	0x20, 0x2e, 0xbf, 0xf3, 0x58, 0x9e, 0xd7, 0x69, 0xe2, 0x39, 0x01, 0x09, 0x67, 0x51, 0x2f, 0xa2,
	0x78, 0xcc, 0xd6, 0xbc, 0x64, 0x9a, 0x4f, 0x0d, 0xef, 0x22, 0x74, 0x09, 0x8b, 0x7e, 0x83, 0x45,
	0xd6, 0x50, 0x17, 0xf6, 0x4e, 0xb8, 0x5c, 0x65, 0xd9, 0x99, 0x4e, 0x53, 0xd9, 0xb6, 0xe9, 0x0a,
	0xf0, 0x46, 0xbc, 0xc8, 0x1a, 0x7c, 0x01, 0xf7, 0x4c, 0xb9, 0xca, 0x23, 0xc1, 0x4e, 0xf8, 0xe0,
	0x68, 0x69, 0x46, 0x70, 0x68, 0x54, 0xa7, 0xf9, 0x57, 0x11, 0x6d, 0x14, 0xf4, 0x39, 0x3c, 0x3a,
	0xe1, 0xd2, 0x90, 0xcd, 0x34, 0x10, 0x9c, 0x5a, 0xf5, 0x46, 0x74, 0x6f, 0xfa, 0x4c, 0xdf, 0xc1,
	0xbc, 0xa7, 0x53, 0x65, 0x8e, 0x61, 0x96, 0x89, 0x98, 0xc9, 0x54, 0xe4, 0x9b, 0x42, 0xfb, 0x83,
	0x42, 0xef, 0x2d, 0x8d, 0x3a, 0x1d, 0xfd, 0x4d, 0x00, 0xba, 0x36, 0xc6, 0x2a, 0xa9, 0x58, 0xce,
	0xd6, 0xdc, 0x4e, 0x4e, 0x9f, 0x75, 0x4c, 0x48, 0xae, 0xe7, 0xaf, 0x62, 0x42, 0x72, 0x7c, 0x05,
	0xbb, 0x19, 0xab, 0xda, 0x2a, 0x7a, 0x07, 0xb7, 0xb6, 0x30, 0x90, 0xd2, 0x0f, 0x30, 0x1f, 0x72,
	0xdc, 0x05, 0x72, 0x65, 0x37, 0x4f, 0xae, 0xd4, 0xad, 0xb1, 0xbb, 0x26, 0x8d, 0xba, 0x5d, 0xeb,
	0xec, 0x93, 0x88, 0x5c, 0xb7, 0x2b, 0x77, 0xba, 0x95, 0xd3, 0x97, 0xe0, 0x9e, 0x17, 0x09, 0x93,
	0xfc, 0xa3, 0xa8, 0x52, 0xfd, 0xd0, 0xff, 0x8d, 0xd2, 0x85, 0xbd, 0x2d, 0x75, 0x91, 0x35, 0x47,
	0x7f, 0x27, 0xb0, 0xf8, 0x6c, 0x4d, 0xfb, 0x89, 0x97, 0xbf, 0xd2, 0x98, 0xe3, 0x1b, 0x80, 0xce,
	0x0c, 0xe8, 0xd9, 0xc7, 0x6d, 0x19, 0xd8, 0x77, 0x47, 0x88, 0x72, 0xce, 0x1d, 0x3c, 0x85, 0x87,
	0x03, 0x8f, 0xe0, 0x81, 0x95, 0x8e, 0x39, 0xca, 0x7f, 0x32, 0x0e, 0x4d, 0xaa, 0xd7, 0x30, 0x6b,
	0x3d, 0x80, 0x8f, 0x3b, 0xe5, 0xc0, 0x3d, 0xfe, 0xfe, 0x36, 0x30, 0xbf, 0x9f, 0xc1, 0xe2, 0xc6,
	0xbb, 0xf1, 0xa9, 0xd5, 0x8e, 0x4f, 0xcf, 0x3f, 0xb8, 0x0d, 0xeb, 0x84, 0x97, 0x77, 0x35, 0x3d,
	0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x5f, 0x63, 0x03, 0xeb, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TrackingServiceClient is the client API for TrackingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TrackingServiceClient interface {
	// Adds a capture from a camera to the database, without calculating 3d coordinates
	AddCapture(ctx context.Context, in *AddCaptureRequest, opts ...grpc.CallOption) (*AddCaptureReply, error)
	GetAllObjects(ctx context.Context, in *GetAllObjectsRequest, opts ...grpc.CallOption) (*GetAllObjectsReply, error)
	GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (*GetObjectReply, error)
	UpdatePositions(ctx context.Context, in *UpdatePositionsRequest, opts ...grpc.CallOption) (*UpdatePositionsReply, error)
}

type trackingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrackingServiceClient(cc grpc.ClientConnInterface) TrackingServiceClient {
	return &trackingServiceClient{cc}
}

func (c *trackingServiceClient) AddCapture(ctx context.Context, in *AddCaptureRequest, opts ...grpc.CallOption) (*AddCaptureReply, error) {
	out := new(AddCaptureReply)
	err := c.cc.Invoke(ctx, "/proto.TrackingService/AddCapture", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackingServiceClient) GetAllObjects(ctx context.Context, in *GetAllObjectsRequest, opts ...grpc.CallOption) (*GetAllObjectsReply, error) {
	out := new(GetAllObjectsReply)
	err := c.cc.Invoke(ctx, "/proto.TrackingService/GetAllObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackingServiceClient) GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (*GetObjectReply, error) {
	out := new(GetObjectReply)
	err := c.cc.Invoke(ctx, "/proto.TrackingService/GetObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackingServiceClient) UpdatePositions(ctx context.Context, in *UpdatePositionsRequest, opts ...grpc.CallOption) (*UpdatePositionsReply, error) {
	out := new(UpdatePositionsReply)
	err := c.cc.Invoke(ctx, "/proto.TrackingService/UpdatePositions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrackingServiceServer is the server API for TrackingService service.
type TrackingServiceServer interface {
	// Adds a capture from a camera to the database, without calculating 3d coordinates
	AddCapture(context.Context, *AddCaptureRequest) (*AddCaptureReply, error)
	GetAllObjects(context.Context, *GetAllObjectsRequest) (*GetAllObjectsReply, error)
	GetObject(context.Context, *GetObjectRequest) (*GetObjectReply, error)
	UpdatePositions(context.Context, *UpdatePositionsRequest) (*UpdatePositionsReply, error)
}

// UnimplementedTrackingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTrackingServiceServer struct {
}

func (*UnimplementedTrackingServiceServer) AddCapture(ctx context.Context, req *AddCaptureRequest) (*AddCaptureReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCapture not implemented")
}
func (*UnimplementedTrackingServiceServer) GetAllObjects(ctx context.Context, req *GetAllObjectsRequest) (*GetAllObjectsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllObjects not implemented")
}
func (*UnimplementedTrackingServiceServer) GetObject(ctx context.Context, req *GetObjectRequest) (*GetObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObject not implemented")
}
func (*UnimplementedTrackingServiceServer) UpdatePositions(ctx context.Context, req *UpdatePositionsRequest) (*UpdatePositionsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePositions not implemented")
}

func RegisterTrackingServiceServer(s *grpc.Server, srv TrackingServiceServer) {
	s.RegisterService(&_TrackingService_serviceDesc, srv)
}

func _TrackingService_AddCapture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCaptureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackingServiceServer).AddCapture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TrackingService/AddCapture",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackingServiceServer).AddCapture(ctx, req.(*AddCaptureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrackingService_GetAllObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackingServiceServer).GetAllObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TrackingService/GetAllObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackingServiceServer).GetAllObjects(ctx, req.(*GetAllObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrackingService_GetObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackingServiceServer).GetObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TrackingService/GetObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackingServiceServer).GetObject(ctx, req.(*GetObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrackingService_UpdatePositions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackingServiceServer).UpdatePositions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TrackingService/UpdatePositions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackingServiceServer).UpdatePositions(ctx, req.(*UpdatePositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TrackingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TrackingService",
	HandlerType: (*TrackingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCapture",
			Handler:    _TrackingService_AddCapture_Handler,
		},
		{
			MethodName: "GetAllObjects",
			Handler:    _TrackingService_GetAllObjects_Handler,
		},
		{
			MethodName: "GetObject",
			Handler:    _TrackingService_GetObject_Handler,
		},
		{
			MethodName: "UpdatePositions",
			Handler:    _TrackingService_UpdatePositions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tracking.proto",
}

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

func init() {
	proto.RegisterType((*AddCaptureRequest)(nil), "proto.AddCaptureRequest")
	proto.RegisterType((*AddCaptureReply)(nil), "proto.AddCaptureReply")
	proto.RegisterType((*GetAllObjectsRequest)(nil), "proto.GetAllObjectsRequest")
	proto.RegisterType((*GetAllObjectsReply)(nil), "proto.GetAllObjectsReply")
	proto.RegisterType((*GetObjectRequest)(nil), "proto.GetObjectRequest")
	proto.RegisterType((*GetObjectReply)(nil), "proto.GetObjectReply")
	proto.RegisterType((*ObjectInfo)(nil), "proto.ObjectInfo")
	proto.RegisterType((*ObjectLocation)(nil), "proto.ObjectLocation")
}

func init() { proto.RegisterFile("tracking.proto", fileDescriptor_62a1d65b4bf956e8) }

var fileDescriptor_62a1d65b4bf956e8 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4f, 0x4b, 0xe3, 0x40,
	0x1c, 0xdd, 0x69, 0xd3, 0xdd, 0xed, 0x6f, 0xbb, 0xed, 0x76, 0xd8, 0x76, 0xb3, 0x59, 0x58, 0xc2,
	0x1c, 0x24, 0x20, 0xf4, 0xd0, 0x9e, 0x3c, 0x08, 0x16, 0x91, 0x52, 0x10, 0x85, 0xa8, 0x60, 0x8f,
	0x69, 0x3a, 0x4a, 0x34, 0xcd, 0xc4, 0x64, 0x22, 0x4d, 0xaf, 0x7e, 0x10, 0x3f, 0x99, 0xdf, 0x45,
	0x66, 0x32, 0xcd, 0x1f, 0x1b, 0x4f, 0x99, 0x79, 0xef, 0x65, 0xde, 0xe3, 0xf7, 0x7e, 0xd0, 0xe5,
	0x91, 0xe3, 0x3e, 0x7a, 0xc1, 0xfd, 0x28, 0x8c, 0x18, 0x67, 0xb8, 0x25, 0x3f, 0xe4, 0x15, 0x41,
	0x7f, 0xba, 0x5a, 0x9d, 0x3a, 0x21, 0x4f, 0x22, 0x6a, 0xd3, 0xa7, 0x84, 0xc6, 0x1c, 0x1b, 0xf0,
	0xdd, 0xcd, 0x90, 0x5b, 0x1d, 0x99, 0xc8, 0x6a, 0xd8, 0xf9, 0xbd, 0xc4, 0x2d, 0xf4, 0x46, 0x85,
	0x5b, 0x60, 0x0c, 0x1a, 0xf7, 0xd6, 0x54, 0x6f, 0x9a, 0xc8, 0x6a, 0xd9, 0xf2, 0x8c, 0xff, 0x03,
	0xb0, 0xe5, 0x03, 0x75, 0xf9, 0x4d, 0xe2, 0xad, 0x74, 0xcd, 0x44, 0x56, 0xdb, 0x2e, 0x21, 0x82,
	0x77, 0x9d, 0x35, 0x8d, 0x1c, 0xc9, 0xb7, 0x32, 0xbe, 0x40, 0x48, 0x1f, 0x7a, 0xe5, 0x80, 0xa1,
	0x9f, 0x92, 0x21, 0xfc, 0x9e, 0x51, 0x3e, 0xf5, 0xfd, 0x4b, 0xf9, 0x4c, 0xac, 0x62, 0x93, 0x29,
	0xe0, 0x0f, 0x78, 0xe8, 0xa7, 0xf8, 0x10, 0xbe, 0x65, 0x76, 0xb1, 0x8e, 0xcc, 0xa6, 0xf5, 0x63,
	0xdc, 0xcf, 0x46, 0x30, 0xca, 0x54, 0xf3, 0xe0, 0x8e, 0xd9, 0x3b, 0x05, 0x39, 0x80, 0x5f, 0x33,
	0xca, 0x33, 0x66, 0x37, 0x0d, 0x0c, 0x5a, 0x22, 0xb2, 0x21, 0x99, 0x4d, 0x9e, 0xc9, 0x19, 0x74,
	0x4b, 0x3a, 0x61, 0x33, 0x81, 0xb6, 0xcf, 0x5c, 0x87, 0x7b, 0x2c, 0xd8, 0x19, 0x0d, 0x2a, 0x46,
	0xe7, 0x8a, 0xb5, 0x0b, 0x1d, 0x79, 0x41, 0x00, 0x45, 0x8c, 0x3a, 0x27, 0x81, 0x05, 0xce, 0x9a,
	0xaa, 0xc9, 0xc9, 0xb3, 0xc4, 0x18, 0xa7, 0x72, 0xfe, 0x02, 0x63, 0x9c, 0xe2, 0x23, 0xe8, 0xf8,
	0x4e, 0x9c, 0xbb, 0xc8, 0x0e, 0x3e, 0x8d, 0x50, 0x91, 0x92, 0x0b, 0xe8, 0x56, 0x79, 0xdc, 0x01,
	0xb4, 0x51, 0xcd, 0xa3, 0x8d, 0xb8, 0xa5, 0xaa, 0x6b, 0x94, 0x8a, 0xdb, 0x56, 0xbe, 0xde, 0xb0,
	0xd1, 0x36, 0xaf, 0x5c, 0x2b, 0x2a, 0x1f, 0xbf, 0x21, 0xe8, 0x5d, 0xab, 0x75, 0xbb, 0xa2, 0xd1,
	0xb3, 0xe7, 0x52, 0x7c, 0x02, 0x50, 0xd4, 0x88, 0x75, 0x15, 0x6b, 0x6f, 0xf5, 0x8c, 0x61, 0x0d,
	0x23, 0x3a, 0xff, 0x82, 0xe7, 0xf0, 0xb3, 0xd2, 0x2e, 0xfe, 0xa7, 0xa4, 0x75, 0xbb, 0x60, 0xfc,
	0xad, 0x27, 0xb3, 0xa7, 0x8e, 0xa1, 0x9d, 0xb7, 0x87, 0xff, 0x14, 0xca, 0x4a, 0xef, 0xc6, 0x60,
	0x9f, 0x90, 0xbf, 0x2f, 0xbf, 0x4a, 0x7c, 0xf2, 0x1e, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xe9, 0x85,
	0xec, 0x54, 0x03, 0x00, 0x00,
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

// TrackingServiceServer is the server API for TrackingService service.
type TrackingServiceServer interface {
	// Adds a capture from a camera to the database, without calculating 3d coordinates
	AddCapture(context.Context, *AddCaptureRequest) (*AddCaptureReply, error)
	GetAllObjects(context.Context, *GetAllObjectsRequest) (*GetAllObjectsReply, error)
	GetObject(context.Context, *GetObjectRequest) (*GetObjectReply, error)
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tracking.proto",
}

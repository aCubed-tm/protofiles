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
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xcd, 0x4a, 0xeb, 0x40,
	0x18, 0xbd, 0xd3, 0xa6, 0xf7, 0xde, 0x7c, 0xb7, 0xb7, 0xb5, 0x83, 0xad, 0x31, 0x82, 0x84, 0x59,
	0x48, 0x40, 0xe8, 0xa2, 0x5d, 0xb9, 0x10, 0x2c, 0x22, 0xa5, 0x20, 0x0a, 0xa3, 0x82, 0x5d, 0xa6,
	0xe9, 0x28, 0xd1, 0x34, 0x13, 0x93, 0x89, 0x34, 0x7d, 0x19, 0x9f, 0xcc, 0x77, 0x91, 0x4c, 0xd2,
	0xfc, 0xd8, 0xb8, 0xca, 0xcc, 0x39, 0x27, 0xdf, 0x39, 0x7c, 0x67, 0xa0, 0x23, 0x02, 0xcb, 0x7e,
	0x75, 0xbc, 0xe7, 0xa1, 0x1f, 0x70, 0xc1, 0x71, 0x4b, 0x7e, 0xc8, 0x07, 0x82, 0xde, 0x64, 0xb9,
	0xbc, 0xb4, 0x7c, 0x11, 0x05, 0x8c, 0xb2, 0xb7, 0x88, 0x85, 0x02, 0xeb, 0xf0, 0xd7, 0x4e, 0x91,
	0x47, 0x0d, 0x19, 0xc8, 0x6c, 0xd0, 0xfc, 0x5e, 0xe2, 0xe6, 0x5a, 0xa3, 0xc2, 0xcd, 0x31, 0x06,
	0x45, 0x38, 0x2b, 0xa6, 0x35, 0x0d, 0x64, 0xb6, 0xa8, 0x3c, 0xe3, 0x63, 0x00, 0xbe, 0x78, 0x61,
	0xb6, 0x78, 0x88, 0x9c, 0xa5, 0xa6, 0x18, 0xc8, 0x54, 0x69, 0x09, 0x49, 0x78, 0xdb, 0x5a, 0xb1,
	0xc0, 0x92, 0x7c, 0x2b, 0xe5, 0x0b, 0x84, 0xf4, 0xa0, 0x5b, 0x0e, 0xe8, 0xbb, 0x31, 0x19, 0xc0,
	0xfe, 0x94, 0x89, 0x89, 0xeb, 0xde, 0xca, 0x31, 0x61, 0x16, 0x9b, 0x4c, 0x00, 0x7f, 0xc3, 0x7d,
	0x37, 0xc6, 0xa7, 0xf0, 0x27, 0xb5, 0x0b, 0x35, 0x64, 0x34, 0xcd, 0x7f, 0xa3, 0x5e, 0xba, 0x82,
	0x61, 0xaa, 0x9a, 0x79, 0x4f, 0x9c, 0x6e, 0x15, 0xe4, 0x04, 0xf6, 0xa6, 0x4c, 0xa4, 0xcc, 0x76,
	0x1b, 0x18, 0x94, 0x28, 0xc9, 0x86, 0x64, 0x36, 0x79, 0x26, 0x57, 0xd0, 0x29, 0xe9, 0x12, 0x9b,
	0x31, 0xa8, 0x2e, 0xb7, 0x2d, 0xe1, 0x70, 0x6f, 0x6b, 0xd4, 0xaf, 0x18, 0x5d, 0x67, 0x2c, 0x2d,
	0x74, 0x84, 0x03, 0x14, 0x29, 0xea, 0x8c, 0x12, 0xcc, 0xe3, 0x82, 0xc9, 0x55, 0xab, 0x54, 0x9e,
	0xf1, 0x19, 0xb4, 0x5d, 0x2b, 0xcc, 0x07, 0xca, 0x75, 0xff, 0xe8, 0x56, 0x91, 0x92, 0x1b, 0xe8,
	0x54, 0x79, 0xdc, 0x06, 0xb4, 0xce, 0x4a, 0x46, 0xeb, 0xe4, 0x16, 0x67, 0xb5, 0xa2, 0x38, 0xb9,
	0x6d, 0xe4, 0xf4, 0x06, 0x45, 0x9b, 0xbc, 0x5d, 0xa5, 0x68, 0x77, 0xf4, 0x89, 0xa0, 0x7b, 0x9f,
	0xbd, 0xac, 0x3b, 0x16, 0xbc, 0x3b, 0x36, 0xc3, 0x17, 0x00, 0x45, 0x63, 0x58, 0xcb, 0x62, 0xed,
	0xbc, 0x32, 0x7d, 0x50, 0xc3, 0x24, 0xf5, 0xfe, 0xc2, 0x33, 0xf8, 0x5f, 0x29, 0x12, 0x1f, 0x65,
	0xd2, 0xba, 0xda, 0xf5, 0xc3, 0x7a, 0x32, 0x1d, 0x75, 0x0e, 0x6a, 0x5e, 0x14, 0x3e, 0x28, 0x94,
	0x95, 0x8a, 0xf5, 0xfe, 0x2e, 0x21, 0x7f, 0x5f, 0xfc, 0x96, 0xf8, 0xf8, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0x0a, 0x75, 0x50, 0xfa, 0x3f, 0x03, 0x00, 0x00,
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

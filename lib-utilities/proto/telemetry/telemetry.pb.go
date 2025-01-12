// Code generated by protoc-gen-go. DO NOT EDIT.
// source: telemetry.proto

package telemetry

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

type TelemetryRequest struct {
	SessionToken         string   `protobuf:"bytes,1,opt,name=SessionToken,proto3" json:"SessionToken,omitempty"`
	URL                  string   `protobuf:"bytes,2,opt,name=URL,proto3" json:"URL,omitempty"`
	RequestBody          []byte   `protobuf:"bytes,3,opt,name=RequestBody,proto3" json:"RequestBody,omitempty"`
	ResourceID           string   `protobuf:"bytes,4,opt,name=resourceID,proto3" json:"resourceID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TelemetryRequest) Reset()         { *m = TelemetryRequest{} }
func (m *TelemetryRequest) String() string { return proto.CompactTextString(m) }
func (*TelemetryRequest) ProtoMessage()    {}
func (*TelemetryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_edbfcf76559f568d, []int{0}
}

func (m *TelemetryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetryRequest.Unmarshal(m, b)
}
func (m *TelemetryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetryRequest.Marshal(b, m, deterministic)
}
func (m *TelemetryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetryRequest.Merge(m, src)
}
func (m *TelemetryRequest) XXX_Size() int {
	return xxx_messageInfo_TelemetryRequest.Size(m)
}
func (m *TelemetryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetryRequest proto.InternalMessageInfo

func (m *TelemetryRequest) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

func (m *TelemetryRequest) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

func (m *TelemetryRequest) GetRequestBody() []byte {
	if m != nil {
		return m.RequestBody
	}
	return nil
}

func (m *TelemetryRequest) GetResourceID() string {
	if m != nil {
		return m.ResourceID
	}
	return ""
}

type TelemetryResponse struct {
	StatusCode           int32             `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	StatusMessage        string            `protobuf:"bytes,2,opt,name=statusMessage,proto3" json:"statusMessage,omitempty"`
	Header               map[string]string `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 []byte            `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TelemetryResponse) Reset()         { *m = TelemetryResponse{} }
func (m *TelemetryResponse) String() string { return proto.CompactTextString(m) }
func (*TelemetryResponse) ProtoMessage()    {}
func (*TelemetryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_edbfcf76559f568d, []int{1}
}

func (m *TelemetryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetryResponse.Unmarshal(m, b)
}
func (m *TelemetryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetryResponse.Marshal(b, m, deterministic)
}
func (m *TelemetryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetryResponse.Merge(m, src)
}
func (m *TelemetryResponse) XXX_Size() int {
	return xxx_messageInfo_TelemetryResponse.Size(m)
}
func (m *TelemetryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetryResponse proto.InternalMessageInfo

func (m *TelemetryResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *TelemetryResponse) GetStatusMessage() string {
	if m != nil {
		return m.StatusMessage
	}
	return ""
}

func (m *TelemetryResponse) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *TelemetryResponse) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*TelemetryRequest)(nil), "TelemetryRequest")
	proto.RegisterType((*TelemetryResponse)(nil), "TelemetryResponse")
	proto.RegisterMapType((map[string]string)(nil), "TelemetryResponse.HeaderEntry")
}

func init() { proto.RegisterFile("telemetry.proto", fileDescriptor_edbfcf76559f568d) }

var fileDescriptor_edbfcf76559f568d = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xc1, 0x6e, 0xda, 0x40,
	0x10, 0xad, 0x31, 0x20, 0x31, 0x80, 0x80, 0x2d, 0x07, 0x17, 0xa9, 0xc8, 0x72, 0x7b, 0xe0, 0xe4,
	0x03, 0x55, 0x2b, 0xca, 0x21, 0x91, 0x02, 0x11, 0x44, 0x0a, 0x39, 0x18, 0xf8, 0x00, 0x63, 0x4f,
	0x88, 0x85, 0xe3, 0x75, 0x76, 0xd7, 0x48, 0xfc, 0x42, 0x7e, 0x22, 0xbf, 0x96, 0x4f, 0x89, 0xd6,
	0x18, 0x62, 0xcc, 0x85, 0xc0, 0x6d, 0xe6, 0x79, 0xde, 0x9b, 0x37, 0xcf, 0xd2, 0x42, 0x4d, 0xa0,
	0x8f, 0xcf, 0x28, 0xd8, 0xc6, 0x0c, 0x19, 0x15, 0xd4, 0x78, 0x55, 0xa0, 0x3e, 0xdb, 0x61, 0x16,
	0xbe, 0x44, 0xc8, 0x05, 0x31, 0xa0, 0x32, 0x45, 0xce, 0x3d, 0x1a, 0xcc, 0xe8, 0x0a, 0x03, 0x4d,
	0xd1, 0x95, 0x4e, 0xc9, 0x3a, 0xc0, 0x48, 0x1d, 0xd4, 0xb9, 0x75, 0xaf, 0xe5, 0xe2, 0x4f, 0xb2,
	0x24, 0x3a, 0x94, 0x13, 0x81, 0x1b, 0xea, 0x6e, 0x34, 0x55, 0x57, 0x3a, 0x15, 0x2b, 0x0d, 0x91,
	0x36, 0x00, 0x43, 0x4e, 0x23, 0xe6, 0xe0, 0xdd, 0x50, 0xcb, 0xc7, 0xd4, 0x14, 0x62, 0xbc, 0x2b,
	0xd0, 0x48, 0x99, 0xe1, 0x21, 0x0d, 0x38, 0x4a, 0x16, 0x17, 0xb6, 0x88, 0xf8, 0x80, 0xba, 0x18,
	0x7b, 0x29, 0x58, 0x29, 0x84, 0xfc, 0x86, 0xea, 0xb6, 0x9b, 0x20, 0xe7, 0xf6, 0x12, 0x13, 0x4f,
	0x87, 0x20, 0xf9, 0x07, 0xc5, 0x27, 0xb4, 0x5d, 0x64, 0x9a, 0xaa, 0xab, 0x9d, 0x72, 0xb7, 0x6d,
	0x1e, 0x6d, 0x32, 0xc7, 0xf1, 0xc0, 0x6d, 0x20, 0xb1, 0x64, 0x9a, 0x10, 0xc8, 0x2f, 0xe4, 0x39,
	0xf9, 0xf8, 0x9c, 0xb8, 0x6e, 0xfd, 0x87, 0x72, 0x6a, 0x54, 0x46, 0xb1, 0xc2, 0x4d, 0x92, 0x92,
	0x2c, 0x49, 0x13, 0x0a, 0x6b, 0xdb, 0x8f, 0x76, 0x56, 0xb6, 0x4d, 0x3f, 0xd7, 0x53, 0xba, 0x6f,
	0x05, 0x28, 0xed, 0x17, 0x93, 0x2b, 0xf8, 0x3e, 0x42, 0xb1, 0xef, 0xa7, 0xc8, 0xd6, 0x9e, 0x83,
	0xa4, 0x61, 0x66, 0x7f, 0x49, 0x8b, 0x1c, 0xdb, 0x35, 0xbe, 0x91, 0x31, 0xfc, 0x1c, 0xa1, 0x98,
	0xa0, 0x60, 0x9e, 0x33, 0xc4, 0x47, 0x2f, 0xf0, 0x84, 0x47, 0x83, 0x01, 0xf5, 0x7d, 0x74, 0x64,
	0x75, 0xba, 0xd2, 0x03, 0xfc, 0xda, 0x2b, 0x59, 0x18, 0x52, 0x26, 0x2e, 0xd3, 0x1b, 0xc2, 0x8f,
	0x8c, 0xde, 0x39, 0x2a, 0xd7, 0xd0, 0x94, 0xf9, 0x30, 0x6f, 0xb9, 0x44, 0x76, 0x8e, 0xc0, 0x36,
	0xe0, 0x6c, 0x40, 0x97, 0x9c, 0x71, 0x8e, 0x4a, 0x1f, 0x6a, 0x19, 0x95, 0xd3, 0xb9, 0x7f, 0x01,
	0x3e, 0x23, 0x38, 0x9d, 0xd6, 0x83, 0xea, 0x3c, 0x74, 0x6d, 0x81, 0x5f, 0x65, 0x2e, 0x8a, 0xf1,
	0xc3, 0xf0, 0xe7, 0x23, 0x00, 0x00, 0xff, 0xff, 0x24, 0x51, 0xf9, 0x42, 0x2b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TelemetryClient is the client API for Telemetry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TelemetryClient interface {
	GetTelemetryService(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error)
	GetMetricDefinitionCollection(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error)
	GetMetricReportDefinitionCollection(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error)
	GetMetricReportCollection(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error)
	GetTriggerCollection(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error)
	GetMetricDefinition(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error)
	GetMetricReportDefinition(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error)
	GetMetricReport(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error)
	GetTrigger(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error)
	UpdateTrigger(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error)
}

type telemetryClient struct {
	cc *grpc.ClientConn
}

func NewTelemetryClient(cc *grpc.ClientConn) TelemetryClient {
	return &telemetryClient{cc}
}

func (c *telemetryClient) GetTelemetryService(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error) {
	out := new(TelemetryResponse)
	err := c.cc.Invoke(ctx, "/Telemetry/GetTelemetryService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryClient) GetMetricDefinitionCollection(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error) {
	out := new(TelemetryResponse)
	err := c.cc.Invoke(ctx, "/Telemetry/GetMetricDefinitionCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryClient) GetMetricReportDefinitionCollection(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error) {
	out := new(TelemetryResponse)
	err := c.cc.Invoke(ctx, "/Telemetry/GetMetricReportDefinitionCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryClient) GetMetricReportCollection(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error) {
	out := new(TelemetryResponse)
	err := c.cc.Invoke(ctx, "/Telemetry/GetMetricReportCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryClient) GetTriggerCollection(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error) {
	out := new(TelemetryResponse)
	err := c.cc.Invoke(ctx, "/Telemetry/GetTriggerCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryClient) GetMetricDefinition(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error) {
	out := new(TelemetryResponse)
	err := c.cc.Invoke(ctx, "/Telemetry/GetMetricDefinition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryClient) GetMetricReportDefinition(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error) {
	out := new(TelemetryResponse)
	err := c.cc.Invoke(ctx, "/Telemetry/GetMetricReportDefinition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryClient) GetMetricReport(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error) {
	out := new(TelemetryResponse)
	err := c.cc.Invoke(ctx, "/Telemetry/GetMetricReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryClient) GetTrigger(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error) {
	out := new(TelemetryResponse)
	err := c.cc.Invoke(ctx, "/Telemetry/GetTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryClient) UpdateTrigger(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error) {
	out := new(TelemetryResponse)
	err := c.cc.Invoke(ctx, "/Telemetry/UpdateTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TelemetryServer is the server API for Telemetry service.
type TelemetryServer interface {
	GetTelemetryService(context.Context, *TelemetryRequest) (*TelemetryResponse, error)
	GetMetricDefinitionCollection(context.Context, *TelemetryRequest) (*TelemetryResponse, error)
	GetMetricReportDefinitionCollection(context.Context, *TelemetryRequest) (*TelemetryResponse, error)
	GetMetricReportCollection(context.Context, *TelemetryRequest) (*TelemetryResponse, error)
	GetTriggerCollection(context.Context, *TelemetryRequest) (*TelemetryResponse, error)
	GetMetricDefinition(context.Context, *TelemetryRequest) (*TelemetryResponse, error)
	GetMetricReportDefinition(context.Context, *TelemetryRequest) (*TelemetryResponse, error)
	GetMetricReport(context.Context, *TelemetryRequest) (*TelemetryResponse, error)
	GetTrigger(context.Context, *TelemetryRequest) (*TelemetryResponse, error)
	UpdateTrigger(context.Context, *TelemetryRequest) (*TelemetryResponse, error)
}

// UnimplementedTelemetryServer can be embedded to have forward compatible implementations.
type UnimplementedTelemetryServer struct {
}

func (*UnimplementedTelemetryServer) GetTelemetryService(ctx context.Context, req *TelemetryRequest) (*TelemetryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTelemetryService not implemented")
}
func (*UnimplementedTelemetryServer) GetMetricDefinitionCollection(ctx context.Context, req *TelemetryRequest) (*TelemetryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricDefinitionCollection not implemented")
}
func (*UnimplementedTelemetryServer) GetMetricReportDefinitionCollection(ctx context.Context, req *TelemetryRequest) (*TelemetryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricReportDefinitionCollection not implemented")
}
func (*UnimplementedTelemetryServer) GetMetricReportCollection(ctx context.Context, req *TelemetryRequest) (*TelemetryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricReportCollection not implemented")
}
func (*UnimplementedTelemetryServer) GetTriggerCollection(ctx context.Context, req *TelemetryRequest) (*TelemetryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTriggerCollection not implemented")
}
func (*UnimplementedTelemetryServer) GetMetricDefinition(ctx context.Context, req *TelemetryRequest) (*TelemetryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricDefinition not implemented")
}
func (*UnimplementedTelemetryServer) GetMetricReportDefinition(ctx context.Context, req *TelemetryRequest) (*TelemetryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricReportDefinition not implemented")
}
func (*UnimplementedTelemetryServer) GetMetricReport(ctx context.Context, req *TelemetryRequest) (*TelemetryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricReport not implemented")
}
func (*UnimplementedTelemetryServer) GetTrigger(ctx context.Context, req *TelemetryRequest) (*TelemetryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrigger not implemented")
}
func (*UnimplementedTelemetryServer) UpdateTrigger(ctx context.Context, req *TelemetryRequest) (*TelemetryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTrigger not implemented")
}

func RegisterTelemetryServer(s *grpc.Server, srv TelemetryServer) {
	s.RegisterService(&_Telemetry_serviceDesc, srv)
}

func _Telemetry_GetTelemetryService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TelemetryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServer).GetTelemetryService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Telemetry/GetTelemetryService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServer).GetTelemetryService(ctx, req.(*TelemetryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Telemetry_GetMetricDefinitionCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TelemetryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServer).GetMetricDefinitionCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Telemetry/GetMetricDefinitionCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServer).GetMetricDefinitionCollection(ctx, req.(*TelemetryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Telemetry_GetMetricReportDefinitionCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TelemetryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServer).GetMetricReportDefinitionCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Telemetry/GetMetricReportDefinitionCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServer).GetMetricReportDefinitionCollection(ctx, req.(*TelemetryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Telemetry_GetMetricReportCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TelemetryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServer).GetMetricReportCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Telemetry/GetMetricReportCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServer).GetMetricReportCollection(ctx, req.(*TelemetryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Telemetry_GetTriggerCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TelemetryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServer).GetTriggerCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Telemetry/GetTriggerCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServer).GetTriggerCollection(ctx, req.(*TelemetryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Telemetry_GetMetricDefinition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TelemetryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServer).GetMetricDefinition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Telemetry/GetMetricDefinition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServer).GetMetricDefinition(ctx, req.(*TelemetryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Telemetry_GetMetricReportDefinition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TelemetryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServer).GetMetricReportDefinition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Telemetry/GetMetricReportDefinition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServer).GetMetricReportDefinition(ctx, req.(*TelemetryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Telemetry_GetMetricReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TelemetryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServer).GetMetricReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Telemetry/GetMetricReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServer).GetMetricReport(ctx, req.(*TelemetryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Telemetry_GetTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TelemetryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServer).GetTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Telemetry/GetTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServer).GetTrigger(ctx, req.(*TelemetryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Telemetry_UpdateTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TelemetryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServer).UpdateTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Telemetry/UpdateTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServer).UpdateTrigger(ctx, req.(*TelemetryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Telemetry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Telemetry",
	HandlerType: (*TelemetryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTelemetryService",
			Handler:    _Telemetry_GetTelemetryService_Handler,
		},
		{
			MethodName: "GetMetricDefinitionCollection",
			Handler:    _Telemetry_GetMetricDefinitionCollection_Handler,
		},
		{
			MethodName: "GetMetricReportDefinitionCollection",
			Handler:    _Telemetry_GetMetricReportDefinitionCollection_Handler,
		},
		{
			MethodName: "GetMetricReportCollection",
			Handler:    _Telemetry_GetMetricReportCollection_Handler,
		},
		{
			MethodName: "GetTriggerCollection",
			Handler:    _Telemetry_GetTriggerCollection_Handler,
		},
		{
			MethodName: "GetMetricDefinition",
			Handler:    _Telemetry_GetMetricDefinition_Handler,
		},
		{
			MethodName: "GetMetricReportDefinition",
			Handler:    _Telemetry_GetMetricReportDefinition_Handler,
		},
		{
			MethodName: "GetMetricReport",
			Handler:    _Telemetry_GetMetricReport_Handler,
		},
		{
			MethodName: "GetTrigger",
			Handler:    _Telemetry_GetTrigger_Handler,
		},
		{
			MethodName: "UpdateTrigger",
			Handler:    _Telemetry_UpdateTrigger_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "telemetry.proto",
}

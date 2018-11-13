// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fim.proto

package fim

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FimdConfig struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NodeName             string               `protobuf:"bytes,2,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	PodName              string               `protobuf:"bytes,3,opt,name=podName,proto3" json:"podName,omitempty"`
	Pid                  []int32              `protobuf:"varint,4,rep,packed,name=pid,proto3" json:"pid,omitempty"`
	Cid                  []string             `protobuf:"bytes,5,rep,name=cid,proto3" json:"cid,omitempty"`
	Subject              []*FimWatcherSubject `protobuf:"bytes,6,rep,name=subject,proto3" json:"subject,omitempty"`
	LogFormat            string               `protobuf:"bytes,7,opt,name=logFormat,proto3" json:"logFormat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FimdConfig) Reset()         { *m = FimdConfig{} }
func (m *FimdConfig) String() string { return proto.CompactTextString(m) }
func (*FimdConfig) ProtoMessage()    {}
func (*FimdConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_fim_78ef4afa0f00b758, []int{0}
}
func (m *FimdConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FimdConfig.Unmarshal(m, b)
}
func (m *FimdConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FimdConfig.Marshal(b, m, deterministic)
}
func (dst *FimdConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FimdConfig.Merge(dst, src)
}
func (m *FimdConfig) XXX_Size() int {
	return xxx_messageInfo_FimdConfig.Size(m)
}
func (m *FimdConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FimdConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FimdConfig proto.InternalMessageInfo

func (m *FimdConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FimdConfig) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *FimdConfig) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

func (m *FimdConfig) GetPid() []int32 {
	if m != nil {
		return m.Pid
	}
	return nil
}

func (m *FimdConfig) GetCid() []string {
	if m != nil {
		return m.Cid
	}
	return nil
}

func (m *FimdConfig) GetSubject() []*FimWatcherSubject {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (m *FimdConfig) GetLogFormat() string {
	if m != nil {
		return m.LogFormat
	}
	return ""
}

type FimWatcherSubject struct {
	Path                 []string          `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
	Event                []string          `protobuf:"bytes,2,rep,name=event,proto3" json:"event,omitempty"`
	Ignore               []string          `protobuf:"bytes,3,rep,name=ignore,proto3" json:"ignore,omitempty"`
	OnlyDir              bool              `protobuf:"varint,4,opt,name=onlyDir,proto3" json:"onlyDir,omitempty"`
	Recursive            bool              `protobuf:"varint,5,opt,name=recursive,proto3" json:"recursive,omitempty"`
	MaxDepth             int32             `protobuf:"varint,6,opt,name=maxDepth,proto3" json:"maxDepth,omitempty"`
	Tags                 map[string]string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FimWatcherSubject) Reset()         { *m = FimWatcherSubject{} }
func (m *FimWatcherSubject) String() string { return proto.CompactTextString(m) }
func (*FimWatcherSubject) ProtoMessage()    {}
func (*FimWatcherSubject) Descriptor() ([]byte, []int) {
	return fileDescriptor_fim_78ef4afa0f00b758, []int{1}
}
func (m *FimWatcherSubject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FimWatcherSubject.Unmarshal(m, b)
}
func (m *FimWatcherSubject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FimWatcherSubject.Marshal(b, m, deterministic)
}
func (dst *FimWatcherSubject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FimWatcherSubject.Merge(dst, src)
}
func (m *FimWatcherSubject) XXX_Size() int {
	return xxx_messageInfo_FimWatcherSubject.Size(m)
}
func (m *FimWatcherSubject) XXX_DiscardUnknown() {
	xxx_messageInfo_FimWatcherSubject.DiscardUnknown(m)
}

var xxx_messageInfo_FimWatcherSubject proto.InternalMessageInfo

func (m *FimWatcherSubject) GetPath() []string {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *FimWatcherSubject) GetEvent() []string {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *FimWatcherSubject) GetIgnore() []string {
	if m != nil {
		return m.Ignore
	}
	return nil
}

func (m *FimWatcherSubject) GetOnlyDir() bool {
	if m != nil {
		return m.OnlyDir
	}
	return false
}

func (m *FimWatcherSubject) GetRecursive() bool {
	if m != nil {
		return m.Recursive
	}
	return false
}

func (m *FimWatcherSubject) GetMaxDepth() int32 {
	if m != nil {
		return m.MaxDepth
	}
	return 0
}

func (m *FimWatcherSubject) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type FimdHandle struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	PodName              string   `protobuf:"bytes,2,opt,name=podName,proto3" json:"podName,omitempty"`
	Pid                  []int32  `protobuf:"varint,3,rep,packed,name=pid,proto3" json:"pid,omitempty"`
	ProcessEventfd       []int32  `protobuf:"varint,4,rep,packed,name=processEventfd,proto3" json:"processEventfd,omitempty"`
	MqFd                 int32    `protobuf:"varint,5,opt,name=mqFd,proto3" json:"mqFd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FimdHandle) Reset()         { *m = FimdHandle{} }
func (m *FimdHandle) String() string { return proto.CompactTextString(m) }
func (*FimdHandle) ProtoMessage()    {}
func (*FimdHandle) Descriptor() ([]byte, []int) {
	return fileDescriptor_fim_78ef4afa0f00b758, []int{2}
}
func (m *FimdHandle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FimdHandle.Unmarshal(m, b)
}
func (m *FimdHandle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FimdHandle.Marshal(b, m, deterministic)
}
func (dst *FimdHandle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FimdHandle.Merge(dst, src)
}
func (m *FimdHandle) XXX_Size() int {
	return xxx_messageInfo_FimdHandle.Size(m)
}
func (m *FimdHandle) XXX_DiscardUnknown() {
	xxx_messageInfo_FimdHandle.DiscardUnknown(m)
}

var xxx_messageInfo_FimdHandle proto.InternalMessageInfo

func (m *FimdHandle) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *FimdHandle) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

func (m *FimdHandle) GetPid() []int32 {
	if m != nil {
		return m.Pid
	}
	return nil
}

func (m *FimdHandle) GetProcessEventfd() []int32 {
	if m != nil {
		return m.ProcessEventfd
	}
	return nil
}

func (m *FimdHandle) GetMqFd() int32 {
	if m != nil {
		return m.MqFd
	}
	return 0
}

type FimdMetricsHandle struct {
	FimWatcher           string   `protobuf:"bytes,1,opt,name=fimWatcher,proto3" json:"fimWatcher,omitempty"`
	Event                string   `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	NodeName             string   `protobuf:"bytes,3,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FimdMetricsHandle) Reset()         { *m = FimdMetricsHandle{} }
func (m *FimdMetricsHandle) String() string { return proto.CompactTextString(m) }
func (*FimdMetricsHandle) ProtoMessage()    {}
func (*FimdMetricsHandle) Descriptor() ([]byte, []int) {
	return fileDescriptor_fim_78ef4afa0f00b758, []int{3}
}
func (m *FimdMetricsHandle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FimdMetricsHandle.Unmarshal(m, b)
}
func (m *FimdMetricsHandle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FimdMetricsHandle.Marshal(b, m, deterministic)
}
func (dst *FimdMetricsHandle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FimdMetricsHandle.Merge(dst, src)
}
func (m *FimdMetricsHandle) XXX_Size() int {
	return xxx_messageInfo_FimdMetricsHandle.Size(m)
}
func (m *FimdMetricsHandle) XXX_DiscardUnknown() {
	xxx_messageInfo_FimdMetricsHandle.DiscardUnknown(m)
}

var xxx_messageInfo_FimdMetricsHandle proto.InternalMessageInfo

func (m *FimdMetricsHandle) GetFimWatcher() string {
	if m != nil {
		return m.FimWatcher
	}
	return ""
}

func (m *FimdMetricsHandle) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *FimdMetricsHandle) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_fim_78ef4afa0f00b758, []int{4}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FimdConfig)(nil), "fim.FimdConfig")
	proto.RegisterType((*FimWatcherSubject)(nil), "fim.FimWatcherSubject")
	proto.RegisterMapType((map[string]string)(nil), "fim.FimWatcherSubject.TagsEntry")
	proto.RegisterType((*FimdHandle)(nil), "fim.FimdHandle")
	proto.RegisterType((*FimdMetricsHandle)(nil), "fim.FimdMetricsHandle")
	proto.RegisterType((*Empty)(nil), "fim.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FimdClient is the client API for Fimd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FimdClient interface {
	CreateWatch(ctx context.Context, in *FimdConfig, opts ...grpc.CallOption) (*FimdHandle, error)
	DestroyWatch(ctx context.Context, in *FimdConfig, opts ...grpc.CallOption) (*Empty, error)
	GetWatchState(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Fimd_GetWatchStateClient, error)
	RecordMetrics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Fimd_RecordMetricsClient, error)
}

type fimdClient struct {
	cc *grpc.ClientConn
}

func NewFimdClient(cc *grpc.ClientConn) FimdClient {
	return &fimdClient{cc}
}

func (c *fimdClient) CreateWatch(ctx context.Context, in *FimdConfig, opts ...grpc.CallOption) (*FimdHandle, error) {
	out := new(FimdHandle)
	err := c.cc.Invoke(ctx, "/fim.Fimd/CreateWatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fimdClient) DestroyWatch(ctx context.Context, in *FimdConfig, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/fim.Fimd/DestroyWatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fimdClient) GetWatchState(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Fimd_GetWatchStateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Fimd_serviceDesc.Streams[0], "/fim.Fimd/GetWatchState", opts...)
	if err != nil {
		return nil, err
	}
	x := &fimdGetWatchStateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Fimd_GetWatchStateClient interface {
	Recv() (*FimdHandle, error)
	grpc.ClientStream
}

type fimdGetWatchStateClient struct {
	grpc.ClientStream
}

func (x *fimdGetWatchStateClient) Recv() (*FimdHandle, error) {
	m := new(FimdHandle)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fimdClient) RecordMetrics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Fimd_RecordMetricsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Fimd_serviceDesc.Streams[1], "/fim.Fimd/RecordMetrics", opts...)
	if err != nil {
		return nil, err
	}
	x := &fimdRecordMetricsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Fimd_RecordMetricsClient interface {
	Recv() (*FimdMetricsHandle, error)
	grpc.ClientStream
}

type fimdRecordMetricsClient struct {
	grpc.ClientStream
}

func (x *fimdRecordMetricsClient) Recv() (*FimdMetricsHandle, error) {
	m := new(FimdMetricsHandle)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FimdServer is the server API for Fimd service.
type FimdServer interface {
	CreateWatch(context.Context, *FimdConfig) (*FimdHandle, error)
	DestroyWatch(context.Context, *FimdConfig) (*Empty, error)
	GetWatchState(*Empty, Fimd_GetWatchStateServer) error
	RecordMetrics(*Empty, Fimd_RecordMetricsServer) error
}

func RegisterFimdServer(s *grpc.Server, srv FimdServer) {
	s.RegisterService(&_Fimd_serviceDesc, srv)
}

func _Fimd_CreateWatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FimdConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FimdServer).CreateWatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fim.Fimd/CreateWatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FimdServer).CreateWatch(ctx, req.(*FimdConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fimd_DestroyWatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FimdConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FimdServer).DestroyWatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fim.Fimd/DestroyWatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FimdServer).DestroyWatch(ctx, req.(*FimdConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fimd_GetWatchState_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FimdServer).GetWatchState(m, &fimdGetWatchStateServer{stream})
}

type Fimd_GetWatchStateServer interface {
	Send(*FimdHandle) error
	grpc.ServerStream
}

type fimdGetWatchStateServer struct {
	grpc.ServerStream
}

func (x *fimdGetWatchStateServer) Send(m *FimdHandle) error {
	return x.ServerStream.SendMsg(m)
}

func _Fimd_RecordMetrics_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FimdServer).RecordMetrics(m, &fimdRecordMetricsServer{stream})
}

type Fimd_RecordMetricsServer interface {
	Send(*FimdMetricsHandle) error
	grpc.ServerStream
}

type fimdRecordMetricsServer struct {
	grpc.ServerStream
}

func (x *fimdRecordMetricsServer) Send(m *FimdMetricsHandle) error {
	return x.ServerStream.SendMsg(m)
}

var _Fimd_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fim.Fimd",
	HandlerType: (*FimdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWatch",
			Handler:    _Fimd_CreateWatch_Handler,
		},
		{
			MethodName: "DestroyWatch",
			Handler:    _Fimd_DestroyWatch_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetWatchState",
			Handler:       _Fimd_GetWatchState_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RecordMetrics",
			Handler:       _Fimd_RecordMetrics_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "fim.proto",
}

func init() { proto.RegisterFile("fim.proto", fileDescriptor_fim_78ef4afa0f00b758) }

var fileDescriptor_fim_78ef4afa0f00b758 = []byte{
	// 499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x5d, 0x37, 0x4d, 0xb3, 0x99, 0x65, 0xf9, 0xb0, 0xd0, 0xca, 0x5a, 0x21, 0x14, 0xe5, 0x80,
	0x72, 0xa1, 0x2a, 0x0b, 0xd2, 0x22, 0xae, 0xdb, 0x16, 0x2e, 0x70, 0xc8, 0x22, 0x71, 0xf6, 0x26,
	0x4e, 0x6a, 0x68, 0xe2, 0xe0, 0xb8, 0x15, 0xf9, 0x17, 0x1c, 0xf9, 0x4d, 0x5c, 0xf9, 0x43, 0xc8,
	0xd3, 0x7c, 0xb4, 0xd0, 0xbd, 0xcd, 0x7b, 0x63, 0x7b, 0xe6, 0xbd, 0x27, 0x83, 0x9f, 0xc9, 0x62,
	0x5a, 0x69, 0x65, 0x14, 0x75, 0x32, 0x59, 0x84, 0xbf, 0x09, 0xc0, 0x52, 0x16, 0xe9, 0x8d, 0x2a,
	0x33, 0x99, 0x53, 0x0a, 0xe3, 0x92, 0x17, 0x82, 0x91, 0x80, 0x44, 0x7e, 0x8c, 0x35, 0xbd, 0x84,
	0xd3, 0x52, 0xa5, 0xe2, 0x93, 0xe5, 0x47, 0xc8, 0xf7, 0x98, 0x32, 0xf0, 0x2a, 0x95, 0x62, 0xcb,
	0xc1, 0x56, 0x07, 0xe9, 0x63, 0x70, 0x2a, 0x99, 0xb2, 0x71, 0xe0, 0x44, 0x6e, 0x6c, 0x4b, 0xcb,
	0x24, 0x32, 0x65, 0x6e, 0xe0, 0x44, 0x7e, 0x6c, 0x4b, 0x3a, 0x03, 0xaf, 0xde, 0xdc, 0x7d, 0x15,
	0x89, 0x61, 0x93, 0xc0, 0x89, 0xce, 0xae, 0x2e, 0xa6, 0x76, 0xbd, 0xa5, 0x2c, 0xbe, 0x70, 0x93,
	0xac, 0x84, 0xbe, 0xdd, 0x75, 0xe3, 0xee, 0x18, 0x7d, 0x06, 0xfe, 0x5a, 0xe5, 0x4b, 0xa5, 0x0b,
	0x6e, 0x98, 0x87, 0x13, 0x07, 0x22, 0xfc, 0x35, 0x82, 0x27, 0xff, 0x5d, 0xb6, 0x9a, 0x2a, 0x6e,
	0x56, 0x8c, 0xe0, 0x60, 0xac, 0xe9, 0x53, 0x70, 0xc5, 0x56, 0x94, 0x86, 0x8d, 0x90, 0xdc, 0x01,
	0x7a, 0x01, 0x13, 0x99, 0x97, 0x4a, 0x5b, 0x31, 0x96, 0x6e, 0x91, 0x55, 0xa9, 0xca, 0x75, 0x33,
	0x97, 0x9a, 0x8d, 0x03, 0x12, 0x9d, 0xc6, 0x1d, 0xb4, 0xfb, 0x68, 0x91, 0x6c, 0x74, 0x2d, 0xb7,
	0x82, 0xb9, 0xd8, 0x1b, 0x08, 0xeb, 0x5c, 0xc1, 0x7f, 0xcc, 0x45, 0x65, 0x56, 0x6c, 0x12, 0x90,
	0xc8, 0x8d, 0x7b, 0x4c, 0xdf, 0xc0, 0xd8, 0xf0, 0xbc, 0x66, 0x1e, 0x0a, 0x0f, 0x8e, 0x0b, 0x9f,
	0x7e, 0xe6, 0x79, 0xbd, 0x28, 0x8d, 0x6e, 0x62, 0x3c, 0x7d, 0x79, 0x0d, 0x7e, 0x4f, 0x59, 0x43,
	0xbf, 0x89, 0xa6, 0xcd, 0xca, 0x96, 0x56, 0xd6, 0x96, 0xaf, 0x37, 0x5d, 0x4e, 0x3b, 0xf0, 0x6e,
	0xf4, 0x96, 0x84, 0x3f, 0xdb, 0x9c, 0x3f, 0xf0, 0x32, 0x5d, 0x1f, 0x66, 0x4a, 0xee, 0xcf, 0x74,
	0x74, 0x34, 0x53, 0x67, 0xc8, 0xf4, 0x05, 0x3c, 0xac, 0xb4, 0x4a, 0x44, 0x5d, 0x2f, 0xac, 0x83,
	0x59, 0x17, 0xf8, 0x3f, 0xac, 0xcd, 0xa0, 0xf8, 0xbe, 0x4c, 0xd1, 0x22, 0x37, 0xc6, 0x3a, 0x14,
	0x18, 0x56, 0xfa, 0x51, 0x18, 0x2d, 0x93, 0xba, 0x5d, 0xec, 0x39, 0x40, 0xd6, 0xbb, 0xd0, 0xae,
	0xb6, 0xc7, 0xec, 0x07, 0x47, 0x86, 0xe0, 0xf6, 0xe5, 0x38, 0x87, 0x72, 0x42, 0x0f, 0xdc, 0x45,
	0x51, 0x99, 0xe6, 0xea, 0x0f, 0x81, 0xb1, 0x1d, 0x48, 0x5f, 0xc1, 0xd9, 0x8d, 0x16, 0xdc, 0x08,
	0x7c, 0x94, 0x3e, 0xea, 0xbc, 0x6f, 0x3f, 0xc1, 0xe5, 0x40, 0xec, 0x96, 0x0a, 0x4f, 0xe8, 0x4b,
	0x78, 0x30, 0x17, 0xb5, 0xd1, 0xaa, 0xb9, 0xe7, 0x0e, 0x20, 0x81, 0x83, 0xc2, 0x13, 0x3a, 0x83,
	0xf3, 0xf7, 0xc2, 0xe0, 0xd1, 0x5b, 0xc3, 0x8d, 0xa0, 0x7b, 0xed, 0x23, 0xcf, 0xcf, 0x08, 0xbd,
	0x86, 0xf3, 0x58, 0x24, 0x4a, 0x77, 0x76, 0x1c, 0xdc, 0xe8, 0xbf, 0xc5, 0xa1, 0x59, 0xf6, 0xe2,
	0xdd, 0x04, 0x3f, 0xf3, 0xeb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x76, 0x40, 0x9e, 0xf4, 0xd9,
	0x03, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package loker

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Level int32

const (
	Level_Level_UNKNOWN Level = 0
	Level_ADMIN         Level = 1
	Level_MEMBER        Level = 2
	Level_PERUSAHAAN    Level = 3
)

var Level_name = map[int32]string{
	0: "Level_UNKNOWN",
	1: "ADMIN",
	2: "MEMBER",
	3: "PERUSAHAAN",
}

var Level_value = map[string]int32{
	"Level_UNKNOWN": 0,
	"ADMIN":         1,
	"MEMBER":        2,
	"PERUSAHAAN":    3,
}

func (x Level) String() string {
	return proto.EnumName(Level_name, int32(x))
}

func (Level) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

type User struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string               `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Foto                 string               `protobuf:"bytes,3,opt,name=foto,proto3" json:"foto,omitempty"`
	TglDibuat            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=tgl_dibuat,json=tglDibuat,proto3" json:"tgl_dibuat,omitempty"`
	Level                Level                `protobuf:"varint,5,opt,name=level,proto3,enum=loker.Level" json:"level,omitempty"`
	Verifikasi           bool                 `protobuf:"varint,6,opt,name=verifikasi,proto3" json:"verifikasi,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetFoto() string {
	if m != nil {
		return m.Foto
	}
	return ""
}

func (m *User) GetTglDibuat() *timestamp.Timestamp {
	if m != nil {
		return m.TglDibuat
	}
	return nil
}

func (m *User) GetLevel() Level {
	if m != nil {
		return m.Level
	}
	return Level_Level_UNKNOWN
}

func (m *User) GetVerifikasi() bool {
	if m != nil {
		return m.Verifikasi
	}
	return false
}

type UserRequestId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequestId) Reset()         { *m = UserRequestId{} }
func (m *UserRequestId) String() string { return proto.CompactTextString(m) }
func (*UserRequestId) ProtoMessage()    {}
func (*UserRequestId) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserRequestId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequestId.Unmarshal(m, b)
}
func (m *UserRequestId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequestId.Marshal(b, m, deterministic)
}
func (m *UserRequestId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequestId.Merge(m, src)
}
func (m *UserRequestId) XXX_Size() int {
	return xxx_messageInfo_UserRequestId.Size(m)
}
func (m *UserRequestId) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequestId.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequestId proto.InternalMessageInfo

func (m *UserRequestId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterEnum("loker.Level", Level_name, Level_value)
	proto.RegisterType((*User)(nil), "loker.User")
	proto.RegisterType((*UserRequestId)(nil), "loker.UserRequestId")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0xd9, 0x42, 0x09, 0x1d, 0x84, 0xd4, 0x89, 0x87, 0xa6, 0x07, 0x69, 0x7a, 0x6a, 0x38,
	0x94, 0x04, 0x4f, 0x1c, 0x6b, 0x20, 0x4a, 0x94, 0x6a, 0x16, 0x89, 0x47, 0x52, 0xec, 0xd0, 0x6c,
	0x28, 0x2e, 0xb6, 0x5b, 0xde, 0xd0, 0xf7, 0x32, 0x5d, 0xc4, 0xe0, 0x6d, 0xf6, 0xdb, 0xc9, 0x7c,
	0xf3, 0x0f, 0x40, 0x55, 0x52, 0x11, 0x1e, 0x0a, 0xa9, 0x24, 0x9a, 0xb9, 0xdc, 0x51, 0xe1, 0x0e,
	0x32, 0x29, 0xb3, 0x9c, 0x46, 0x1a, 0x6e, 0xaa, 0xed, 0x48, 0x89, 0x3d, 0x95, 0x2a, 0xd9, 0x1f,
	0x4e, 0x7d, 0xfe, 0x37, 0x83, 0xd6, 0xaa, 0xa4, 0x02, 0xfb, 0x60, 0x88, 0xd4, 0x61, 0x1e, 0x0b,
	0x4c, 0x6e, 0x88, 0x14, 0x5d, 0xe8, 0xd4, 0xe3, 0x3e, 0x93, 0x3d, 0x39, 0x86, 0xc7, 0x02, 0x8b,
	0xff, 0xbd, 0x11, 0xa1, 0xb5, 0x95, 0x4a, 0x3a, 0x4d, 0xcd, 0x75, 0x8d, 0x13, 0x00, 0x95, 0xe5,
	0xeb, 0x54, 0x6c, 0xaa, 0x44, 0x39, 0x2d, 0x8f, 0x05, 0xdd, 0xb1, 0x1b, 0x9e, 0xf4, 0xe1, 0x59,
	0x1f, 0xbe, 0x9d, 0xf5, 0xdc, 0x52, 0x59, 0x3e, 0xd5, 0xcd, 0xe8, 0x83, 0x99, 0xd3, 0x91, 0x72,
	0xc7, 0xf4, 0x58, 0xd0, 0x1f, 0x5f, 0x85, 0x7a, 0xf7, 0xf0, 0xb9, 0x66, 0xfc, 0xf4, 0x85, 0xb7,
	0x00, 0x47, 0x2a, 0xc4, 0x56, 0xec, 0x92, 0x52, 0x38, 0x6d, 0x8f, 0x05, 0x1d, 0x7e, 0x41, 0xfc,
	0x01, 0xf4, 0xea, 0x18, 0x9c, 0xbe, 0x2a, 0x2a, 0xd5, 0x3c, 0xbd, 0xc8, 0x63, 0xd5, 0x79, 0x86,
	0x11, 0x98, 0x7a, 0x20, 0x5e, 0x43, 0x4f, 0x17, 0xeb, 0x55, 0xfc, 0x14, 0xbf, 0xbc, 0xc7, 0x76,
	0x03, 0x2d, 0x30, 0xa3, 0xe9, 0x62, 0x1e, 0xdb, 0x0c, 0x01, 0xda, 0x8b, 0xd9, 0xe2, 0x7e, 0xc6,
	0x6d, 0x03, 0xfb, 0x00, 0xaf, 0x33, 0xbe, 0x5a, 0x46, 0x8f, 0x51, 0x14, 0xdb, 0xcd, 0xf1, 0x04,
	0xba, 0xb5, 0x63, 0x49, 0xc5, 0x51, 0x7c, 0x10, 0x0e, 0xa1, 0xf9, 0x40, 0x0a, 0x6f, 0x7e, 0xd7,
	0xfd, 0xa7, 0x77, 0xbb, 0x17, 0xd4, 0x6f, 0x6c, 0xda, 0xfa, 0x02, 0x77, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xdc, 0xce, 0x26, 0xb5, 0xa3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Get(ctx context.Context, in *UserRequestId, opts ...grpc.CallOption) (*User, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Get(ctx context.Context, in *UserRequestId, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/loker.UserService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Get(context.Context, *UserRequestId) (*User, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) Get(ctx context.Context, req *UserRequestId) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loker.UserService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Get(ctx, req.(*UserRequestId))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "loker.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _UserService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

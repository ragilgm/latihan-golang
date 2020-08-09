// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: user.proto

package UserService

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	FIRSTNAME string `protobuf:"bytes,2,opt,name=FIRSTNAME,proto3" json:"FIRSTNAME,omitempty"`
	LASTNAME  string `protobuf:"bytes,3,opt,name=LASTNAME,proto3" json:"LASTNAME,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *User) GetFIRSTNAME() string {
	if x != nil {
		return x.FIRSTNAME
	}
	return ""
}

func (x *User) GetLASTNAME() string {
	if x != nil {
		return x.LASTNAME
	}
	return ""
}

type Users struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User []*User `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
}

func (x *Users) Reset() {
	*x = Users{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Users) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Users) ProtoMessage() {}

func (x *Users) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Users.ProtoReflect.Descriptor instead.
func (*Users) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *Users) GetUser() []*User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x50, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x49,
	0x52, 0x53, 0x54, 0x4e, 0x41, 0x4d, 0x45, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46,
	0x49, 0x52, 0x53, 0x54, 0x4e, 0x41, 0x4d, 0x45, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x41, 0x53, 0x54,
	0x4e, 0x41, 0x4d, 0x45, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x41, 0x53, 0x54,
	0x4e, 0x41, 0x4d, 0x45, 0x22, 0x2e, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x25, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x32, 0xce, 0x01, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x32, 0x0a, 0x0a,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x11, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x11, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x30, 0x0a, 0x08, 0x45, 0x64, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x11, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a,
	0x11, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x2e, 0x0a, 0x06, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x44, 0x12, 0x11, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a,
	0x11, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x31, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x12, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x12, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_user_proto_goTypes = []interface{}{
	(*Empty)(nil), // 0: UserService.Empty
	(*User)(nil),  // 1: UserService.User
	(*Users)(nil), // 2: UserService.Users
}
var file_user_proto_depIdxs = []int32{
	1, // 0: UserService.Users.user:type_name -> UserService.User
	1, // 1: UserService.Add.InsertData:input_type -> UserService.User
	1, // 2: UserService.Add.EditData:input_type -> UserService.User
	1, // 3: UserService.Add.FindID:input_type -> UserService.User
	0, // 4: UserService.Add.FindAll:input_type -> UserService.Empty
	1, // 5: UserService.Add.InsertData:output_type -> UserService.User
	1, // 6: UserService.Add.EditData:output_type -> UserService.User
	1, // 7: UserService.Add.FindID:output_type -> UserService.User
	2, // 8: UserService.Add.FindAll:output_type -> UserService.Users
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Users); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AddClient is the client API for Add service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddClient interface {
	InsertData(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	EditData(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	FindID(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	FindAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Users, error)
}

type addClient struct {
	cc grpc.ClientConnInterface
}

func NewAddClient(cc grpc.ClientConnInterface) AddClient {
	return &addClient{cc}
}

func (c *addClient) InsertData(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/UserService.Add/InsertData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addClient) EditData(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/UserService.Add/EditData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addClient) FindID(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/UserService.Add/FindID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addClient) FindAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/UserService.Add/FindAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddServer is the server API for Add service.
type AddServer interface {
	InsertData(context.Context, *User) (*User, error)
	EditData(context.Context, *User) (*User, error)
	FindID(context.Context, *User) (*User, error)
	FindAll(context.Context, *Empty) (*Users, error)
}

// UnimplementedAddServer can be embedded to have forward compatible implementations.
type UnimplementedAddServer struct {
}

func (*UnimplementedAddServer) InsertData(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertData not implemented")
}
func (*UnimplementedAddServer) EditData(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditData not implemented")
}
func (*UnimplementedAddServer) FindID(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindID not implemented")
}
func (*UnimplementedAddServer) FindAll(context.Context, *Empty) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}

func RegisterAddServer(s *grpc.Server, srv AddServer) {
	s.RegisterService(&_Add_serviceDesc, srv)
}

func _Add_InsertData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServer).InsertData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService.Add/InsertData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServer).InsertData(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Add_EditData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServer).EditData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService.Add/EditData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServer).EditData(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Add_FindID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServer).FindID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService.Add/FindID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServer).FindID(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Add_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService.Add/FindAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServer).FindAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Add_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UserService.Add",
	HandlerType: (*AddServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertData",
			Handler:    _Add_InsertData_Handler,
		},
		{
			MethodName: "EditData",
			Handler:    _Add_EditData_Handler,
		},
		{
			MethodName: "FindID",
			Handler:    _Add_FindID_Handler,
		},
		{
			MethodName: "FindAll",
			Handler:    _Add_FindAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
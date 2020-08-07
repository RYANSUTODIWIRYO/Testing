// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: parkir.proto

package parkir

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
		mi := &file_parkir_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_parkir_proto_msgTypes[0]
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
	return file_parkir_proto_rawDescGZIP(), []int{0}
}

type KarcisParkir struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdKendaraan string `protobuf:"bytes,1,opt,name=idKendaraan,proto3" json:"idKendaraan,omitempty"`
}

func (x *KarcisParkir) Reset() {
	*x = KarcisParkir{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkir_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KarcisParkir) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KarcisParkir) ProtoMessage() {}

func (x *KarcisParkir) ProtoReflect() protoreflect.Message {
	mi := &file_parkir_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KarcisParkir.ProtoReflect.Descriptor instead.
func (*KarcisParkir) Descriptor() ([]byte, []int) {
	return file_parkir_proto_rawDescGZIP(), []int{1}
}

func (x *KarcisParkir) GetIdKendaraan() string {
	if x != nil {
		return x.IdKendaraan
	}
	return ""
}

type Keluar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JenisKendaraan string `protobuf:"bytes,1,opt,name=jenisKendaraan,proto3" json:"jenisKendaraan,omitempty"`
	PlatNomor      string `protobuf:"bytes,2,opt,name=platNomor,proto3" json:"platNomor,omitempty"`
	IdKendaraan    string `protobuf:"bytes,3,opt,name=idKendaraan,proto3" json:"idKendaraan,omitempty"`
}

func (x *Keluar) Reset() {
	*x = Keluar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkir_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Keluar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Keluar) ProtoMessage() {}

func (x *Keluar) ProtoReflect() protoreflect.Message {
	mi := &file_parkir_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Keluar.ProtoReflect.Descriptor instead.
func (*Keluar) Descriptor() ([]byte, []int) {
	return file_parkir_proto_rawDescGZIP(), []int{2}
}

func (x *Keluar) GetJenisKendaraan() string {
	if x != nil {
		return x.JenisKendaraan
	}
	return ""
}

func (x *Keluar) GetPlatNomor() string {
	if x != nil {
		return x.PlatNomor
	}
	return ""
}

func (x *Keluar) GetIdKendaraan() string {
	if x != nil {
		return x.IdKendaraan
	}
	return ""
}

type StrukParkir struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdKendaraan    string `protobuf:"bytes,1,opt,name=idKendaraan,proto3" json:"idKendaraan,omitempty"`
	JenisKendaraan string `protobuf:"bytes,2,opt,name=jenisKendaraan,proto3" json:"jenisKendaraan,omitempty"`
	PlatNomor      string `protobuf:"bytes,3,opt,name=platNomor,proto3" json:"platNomor,omitempty"`
	JamMasuk       string `protobuf:"bytes,4,opt,name=jamMasuk,proto3" json:"jamMasuk,omitempty"`
	JamKeluar      string `protobuf:"bytes,5,opt,name=jamKeluar,proto3" json:"jamKeluar,omitempty"`
	Durasi         int64  `protobuf:"varint,6,opt,name=durasi,proto3" json:"durasi,omitempty"`
	TotalTarif     int64  `protobuf:"varint,7,opt,name=totalTarif,proto3" json:"totalTarif,omitempty"`
}

func (x *StrukParkir) Reset() {
	*x = StrukParkir{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkir_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrukParkir) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrukParkir) ProtoMessage() {}

func (x *StrukParkir) ProtoReflect() protoreflect.Message {
	mi := &file_parkir_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrukParkir.ProtoReflect.Descriptor instead.
func (*StrukParkir) Descriptor() ([]byte, []int) {
	return file_parkir_proto_rawDescGZIP(), []int{3}
}

func (x *StrukParkir) GetIdKendaraan() string {
	if x != nil {
		return x.IdKendaraan
	}
	return ""
}

func (x *StrukParkir) GetJenisKendaraan() string {
	if x != nil {
		return x.JenisKendaraan
	}
	return ""
}

func (x *StrukParkir) GetPlatNomor() string {
	if x != nil {
		return x.PlatNomor
	}
	return ""
}

func (x *StrukParkir) GetJamMasuk() string {
	if x != nil {
		return x.JamMasuk
	}
	return ""
}

func (x *StrukParkir) GetJamKeluar() string {
	if x != nil {
		return x.JamKeluar
	}
	return ""
}

func (x *StrukParkir) GetDurasi() int64 {
	if x != nil {
		return x.Durasi
	}
	return 0
}

func (x *StrukParkir) GetTotalTarif() int64 {
	if x != nil {
		return x.TotalTarif
	}
	return 0
}

var File_parkir_proto protoreflect.FileDescriptor

var file_parkir_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x61, 0x72, 0x6b, 0x69, 0x72, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x30, 0x0a, 0x0c, 0x4b, 0x61, 0x72, 0x63, 0x69, 0x73, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x72, 0x12,
	0x20, 0x0a, 0x0b, 0x69, 0x64, 0x4b, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x61, 0x61, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x64, 0x4b, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x61, 0x61,
	0x6e, 0x22, 0x70, 0x0a, 0x06, 0x4b, 0x65, 0x6c, 0x75, 0x61, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x6a,
	0x65, 0x6e, 0x69, 0x73, 0x4b, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x61, 0x61, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x6a, 0x65, 0x6e, 0x69, 0x73, 0x4b, 0x65, 0x6e, 0x64, 0x61, 0x72,
	0x61, 0x61, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x74, 0x4e, 0x6f, 0x6d, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x74, 0x4e, 0x6f, 0x6d, 0x6f,
	0x72, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x64, 0x4b, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x61, 0x61, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x64, 0x4b, 0x65, 0x6e, 0x64, 0x61, 0x72,
	0x61, 0x61, 0x6e, 0x22, 0xe7, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x75, 0x6b, 0x50, 0x61, 0x72,
	0x6b, 0x69, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x64, 0x4b, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x61,
	0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x64, 0x4b, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x61, 0x61, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x6a, 0x65, 0x6e, 0x69, 0x73, 0x4b, 0x65,
	0x6e, 0x64, 0x61, 0x72, 0x61, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6a,
	0x65, 0x6e, 0x69, 0x73, 0x4b, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x61, 0x61, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x6c, 0x61, 0x74, 0x4e, 0x6f, 0x6d, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x6c, 0x61, 0x74, 0x4e, 0x6f, 0x6d, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6a,
	0x61, 0x6d, 0x4d, 0x61, 0x73, 0x75, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a,
	0x61, 0x6d, 0x4d, 0x61, 0x73, 0x75, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x6a, 0x61, 0x6d, 0x4b, 0x65,
	0x6c, 0x75, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x61, 0x6d, 0x4b,
	0x65, 0x6c, 0x75, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x75, 0x72, 0x61, 0x73, 0x69, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x64, 0x75, 0x72, 0x61, 0x73, 0x69, 0x12, 0x1e, 0x0a,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x61, 0x72, 0x69, 0x66, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x61, 0x72, 0x69, 0x66, 0x32, 0x7c, 0x0a,
	0x0d, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34,
	0x0a, 0x0b, 0x4d, 0x61, 0x73, 0x75, 0x6b, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x72, 0x12, 0x0d, 0x2e,
	0x70, 0x61, 0x72, 0x6b, 0x69, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x70,
	0x61, 0x72, 0x6b, 0x69, 0x72, 0x2e, 0x4b, 0x61, 0x72, 0x63, 0x69, 0x73, 0x50, 0x61, 0x72, 0x6b,
	0x69, 0x72, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0c, 0x4b, 0x65, 0x6c, 0x75, 0x61, 0x72, 0x50, 0x61,
	0x72, 0x6b, 0x69, 0x72, 0x12, 0x0e, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x72, 0x2e, 0x4b, 0x65,
	0x6c, 0x75, 0x61, 0x72, 0x1a, 0x13, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x72, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x6b, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x72, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_parkir_proto_rawDescOnce sync.Once
	file_parkir_proto_rawDescData = file_parkir_proto_rawDesc
)

func file_parkir_proto_rawDescGZIP() []byte {
	file_parkir_proto_rawDescOnce.Do(func() {
		file_parkir_proto_rawDescData = protoimpl.X.CompressGZIP(file_parkir_proto_rawDescData)
	})
	return file_parkir_proto_rawDescData
}

var file_parkir_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_parkir_proto_goTypes = []interface{}{
	(*Empty)(nil),        // 0: parkir.Empty
	(*KarcisParkir)(nil), // 1: parkir.KarcisParkir
	(*Keluar)(nil),       // 2: parkir.Keluar
	(*StrukParkir)(nil),  // 3: parkir.StrukParkir
}
var file_parkir_proto_depIdxs = []int32{
	0, // 0: parkir.ParkirService.MasukParkir:input_type -> parkir.Empty
	2, // 1: parkir.ParkirService.KeluarParkir:input_type -> parkir.Keluar
	1, // 2: parkir.ParkirService.MasukParkir:output_type -> parkir.KarcisParkir
	3, // 3: parkir.ParkirService.KeluarParkir:output_type -> parkir.StrukParkir
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_parkir_proto_init() }
func file_parkir_proto_init() {
	if File_parkir_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_parkir_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_parkir_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KarcisParkir); i {
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
		file_parkir_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Keluar); i {
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
		file_parkir_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrukParkir); i {
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
			RawDescriptor: file_parkir_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_parkir_proto_goTypes,
		DependencyIndexes: file_parkir_proto_depIdxs,
		MessageInfos:      file_parkir_proto_msgTypes,
	}.Build()
	File_parkir_proto = out.File
	file_parkir_proto_rawDesc = nil
	file_parkir_proto_goTypes = nil
	file_parkir_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ParkirServiceClient is the client API for ParkirService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ParkirServiceClient interface {
	MasukParkir(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*KarcisParkir, error)
	KeluarParkir(ctx context.Context, in *Keluar, opts ...grpc.CallOption) (*StrukParkir, error)
}

type parkirServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewParkirServiceClient(cc grpc.ClientConnInterface) ParkirServiceClient {
	return &parkirServiceClient{cc}
}

func (c *parkirServiceClient) MasukParkir(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*KarcisParkir, error) {
	out := new(KarcisParkir)
	err := c.cc.Invoke(ctx, "/parkir.ParkirService/MasukParkir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parkirServiceClient) KeluarParkir(ctx context.Context, in *Keluar, opts ...grpc.CallOption) (*StrukParkir, error) {
	out := new(StrukParkir)
	err := c.cc.Invoke(ctx, "/parkir.ParkirService/KeluarParkir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParkirServiceServer is the server API for ParkirService service.
type ParkirServiceServer interface {
	MasukParkir(context.Context, *Empty) (*KarcisParkir, error)
	KeluarParkir(context.Context, *Keluar) (*StrukParkir, error)
}

// UnimplementedParkirServiceServer can be embedded to have forward compatible implementations.
type UnimplementedParkirServiceServer struct {
}

func (*UnimplementedParkirServiceServer) MasukParkir(context.Context, *Empty) (*KarcisParkir, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MasukParkir not implemented")
}
func (*UnimplementedParkirServiceServer) KeluarParkir(context.Context, *Keluar) (*StrukParkir, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeluarParkir not implemented")
}

func RegisterParkirServiceServer(s *grpc.Server, srv ParkirServiceServer) {
	s.RegisterService(&_ParkirService_serviceDesc, srv)
}

func _ParkirService_MasukParkir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParkirServiceServer).MasukParkir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parkir.ParkirService/MasukParkir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParkirServiceServer).MasukParkir(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParkirService_KeluarParkir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Keluar)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParkirServiceServer).KeluarParkir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parkir.ParkirService/KeluarParkir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParkirServiceServer).KeluarParkir(ctx, req.(*Keluar))
	}
	return interceptor(ctx, in, info, handler)
}

var _ParkirService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "parkir.ParkirService",
	HandlerType: (*ParkirServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MasukParkir",
			Handler:    _ParkirService_MasukParkir_Handler,
		},
		{
			MethodName: "KeluarParkir",
			Handler:    _ParkirService_KeluarParkir_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "parkir.proto",
}

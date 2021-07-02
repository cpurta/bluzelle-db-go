// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: crud/lease.proto

package types

import (
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

type Lease struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seconds uint32 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Minutes uint32 `protobuf:"varint,2,opt,name=minutes,proto3" json:"minutes,omitempty"`
	Hours   uint32 `protobuf:"varint,3,opt,name=hours,proto3" json:"hours,omitempty"`
	Days    uint32 `protobuf:"varint,4,opt,name=days,proto3" json:"days,omitempty"`
	Years   uint32 `protobuf:"varint,5,opt,name=years,proto3" json:"years,omitempty"`
}

func (x *Lease) Reset() {
	*x = Lease{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_lease_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lease) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lease) ProtoMessage() {}

func (x *Lease) ProtoReflect() protoreflect.Message {
	mi := &file_crud_lease_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lease.ProtoReflect.Descriptor instead.
func (*Lease) Descriptor() ([]byte, []int) {
	return file_crud_lease_proto_rawDescGZIP(), []int{0}
}

func (x *Lease) GetSeconds() uint32 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *Lease) GetMinutes() uint32 {
	if x != nil {
		return x.Minutes
	}
	return 0
}

func (x *Lease) GetHours() uint32 {
	if x != nil {
		return x.Hours
	}
	return 0
}

func (x *Lease) GetDays() uint32 {
	if x != nil {
		return x.Days
	}
	return 0
}

func (x *Lease) GetYears() uint32 {
	if x != nil {
		return x.Years
	}
	return 0
}

var File_crud_lease_proto protoreflect.FileDescriptor

var file_crud_lease_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x72, 0x75, 0x64, 0x2f, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x62, 0x6c, 0x75, 0x7a, 0x65, 0x6c, 0x6c, 0x65, 0x2e, 0x63, 0x75, 0x72,
	0x69, 0x75, 0x6d, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x22, 0x7b, 0x0a, 0x05, 0x4c, 0x65, 0x61, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x69,
	0x6e, 0x75, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x79, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x64, 0x61, 0x79, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x79, 0x65, 0x61, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x79, 0x65, 0x61, 0x72, 0x73, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6c, 0x75, 0x7a, 0x65, 0x6c, 0x6c, 0x65, 0x2f, 0x63, 0x75, 0x72,
	0x69, 0x75, 0x6d, 0x2f, 0x78, 0x2f, 0x63, 0x72, 0x75, 0x64, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crud_lease_proto_rawDescOnce sync.Once
	file_crud_lease_proto_rawDescData = file_crud_lease_proto_rawDesc
)

func file_crud_lease_proto_rawDescGZIP() []byte {
	file_crud_lease_proto_rawDescOnce.Do(func() {
		file_crud_lease_proto_rawDescData = protoimpl.X.CompressGZIP(file_crud_lease_proto_rawDescData)
	})
	return file_crud_lease_proto_rawDescData
}

var file_crud_lease_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_crud_lease_proto_goTypes = []interface{}{
	(*Lease)(nil), // 0: bluzelle.curium.crud.Lease
}
var file_crud_lease_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_crud_lease_proto_init() }
func file_crud_lease_proto_init() {
	if File_crud_lease_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crud_lease_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lease); i {
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
			RawDescriptor: file_crud_lease_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_crud_lease_proto_goTypes,
		DependencyIndexes: file_crud_lease_proto_depIdxs,
		MessageInfos:      file_crud_lease_proto_msgTypes,
	}.Build()
	File_crud_lease_proto = out.File
	file_crud_lease_proto_rawDesc = nil
	file_crud_lease_proto_goTypes = nil
	file_crud_lease_proto_depIdxs = nil
}

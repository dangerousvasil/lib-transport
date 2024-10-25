// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.28.2
// source: tag_message.proto

package ptransport

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TagLite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Color       string `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	IsDeleted   bool   `protobuf:"varint,5,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
}

func (x *TagLite) Reset() {
	*x = TagLite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagLite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagLite) ProtoMessage() {}

func (x *TagLite) ProtoReflect() protoreflect.Message {
	mi := &file_tag_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagLite.ProtoReflect.Descriptor instead.
func (*TagLite) Descriptor() ([]byte, []int) {
	return file_tag_message_proto_rawDescGZIP(), []int{0}
}

func (x *TagLite) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TagLite) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TagLite) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *TagLite) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TagLite) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

type TagFull struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Name        string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Color       string                 `protobuf:"bytes,5,opt,name=color,proto3" json:"color,omitempty"`
	Description string                 `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	IsDeleted   bool                   `protobuf:"varint,7,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
}

func (x *TagFull) Reset() {
	*x = TagFull{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagFull) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagFull) ProtoMessage() {}

func (x *TagFull) ProtoReflect() protoreflect.Message {
	mi := &file_tag_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagFull.ProtoReflect.Descriptor instead.
func (*TagFull) Descriptor() ([]byte, []int) {
	return file_tag_message_proto_rawDescGZIP(), []int{1}
}

func (x *TagFull) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TagFull) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *TagFull) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *TagFull) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TagFull) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *TagFull) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TagFull) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

var File_tag_message_proto protoreflect.FileDescriptor

var file_tag_message_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x61, 0x67, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x1a, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84,
	0x01, 0x0a, 0x07, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0xfa, 0x01, 0x0a, 0x07, 0x54, 0x61, 0x67, 0x46, 0x75, 0x6c,
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x3b, 0x70, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tag_message_proto_rawDescOnce sync.Once
	file_tag_message_proto_rawDescData = file_tag_message_proto_rawDesc
)

func file_tag_message_proto_rawDescGZIP() []byte {
	file_tag_message_proto_rawDescOnce.Do(func() {
		file_tag_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_tag_message_proto_rawDescData)
	})
	return file_tag_message_proto_rawDescData
}

var file_tag_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tag_message_proto_goTypes = []interface{}{
	(*TagLite)(nil),               // 0: transport.api.TagLite
	(*TagFull)(nil),               // 1: transport.api.TagFull
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_tag_message_proto_depIdxs = []int32{
	2, // 0: transport.api.TagFull.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: transport.api.TagFull.updated_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tag_message_proto_init() }
func file_tag_message_proto_init() {
	if File_tag_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tag_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagLite); i {
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
		file_tag_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagFull); i {
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
			RawDescriptor: file_tag_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tag_message_proto_goTypes,
		DependencyIndexes: file_tag_message_proto_depIdxs,
		MessageInfos:      file_tag_message_proto_msgTypes,
	}.Build()
	File_tag_message_proto = out.File
	file_tag_message_proto_rawDesc = nil
	file_tag_message_proto_goTypes = nil
	file_tag_message_proto_depIdxs = nil
}

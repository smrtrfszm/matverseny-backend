// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: competition.proto

package proto

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

type GetSolutionsResponse_Modification int32

const (
	GetSolutionsResponse_k_CHANGE GetSolutionsResponse_Modification = 0
	GetSolutionsResponse_k_DELETE GetSolutionsResponse_Modification = 1
)

// Enum value maps for GetSolutionsResponse_Modification.
var (
	GetSolutionsResponse_Modification_name = map[int32]string{
		0: "k_CHANGE",
		1: "k_DELETE",
	}
	GetSolutionsResponse_Modification_value = map[string]int32{
		"k_CHANGE": 0,
		"k_DELETE": 1,
	}
)

func (x GetSolutionsResponse_Modification) Enum() *GetSolutionsResponse_Modification {
	p := new(GetSolutionsResponse_Modification)
	*p = x
	return p
}

func (x GetSolutionsResponse_Modification) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetSolutionsResponse_Modification) Descriptor() protoreflect.EnumDescriptor {
	return file_competition_proto_enumTypes[0].Descriptor()
}

func (GetSolutionsResponse_Modification) Type() protoreflect.EnumType {
	return &file_competition_proto_enumTypes[0]
}

func (x GetSolutionsResponse_Modification) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetSolutionsResponse_Modification.Descriptor instead.
func (GetSolutionsResponse_Modification) EnumDescriptor() ([]byte, []int) {
	return file_competition_proto_rawDescGZIP(), []int{1, 0}
}

type GetSolutionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSolutionsRequest) Reset() {
	*x = GetSolutionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_competition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSolutionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSolutionsRequest) ProtoMessage() {}

func (x *GetSolutionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_competition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSolutionsRequest.ProtoReflect.Descriptor instead.
func (*GetSolutionsRequest) Descriptor() ([]byte, []int) {
	return file_competition_proto_rawDescGZIP(), []int{0}
}

type GetSolutionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  GetSolutionsResponse_Modification `protobuf:"varint,1,opt,name=type,proto3,enum=competition.GetSolutionsResponse_Modification" json:"type,omitempty"`
	Value int64                             `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetSolutionsResponse) Reset() {
	*x = GetSolutionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_competition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSolutionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSolutionsResponse) ProtoMessage() {}

func (x *GetSolutionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_competition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSolutionsResponse.ProtoReflect.Descriptor instead.
func (*GetSolutionsResponse) Descriptor() ([]byte, []int) {
	return file_competition_proto_rawDescGZIP(), []int{1}
}

func (x *GetSolutionsResponse) GetType() GetSolutionsResponse_Modification {
	if x != nil {
		return x.Type
	}
	return GetSolutionsResponse_k_CHANGE
}

func (x *GetSolutionsResponse) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type SetSolutionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Value int64 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetSolutionsRequest) Reset() {
	*x = SetSolutionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_competition_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetSolutionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSolutionsRequest) ProtoMessage() {}

func (x *SetSolutionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_competition_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSolutionsRequest.ProtoReflect.Descriptor instead.
func (*SetSolutionsRequest) Descriptor() ([]byte, []int) {
	return file_competition_proto_rawDescGZIP(), []int{2}
}

func (x *SetSolutionsRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SetSolutionsRequest) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type SetSolutionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetSolutionsResponse) Reset() {
	*x = SetSolutionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_competition_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetSolutionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSolutionsResponse) ProtoMessage() {}

func (x *SetSolutionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_competition_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSolutionsResponse.ProtoReflect.Descriptor instead.
func (*SetSolutionsResponse) Descriptor() ([]byte, []int) {
	return file_competition_proto_rawDescGZIP(), []int{3}
}

var File_competition_proto protoreflect.FileDescriptor

var file_competition_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x9c, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x42, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2a, 0x0a, 0x0c, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x08, 0x6b, 0x5f,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x6b, 0x5f, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x10, 0x01, 0x22, 0x3b, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x53, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xbd, 0x01, 0x0a, 0x0b,
	0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x57, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x55, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x53, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_competition_proto_rawDescOnce sync.Once
	file_competition_proto_rawDescData = file_competition_proto_rawDesc
)

func file_competition_proto_rawDescGZIP() []byte {
	file_competition_proto_rawDescOnce.Do(func() {
		file_competition_proto_rawDescData = protoimpl.X.CompressGZIP(file_competition_proto_rawDescData)
	})
	return file_competition_proto_rawDescData
}

var file_competition_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_competition_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_competition_proto_goTypes = []interface{}{
	(GetSolutionsResponse_Modification)(0), // 0: competition.GetSolutionsResponse.Modification
	(*GetSolutionsRequest)(nil),            // 1: competition.GetSolutionsRequest
	(*GetSolutionsResponse)(nil),           // 2: competition.GetSolutionsResponse
	(*SetSolutionsRequest)(nil),            // 3: competition.SetSolutionsRequest
	(*SetSolutionsResponse)(nil),           // 4: competition.SetSolutionsResponse
}
var file_competition_proto_depIdxs = []int32{
	0, // 0: competition.GetSolutionsResponse.type:type_name -> competition.GetSolutionsResponse.Modification
	1, // 1: competition.Competition.GetSolutions:input_type -> competition.GetSolutionsRequest
	3, // 2: competition.Competition.SetSolutions:input_type -> competition.SetSolutionsRequest
	2, // 3: competition.Competition.GetSolutions:output_type -> competition.GetSolutionsResponse
	4, // 4: competition.Competition.SetSolutions:output_type -> competition.SetSolutionsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_competition_proto_init() }
func file_competition_proto_init() {
	if File_competition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_competition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSolutionsRequest); i {
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
		file_competition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSolutionsResponse); i {
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
		file_competition_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetSolutionsRequest); i {
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
		file_competition_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetSolutionsResponse); i {
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
			RawDescriptor: file_competition_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_competition_proto_goTypes,
		DependencyIndexes: file_competition_proto_depIdxs,
		EnumInfos:         file_competition_proto_enumTypes,
		MessageInfos:      file_competition_proto_msgTypes,
	}.Build()
	File_competition_proto = out.File
	file_competition_proto_rawDesc = nil
	file_competition_proto_goTypes = nil
	file_competition_proto_depIdxs = nil
}

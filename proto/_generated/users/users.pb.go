// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: users/users.proto

package users

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

type CreateProfilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string  `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	FirstName *string `protobuf:"bytes,2,opt,name=firstName,proto3,oneof" json:"firstName,omitempty"`
	LastName  *string `protobuf:"bytes,3,opt,name=lastName,proto3,oneof" json:"lastName,omitempty"`
	Email     string  `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone     *string `protobuf:"bytes,5,opt,name=phone,proto3,oneof" json:"phone,omitempty"`
}

func (x *CreateProfilesRequest) Reset() {
	*x = CreateProfilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProfilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProfilesRequest) ProtoMessage() {}

func (x *CreateProfilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProfilesRequest.ProtoReflect.Descriptor instead.
func (*CreateProfilesRequest) Descriptor() ([]byte, []int) {
	return file_users_users_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProfilesRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateProfilesRequest) GetFirstName() string {
	if x != nil && x.FirstName != nil {
		return *x.FirstName
	}
	return ""
}

func (x *CreateProfilesRequest) GetLastName() string {
	if x != nil && x.LastName != nil {
		return *x.LastName
	}
	return ""
}

func (x *CreateProfilesRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateProfilesRequest) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

type CreateProfilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *string `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *CreateProfilesResponse) Reset() {
	*x = CreateProfilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProfilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProfilesResponse) ProtoMessage() {}

func (x *CreateProfilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProfilesResponse.ProtoReflect.Descriptor instead.
func (*CreateProfilesResponse) Descriptor() ([]byte, []int) {
	return file_users_users_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProfilesResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateProfilesResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateProfilesResponse) GetData() string {
	if x != nil && x.Data != nil {
		return *x.Data
	}
	return ""
}

var File_users_users_proto protoreflect.FileDescriptor

var file_users_users_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0xc9, 0x01, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x1f, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x6c, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x32, 0x5e, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x4d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x2d, 0x6a, 0x69, 0x72, 0x61, 0x79, 0x75, 0x73, 0x61, 0x6b, 0x75, 0x6c,
	0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2d, 0x65, 0x63, 0x68, 0x6f, 0x2d, 0x68, 0x6f, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2d, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x5f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_users_users_proto_rawDescOnce sync.Once
	file_users_users_proto_rawDescData = file_users_users_proto_rawDesc
)

func file_users_users_proto_rawDescGZIP() []byte {
	file_users_users_proto_rawDescOnce.Do(func() {
		file_users_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_users_users_proto_rawDescData)
	})
	return file_users_users_proto_rawDescData
}

var file_users_users_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_users_users_proto_goTypes = []interface{}{
	(*CreateProfilesRequest)(nil),  // 0: users.CreateProfilesRequest
	(*CreateProfilesResponse)(nil), // 1: users.CreateProfilesResponse
}
var file_users_users_proto_depIdxs = []int32{
	0, // 0: users.UsersServices.CreateProfiles:input_type -> users.CreateProfilesRequest
	1, // 1: users.UsersServices.CreateProfiles:output_type -> users.CreateProfilesResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_users_users_proto_init() }
func file_users_users_proto_init() {
	if File_users_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_users_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProfilesRequest); i {
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
		file_users_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProfilesResponse); i {
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
	file_users_users_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_users_users_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_users_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_users_users_proto_goTypes,
		DependencyIndexes: file_users_users_proto_depIdxs,
		MessageInfos:      file_users_users_proto_msgTypes,
	}.Build()
	File_users_users_proto = out.File
	file_users_users_proto_rawDesc = nil
	file_users_users_proto_goTypes = nil
	file_users_users_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: friend_service.proto

package friend_service

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Account       string `protobuf:"bytes,2,opt,name=Account,proto3" json:"Account,omitempty"`
	NickName      string `protobuf:"bytes,3,opt,name=NickName,proto3" json:"NickName,omitempty"`
	Pwd           string `protobuf:"bytes,4,opt,name=Pwd,proto3" json:"Pwd,omitempty"`
	Email         string `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
	Phone         string `protobuf:"bytes,6,opt,name=Phone,proto3" json:"Phone,omitempty"`
	PicUrl        string `protobuf:"bytes,7,opt,name=PicUrl,proto3" json:"PicUrl,omitempty"`
	BirthDay      int64  `protobuf:"varint,8,opt,name=BirthDay,proto3" json:"BirthDay,omitempty"`
	CreateTime    int64  `protobuf:"varint,9,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	Status        int32  `protobuf:"varint,10,opt,name=Status,proto3" json:"Status,omitempty"`
	LastLoginTime int64  `protobuf:"varint,11,opt,name=LastLoginTime,proto3" json:"LastLoginTime,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_friend_service_proto_msgTypes[0]
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
	return file_friend_service_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *User) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *User) GetPwd() string {
	if x != nil {
		return x.Pwd
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *User) GetPicUrl() string {
	if x != nil {
		return x.PicUrl
	}
	return ""
}

func (x *User) GetBirthDay() int64 {
	if x != nil {
		return x.BirthDay
	}
	return 0
}

func (x *User) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *User) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *User) GetLastLoginTime() int64 {
	if x != nil {
		return x.LastLoginTime
	}
	return 0
}

type MsgRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *MsgRsp) Reset() {
	*x = MsgRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgRsp) ProtoMessage() {}

func (x *MsgRsp) ProtoReflect() protoreflect.Message {
	mi := &file_friend_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgRsp.ProtoReflect.Descriptor instead.
func (*MsgRsp) Descriptor() ([]byte, []int) {
	return file_friend_service_proto_rawDescGZIP(), []int{1}
}

func (x *MsgRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GroupReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupName string `protobuf:"bytes,1,opt,name=groupName,proto3" json:"groupName,omitempty"`
}

func (x *GroupReq) Reset() {
	*x = GroupReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupReq) ProtoMessage() {}

func (x *GroupReq) ProtoReflect() protoreflect.Message {
	mi := &file_friend_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupReq.ProtoReflect.Descriptor instead.
func (*GroupReq) Descriptor() ([]byte, []int) {
	return file_friend_service_proto_rawDescGZIP(), []int{2}
}

func (x *GroupReq) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

type NameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NameReq) Reset() {
	*x = NameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameReq) ProtoMessage() {}

func (x *NameReq) ProtoReflect() protoreflect.Message {
	mi := &file_friend_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameReq.ProtoReflect.Descriptor instead.
func (*NameReq) Descriptor() ([]byte, []int) {
	return file_friend_service_proto_rawDescGZIP(), []int{3}
}

func (x *NameReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_friend_service_proto protoreflect.FileDescriptor

var file_friend_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69,
	0x64, 0x6c, 0x22, 0x9c, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x77, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x50, 0x77, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x50, 0x69, 0x63, 0x55, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x50, 0x69, 0x63, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x69, 0x72, 0x74, 0x68,
	0x44, 0x61, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x42, 0x69, 0x72, 0x74, 0x68,
	0x44, 0x61, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x4c,
	0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x4c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x1a, 0x0a, 0x06, 0x4d, 0x73, 0x67, 0x52, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x28, 0x0a,
	0x08, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x81, 0x01, 0x0a, 0x0d, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x14, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64,
	0x6c, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x66, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x6c, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x38, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x13, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x6c, 0x2e, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64,
	0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x30, 0x01, 0x42, 0x23, 0x5a, 0x21, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x3b, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_friend_service_proto_rawDescOnce sync.Once
	file_friend_service_proto_rawDescData = file_friend_service_proto_rawDesc
)

func file_friend_service_proto_rawDescGZIP() []byte {
	file_friend_service_proto_rawDescOnce.Do(func() {
		file_friend_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_friend_service_proto_rawDescData)
	})
	return file_friend_service_proto_rawDescData
}

var file_friend_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_friend_service_proto_goTypes = []interface{}{
	(*User)(nil),     // 0: friend_idl.User
	(*MsgRsp)(nil),   // 1: friend_idl.MsgRsp
	(*GroupReq)(nil), // 2: friend_idl.GroupReq
	(*NameReq)(nil),  // 3: friend_idl.NameReq
}
var file_friend_service_proto_depIdxs = []int32{
	2, // 0: friend_idl.FriendService.AddGroup:input_type -> friend_idl.GroupReq
	3, // 1: friend_idl.FriendService.SearchUsers:input_type -> friend_idl.NameReq
	1, // 2: friend_idl.FriendService.AddGroup:output_type -> friend_idl.MsgRsp
	0, // 3: friend_idl.FriendService.SearchUsers:output_type -> friend_idl.User
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_friend_service_proto_init() }
func file_friend_service_proto_init() {
	if File_friend_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_friend_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_friend_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgRsp); i {
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
		file_friend_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupReq); i {
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
		file_friend_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameReq); i {
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
			RawDescriptor: file_friend_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_friend_service_proto_goTypes,
		DependencyIndexes: file_friend_service_proto_depIdxs,
		MessageInfos:      file_friend_service_proto_msgTypes,
	}.Build()
	File_friend_service_proto = out.File
	file_friend_service_proto_rawDesc = nil
	file_friend_service_proto_goTypes = nil
	file_friend_service_proto_depIdxs = nil
}

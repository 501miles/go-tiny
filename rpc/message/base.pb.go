// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: base.proto

package message

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

type ReqMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     int32  `protobuf:"zigzag32,1,opt,name=version,proto3" json:"version,omitempty"`
	T           int64  `protobuf:"varint,2,opt,name=t,proto3" json:"t,omitempty"`
	MsgId       int64  `protobuf:"varint,3,opt,name=msgId,proto3" json:"msgId,omitempty"`
	RequestData []byte `protobuf:"bytes,4,opt,name=requestData,proto3" json:"requestData,omitempty"` //请求参数
	UserId      int64  `protobuf:"varint,5,opt,name=userId,proto3" json:"userId,omitempty"`
	Source      uint32 `protobuf:"varint,6,opt,name=source,proto3" json:"source,omitempty"` //调用来源 api gateway or 内部rpc
	Token       string `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ReqMsg) Reset() {
	*x = ReqMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqMsg) ProtoMessage() {}

func (x *ReqMsg) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqMsg.ProtoReflect.Descriptor instead.
func (*ReqMsg) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{0}
}

func (x *ReqMsg) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ReqMsg) GetT() int64 {
	if x != nil {
		return x.T
	}
	return 0
}

func (x *ReqMsg) GetMsgId() int64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *ReqMsg) GetRequestData() []byte {
	if x != nil {
		return x.RequestData
	}
	return nil
}

func (x *ReqMsg) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ReqMsg) GetSource() uint32 {
	if x != nil {
		return x.Source
	}
	return 0
}

func (x *ReqMsg) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ReqDataMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestData []byte `protobuf:"bytes,1,opt,name=requestData,proto3" json:"requestData,omitempty"`
	UserId      int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *ReqDataMsg) Reset() {
	*x = ReqDataMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqDataMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqDataMsg) ProtoMessage() {}

func (x *ReqDataMsg) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqDataMsg.ProtoReflect.Descriptor instead.
func (*ReqDataMsg) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{1}
}

func (x *ReqDataMsg) GetRequestData() []byte {
	if x != nil {
		return x.RequestData
	}
	return nil
}

func (x *ReqDataMsg) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ResMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgId        int64  `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	T            int64  `protobuf:"varint,2,opt,name=t,proto3" json:"t,omitempty"`
	ResponseData []byte `protobuf:"bytes,3,opt,name=responseData,proto3" json:"responseData,omitempty"` //相应数据
}

func (x *ResMsg) Reset() {
	*x = ResMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResMsg) ProtoMessage() {}

func (x *ResMsg) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResMsg.ProtoReflect.Descriptor instead.
func (*ResMsg) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{2}
}

func (x *ResMsg) GetMsgId() int64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *ResMsg) GetT() int64 {
	if x != nil {
		return x.T
	}
	return 0
}

func (x *ResMsg) GetResponseData() []byte {
	if x != nil {
		return x.ResponseData
	}
	return nil
}

var File_base_proto protoreflect.FileDescriptor

var file_base_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xae, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x71, 0x4d, 0x73, 0x67,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x11, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x46, 0x0a, 0x0a, 0x52, 0x65, 0x71, 0x44, 0x61, 0x74,
	0x61, 0x4d, 0x73, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x50,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x0c,
	0x0a, 0x01, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x74, 0x12, 0x22, 0x0a, 0x0c,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_base_proto_rawDescOnce sync.Once
	file_base_proto_rawDescData = file_base_proto_rawDesc
)

func file_base_proto_rawDescGZIP() []byte {
	file_base_proto_rawDescOnce.Do(func() {
		file_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_proto_rawDescData)
	})
	return file_base_proto_rawDescData
}

var file_base_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_base_proto_goTypes = []interface{}{
	(*ReqMsg)(nil),     // 0: message.ReqMsg
	(*ReqDataMsg)(nil), // 1: message.ReqDataMsg
	(*ResMsg)(nil),     // 2: message.ResMsg
}
var file_base_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_base_proto_init() }
func file_base_proto_init() {
	if File_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqMsg); i {
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
		file_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqDataMsg); i {
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
		file_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResMsg); i {
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
			RawDescriptor: file_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_proto_goTypes,
		DependencyIndexes: file_base_proto_depIdxs,
		MessageInfos:      file_base_proto_msgTypes,
	}.Build()
	File_base_proto = out.File
	file_base_proto_rawDesc = nil
	file_base_proto_goTypes = nil
	file_base_proto_depIdxs = nil
}

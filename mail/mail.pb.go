// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: mail/mail.proto

package mail

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

type TransactionalAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color string `protobuf:"bytes,1,opt,name=Color,proto3" json:"Color,omitempty"`
	Link  string `protobuf:"bytes,2,opt,name=Link,proto3" json:"Link,omitempty"`
	Text  string `protobuf:"bytes,3,opt,name=Text,proto3" json:"Text,omitempty"`
}

func (x *TransactionalAction) Reset() {
	*x = TransactionalAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionalAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionalAction) ProtoMessage() {}

func (x *TransactionalAction) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionalAction.ProtoReflect.Descriptor instead.
func (*TransactionalAction) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionalAction) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *TransactionalAction) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *TransactionalAction) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type TransactionalBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string                 `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Intros  []string               `protobuf:"bytes,2,rep,name=Intros,proto3" json:"Intros,omitempty"`
	Actions []*TransactionalAction `protobuf:"bytes,3,rep,name=Actions,proto3" json:"Actions,omitempty"`
	Outros  []string               `protobuf:"bytes,4,rep,name=Outros,proto3" json:"Outros,omitempty"`
}

func (x *TransactionalBody) Reset() {
	*x = TransactionalBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionalBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionalBody) ProtoMessage() {}

func (x *TransactionalBody) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionalBody.ProtoReflect.Descriptor instead.
func (*TransactionalBody) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionalBody) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TransactionalBody) GetIntros() []string {
	if x != nil {
		return x.Intros
	}
	return nil
}

func (x *TransactionalBody) GetActions() []*TransactionalAction {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *TransactionalBody) GetOutros() []string {
	if x != nil {
		return x.Outros
	}
	return nil
}

type SendTransactionalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From    string             `protobuf:"bytes,1,opt,name=From,proto3" json:"From,omitempty"`
	To      string             `protobuf:"bytes,2,opt,name=To,proto3" json:"To,omitempty"`
	Subject string             `protobuf:"bytes,3,opt,name=Subject,proto3" json:"Subject,omitempty"`
	Body    *TransactionalBody `protobuf:"bytes,4,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (x *SendTransactionalRequest) Reset() {
	*x = SendTransactionalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTransactionalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTransactionalRequest) ProtoMessage() {}

func (x *SendTransactionalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTransactionalRequest.ProtoReflect.Descriptor instead.
func (*SendTransactionalRequest) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{2}
}

func (x *SendTransactionalRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SendTransactionalRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SendTransactionalRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SendTransactionalRequest) GetBody() *TransactionalBody {
	if x != nil {
		return x.Body
	}
	return nil
}

type SendTransactionalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *SendTransactionalResponse) Reset() {
	*x = SendTransactionalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTransactionalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTransactionalResponse) ProtoMessage() {}

func (x *SendTransactionalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mail_mail_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTransactionalResponse.ProtoReflect.Descriptor instead.
func (*SendTransactionalResponse) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{3}
}

func (x *SendTransactionalResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_mail_mail_proto protoreflect.FileDescriptor

var file_mail_mail_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x61, 0x69, 0x6c, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x53, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43,
	0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x22, 0x8c, 0x01, 0x0a,
	0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x42, 0x6f,
	0x64, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x73, 0x12, 0x33,
	0x0a, 0x07, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x72, 0x6f, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x4f, 0x75, 0x74, 0x72, 0x6f, 0x73, 0x22, 0x85, 0x01, 0x0a, 0x18,
	0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x54, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04, 0x42,
	0x6f, 0x64, 0x79, 0x22, 0x35, 0x0a, 0x19, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x5e, 0x0a, 0x04, 0x4d, 0x61,
	0x69, 0x6c, 0x12, 0x56, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x1e, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x69, 0x74, 0x72, 0x61, 0x2d,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mail_mail_proto_rawDescOnce sync.Once
	file_mail_mail_proto_rawDescData = file_mail_mail_proto_rawDesc
)

func file_mail_mail_proto_rawDescGZIP() []byte {
	file_mail_mail_proto_rawDescOnce.Do(func() {
		file_mail_mail_proto_rawDescData = protoimpl.X.CompressGZIP(file_mail_mail_proto_rawDescData)
	})
	return file_mail_mail_proto_rawDescData
}

var file_mail_mail_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mail_mail_proto_goTypes = []any{
	(*TransactionalAction)(nil),       // 0: mail.TransactionalAction
	(*TransactionalBody)(nil),         // 1: mail.TransactionalBody
	(*SendTransactionalRequest)(nil),  // 2: mail.SendTransactionalRequest
	(*SendTransactionalResponse)(nil), // 3: mail.SendTransactionalResponse
}
var file_mail_mail_proto_depIdxs = []int32{
	0, // 0: mail.TransactionalBody.Actions:type_name -> mail.TransactionalAction
	1, // 1: mail.SendTransactionalRequest.Body:type_name -> mail.TransactionalBody
	2, // 2: mail.Mail.SendTransactional:input_type -> mail.SendTransactionalRequest
	3, // 3: mail.Mail.SendTransactional:output_type -> mail.SendTransactionalResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mail_mail_proto_init() }
func file_mail_mail_proto_init() {
	if File_mail_mail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mail_mail_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TransactionalAction); i {
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
		file_mail_mail_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TransactionalBody); i {
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
		file_mail_mail_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SendTransactionalRequest); i {
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
		file_mail_mail_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SendTransactionalResponse); i {
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
			RawDescriptor: file_mail_mail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mail_mail_proto_goTypes,
		DependencyIndexes: file_mail_mail_proto_depIdxs,
		MessageInfos:      file_mail_mail_proto_msgTypes,
	}.Build()
	File_mail_mail_proto = out.File
	file_mail_mail_proto_rawDesc = nil
	file_mail_mail_proto_goTypes = nil
	file_mail_mail_proto_depIdxs = nil
}

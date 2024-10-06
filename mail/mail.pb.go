// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: mail/mail.proto

package mail

import (
	common "github.com/emitra-labs/pb/common"
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

type SendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From    string `protobuf:"bytes,1,opt,name=From,proto3" json:"From,omitempty"`
	To      string `protobuf:"bytes,2,opt,name=To,proto3" json:"To,omitempty"`
	Subject string `protobuf:"bytes,3,opt,name=Subject,proto3" json:"Subject,omitempty"`
	Body    *Body  `protobuf:"bytes,4,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (x *SendRequest) Reset() {
	*x = SendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequest) ProtoMessage() {}

func (x *SendRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SendRequest.ProtoReflect.Descriptor instead.
func (*SendRequest) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{0}
}

func (x *SendRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SendRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SendRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SendRequest) GetBody() *Body {
	if x != nil {
		return x.Body
	}
	return nil
}

type Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string    `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Intros    []string  `protobuf:"bytes,2,rep,name=Intros,proto3" json:"Intros,omitempty"`
	Actions   []*Action `protobuf:"bytes,3,rep,name=Actions,proto3" json:"Actions,omitempty"`
	Outros    []string  `protobuf:"bytes,4,rep,name=Outros,proto3" json:"Outros,omitempty"`
	Greeting  string    `protobuf:"bytes,5,opt,name=Greeting,proto3" json:"Greeting,omitempty"`
	Signature string    `protobuf:"bytes,6,opt,name=Signature,proto3" json:"Signature,omitempty"`
	Title     string    `protobuf:"bytes,7,opt,name=Title,proto3" json:"Title,omitempty"`
}

func (x *Body) Reset() {
	*x = Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Body) ProtoMessage() {}

func (x *Body) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Body.ProtoReflect.Descriptor instead.
func (*Body) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{1}
}

func (x *Body) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Body) GetIntros() []string {
	if x != nil {
		return x.Intros
	}
	return nil
}

func (x *Body) GetActions() []*Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *Body) GetOutros() []string {
	if x != nil {
		return x.Outros
	}
	return nil
}

func (x *Body) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

func (x *Body) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *Body) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instructions string  `protobuf:"bytes,1,opt,name=Instructions,proto3" json:"Instructions,omitempty"`
	Button       *Button `protobuf:"bytes,2,opt,name=Button,proto3" json:"Button,omitempty"`
	InviteCode   string  `protobuf:"bytes,3,opt,name=InviteCode,proto3" json:"InviteCode,omitempty"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{2}
}

func (x *Action) GetInstructions() string {
	if x != nil {
		return x.Instructions
	}
	return ""
}

func (x *Action) GetButton() *Button {
	if x != nil {
		return x.Button
	}
	return nil
}

func (x *Action) GetInviteCode() string {
	if x != nil {
		return x.InviteCode
	}
	return ""
}

type Button struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color     string `protobuf:"bytes,1,opt,name=Color,proto3" json:"Color,omitempty"`
	TextColor string `protobuf:"bytes,2,opt,name=TextColor,proto3" json:"TextColor,omitempty"`
	Text      string `protobuf:"bytes,3,opt,name=Text,proto3" json:"Text,omitempty"`
	Link      string `protobuf:"bytes,4,opt,name=Link,proto3" json:"Link,omitempty"`
}

func (x *Button) Reset() {
	*x = Button{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_mail_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Button) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Button) ProtoMessage() {}

func (x *Button) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Button.ProtoReflect.Descriptor instead.
func (*Button) Descriptor() ([]byte, []int) {
	return file_mail_mail_proto_rawDescGZIP(), []int{3}
}

func (x *Button) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Button) GetTextColor() string {
	if x != nil {
		return x.TextColor
	}
	return ""
}

func (x *Button) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Button) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

var File_mail_mail_proto protoreflect.FileDescriptor

var file_mail_mail_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x61, 0x69, 0x6c, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x0b,
	0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x46,
	0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x54, 0x6f, 0x12,
	0x18, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x42, 0x6f, 0x64,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x42,
	0x6f, 0x64, 0x79, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x22, 0xc2, 0x01, 0x0a, 0x04, 0x42, 0x6f,
	0x64, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x73, 0x12, 0x26,
	0x0a, 0x07, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x72, 0x6f, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x4f, 0x75, 0x74, 0x72, 0x6f, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x72,
	0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x06,
	0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d,
	0x61, 0x69, 0x6c, 0x2e, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x52, 0x06, 0x42, 0x75, 0x74, 0x74,
	0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x64, 0x0a, 0x06, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x54, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x32, 0x3a, 0x0a, 0x04, 0x4d, 0x61, 0x69, 0x6c,
	0x12, 0x32, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x11, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x69, 0x74, 0x72, 0x61, 0x2d, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70,
	0x62, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(*SendRequest)(nil),          // 0: mail.SendRequest
	(*Body)(nil),                 // 1: mail.Body
	(*Action)(nil),               // 2: mail.Action
	(*Button)(nil),               // 3: mail.Button
	(*common.BasicResponse)(nil), // 4: common.BasicResponse
}
var file_mail_mail_proto_depIdxs = []int32{
	1, // 0: mail.SendRequest.Body:type_name -> mail.Body
	2, // 1: mail.Body.Actions:type_name -> mail.Action
	3, // 2: mail.Action.Button:type_name -> mail.Button
	0, // 3: mail.Mail.Send:input_type -> mail.SendRequest
	4, // 4: mail.Mail.Send:output_type -> common.BasicResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mail_mail_proto_init() }
func file_mail_mail_proto_init() {
	if File_mail_mail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mail_mail_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SendRequest); i {
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
			switch v := v.(*Body); i {
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
			switch v := v.(*Action); i {
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
			switch v := v.(*Button); i {
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

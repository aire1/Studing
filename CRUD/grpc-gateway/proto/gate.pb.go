// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: proto/gate.proto

package crud

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

type AuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ExpiresIn int32  `protobuf:"varint,2,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_proto_gate_proto_rawDescGZIP(), []int{0}
}

func (x *AuthResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AuthResponse) GetExpiresIn() int32 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Passhash string `protobuf:"bytes,2,opt,name=passhash,proto3" json:"passhash,omitempty"`
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_proto_gate_proto_rawDescGZIP(), []int{1}
}

func (x *AuthRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *AuthRequest) GetPasshash() string {
	if x != nil {
		return x.Passhash
	}
	return ""
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Passhash string `protobuf:"bytes,2,opt,name=passhash,proto3" json:"passhash,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_proto_gate_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *RegisterRequest) GetPasshash() string {
	if x != nil {
		return x.Passhash
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_proto_gate_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type NoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoteRequest) Reset() {
	*x = NoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteRequest) ProtoMessage() {}

func (x *NoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteRequest.ProtoReflect.Descriptor instead.
func (*NoteRequest) Descriptor() ([]byte, []int) {
	return file_proto_gate_proto_rawDescGZIP(), []int{4}
}

type NoteResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Noted []*Note `protobuf:"bytes,1,rep,name=noted,proto3" json:"noted,omitempty"`
}

func (x *NoteResponce) Reset() {
	*x = NoteResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteResponce) ProtoMessage() {}

func (x *NoteResponce) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gate_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteResponce.ProtoReflect.Descriptor instead.
func (*NoteResponce) Descriptor() ([]byte, []int) {
	return file_proto_gate_proto_rawDescGZIP(), []int{5}
}

func (x *NoteResponce) GetNoted() []*Note {
	if x != nil {
		return x.Noted
	}
	return nil
}

type Note struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Text      string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Note) Reset() {
	*x = Note{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gate_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gate_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Note.ProtoReflect.Descriptor instead.
func (*Note) Descriptor() ([]byte, []int) {
	return file_proto_gate_proto_rawDescGZIP(), []int{6}
}

func (x *Note) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Note) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Note) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Note) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Note) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_proto_gate_proto protoreflect.FileDescriptor

var file_proto_gate_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x67, 0x61, 0x74, 0x65, 0x22, 0x43, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x22, 0x3f, 0x0a,
	0x0b, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x68, 0x61, 0x73, 0x68, 0x22, 0x43,
	0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x68,
	0x61, 0x73, 0x68, 0x22, 0x2c, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x30, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65,
	0x12, 0x20, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x74,
	0x65, 0x64, 0x22, 0x7c, 0x0a, 0x04, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x32, 0xad, 0x01, 0x0a, 0x04, 0x47, 0x61, 0x74, 0x65, 0x12, 0x37, 0x0a, 0x0e, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x11, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x15,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x11, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65,
	0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x63, 0x72, 0x75, 0x64, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_gate_proto_rawDescOnce sync.Once
	file_proto_gate_proto_rawDescData = file_proto_gate_proto_rawDesc
)

func file_proto_gate_proto_rawDescGZIP() []byte {
	file_proto_gate_proto_rawDescOnce.Do(func() {
		file_proto_gate_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_gate_proto_rawDescData)
	})
	return file_proto_gate_proto_rawDescData
}

var file_proto_gate_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_gate_proto_goTypes = []interface{}{
	(*AuthResponse)(nil),     // 0: gate.AuthResponse
	(*AuthRequest)(nil),      // 1: gate.AuthRequest
	(*RegisterRequest)(nil),  // 2: gate.RegisterRequest
	(*RegisterResponse)(nil), // 3: gate.RegisterResponse
	(*NoteRequest)(nil),      // 4: gate.NoteRequest
	(*NoteResponce)(nil),     // 5: gate.NoteResponce
	(*Note)(nil),             // 6: gate.Note
}
var file_proto_gate_proto_depIdxs = []int32{
	6, // 0: gate.NoteResponce.noted:type_name -> gate.Note
	1, // 1: gate.Gate.Authentication:input_type -> gate.AuthRequest
	2, // 2: gate.Gate.Register:input_type -> gate.RegisterRequest
	4, // 3: gate.Gate.GetNotes:input_type -> gate.NoteRequest
	0, // 4: gate.Gate.Authentication:output_type -> gate.AuthResponse
	3, // 5: gate.Gate.Register:output_type -> gate.RegisterResponse
	5, // 6: gate.Gate.GetNotes:output_type -> gate.NoteResponce
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_gate_proto_init() }
func file_proto_gate_proto_init() {
	if File_proto_gate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_gate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResponse); i {
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
		file_proto_gate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest); i {
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
		file_proto_gate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_proto_gate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_proto_gate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteRequest); i {
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
		file_proto_gate_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteResponce); i {
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
		file_proto_gate_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Note); i {
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
			RawDescriptor: file_proto_gate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_gate_proto_goTypes,
		DependencyIndexes: file_proto_gate_proto_depIdxs,
		MessageInfos:      file_proto_gate_proto_msgTypes,
	}.Build()
	File_proto_gate_proto = out.File
	file_proto_gate_proto_rawDesc = nil
	file_proto_gate_proto_goTypes = nil
	file_proto_gate_proto_depIdxs = nil
}

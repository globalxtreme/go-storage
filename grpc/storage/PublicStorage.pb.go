// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.28.1
// source: PublicStorage.proto

package storage

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

type PublicStorageCredential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientID     string `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
	ClientSecret string `protobuf:"bytes,2,opt,name=clientSecret,proto3" json:"clientSecret,omitempty"`
}

func (x *PublicStorageCredential) Reset() {
	*x = PublicStorageCredential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PublicStorage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicStorageCredential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicStorageCredential) ProtoMessage() {}

func (x *PublicStorageCredential) ProtoReflect() protoreflect.Message {
	mi := &file_PublicStorage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicStorageCredential.ProtoReflect.Descriptor instead.
func (*PublicStorageCredential) Descriptor() ([]byte, []int) {
	return file_PublicStorage_proto_rawDescGZIP(), []int{0}
}

func (x *PublicStorageCredential) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

func (x *PublicStorageCredential) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

type PublicStorageStoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content      []byte                   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Path         string                   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Filename     string                   `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
	Title        string                   `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	MimeType     string                   `protobuf:"bytes,5,opt,name=mimeType,proto3" json:"mimeType,omitempty"`
	OriginalName string                   `protobuf:"bytes,6,opt,name=originalName,proto3" json:"originalName,omitempty"`
	Credential   *PublicStorageCredential `protobuf:"bytes,7,opt,name=credential,proto3" json:"credential,omitempty"`
}

func (x *PublicStorageStoreRequest) Reset() {
	*x = PublicStorageStoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PublicStorage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicStorageStoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicStorageStoreRequest) ProtoMessage() {}

func (x *PublicStorageStoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_PublicStorage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicStorageStoreRequest.ProtoReflect.Descriptor instead.
func (*PublicStorageStoreRequest) Descriptor() ([]byte, []int) {
	return file_PublicStorage_proto_rawDescGZIP(), []int{1}
}

func (x *PublicStorageStoreRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *PublicStorageStoreRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *PublicStorageStoreRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *PublicStorageStoreRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PublicStorageStoreRequest) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *PublicStorageStoreRequest) GetOriginalName() string {
	if x != nil {
		return x.OriginalName
	}
	return ""
}

func (x *PublicStorageStoreRequest) GetCredential() *PublicStorageCredential {
	if x != nil {
		return x.Credential
	}
	return nil
}

type PublicStorageDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path       string                   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Credential *PublicStorageCredential `protobuf:"bytes,2,opt,name=credential,proto3" json:"credential,omitempty"`
}

func (x *PublicStorageDeleteRequest) Reset() {
	*x = PublicStorageDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PublicStorage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicStorageDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicStorageDeleteRequest) ProtoMessage() {}

func (x *PublicStorageDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_PublicStorage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicStorageDeleteRequest.ProtoReflect.Descriptor instead.
func (*PublicStorageDeleteRequest) Descriptor() ([]byte, []int) {
	return file_PublicStorage_proto_rawDescGZIP(), []int{2}
}

func (x *PublicStorageDeleteRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *PublicStorageDeleteRequest) GetCredential() *PublicStorageCredential {
	if x != nil {
		return x.Credential
	}
	return nil
}

type PublicStorageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string               `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Result  *PublicStorageResult `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PublicStorageResponse) Reset() {
	*x = PublicStorageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PublicStorage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicStorageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicStorageResponse) ProtoMessage() {}

func (x *PublicStorageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_PublicStorage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicStorageResponse.ProtoReflect.Descriptor instead.
func (*PublicStorageResponse) Descriptor() ([]byte, []int) {
	return file_PublicStorage_proto_rawDescGZIP(), []int{3}
}

func (x *PublicStorageResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PublicStorageResponse) GetResult() *PublicStorageResult {
	if x != nil {
		return x.Result
	}
	return nil
}

type PublicStorageResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path     string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	FullPath string `protobuf:"bytes,2,opt,name=fullPath,proto3" json:"fullPath,omitempty"`
}

func (x *PublicStorageResult) Reset() {
	*x = PublicStorageResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PublicStorage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicStorageResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicStorageResult) ProtoMessage() {}

func (x *PublicStorageResult) ProtoReflect() protoreflect.Message {
	mi := &file_PublicStorage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicStorageResult.ProtoReflect.Descriptor instead.
func (*PublicStorageResult) Descriptor() ([]byte, []int) {
	return file_PublicStorage_proto_rawDescGZIP(), []int{4}
}

func (x *PublicStorageResult) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *PublicStorageResult) GetFullPath() string {
	if x != nil {
		return x.FullPath
	}
	return ""
}

var File_PublicStorage_proto protoreflect.FileDescriptor

var file_PublicStorage_proto_rawDesc = []byte{
	0x0a, 0x13, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0x59,
	0x0a, 0x17, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0xfd, 0x01, 0x0a, 0x19, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0x72, 0x0a, 0x1a, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x40, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0x67, 0x0a,
	0x15, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x34, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x45, 0x0a, 0x13, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x32, 0xb1, 0x01,
	0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12,
	0x4f, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x22, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x12, 0x4f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PublicStorage_proto_rawDescOnce sync.Once
	file_PublicStorage_proto_rawDescData = file_PublicStorage_proto_rawDesc
)

func file_PublicStorage_proto_rawDescGZIP() []byte {
	file_PublicStorage_proto_rawDescOnce.Do(func() {
		file_PublicStorage_proto_rawDescData = protoimpl.X.CompressGZIP(file_PublicStorage_proto_rawDescData)
	})
	return file_PublicStorage_proto_rawDescData
}

var file_PublicStorage_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_PublicStorage_proto_goTypes = []interface{}{
	(*PublicStorageCredential)(nil),    // 0: storage.PublicStorageCredential
	(*PublicStorageStoreRequest)(nil),  // 1: storage.PublicStorageStoreRequest
	(*PublicStorageDeleteRequest)(nil), // 2: storage.PublicStorageDeleteRequest
	(*PublicStorageResponse)(nil),      // 3: storage.PublicStorageResponse
	(*PublicStorageResult)(nil),        // 4: storage.PublicStorageResult
}
var file_PublicStorage_proto_depIdxs = []int32{
	0, // 0: storage.PublicStorageStoreRequest.credential:type_name -> storage.PublicStorageCredential
	0, // 1: storage.PublicStorageDeleteRequest.credential:type_name -> storage.PublicStorageCredential
	4, // 2: storage.PublicStorageResponse.result:type_name -> storage.PublicStorageResult
	1, // 3: storage.PublicStorage.Store:input_type -> storage.PublicStorageStoreRequest
	2, // 4: storage.PublicStorage.Delete:input_type -> storage.PublicStorageDeleteRequest
	3, // 5: storage.PublicStorage.Store:output_type -> storage.PublicStorageResponse
	3, // 6: storage.PublicStorage.Delete:output_type -> storage.PublicStorageResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_PublicStorage_proto_init() }
func file_PublicStorage_proto_init() {
	if File_PublicStorage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PublicStorage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicStorageCredential); i {
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
		file_PublicStorage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicStorageStoreRequest); i {
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
		file_PublicStorage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicStorageDeleteRequest); i {
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
		file_PublicStorage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicStorageResponse); i {
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
		file_PublicStorage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicStorageResult); i {
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
			RawDescriptor: file_PublicStorage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_PublicStorage_proto_goTypes,
		DependencyIndexes: file_PublicStorage_proto_depIdxs,
		MessageInfos:      file_PublicStorage_proto_msgTypes,
	}.Build()
	File_PublicStorage_proto = out.File
	file_PublicStorage_proto_rawDesc = nil
	file_PublicStorage_proto_goTypes = nil
	file_PublicStorage_proto_depIdxs = nil
}

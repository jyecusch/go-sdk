// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: v1/documents.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This represents the high level document collection
	// e.g. a table in SQL
	Collection string `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	Key        string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// A simple JSON object to store
	Document *structpb.Struct `protobuf:"bytes,3,opt,name=document,proto3" json:"document,omitempty"`
}

func (x *CreateDocumentRequest) Reset() {
	*x = CreateDocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_documents_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDocumentRequest) ProtoMessage() {}

func (x *CreateDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_documents_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDocumentRequest.ProtoReflect.Descriptor instead.
func (*CreateDocumentRequest) Descriptor() ([]byte, []int) {
	return file_v1_documents_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDocumentRequest) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *CreateDocumentRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CreateDocumentRequest) GetDocument() *structpb.Struct {
	if x != nil {
		return x.Document
	}
	return nil
}

type GetDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collection string `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	Key        string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetDocumentRequest) Reset() {
	*x = GetDocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_documents_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDocumentRequest) ProtoMessage() {}

func (x *GetDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_documents_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDocumentRequest.ProtoReflect.Descriptor instead.
func (*GetDocumentRequest) Descriptor() ([]byte, []int) {
	return file_v1_documents_proto_rawDescGZIP(), []int{1}
}

func (x *GetDocumentRequest) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *GetDocumentRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetDocumentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Document *structpb.Struct `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
}

func (x *GetDocumentReply) Reset() {
	*x = GetDocumentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_documents_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDocumentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDocumentReply) ProtoMessage() {}

func (x *GetDocumentReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_documents_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDocumentReply.ProtoReflect.Descriptor instead.
func (*GetDocumentReply) Descriptor() ([]byte, []int) {
	return file_v1_documents_proto_rawDescGZIP(), []int{2}
}

func (x *GetDocumentReply) GetDocument() *structpb.Struct {
	if x != nil {
		return x.Document
	}
	return nil
}

type UpdateDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This represents the high level document collection
	// e.g. a table in SQL
	Collection string `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	// The key of the document
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// A simple JSON object to store
	Document *structpb.Struct `protobuf:"bytes,3,opt,name=document,proto3" json:"document,omitempty"`
}

func (x *UpdateDocumentRequest) Reset() {
	*x = UpdateDocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_documents_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDocumentRequest) ProtoMessage() {}

func (x *UpdateDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_documents_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDocumentRequest.ProtoReflect.Descriptor instead.
func (*UpdateDocumentRequest) Descriptor() ([]byte, []int) {
	return file_v1_documents_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateDocumentRequest) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *UpdateDocumentRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UpdateDocumentRequest) GetDocument() *structpb.Struct {
	if x != nil {
		return x.Document
	}
	return nil
}

type DeleteDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collection string `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	Key        string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteDocumentRequest) Reset() {
	*x = DeleteDocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_documents_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDocumentRequest) ProtoMessage() {}

func (x *DeleteDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_documents_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDocumentRequest.ProtoReflect.Descriptor instead.
func (*DeleteDocumentRequest) Descriptor() ([]byte, []int) {
	return file_v1_documents_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteDocumentRequest) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *DeleteDocumentRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_v1_documents_proto protoreflect.FileDescriptor

var file_v1_documents_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x33, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0x46, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x47, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x33, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x7e, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x33, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x49, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x32, 0xec, 0x02, 0x0a, 0x09, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x54,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x2a, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x5d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x27, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6e,
	0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x54, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x54, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x2e, 0x6e, 0x69,
	0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42,
	0x39, 0x0a, 0x16, 0x69, 0x6f, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x0f, 0x4e, 0x69, 0x74, 0x72, 0x69,
	0x63, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x01, 0x5a, 0x0c, 0x6e, 0x69,
	0x74, 0x72, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_v1_documents_proto_rawDescOnce sync.Once
	file_v1_documents_proto_rawDescData = file_v1_documents_proto_rawDesc
)

func file_v1_documents_proto_rawDescGZIP() []byte {
	file_v1_documents_proto_rawDescOnce.Do(func() {
		file_v1_documents_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_documents_proto_rawDescData)
	})
	return file_v1_documents_proto_rawDescData
}

var file_v1_documents_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_documents_proto_goTypes = []interface{}{
	(*CreateDocumentRequest)(nil), // 0: nitric.v1.documents.CreateDocumentRequest
	(*GetDocumentRequest)(nil),    // 1: nitric.v1.documents.GetDocumentRequest
	(*GetDocumentReply)(nil),      // 2: nitric.v1.documents.GetDocumentReply
	(*UpdateDocumentRequest)(nil), // 3: nitric.v1.documents.UpdateDocumentRequest
	(*DeleteDocumentRequest)(nil), // 4: nitric.v1.documents.DeleteDocumentRequest
	(*structpb.Struct)(nil),       // 5: google.protobuf.Struct
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
}
var file_v1_documents_proto_depIdxs = []int32{
	5, // 0: nitric.v1.documents.CreateDocumentRequest.document:type_name -> google.protobuf.Struct
	5, // 1: nitric.v1.documents.GetDocumentReply.document:type_name -> google.protobuf.Struct
	5, // 2: nitric.v1.documents.UpdateDocumentRequest.document:type_name -> google.protobuf.Struct
	0, // 3: nitric.v1.documents.Documents.CreateDocument:input_type -> nitric.v1.documents.CreateDocumentRequest
	1, // 4: nitric.v1.documents.Documents.GetDocument:input_type -> nitric.v1.documents.GetDocumentRequest
	3, // 5: nitric.v1.documents.Documents.UpdateDocument:input_type -> nitric.v1.documents.UpdateDocumentRequest
	4, // 6: nitric.v1.documents.Documents.DeleteDocument:input_type -> nitric.v1.documents.DeleteDocumentRequest
	6, // 7: nitric.v1.documents.Documents.CreateDocument:output_type -> google.protobuf.Empty
	2, // 8: nitric.v1.documents.Documents.GetDocument:output_type -> nitric.v1.documents.GetDocumentReply
	6, // 9: nitric.v1.documents.Documents.UpdateDocument:output_type -> google.protobuf.Empty
	6, // 10: nitric.v1.documents.Documents.DeleteDocument:output_type -> google.protobuf.Empty
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_documents_proto_init() }
func file_v1_documents_proto_init() {
	if File_v1_documents_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_documents_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDocumentRequest); i {
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
		file_v1_documents_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDocumentRequest); i {
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
		file_v1_documents_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDocumentReply); i {
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
		file_v1_documents_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDocumentRequest); i {
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
		file_v1_documents_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDocumentRequest); i {
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
			RawDescriptor: file_v1_documents_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_documents_proto_goTypes,
		DependencyIndexes: file_v1_documents_proto_depIdxs,
		MessageInfos:      file_v1_documents_proto_msgTypes,
	}.Build()
	File_v1_documents_proto = out.File
	file_v1_documents_proto_rawDesc = nil
	file_v1_documents_proto_goTypes = nil
	file_v1_documents_proto_depIdxs = nil
}

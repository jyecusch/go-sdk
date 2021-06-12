// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: event/v1/event.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request to publish an event to a topic
type EventPublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the topic to publish the event to
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	// The event to be published
	Event *NitricEvent `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *EventPublishRequest) Reset() {
	*x = EventPublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_v1_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventPublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventPublishRequest) ProtoMessage() {}

func (x *EventPublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_v1_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventPublishRequest.ProtoReflect.Descriptor instead.
func (*EventPublishRequest) Descriptor() ([]byte, []int) {
	return file_event_v1_event_proto_rawDescGZIP(), []int{0}
}

func (x *EventPublishRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *EventPublishRequest) GetEvent() *NitricEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

// Result of publishing an event
type EventPublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the published message
	// When an id was not supplied
	// one should be automatically generated
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EventPublishResponse) Reset() {
	*x = EventPublishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_v1_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventPublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventPublishResponse) ProtoMessage() {}

func (x *EventPublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_v1_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventPublishResponse.ProtoReflect.Descriptor instead.
func (*EventPublishResponse) Descriptor() ([]byte, []int) {
	return file_event_v1_event_proto_rawDescGZIP(), []int{1}
}

func (x *EventPublishResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Request for the Topic List method
type TopicListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TopicListRequest) Reset() {
	*x = TopicListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_v1_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicListRequest) ProtoMessage() {}

func (x *TopicListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_v1_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicListRequest.ProtoReflect.Descriptor instead.
func (*TopicListRequest) Descriptor() ([]byte, []int) {
	return file_event_v1_event_proto_rawDescGZIP(), []int{2}
}

// Topic List Response
type TopicListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of found topics
	Topics []*NitricTopic `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (x *TopicListResponse) Reset() {
	*x = TopicListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_v1_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicListResponse) ProtoMessage() {}

func (x *TopicListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_event_v1_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicListResponse.ProtoReflect.Descriptor instead.
func (*TopicListResponse) Descriptor() ([]byte, []int) {
	return file_event_v1_event_proto_rawDescGZIP(), []int{3}
}

func (x *TopicListResponse) GetTopics() []*NitricTopic {
	if x != nil {
		return x.Topics
	}
	return nil
}

// Represents an event topic
type NitricTopic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Nitric name for the topic
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NitricTopic) Reset() {
	*x = NitricTopic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_v1_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NitricTopic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NitricTopic) ProtoMessage() {}

func (x *NitricTopic) ProtoReflect() protoreflect.Message {
	mi := &file_event_v1_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NitricTopic.ProtoReflect.Descriptor instead.
func (*NitricTopic) Descriptor() ([]byte, []int) {
	return file_event_v1_event_proto_rawDescGZIP(), []int{4}
}

func (x *NitricTopic) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Nitric Event Model
type NitricEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A Unique ID for the Nitric Event
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// A content hint for the events payload
	PayloadType string `protobuf:"bytes,2,opt,name=payloadType,proto3" json:"payloadType,omitempty"`
	// The payload of the event
	Payload *structpb.Struct `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *NitricEvent) Reset() {
	*x = NitricEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_v1_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NitricEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NitricEvent) ProtoMessage() {}

func (x *NitricEvent) ProtoReflect() protoreflect.Message {
	mi := &file_event_v1_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NitricEvent.ProtoReflect.Descriptor instead.
func (*NitricEvent) Descriptor() ([]byte, []int) {
	return file_event_v1_event_proto_rawDescGZIP(), []int{5}
}

func (x *NitricEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NitricEvent) GetPayloadType() string {
	if x != nil {
		return x.PayloadType
	}
	return ""
}

func (x *NitricEvent) GetPayload() *structpb.Struct {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_event_v1_event_proto protoreflect.FileDescriptor

var file_event_v1_event_proto_rawDesc = []byte{
	0x0a, 0x14, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x13, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x12, 0x32, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x12,
	0x0a, 0x10, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x49, 0x0a, 0x11, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x69, 0x74, 0x72, 0x69, 0x63,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x22, 0x21, 0x0a,
	0x0b, 0x4e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x72, 0x0a, 0x0b, 0x4e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x32, 0x5f, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x56, 0x0a,
	0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x24, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69,
	0x63, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x56, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x4d,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6e, 0x69, 0x74, 0x72,
	0x69, 0x63, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x62, 0x0a,
	0x18, 0x69, 0x6f, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x50, 0x01, 0x5a, 0x0c, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0xaa, 0x02, 0x15, 0x4e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0xca, 0x02, 0x15, 0x4e, 0x69, 0x74, 0x72,
	0x69, 0x63, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x5c, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_v1_event_proto_rawDescOnce sync.Once
	file_event_v1_event_proto_rawDescData = file_event_v1_event_proto_rawDesc
)

func file_event_v1_event_proto_rawDescGZIP() []byte {
	file_event_v1_event_proto_rawDescOnce.Do(func() {
		file_event_v1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_v1_event_proto_rawDescData)
	})
	return file_event_v1_event_proto_rawDescData
}

var file_event_v1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_event_v1_event_proto_goTypes = []interface{}{
	(*EventPublishRequest)(nil),  // 0: nitric.event.v1.EventPublishRequest
	(*EventPublishResponse)(nil), // 1: nitric.event.v1.EventPublishResponse
	(*TopicListRequest)(nil),     // 2: nitric.event.v1.TopicListRequest
	(*TopicListResponse)(nil),    // 3: nitric.event.v1.TopicListResponse
	(*NitricTopic)(nil),          // 4: nitric.event.v1.NitricTopic
	(*NitricEvent)(nil),          // 5: nitric.event.v1.NitricEvent
	(*structpb.Struct)(nil),      // 6: google.protobuf.Struct
}
var file_event_v1_event_proto_depIdxs = []int32{
	5, // 0: nitric.event.v1.EventPublishRequest.event:type_name -> nitric.event.v1.NitricEvent
	4, // 1: nitric.event.v1.TopicListResponse.topics:type_name -> nitric.event.v1.NitricTopic
	6, // 2: nitric.event.v1.NitricEvent.payload:type_name -> google.protobuf.Struct
	0, // 3: nitric.event.v1.Event.Publish:input_type -> nitric.event.v1.EventPublishRequest
	2, // 4: nitric.event.v1.Topic.List:input_type -> nitric.event.v1.TopicListRequest
	1, // 5: nitric.event.v1.Event.Publish:output_type -> nitric.event.v1.EventPublishResponse
	3, // 6: nitric.event.v1.Topic.List:output_type -> nitric.event.v1.TopicListResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_event_v1_event_proto_init() }
func file_event_v1_event_proto_init() {
	if File_event_v1_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_v1_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventPublishRequest); i {
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
		file_event_v1_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventPublishResponse); i {
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
		file_event_v1_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicListRequest); i {
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
		file_event_v1_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicListResponse); i {
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
		file_event_v1_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NitricTopic); i {
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
		file_event_v1_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NitricEvent); i {
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
			RawDescriptor: file_event_v1_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_event_v1_event_proto_goTypes,
		DependencyIndexes: file_event_v1_event_proto_depIdxs,
		MessageInfos:      file_event_v1_event_proto_msgTypes,
	}.Build()
	File_event_v1_event_proto = out.File
	file_event_v1_event_proto_rawDesc = nil
	file_event_v1_event_proto_goTypes = nil
	file_event_v1_event_proto_depIdxs = nil
}

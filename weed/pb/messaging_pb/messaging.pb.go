// Code generated by protoc-gen-go.
// source: messaging.proto
// DO NOT EDIT!

/*
Package messaging_pb is a generated protocol buffer package.

It is generated from these files:
	messaging.proto

It has these top-level messages:
	SubscriberMessage
	RawData
	Message
	BrokerMessage
	PublishRequest
	PublishResponse
	ConfigureTopicRequest
	ConfigureTopicResponse
	GetTopicConfigurationRequest
	GetTopicConfigurationResponse
	TopicConfiguration
*/
package messaging_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SubscriberMessage_InitMessage_StartPosition int32

const (
	SubscriberMessage_InitMessage_LATEST    SubscriberMessage_InitMessage_StartPosition = 0
	SubscriberMessage_InitMessage_EARLIEST  SubscriberMessage_InitMessage_StartPosition = 1
	SubscriberMessage_InitMessage_TIMESTAMP SubscriberMessage_InitMessage_StartPosition = 2
)

var SubscriberMessage_InitMessage_StartPosition_name = map[int32]string{
	0: "LATEST",
	1: "EARLIEST",
	2: "TIMESTAMP",
}
var SubscriberMessage_InitMessage_StartPosition_value = map[string]int32{
	"LATEST":    0,
	"EARLIEST":  1,
	"TIMESTAMP": 2,
}

func (x SubscriberMessage_InitMessage_StartPosition) String() string {
	return proto.EnumName(SubscriberMessage_InitMessage_StartPosition_name, int32(x))
}
func (SubscriberMessage_InitMessage_StartPosition) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 0}
}

type SubscriberMessage struct {
	Init *SubscriberMessage_InitMessage `protobuf:"bytes,1,opt,name=init" json:"init,omitempty"`
	Ack  *SubscriberMessage_AckMessage  `protobuf:"bytes,2,opt,name=ack" json:"ack,omitempty"`
}

func (m *SubscriberMessage) Reset()                    { *m = SubscriberMessage{} }
func (m *SubscriberMessage) String() string            { return proto.CompactTextString(m) }
func (*SubscriberMessage) ProtoMessage()               {}
func (*SubscriberMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SubscriberMessage) GetInit() *SubscriberMessage_InitMessage {
	if m != nil {
		return m.Init
	}
	return nil
}

func (m *SubscriberMessage) GetAck() *SubscriberMessage_AckMessage {
	if m != nil {
		return m.Ack
	}
	return nil
}

type SubscriberMessage_InitMessage struct {
	Namespace     string                                      `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Topic         string                                      `protobuf:"bytes,2,opt,name=topic" json:"topic,omitempty"`
	Partition     int32                                       `protobuf:"varint,3,opt,name=partition" json:"partition,omitempty"`
	StartPosition SubscriberMessage_InitMessage_StartPosition `protobuf:"varint,4,opt,name=startPosition,enum=messaging_pb.SubscriberMessage_InitMessage_StartPosition" json:"startPosition,omitempty"`
	TimestampNs   int64                                       `protobuf:"varint,5,opt,name=timestampNs" json:"timestampNs,omitempty"`
	SubscriberId  string                                      `protobuf:"bytes,6,opt,name=subscriber_id,json=subscriberId" json:"subscriber_id,omitempty"`
}

func (m *SubscriberMessage_InitMessage) Reset()         { *m = SubscriberMessage_InitMessage{} }
func (m *SubscriberMessage_InitMessage) String() string { return proto.CompactTextString(m) }
func (*SubscriberMessage_InitMessage) ProtoMessage()    {}
func (*SubscriberMessage_InitMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

func (m *SubscriberMessage_InitMessage) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *SubscriberMessage_InitMessage) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *SubscriberMessage_InitMessage) GetPartition() int32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

func (m *SubscriberMessage_InitMessage) GetStartPosition() SubscriberMessage_InitMessage_StartPosition {
	if m != nil {
		return m.StartPosition
	}
	return SubscriberMessage_InitMessage_LATEST
}

func (m *SubscriberMessage_InitMessage) GetTimestampNs() int64 {
	if m != nil {
		return m.TimestampNs
	}
	return 0
}

func (m *SubscriberMessage_InitMessage) GetSubscriberId() string {
	if m != nil {
		return m.SubscriberId
	}
	return ""
}

type SubscriberMessage_AckMessage struct {
	MessageId int64 `protobuf:"varint,1,opt,name=message_id,json=messageId" json:"message_id,omitempty"`
}

func (m *SubscriberMessage_AckMessage) Reset()                    { *m = SubscriberMessage_AckMessage{} }
func (m *SubscriberMessage_AckMessage) String() string            { return proto.CompactTextString(m) }
func (*SubscriberMessage_AckMessage) ProtoMessage()               {}
func (*SubscriberMessage_AckMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *SubscriberMessage_AckMessage) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

type RawData struct {
	Key     []byte            `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value   []byte            `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Headers map[string][]byte `protobuf:"bytes,3,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *RawData) Reset()                    { *m = RawData{} }
func (m *RawData) String() string            { return proto.CompactTextString(m) }
func (*RawData) ProtoMessage()               {}
func (*RawData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RawData) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RawData) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *RawData) GetHeaders() map[string][]byte {
	if m != nil {
		return m.Headers
	}
	return nil
}

type Message struct {
	Timestamp int64             `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Key       []byte            `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value     []byte            `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Headers   map[string][]byte `protobuf:"bytes,4,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Message) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Message) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Message) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Message) GetHeaders() map[string][]byte {
	if m != nil {
		return m.Headers
	}
	return nil
}

type BrokerMessage struct {
	Data     *Message                       `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	Redirect *BrokerMessage_RedirectMessage `protobuf:"bytes,2,opt,name=redirect" json:"redirect,omitempty"`
}

func (m *BrokerMessage) Reset()                    { *m = BrokerMessage{} }
func (m *BrokerMessage) String() string            { return proto.CompactTextString(m) }
func (*BrokerMessage) ProtoMessage()               {}
func (*BrokerMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BrokerMessage) GetData() *Message {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BrokerMessage) GetRedirect() *BrokerMessage_RedirectMessage {
	if m != nil {
		return m.Redirect
	}
	return nil
}

type BrokerMessage_RedirectMessage struct {
	NewBroker string `protobuf:"bytes,1,opt,name=new_broker,json=newBroker" json:"new_broker,omitempty"`
}

func (m *BrokerMessage_RedirectMessage) Reset()         { *m = BrokerMessage_RedirectMessage{} }
func (m *BrokerMessage_RedirectMessage) String() string { return proto.CompactTextString(m) }
func (*BrokerMessage_RedirectMessage) ProtoMessage()    {}
func (*BrokerMessage_RedirectMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

func (m *BrokerMessage_RedirectMessage) GetNewBroker() string {
	if m != nil {
		return m.NewBroker
	}
	return ""
}

type PublishRequest struct {
	Init *PublishRequest_InitMessage `protobuf:"bytes,1,opt,name=init" json:"init,omitempty"`
	Data *PublishRequest_RawData     `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *PublishRequest) Reset()                    { *m = PublishRequest{} }
func (m *PublishRequest) String() string            { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()               {}
func (*PublishRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PublishRequest) GetInit() *PublishRequest_InitMessage {
	if m != nil {
		return m.Init
	}
	return nil
}

func (m *PublishRequest) GetData() *PublishRequest_RawData {
	if m != nil {
		return m.Data
	}
	return nil
}

type PublishRequest_InitMessage struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Topic     string `protobuf:"bytes,2,opt,name=topic" json:"topic,omitempty"`
	Partition int32  `protobuf:"varint,3,opt,name=partition" json:"partition,omitempty"`
}

func (m *PublishRequest_InitMessage) Reset()                    { *m = PublishRequest_InitMessage{} }
func (m *PublishRequest_InitMessage) String() string            { return proto.CompactTextString(m) }
func (*PublishRequest_InitMessage) ProtoMessage()               {}
func (*PublishRequest_InitMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func (m *PublishRequest_InitMessage) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *PublishRequest_InitMessage) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PublishRequest_InitMessage) GetPartition() int32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

type PublishRequest_RawData struct {
	Key     []byte            `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value   []byte            `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Headers map[string][]byte `protobuf:"bytes,3,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *PublishRequest_RawData) Reset()                    { *m = PublishRequest_RawData{} }
func (m *PublishRequest_RawData) String() string            { return proto.CompactTextString(m) }
func (*PublishRequest_RawData) ProtoMessage()               {}
func (*PublishRequest_RawData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 1} }

func (m *PublishRequest_RawData) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *PublishRequest_RawData) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *PublishRequest_RawData) GetHeaders() map[string][]byte {
	if m != nil {
		return m.Headers
	}
	return nil
}

type PublishResponse struct {
	Config   *PublishResponse_ConfigMessage   `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	Redirect *PublishResponse_RedirectMessage `protobuf:"bytes,2,opt,name=redirect" json:"redirect,omitempty"`
}

func (m *PublishResponse) Reset()                    { *m = PublishResponse{} }
func (m *PublishResponse) String() string            { return proto.CompactTextString(m) }
func (*PublishResponse) ProtoMessage()               {}
func (*PublishResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PublishResponse) GetConfig() *PublishResponse_ConfigMessage {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *PublishResponse) GetRedirect() *PublishResponse_RedirectMessage {
	if m != nil {
		return m.Redirect
	}
	return nil
}

type PublishResponse_ConfigMessage struct {
	PartitionCount int32 `protobuf:"varint,1,opt,name=partition_count,json=partitionCount" json:"partition_count,omitempty"`
}

func (m *PublishResponse_ConfigMessage) Reset()         { *m = PublishResponse_ConfigMessage{} }
func (m *PublishResponse_ConfigMessage) String() string { return proto.CompactTextString(m) }
func (*PublishResponse_ConfigMessage) ProtoMessage()    {}
func (*PublishResponse_ConfigMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 0}
}

func (m *PublishResponse_ConfigMessage) GetPartitionCount() int32 {
	if m != nil {
		return m.PartitionCount
	}
	return 0
}

type PublishResponse_RedirectMessage struct {
	NewBroker string `protobuf:"bytes,1,opt,name=new_broker,json=newBroker" json:"new_broker,omitempty"`
}

func (m *PublishResponse_RedirectMessage) Reset()         { *m = PublishResponse_RedirectMessage{} }
func (m *PublishResponse_RedirectMessage) String() string { return proto.CompactTextString(m) }
func (*PublishResponse_RedirectMessage) ProtoMessage()    {}
func (*PublishResponse_RedirectMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 1}
}

func (m *PublishResponse_RedirectMessage) GetNewBroker() string {
	if m != nil {
		return m.NewBroker
	}
	return ""
}

type ConfigureTopicRequest struct {
	Namespace     string              `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Topic         string              `protobuf:"bytes,2,opt,name=topic" json:"topic,omitempty"`
	Configuration *TopicConfiguration `protobuf:"bytes,3,opt,name=configuration" json:"configuration,omitempty"`
}

func (m *ConfigureTopicRequest) Reset()                    { *m = ConfigureTopicRequest{} }
func (m *ConfigureTopicRequest) String() string            { return proto.CompactTextString(m) }
func (*ConfigureTopicRequest) ProtoMessage()               {}
func (*ConfigureTopicRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ConfigureTopicRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ConfigureTopicRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *ConfigureTopicRequest) GetConfiguration() *TopicConfiguration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

type ConfigureTopicResponse struct {
}

func (m *ConfigureTopicResponse) Reset()                    { *m = ConfigureTopicResponse{} }
func (m *ConfigureTopicResponse) String() string            { return proto.CompactTextString(m) }
func (*ConfigureTopicResponse) ProtoMessage()               {}
func (*ConfigureTopicResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type GetTopicConfigurationRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Topic     string `protobuf:"bytes,2,opt,name=topic" json:"topic,omitempty"`
}

func (m *GetTopicConfigurationRequest) Reset()                    { *m = GetTopicConfigurationRequest{} }
func (m *GetTopicConfigurationRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTopicConfigurationRequest) ProtoMessage()               {}
func (*GetTopicConfigurationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetTopicConfigurationRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetTopicConfigurationRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

type GetTopicConfigurationResponse struct {
	Configuration *TopicConfiguration `protobuf:"bytes,1,opt,name=configuration" json:"configuration,omitempty"`
}

func (m *GetTopicConfigurationResponse) Reset()                    { *m = GetTopicConfigurationResponse{} }
func (m *GetTopicConfigurationResponse) String() string            { return proto.CompactTextString(m) }
func (*GetTopicConfigurationResponse) ProtoMessage()               {}
func (*GetTopicConfigurationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetTopicConfigurationResponse) GetConfiguration() *TopicConfiguration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

type TopicConfiguration struct {
	PartitionCount int32  `protobuf:"varint,1,opt,name=partition_count,json=partitionCount" json:"partition_count,omitempty"`
	Collection     string `protobuf:"bytes,2,opt,name=collection" json:"collection,omitempty"`
	Replication    string `protobuf:"bytes,3,opt,name=replication" json:"replication,omitempty"`
}

func (m *TopicConfiguration) Reset()                    { *m = TopicConfiguration{} }
func (m *TopicConfiguration) String() string            { return proto.CompactTextString(m) }
func (*TopicConfiguration) ProtoMessage()               {}
func (*TopicConfiguration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *TopicConfiguration) GetPartitionCount() int32 {
	if m != nil {
		return m.PartitionCount
	}
	return 0
}

func (m *TopicConfiguration) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *TopicConfiguration) GetReplication() string {
	if m != nil {
		return m.Replication
	}
	return ""
}

func init() {
	proto.RegisterType((*SubscriberMessage)(nil), "messaging_pb.SubscriberMessage")
	proto.RegisterType((*SubscriberMessage_InitMessage)(nil), "messaging_pb.SubscriberMessage.InitMessage")
	proto.RegisterType((*SubscriberMessage_AckMessage)(nil), "messaging_pb.SubscriberMessage.AckMessage")
	proto.RegisterType((*RawData)(nil), "messaging_pb.RawData")
	proto.RegisterType((*Message)(nil), "messaging_pb.Message")
	proto.RegisterType((*BrokerMessage)(nil), "messaging_pb.BrokerMessage")
	proto.RegisterType((*BrokerMessage_RedirectMessage)(nil), "messaging_pb.BrokerMessage.RedirectMessage")
	proto.RegisterType((*PublishRequest)(nil), "messaging_pb.PublishRequest")
	proto.RegisterType((*PublishRequest_InitMessage)(nil), "messaging_pb.PublishRequest.InitMessage")
	proto.RegisterType((*PublishRequest_RawData)(nil), "messaging_pb.PublishRequest.RawData")
	proto.RegisterType((*PublishResponse)(nil), "messaging_pb.PublishResponse")
	proto.RegisterType((*PublishResponse_ConfigMessage)(nil), "messaging_pb.PublishResponse.ConfigMessage")
	proto.RegisterType((*PublishResponse_RedirectMessage)(nil), "messaging_pb.PublishResponse.RedirectMessage")
	proto.RegisterType((*ConfigureTopicRequest)(nil), "messaging_pb.ConfigureTopicRequest")
	proto.RegisterType((*ConfigureTopicResponse)(nil), "messaging_pb.ConfigureTopicResponse")
	proto.RegisterType((*GetTopicConfigurationRequest)(nil), "messaging_pb.GetTopicConfigurationRequest")
	proto.RegisterType((*GetTopicConfigurationResponse)(nil), "messaging_pb.GetTopicConfigurationResponse")
	proto.RegisterType((*TopicConfiguration)(nil), "messaging_pb.TopicConfiguration")
	proto.RegisterEnum("messaging_pb.SubscriberMessage_InitMessage_StartPosition", SubscriberMessage_InitMessage_StartPosition_name, SubscriberMessage_InitMessage_StartPosition_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SeaweedMessaging service

type SeaweedMessagingClient interface {
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (SeaweedMessaging_SubscribeClient, error)
	Publish(ctx context.Context, opts ...grpc.CallOption) (SeaweedMessaging_PublishClient, error)
	ConfigureTopic(ctx context.Context, in *ConfigureTopicRequest, opts ...grpc.CallOption) (*ConfigureTopicResponse, error)
	GetTopicConfiguration(ctx context.Context, in *GetTopicConfigurationRequest, opts ...grpc.CallOption) (*GetTopicConfigurationResponse, error)
}

type seaweedMessagingClient struct {
	cc *grpc.ClientConn
}

func NewSeaweedMessagingClient(cc *grpc.ClientConn) SeaweedMessagingClient {
	return &seaweedMessagingClient{cc}
}

func (c *seaweedMessagingClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (SeaweedMessaging_SubscribeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SeaweedMessaging_serviceDesc.Streams[0], c.cc, "/messaging_pb.SeaweedMessaging/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &seaweedMessagingSubscribeClient{stream}
	return x, nil
}

type SeaweedMessaging_SubscribeClient interface {
	Send(*SubscriberMessage) error
	Recv() (*BrokerMessage, error)
	grpc.ClientStream
}

type seaweedMessagingSubscribeClient struct {
	grpc.ClientStream
}

func (x *seaweedMessagingSubscribeClient) Send(m *SubscriberMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *seaweedMessagingSubscribeClient) Recv() (*BrokerMessage, error) {
	m := new(BrokerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *seaweedMessagingClient) Publish(ctx context.Context, opts ...grpc.CallOption) (SeaweedMessaging_PublishClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SeaweedMessaging_serviceDesc.Streams[1], c.cc, "/messaging_pb.SeaweedMessaging/Publish", opts...)
	if err != nil {
		return nil, err
	}
	x := &seaweedMessagingPublishClient{stream}
	return x, nil
}

type SeaweedMessaging_PublishClient interface {
	Send(*PublishRequest) error
	Recv() (*PublishResponse, error)
	grpc.ClientStream
}

type seaweedMessagingPublishClient struct {
	grpc.ClientStream
}

func (x *seaweedMessagingPublishClient) Send(m *PublishRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *seaweedMessagingPublishClient) Recv() (*PublishResponse, error) {
	m := new(PublishResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *seaweedMessagingClient) ConfigureTopic(ctx context.Context, in *ConfigureTopicRequest, opts ...grpc.CallOption) (*ConfigureTopicResponse, error) {
	out := new(ConfigureTopicResponse)
	err := grpc.Invoke(ctx, "/messaging_pb.SeaweedMessaging/ConfigureTopic", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedMessagingClient) GetTopicConfiguration(ctx context.Context, in *GetTopicConfigurationRequest, opts ...grpc.CallOption) (*GetTopicConfigurationResponse, error) {
	out := new(GetTopicConfigurationResponse)
	err := grpc.Invoke(ctx, "/messaging_pb.SeaweedMessaging/GetTopicConfiguration", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SeaweedMessaging service

type SeaweedMessagingServer interface {
	Subscribe(SeaweedMessaging_SubscribeServer) error
	Publish(SeaweedMessaging_PublishServer) error
	ConfigureTopic(context.Context, *ConfigureTopicRequest) (*ConfigureTopicResponse, error)
	GetTopicConfiguration(context.Context, *GetTopicConfigurationRequest) (*GetTopicConfigurationResponse, error)
}

func RegisterSeaweedMessagingServer(s *grpc.Server, srv SeaweedMessagingServer) {
	s.RegisterService(&_SeaweedMessaging_serviceDesc, srv)
}

func _SeaweedMessaging_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SeaweedMessagingServer).Subscribe(&seaweedMessagingSubscribeServer{stream})
}

type SeaweedMessaging_SubscribeServer interface {
	Send(*BrokerMessage) error
	Recv() (*SubscriberMessage, error)
	grpc.ServerStream
}

type seaweedMessagingSubscribeServer struct {
	grpc.ServerStream
}

func (x *seaweedMessagingSubscribeServer) Send(m *BrokerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *seaweedMessagingSubscribeServer) Recv() (*SubscriberMessage, error) {
	m := new(SubscriberMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SeaweedMessaging_Publish_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SeaweedMessagingServer).Publish(&seaweedMessagingPublishServer{stream})
}

type SeaweedMessaging_PublishServer interface {
	Send(*PublishResponse) error
	Recv() (*PublishRequest, error)
	grpc.ServerStream
}

type seaweedMessagingPublishServer struct {
	grpc.ServerStream
}

func (x *seaweedMessagingPublishServer) Send(m *PublishResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *seaweedMessagingPublishServer) Recv() (*PublishRequest, error) {
	m := new(PublishRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SeaweedMessaging_ConfigureTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedMessagingServer).ConfigureTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging_pb.SeaweedMessaging/ConfigureTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedMessagingServer).ConfigureTopic(ctx, req.(*ConfigureTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedMessaging_GetTopicConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedMessagingServer).GetTopicConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging_pb.SeaweedMessaging/GetTopicConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedMessagingServer).GetTopicConfiguration(ctx, req.(*GetTopicConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SeaweedMessaging_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messaging_pb.SeaweedMessaging",
	HandlerType: (*SeaweedMessagingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConfigureTopic",
			Handler:    _SeaweedMessaging_ConfigureTopic_Handler,
		},
		{
			MethodName: "GetTopicConfiguration",
			Handler:    _SeaweedMessaging_GetTopicConfiguration_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _SeaweedMessaging_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Publish",
			Handler:       _SeaweedMessaging_Publish_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "messaging.proto",
}

func init() { proto.RegisterFile("messaging.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 847 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x56, 0xcd, 0x8e, 0xe3, 0x44,
	0x10, 0xde, 0x8e, 0x33, 0x93, 0x75, 0xe5, 0x97, 0x16, 0x83, 0x22, 0x33, 0x03, 0x96, 0x77, 0x25,
	0x02, 0x23, 0xac, 0x21, 0x5c, 0xc2, 0x6a, 0x25, 0x94, 0x84, 0xb0, 0x44, 0x4c, 0x20, 0xea, 0xe4,
	0x8a, 0xa2, 0x8e, 0xd3, 0x9b, 0xb5, 0x92, 0xd8, 0xc6, 0xdd, 0x21, 0xda, 0x13, 0x07, 0xb8, 0x72,
	0xe2, 0x1d, 0xb8, 0x70, 0xe6, 0xc2, 0x8d, 0x07, 0xe0, 0x9d, 0x90, 0x7f, 0x63, 0x27, 0x1e, 0xef,
	0x10, 0xed, 0xdc, 0xec, 0x72, 0xd5, 0x57, 0xf5, 0x55, 0x7d, 0xd5, 0x6e, 0xa8, 0x6f, 0x18, 0xe7,
	0x74, 0x69, 0x5a, 0x4b, 0xdd, 0x71, 0x6d, 0x61, 0xe3, 0x4a, 0x6c, 0x98, 0x39, 0x73, 0xed, 0x97,
	0x22, 0xbc, 0x33, 0xd9, 0xce, 0xb9, 0xe1, 0x9a, 0x73, 0xe6, 0x8e, 0xfc, 0x4f, 0x0c, 0x7f, 0x09,
	0x45, 0xd3, 0x32, 0x45, 0x13, 0xa9, 0xa8, 0x55, 0x6e, 0x5f, 0xeb, 0xc9, 0x10, 0xfd, 0xc8, 0x5d,
	0x1f, 0x5a, 0xa6, 0x08, 0x9f, 0x89, 0x1f, 0x88, 0x9f, 0x83, 0x44, 0x8d, 0x55, 0xb3, 0xe0, 0xc7,
	0x7f, 0xf2, 0xa6, 0xf8, 0xae, 0xb1, 0x8a, 0xc2, 0xbd, 0x30, 0xe5, 0x9f, 0x02, 0x94, 0x13, 0x98,
	0xf8, 0x12, 0x64, 0x8b, 0x6e, 0x18, 0x77, 0xa8, 0xc1, 0xfc, 0x9a, 0x64, 0xb2, 0x37, 0xe0, 0x77,
	0xe1, 0x4c, 0xd8, 0x8e, 0x69, 0xf8, 0xd9, 0x64, 0x12, 0xbc, 0x78, 0x31, 0x0e, 0x75, 0x85, 0x29,
	0x4c, 0xdb, 0x6a, 0x4a, 0x2a, 0x6a, 0x9d, 0x91, 0xbd, 0x01, 0xcf, 0xa0, 0xca, 0x05, 0x75, 0xc5,
	0xd8, 0xe6, 0x81, 0x47, 0x51, 0x45, 0xad, 0x5a, 0xfb, 0x8b, 0xff, 0xc1, 0x54, 0x9f, 0x24, 0x01,
	0x48, 0x1a, 0x0f, 0xab, 0x50, 0x16, 0xe6, 0x86, 0x71, 0x41, 0x37, 0xce, 0x77, 0xbc, 0x79, 0xa6,
	0xa2, 0x96, 0x44, 0x92, 0x26, 0xfc, 0x04, 0xaa, 0x3c, 0xc6, 0x9f, 0x99, 0x8b, 0xe6, 0xb9, 0x5f,
	0x7e, 0x65, 0x6f, 0x1c, 0x2e, 0xb4, 0x0e, 0x54, 0x53, 0x69, 0x30, 0xc0, 0xf9, 0x6d, 0x77, 0x3a,
	0x98, 0x4c, 0x1b, 0x8f, 0x70, 0x05, 0x1e, 0x0f, 0xba, 0xe4, 0x76, 0xe8, 0xbd, 0x21, 0x5c, 0x05,
	0x79, 0x3a, 0x1c, 0x0d, 0x26, 0xd3, 0xee, 0x68, 0xdc, 0x28, 0x28, 0xd7, 0x00, 0xfb, 0xb6, 0xe2,
	0x2b, 0x80, 0x80, 0x19, 0xf3, 0x32, 0x21, 0xbf, 0x1a, 0x39, 0xb4, 0x0c, 0x17, 0xda, 0x9f, 0x08,
	0x4a, 0x84, 0xee, 0xbe, 0xa2, 0x82, 0xe2, 0x06, 0x48, 0x2b, 0xf6, 0xda, 0xf7, 0xa9, 0x10, 0xef,
	0xd1, 0x6b, 0xf0, 0x4f, 0x74, 0xbd, 0x65, 0x7e, 0x83, 0x2b, 0x24, 0x78, 0xc1, 0xcf, 0xa1, 0xf4,
	0x8a, 0xd1, 0x05, 0x73, 0x79, 0x53, 0x52, 0xa5, 0x56, 0xb9, 0xad, 0xa5, 0x9b, 0x17, 0xe2, 0xe9,
	0xdf, 0x04, 0x4e, 0x03, 0x4b, 0xb8, 0xaf, 0x49, 0x14, 0xa2, 0x3c, 0x83, 0x4a, 0xf2, 0x43, 0x32,
	0xab, 0x9c, 0x93, 0xf5, 0x59, 0xa1, 0x83, 0xb4, 0x7f, 0x11, 0x94, 0x22, 0x62, 0x2a, 0xc8, 0x71,
	0x53, 0x03, 0x5e, 0xbd, 0xc2, 0x0d, 0x22, 0x7b, 0x63, 0x84, 0x5c, 0xc8, 0xe0, 0x23, 0xdd, 0xc1,
	0xa7, 0x98, 0xc5, 0x27, 0x1a, 0xfb, 0xdb, 0xe7, 0xf3, 0x17, 0x82, 0x6a, 0xcf, 0xb5, 0x57, 0xfb,
	0xfd, 0xfb, 0x18, 0x8a, 0x0b, 0x2a, 0x68, 0xb8, 0x7f, 0x17, 0x99, 0x85, 0x10, 0xdf, 0x05, 0xbf,
	0x80, 0xc7, 0x2e, 0x5b, 0x98, 0x2e, 0x33, 0x44, 0xb8, 0x6e, 0x07, 0xeb, 0x9a, 0x42, 0xd6, 0x49,
	0xe8, 0x1b, 0x81, 0xc4, 0xc1, 0xca, 0x0d, 0xd4, 0x0f, 0x3e, 0x7a, 0xaa, 0xb1, 0xd8, 0x6e, 0x36,
	0xf7, 0x11, 0xe2, 0xc5, 0x63, 0xbb, 0x00, 0x52, 0xfb, 0x43, 0x82, 0xda, 0x78, 0x3b, 0x5f, 0x9b,
	0xfc, 0x15, 0x61, 0x3f, 0x6e, 0x19, 0xf7, 0xf6, 0x3e, 0x79, 0x70, 0xb4, 0xd2, 0x95, 0xa4, 0x7d,
	0x33, 0x4e, 0x8d, 0x4e, 0x48, 0x3b, 0xe0, 0xf1, 0x34, 0x37, 0x3a, 0x94, 0x57, 0xd0, 0x05, 0x65,
	0xf6, 0xc0, 0x07, 0x86, 0xf2, 0xf7, 0x09, 0x1b, 0xf2, 0xed, 0xe1, 0x86, 0x7c, 0x76, 0x1f, 0x46,
	0x0f, 0x20, 0xb0, 0xdf, 0x0a, 0x50, 0x8f, 0x93, 0x71, 0xc7, 0xb6, 0x38, 0xc3, 0x7d, 0x38, 0x37,
	0x6c, 0xeb, 0xa5, 0xb9, 0xcc, 0x3e, 0xe4, 0x0f, 0xdc, 0xf5, 0xbe, 0xef, 0x1b, 0x8d, 0x2b, 0x0c,
	0xc5, 0xc3, 0x23, 0xf1, 0x7d, 0x9a, 0x0f, 0x73, 0xb7, 0xfc, 0x3a, 0x50, 0x4d, 0xe5, 0xc0, 0x1f,
	0x41, 0x3d, 0x6e, 0xff, 0xcc, 0xb0, 0xb7, 0x56, 0xa0, 0xaa, 0x33, 0x52, 0x8b, 0xcd, 0x7d, 0xcf,
	0x7a, 0x82, 0x70, 0x7f, 0x47, 0x70, 0x11, 0x24, 0xdb, 0xba, 0x6c, 0xea, 0x4d, 0x3f, 0xd2, 0xef,
	0x29, 0xc2, 0xf9, 0x1a, 0xaa, 0x46, 0x08, 0x46, 0x63, 0xf1, 0x94, 0xdb, 0x6a, 0xba, 0x13, 0x7e,
	0x9a, 0x7e, 0xd2, 0x8f, 0xa4, 0xc3, 0xb4, 0x26, 0xbc, 0x77, 0x58, 0x54, 0xd0, 0x35, 0x8d, 0xc0,
	0xe5, 0x0b, 0x26, 0x32, 0x10, 0x4e, 0xaf, 0x5a, 0x5b, 0xc2, 0xd5, 0x1d, 0x98, 0xa1, 0x40, 0x8e,
	0x68, 0xa1, 0xd3, 0x68, 0xfd, 0x0c, 0xf8, 0xd8, 0xe9, 0xde, 0xd3, 0xc5, 0x1f, 0x00, 0x18, 0xf6,
	0x7a, 0xcd, 0x0c, 0xbf, 0x86, 0x80, 0x42, 0xc2, 0xe2, 0xfd, 0x68, 0x5d, 0xe6, 0xac, 0x4d, 0x63,
	0xdf, 0x7b, 0x99, 0x24, 0x4d, 0xed, 0x5f, 0x25, 0x68, 0x4c, 0x18, 0xdd, 0x31, 0xb6, 0x18, 0x45,
	0xa5, 0xe3, 0xef, 0x41, 0x8e, 0xff, 0xee, 0xf8, 0xc3, 0x37, 0xfc, 0xf6, 0x95, 0xf7, 0x73, 0x8e,
	0x54, 0xed, 0x51, 0x0b, 0xdd, 0x20, 0x7c, 0x0b, 0xa5, 0x50, 0xec, 0xf8, 0x32, 0x6f, 0xcd, 0x95,
	0xab, 0xdc, 0x0d, 0x09, 0xd1, 0x7e, 0x80, 0x5a, 0x5a, 0x0b, 0xf8, 0x49, 0x3a, 0x2c, 0x53, 0xbe,
	0xca, 0xd3, 0x7c, 0xa7, 0x28, 0x05, 0x76, 0xe1, 0x22, 0x73, 0xf8, 0xf8, 0xe0, 0xaa, 0x96, 0xa7,
	0x3a, 0xe5, 0xfa, 0x5e, 0xbe, 0x51, 0xce, 0x9e, 0x06, 0x0d, 0x1e, 0x4c, 0xe1, 0x25, 0xd7, 0x8d,
	0xb5, 0xc9, 0x2c, 0xd1, 0xab, 0xc5, 0x03, 0x19, 0x7b, 0x77, 0xd3, 0xf9, 0xb9, 0x7f, 0x45, 0xfd,
	0xfc, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdf, 0xc7, 0xa8, 0xba, 0xb5, 0x0a, 0x00, 0x00,
}
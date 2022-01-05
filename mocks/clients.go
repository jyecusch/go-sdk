// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nitrictech/apis/go/nitric/v1 (interfaces: DocumentServiceClient,EventServiceClient,TopicServiceClient,QueueServiceClient,StorageServiceClient,FaasServiceClient,FaasService_TriggerStreamClient,DocumentService_QueryStreamClient,SecretServiceClient)

// Package mock_v1 is a generated GoMock package.
package mock_v1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/nitrictech/apis/go/nitric/v1"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockDocumentServiceClient is a mock of DocumentServiceClient interface.
type MockDocumentServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockDocumentServiceClientMockRecorder
}

// MockDocumentServiceClientMockRecorder is the mock recorder for MockDocumentServiceClient.
type MockDocumentServiceClientMockRecorder struct {
	mock *MockDocumentServiceClient
}

// NewMockDocumentServiceClient creates a new mock instance.
func NewMockDocumentServiceClient(ctrl *gomock.Controller) *MockDocumentServiceClient {
	mock := &MockDocumentServiceClient{ctrl: ctrl}
	mock.recorder = &MockDocumentServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDocumentServiceClient) EXPECT() *MockDocumentServiceClientMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockDocumentServiceClient) Delete(arg0 context.Context, arg1 *v1.DocumentDeleteRequest, arg2 ...grpc.CallOption) (*v1.DocumentDeleteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*v1.DocumentDeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockDocumentServiceClientMockRecorder) Delete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDocumentServiceClient)(nil).Delete), varargs...)
}

// Get mocks base method.
func (m *MockDocumentServiceClient) Get(arg0 context.Context, arg1 *v1.DocumentGetRequest, arg2 ...grpc.CallOption) (*v1.DocumentGetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*v1.DocumentGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDocumentServiceClientMockRecorder) Get(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDocumentServiceClient)(nil).Get), varargs...)
}

// Query mocks base method.
func (m *MockDocumentServiceClient) Query(arg0 context.Context, arg1 *v1.DocumentQueryRequest, arg2 ...grpc.CallOption) (*v1.DocumentQueryResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(*v1.DocumentQueryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockDocumentServiceClientMockRecorder) Query(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockDocumentServiceClient)(nil).Query), varargs...)
}

// QueryStream mocks base method.
func (m *MockDocumentServiceClient) QueryStream(arg0 context.Context, arg1 *v1.DocumentQueryStreamRequest, arg2 ...grpc.CallOption) (v1.DocumentService_QueryStreamClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryStream", varargs...)
	ret0, _ := ret[0].(v1.DocumentService_QueryStreamClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryStream indicates an expected call of QueryStream.
func (mr *MockDocumentServiceClientMockRecorder) QueryStream(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryStream", reflect.TypeOf((*MockDocumentServiceClient)(nil).QueryStream), varargs...)
}

// Set mocks base method.
func (m *MockDocumentServiceClient) Set(arg0 context.Context, arg1 *v1.DocumentSetRequest, arg2 ...grpc.CallOption) (*v1.DocumentSetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Set", varargs...)
	ret0, _ := ret[0].(*v1.DocumentSetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Set indicates an expected call of Set.
func (mr *MockDocumentServiceClientMockRecorder) Set(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockDocumentServiceClient)(nil).Set), varargs...)
}

// MockEventServiceClient is a mock of EventServiceClient interface.
type MockEventServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockEventServiceClientMockRecorder
}

// MockEventServiceClientMockRecorder is the mock recorder for MockEventServiceClient.
type MockEventServiceClientMockRecorder struct {
	mock *MockEventServiceClient
}

// NewMockEventServiceClient creates a new mock instance.
func NewMockEventServiceClient(ctrl *gomock.Controller) *MockEventServiceClient {
	mock := &MockEventServiceClient{ctrl: ctrl}
	mock.recorder = &MockEventServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventServiceClient) EXPECT() *MockEventServiceClientMockRecorder {
	return m.recorder
}

// Publish mocks base method.
func (m *MockEventServiceClient) Publish(arg0 context.Context, arg1 *v1.EventPublishRequest, arg2 ...grpc.CallOption) (*v1.EventPublishResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Publish", varargs...)
	ret0, _ := ret[0].(*v1.EventPublishResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Publish indicates an expected call of Publish.
func (mr *MockEventServiceClientMockRecorder) Publish(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockEventServiceClient)(nil).Publish), varargs...)
}

// MockTopicServiceClient is a mock of TopicServiceClient interface.
type MockTopicServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockTopicServiceClientMockRecorder
}

// MockTopicServiceClientMockRecorder is the mock recorder for MockTopicServiceClient.
type MockTopicServiceClientMockRecorder struct {
	mock *MockTopicServiceClient
}

// NewMockTopicServiceClient creates a new mock instance.
func NewMockTopicServiceClient(ctrl *gomock.Controller) *MockTopicServiceClient {
	mock := &MockTopicServiceClient{ctrl: ctrl}
	mock.recorder = &MockTopicServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTopicServiceClient) EXPECT() *MockTopicServiceClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockTopicServiceClient) List(arg0 context.Context, arg1 *v1.TopicListRequest, arg2 ...grpc.CallOption) (*v1.TopicListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*v1.TopicListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockTopicServiceClientMockRecorder) List(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTopicServiceClient)(nil).List), varargs...)
}

// MockQueueServiceClient is a mock of QueueServiceClient interface.
type MockQueueServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockQueueServiceClientMockRecorder
}

// MockQueueServiceClientMockRecorder is the mock recorder for MockQueueServiceClient.
type MockQueueServiceClientMockRecorder struct {
	mock *MockQueueServiceClient
}

// NewMockQueueServiceClient creates a new mock instance.
func NewMockQueueServiceClient(ctrl *gomock.Controller) *MockQueueServiceClient {
	mock := &MockQueueServiceClient{ctrl: ctrl}
	mock.recorder = &MockQueueServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueueServiceClient) EXPECT() *MockQueueServiceClientMockRecorder {
	return m.recorder
}

// Complete mocks base method.
func (m *MockQueueServiceClient) Complete(arg0 context.Context, arg1 *v1.QueueCompleteRequest, arg2 ...grpc.CallOption) (*v1.QueueCompleteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Complete", varargs...)
	ret0, _ := ret[0].(*v1.QueueCompleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Complete indicates an expected call of Complete.
func (mr *MockQueueServiceClientMockRecorder) Complete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Complete", reflect.TypeOf((*MockQueueServiceClient)(nil).Complete), varargs...)
}

// Receive mocks base method.
func (m *MockQueueServiceClient) Receive(arg0 context.Context, arg1 *v1.QueueReceiveRequest, arg2 ...grpc.CallOption) (*v1.QueueReceiveResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Receive", varargs...)
	ret0, _ := ret[0].(*v1.QueueReceiveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Receive indicates an expected call of Receive.
func (mr *MockQueueServiceClientMockRecorder) Receive(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*MockQueueServiceClient)(nil).Receive), varargs...)
}

// Send mocks base method.
func (m *MockQueueServiceClient) Send(arg0 context.Context, arg1 *v1.QueueSendRequest, arg2 ...grpc.CallOption) (*v1.QueueSendResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Send", varargs...)
	ret0, _ := ret[0].(*v1.QueueSendResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send.
func (mr *MockQueueServiceClientMockRecorder) Send(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockQueueServiceClient)(nil).Send), varargs...)
}

// SendBatch mocks base method.
func (m *MockQueueServiceClient) SendBatch(arg0 context.Context, arg1 *v1.QueueSendBatchRequest, arg2 ...grpc.CallOption) (*v1.QueueSendBatchResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendBatch", varargs...)
	ret0, _ := ret[0].(*v1.QueueSendBatchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendBatch indicates an expected call of SendBatch.
func (mr *MockQueueServiceClientMockRecorder) SendBatch(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendBatch", reflect.TypeOf((*MockQueueServiceClient)(nil).SendBatch), varargs...)
}

// MockStorageServiceClient is a mock of StorageServiceClient interface.
type MockStorageServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockStorageServiceClientMockRecorder
}

// MockStorageServiceClientMockRecorder is the mock recorder for MockStorageServiceClient.
type MockStorageServiceClientMockRecorder struct {
	mock *MockStorageServiceClient
}

// NewMockStorageServiceClient creates a new mock instance.
func NewMockStorageServiceClient(ctrl *gomock.Controller) *MockStorageServiceClient {
	mock := &MockStorageServiceClient{ctrl: ctrl}
	mock.recorder = &MockStorageServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageServiceClient) EXPECT() *MockStorageServiceClientMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockStorageServiceClient) Delete(arg0 context.Context, arg1 *v1.StorageDeleteRequest, arg2 ...grpc.CallOption) (*v1.StorageDeleteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*v1.StorageDeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockStorageServiceClientMockRecorder) Delete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStorageServiceClient)(nil).Delete), varargs...)
}

// PreSignUrl mocks base method.
func (m *MockStorageServiceClient) PreSignUrl(arg0 context.Context, arg1 *v1.StoragePreSignUrlRequest, arg2 ...grpc.CallOption) (*v1.StoragePreSignUrlResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PreSignUrl", varargs...)
	ret0, _ := ret[0].(*v1.StoragePreSignUrlResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PreSignUrl indicates an expected call of PreSignUrl.
func (mr *MockStorageServiceClientMockRecorder) PreSignUrl(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreSignUrl", reflect.TypeOf((*MockStorageServiceClient)(nil).PreSignUrl), varargs...)
}

// Read mocks base method.
func (m *MockStorageServiceClient) Read(arg0 context.Context, arg1 *v1.StorageReadRequest, arg2 ...grpc.CallOption) (*v1.StorageReadResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Read", varargs...)
	ret0, _ := ret[0].(*v1.StorageReadResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockStorageServiceClientMockRecorder) Read(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockStorageServiceClient)(nil).Read), varargs...)
}

// Write mocks base method.
func (m *MockStorageServiceClient) Write(arg0 context.Context, arg1 *v1.StorageWriteRequest, arg2 ...grpc.CallOption) (*v1.StorageWriteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Write", varargs...)
	ret0, _ := ret[0].(*v1.StorageWriteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockStorageServiceClientMockRecorder) Write(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockStorageServiceClient)(nil).Write), varargs...)
}

// MockFaasServiceClient is a mock of FaasServiceClient interface.
type MockFaasServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockFaasServiceClientMockRecorder
}

// MockFaasServiceClientMockRecorder is the mock recorder for MockFaasServiceClient.
type MockFaasServiceClientMockRecorder struct {
	mock *MockFaasServiceClient
}

// NewMockFaasServiceClient creates a new mock instance.
func NewMockFaasServiceClient(ctrl *gomock.Controller) *MockFaasServiceClient {
	mock := &MockFaasServiceClient{ctrl: ctrl}
	mock.recorder = &MockFaasServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFaasServiceClient) EXPECT() *MockFaasServiceClientMockRecorder {
	return m.recorder
}

// TriggerStream mocks base method.
func (m *MockFaasServiceClient) TriggerStream(arg0 context.Context, arg1 ...grpc.CallOption) (v1.FaasService_TriggerStreamClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TriggerStream", varargs...)
	ret0, _ := ret[0].(v1.FaasService_TriggerStreamClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TriggerStream indicates an expected call of TriggerStream.
func (mr *MockFaasServiceClientMockRecorder) TriggerStream(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerStream", reflect.TypeOf((*MockFaasServiceClient)(nil).TriggerStream), varargs...)
}

// MockFaasService_TriggerStreamClient is a mock of FaasService_TriggerStreamClient interface.
type MockFaasService_TriggerStreamClient struct {
	ctrl     *gomock.Controller
	recorder *MockFaasService_TriggerStreamClientMockRecorder
}

// MockFaasService_TriggerStreamClientMockRecorder is the mock recorder for MockFaasService_TriggerStreamClient.
type MockFaasService_TriggerStreamClientMockRecorder struct {
	mock *MockFaasService_TriggerStreamClient
}

// NewMockFaasService_TriggerStreamClient creates a new mock instance.
func NewMockFaasService_TriggerStreamClient(ctrl *gomock.Controller) *MockFaasService_TriggerStreamClient {
	mock := &MockFaasService_TriggerStreamClient{ctrl: ctrl}
	mock.recorder = &MockFaasService_TriggerStreamClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFaasService_TriggerStreamClient) EXPECT() *MockFaasService_TriggerStreamClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockFaasService_TriggerStreamClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockFaasService_TriggerStreamClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockFaasService_TriggerStreamClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockFaasService_TriggerStreamClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockFaasService_TriggerStreamClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockFaasService_TriggerStreamClient)(nil).Context))
}

// Header mocks base method.
func (m *MockFaasService_TriggerStreamClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockFaasService_TriggerStreamClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockFaasService_TriggerStreamClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockFaasService_TriggerStreamClient) Recv() (*v1.ServerMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*v1.ServerMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockFaasService_TriggerStreamClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockFaasService_TriggerStreamClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockFaasService_TriggerStreamClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockFaasService_TriggerStreamClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockFaasService_TriggerStreamClient)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockFaasService_TriggerStreamClient) Send(arg0 *v1.ClientMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockFaasService_TriggerStreamClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockFaasService_TriggerStreamClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m *MockFaasService_TriggerStreamClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockFaasService_TriggerStreamClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockFaasService_TriggerStreamClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockFaasService_TriggerStreamClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockFaasService_TriggerStreamClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockFaasService_TriggerStreamClient)(nil).Trailer))
}

// MockDocumentService_QueryStreamClient is a mock of DocumentService_QueryStreamClient interface.
type MockDocumentService_QueryStreamClient struct {
	ctrl     *gomock.Controller
	recorder *MockDocumentService_QueryStreamClientMockRecorder
}

// MockDocumentService_QueryStreamClientMockRecorder is the mock recorder for MockDocumentService_QueryStreamClient.
type MockDocumentService_QueryStreamClientMockRecorder struct {
	mock *MockDocumentService_QueryStreamClient
}

// NewMockDocumentService_QueryStreamClient creates a new mock instance.
func NewMockDocumentService_QueryStreamClient(ctrl *gomock.Controller) *MockDocumentService_QueryStreamClient {
	mock := &MockDocumentService_QueryStreamClient{ctrl: ctrl}
	mock.recorder = &MockDocumentService_QueryStreamClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDocumentService_QueryStreamClient) EXPECT() *MockDocumentService_QueryStreamClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockDocumentService_QueryStreamClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockDocumentService_QueryStreamClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockDocumentService_QueryStreamClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockDocumentService_QueryStreamClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockDocumentService_QueryStreamClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockDocumentService_QueryStreamClient)(nil).Context))
}

// Header mocks base method.
func (m *MockDocumentService_QueryStreamClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockDocumentService_QueryStreamClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockDocumentService_QueryStreamClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockDocumentService_QueryStreamClient) Recv() (*v1.DocumentQueryStreamResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*v1.DocumentQueryStreamResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockDocumentService_QueryStreamClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockDocumentService_QueryStreamClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockDocumentService_QueryStreamClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockDocumentService_QueryStreamClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockDocumentService_QueryStreamClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method.
func (m *MockDocumentService_QueryStreamClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockDocumentService_QueryStreamClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockDocumentService_QueryStreamClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockDocumentService_QueryStreamClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockDocumentService_QueryStreamClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockDocumentService_QueryStreamClient)(nil).Trailer))
}

// MockSecretServiceClient is a mock of SecretServiceClient interface.
type MockSecretServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecretServiceClientMockRecorder
}

// MockSecretServiceClientMockRecorder is the mock recorder for MockSecretServiceClient.
type MockSecretServiceClientMockRecorder struct {
	mock *MockSecretServiceClient
}

// NewMockSecretServiceClient creates a new mock instance.
func NewMockSecretServiceClient(ctrl *gomock.Controller) *MockSecretServiceClient {
	mock := &MockSecretServiceClient{ctrl: ctrl}
	mock.recorder = &MockSecretServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretServiceClient) EXPECT() *MockSecretServiceClientMockRecorder {
	return m.recorder
}

// Access mocks base method.
func (m *MockSecretServiceClient) Access(arg0 context.Context, arg1 *v1.SecretAccessRequest, arg2 ...grpc.CallOption) (*v1.SecretAccessResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Access", varargs...)
	ret0, _ := ret[0].(*v1.SecretAccessResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Access indicates an expected call of Access.
func (mr *MockSecretServiceClientMockRecorder) Access(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Access", reflect.TypeOf((*MockSecretServiceClient)(nil).Access), varargs...)
}

// Put mocks base method.
func (m *MockSecretServiceClient) Put(arg0 context.Context, arg1 *v1.SecretPutRequest, arg2 ...grpc.CallOption) (*v1.SecretPutResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Put", varargs...)
	ret0, _ := ret[0].(*v1.SecretPutResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockSecretServiceClientMockRecorder) Put(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockSecretServiceClient)(nil).Put), varargs...)
}

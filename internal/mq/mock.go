// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_mq is a generated GoMock package.
package mq

import (
	gomock "github.com/golang/mock/gomock"
	amqp "github.com/streadway/amqp"
	reflect "reflect"
)

// MockIMessage is a mock of IMessage interface
type MockIMessage struct {
	ctrl     *gomock.Controller
	recorder *MockIMessageMockRecorder
}

// MockIMessageMockRecorder is the mock recorder for MockIMessage
type MockIMessageMockRecorder struct {
	mock *MockIMessage
}

// NewMockIMessage creates a new mock instance
func NewMockIMessage(ctrl *gomock.Controller) *MockIMessage {
	mock := &MockIMessage{ctrl: ctrl}
	mock.recorder = &MockIMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIMessage) EXPECT() *MockIMessageMockRecorder {
	return m.recorder
}

// GetQueues mocks base method
func (m *MockIMessage) GetQueues() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueues")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetQueues indicates an expected call of GetQueues
func (mr *MockIMessageMockRecorder) GetQueues() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueues", reflect.TypeOf((*MockIMessage)(nil).GetQueues))
}

// MarshalToQueue mocks base method
func (m *MockIMessage) MarshalToQueue() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalToQueue")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalToQueue indicates an expected call of MarshalToQueue
func (mr *MockIMessageMockRecorder) MarshalToQueue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalToQueue", reflect.TypeOf((*MockIMessage)(nil).MarshalToQueue))
}

// MockPublisher is a mock of Publisher interface
type MockPublisher struct {
	ctrl     *gomock.Controller
	recorder *MockPublisherMockRecorder
}

// MockPublisherMockRecorder is the mock recorder for MockPublisher
type MockPublisherMockRecorder struct {
	mock *MockPublisher
}

// NewMockPublisher creates a new mock instance
func NewMockPublisher(ctrl *gomock.Controller) *MockPublisher {
	mock := &MockPublisher{ctrl: ctrl}
	mock.recorder = &MockPublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPublisher) EXPECT() *MockPublisherMockRecorder {
	return m.recorder
}

// SendRaw mocks base method
func (m *MockPublisher) SendRaw(queue string, body []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendRaw", queue, body)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendRaw indicates an expected call of SendRaw
func (mr *MockPublisherMockRecorder) SendRaw(queue, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRaw", reflect.TypeOf((*MockPublisher)(nil).SendRaw), queue, body)
}

// Send mocks base method
func (m *MockPublisher) Send(queue IMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", queue)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockPublisherMockRecorder) Send(queue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockPublisher)(nil).Send), queue)
}

// MockIMessagePublisher is a mock of IMessagePublisher interface
type MockIMessagePublisher struct {
	ctrl     *gomock.Controller
	recorder *MockIMessagePublisherMockRecorder
}

// MockIMessagePublisherMockRecorder is the mock recorder for MockIMessagePublisher
type MockIMessagePublisherMockRecorder struct {
	mock *MockIMessagePublisher
}

// NewMockIMessagePublisher creates a new mock instance
func NewMockIMessagePublisher(ctrl *gomock.Controller) *MockIMessagePublisher {
	mock := &MockIMessagePublisher{ctrl: ctrl}
	mock.recorder = &MockIMessagePublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIMessagePublisher) EXPECT() *MockIMessagePublisherMockRecorder {
	return m.recorder
}

// SendRaw mocks base method
func (m *MockIMessagePublisher) SendRaw(queue string, body []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendRaw", queue, body)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendRaw indicates an expected call of SendRaw
func (mr *MockIMessagePublisherMockRecorder) SendRaw(queue, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRaw", reflect.TypeOf((*MockIMessagePublisher)(nil).SendRaw), queue, body)
}

// Send mocks base method
func (m *MockIMessagePublisher) Send(queue IMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", queue)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockIMessagePublisherMockRecorder) Send(queue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockIMessagePublisher)(nil).Send), queue)
}

// Close mocks base method
func (m *MockIMessagePublisher) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockIMessagePublisherMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIMessagePublisher)(nil).Close))
}

// MockReceiver is a mock of Receiver interface
type MockReceiver struct {
	ctrl     *gomock.Controller
	recorder *MockReceiverMockRecorder
}

// MockReceiverMockRecorder is the mock recorder for MockReceiver
type MockReceiverMockRecorder struct {
	mock *MockReceiver
}

// NewMockReceiver creates a new mock instance
func NewMockReceiver(ctrl *gomock.Controller) *MockReceiver {
	mock := &MockReceiver{ctrl: ctrl}
	mock.recorder = &MockReceiverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReceiver) EXPECT() *MockReceiverMockRecorder {
	return m.recorder
}

// Consume mocks base method
func (m *MockReceiver) Consume(queue string) (<-chan amqp.Delivery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", queue)
	ret0, _ := ret[0].(<-chan amqp.Delivery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume
func (mr *MockReceiverMockRecorder) Consume(queue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockReceiver)(nil).Consume), queue)
}

// GetQueues mocks base method
func (m *MockReceiver) GetQueues() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueues")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetQueues indicates an expected call of GetQueues
func (mr *MockReceiverMockRecorder) GetQueues() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueues", reflect.TypeOf((*MockReceiver)(nil).GetQueues))
}

// MockIMessageReceiver is a mock of IMessageReceiver interface
type MockIMessageReceiver struct {
	ctrl     *gomock.Controller
	recorder *MockIMessageReceiverMockRecorder
}

// MockIMessageReceiverMockRecorder is the mock recorder for MockIMessageReceiver
type MockIMessageReceiverMockRecorder struct {
	mock *MockIMessageReceiver
}

// NewMockIMessageReceiver creates a new mock instance
func NewMockIMessageReceiver(ctrl *gomock.Controller) *MockIMessageReceiver {
	mock := &MockIMessageReceiver{ctrl: ctrl}
	mock.recorder = &MockIMessageReceiverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIMessageReceiver) EXPECT() *MockIMessageReceiverMockRecorder {
	return m.recorder
}

// Consume mocks base method
func (m *MockIMessageReceiver) Consume(queue string) (<-chan amqp.Delivery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", queue)
	ret0, _ := ret[0].(<-chan amqp.Delivery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume
func (mr *MockIMessageReceiverMockRecorder) Consume(queue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockIMessageReceiver)(nil).Consume), queue)
}

// GetQueues mocks base method
func (m *MockIMessageReceiver) GetQueues() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueues")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetQueues indicates an expected call of GetQueues
func (mr *MockIMessageReceiverMockRecorder) GetQueues() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueues", reflect.TypeOf((*MockIMessageReceiver)(nil).GetQueues))
}

// Close mocks base method
func (m *MockIMessageReceiver) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockIMessageReceiverMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIMessageReceiver)(nil).Close))
}

// MockMediator is a mock of Mediator interface
type MockMediator struct {
	ctrl     *gomock.Controller
	recorder *MockMediatorMockRecorder
}

// MockMediatorMockRecorder is the mock recorder for MockMediator
type MockMediatorMockRecorder struct {
	mock *MockMediator
}

// NewMockMediator creates a new mock instance
func NewMockMediator(ctrl *gomock.Controller) *MockMediator {
	mock := &MockMediator{ctrl: ctrl}
	mock.recorder = &MockMediatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMediator) EXPECT() *MockMediatorMockRecorder {
	return m.recorder
}

// SendRaw mocks base method
func (m *MockMediator) SendRaw(queue string, body []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendRaw", queue, body)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendRaw indicates an expected call of SendRaw
func (mr *MockMediatorMockRecorder) SendRaw(queue, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRaw", reflect.TypeOf((*MockMediator)(nil).SendRaw), queue, body)
}

// Send mocks base method
func (m *MockMediator) Send(queue IMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", queue)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockMediatorMockRecorder) Send(queue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockMediator)(nil).Send), queue)
}

// Consume mocks base method
func (m *MockMediator) Consume(queue string) (<-chan amqp.Delivery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", queue)
	ret0, _ := ret[0].(<-chan amqp.Delivery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume
func (mr *MockMediatorMockRecorder) Consume(queue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockMediator)(nil).Consume), queue)
}

// GetQueues mocks base method
func (m *MockMediator) GetQueues() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueues")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetQueues indicates an expected call of GetQueues
func (mr *MockMediatorMockRecorder) GetQueues() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueues", reflect.TypeOf((*MockMediator)(nil).GetQueues))
}

// Close mocks base method
func (m *MockMediator) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockMediatorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMediator)(nil).Close))
}

// MockData is a mock of Data interface
type MockData struct {
	ctrl     *gomock.Controller
	recorder *MockDataMockRecorder
}

// MockDataMockRecorder is the mock recorder for MockData
type MockDataMockRecorder struct {
	mock *MockData
}

// NewMockData creates a new mock instance
func NewMockData(ctrl *gomock.Controller) *MockData {
	mock := &MockData{ctrl: ctrl}
	mock.recorder = &MockDataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockData) EXPECT() *MockDataMockRecorder {
	return m.recorder
}

// GetBody mocks base method
func (m *MockData) GetBody() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBody")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetBody indicates an expected call of GetBody
func (mr *MockDataMockRecorder) GetBody() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBody", reflect.TypeOf((*MockData)(nil).GetBody))
}

// GetKey mocks base method
func (m *MockData) GetKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetKey indicates an expected call of GetKey
func (mr *MockDataMockRecorder) GetKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKey", reflect.TypeOf((*MockData)(nil).GetKey))
}

// Ack mocks base method
func (m *MockData) Ack(arg0 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ack", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ack indicates an expected call of Ack
func (mr *MockDataMockRecorder) Ack(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ack", reflect.TypeOf((*MockData)(nil).Ack), arg0)
}

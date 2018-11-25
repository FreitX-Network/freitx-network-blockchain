// Code generated by MockGen. DO NOT EDIT.
// Source: ./dispatcher/dispatcher.go

// Package mock_dispatcher is a generated GoMock package.
package mock_dispatcher

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	proto "github.com/golang/protobuf/proto"
	dispatcher "github.com/freitx-project/freitx-network-blockchain/dispatcher"
	proto0 "github.com/freitx-project/freitx-network-blockchain/proto"
	net "net"
	reflect "reflect"
)

// MockSubscriber is a mock of Subscriber interface
type MockSubscriber struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriberMockRecorder
}

// MockSubscriberMockRecorder is the mock recorder for MockSubscriber
type MockSubscriberMockRecorder struct {
	mock *MockSubscriber
}

// NewMockSubscriber creates a new mock instance
func NewMockSubscriber(ctrl *gomock.Controller) *MockSubscriber {
	mock := &MockSubscriber{ctrl: ctrl}
	mock.recorder = &MockSubscriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSubscriber) EXPECT() *MockSubscriberMockRecorder {
	return m.recorder
}

// HandleAction mocks base method
func (m *MockSubscriber) HandleAction(arg0 *proto0.ActionPb) error {
	ret := m.ctrl.Call(m, "HandleAction", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleAction indicates an expected call of HandleAction
func (mr *MockSubscriberMockRecorder) HandleAction(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleAction", reflect.TypeOf((*MockSubscriber)(nil).HandleAction), arg0)
}

// HandleBlock mocks base method
func (m *MockSubscriber) HandleBlock(arg0 *proto0.BlockPb) error {
	ret := m.ctrl.Call(m, "HandleBlock", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleBlock indicates an expected call of HandleBlock
func (mr *MockSubscriberMockRecorder) HandleBlock(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleBlock", reflect.TypeOf((*MockSubscriber)(nil).HandleBlock), arg0)
}

// HandleBlockSync mocks base method
func (m *MockSubscriber) HandleBlockSync(arg0 *proto0.BlockPb) error {
	ret := m.ctrl.Call(m, "HandleBlockSync", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleBlockSync indicates an expected call of HandleBlockSync
func (mr *MockSubscriberMockRecorder) HandleBlockSync(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleBlockSync", reflect.TypeOf((*MockSubscriber)(nil).HandleBlockSync), arg0)
}

// HandleSyncRequest mocks base method
func (m *MockSubscriber) HandleSyncRequest(arg0 string, arg1 *proto0.BlockSync) error {
	ret := m.ctrl.Call(m, "HandleSyncRequest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleSyncRequest indicates an expected call of HandleSyncRequest
func (mr *MockSubscriberMockRecorder) HandleSyncRequest(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleSyncRequest", reflect.TypeOf((*MockSubscriber)(nil).HandleSyncRequest), arg0, arg1)
}

// HandleBlockPropose mocks base method
func (m *MockSubscriber) HandleBlockPropose(arg0 *proto0.ProposePb) error {
	ret := m.ctrl.Call(m, "HandleBlockPropose", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleBlockPropose indicates an expected call of HandleBlockPropose
func (mr *MockSubscriberMockRecorder) HandleBlockPropose(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleBlockPropose", reflect.TypeOf((*MockSubscriber)(nil).HandleBlockPropose), arg0)
}

// HandleEndorse mocks base method
func (m *MockSubscriber) HandleEndorse(arg0 *proto0.EndorsePb) error {
	ret := m.ctrl.Call(m, "HandleEndorse", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleEndorse indicates an expected call of HandleEndorse
func (mr *MockSubscriberMockRecorder) HandleEndorse(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleEndorse", reflect.TypeOf((*MockSubscriber)(nil).HandleEndorse), arg0)
}

// MockDispatcher is a mock of Dispatcher interface
type MockDispatcher struct {
	ctrl     *gomock.Controller
	recorder *MockDispatcherMockRecorder
}

// MockDispatcherMockRecorder is the mock recorder for MockDispatcher
type MockDispatcherMockRecorder struct {
	mock *MockDispatcher
}

// NewMockDispatcher creates a new mock instance
func NewMockDispatcher(ctrl *gomock.Controller) *MockDispatcher {
	mock := &MockDispatcher{ctrl: ctrl}
	mock.recorder = &MockDispatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDispatcher) EXPECT() *MockDispatcherMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockDispatcher) Start(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockDispatcherMockRecorder) Start(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockDispatcher)(nil).Start), arg0)
}

// Stop mocks base method
func (m *MockDispatcher) Stop(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Stop", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockDispatcherMockRecorder) Stop(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockDispatcher)(nil).Stop), arg0)
}

// AddSubscriber mocks base method
func (m *MockDispatcher) AddSubscriber(arg0 uint32, arg1 dispatcher.Subscriber) {
	m.ctrl.Call(m, "AddSubscriber", arg0, arg1)
}

// AddSubscriber indicates an expected call of AddSubscriber
func (mr *MockDispatcherMockRecorder) AddSubscriber(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSubscriber", reflect.TypeOf((*MockDispatcher)(nil).AddSubscriber), arg0, arg1)
}

// HandleBroadcast mocks base method
func (m *MockDispatcher) HandleBroadcast(arg0 uint32, arg1 proto.Message, arg2 chan bool) {
	m.ctrl.Call(m, "HandleBroadcast", arg0, arg1, arg2)
}

// HandleBroadcast indicates an expected call of HandleBroadcast
func (mr *MockDispatcherMockRecorder) HandleBroadcast(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleBroadcast", reflect.TypeOf((*MockDispatcher)(nil).HandleBroadcast), arg0, arg1, arg2)
}

// HandleTell mocks base method
func (m *MockDispatcher) HandleTell(arg0 uint32, arg1 net.Addr, arg2 proto.Message, arg3 chan bool) {
	m.ctrl.Call(m, "HandleTell", arg0, arg1, arg2, arg3)
}

// HandleTell indicates an expected call of HandleTell
func (mr *MockDispatcherMockRecorder) HandleTell(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleTell", reflect.TypeOf((*MockDispatcher)(nil).HandleTell), arg0, arg1, arg2, arg3)
}

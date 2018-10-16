// Source: github.com/clustergarage/fim-proto/golang (interfaces: FimdClient,Fimd_GetWatchStateClient)
package fim_mock

import (
	context "context"
	pb "github.com/clustergarage/fim-proto/golang"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockFimdClient is a mock of FimdClient interface
type MockFimdClient struct {
	ctrl     *gomock.Controller
	recorder *MockFimdClientMockRecorder
}

// MockFimdClientMockRecorder is the mock recorder for MockFimdClient
type MockFimdClientMockRecorder struct {
	mock *MockFimdClient
}

// NewMockFimdClient creates a new mock instance
func NewMockFimdClient(ctrl *gomock.Controller) *MockFimdClient {
	mock := &MockFimdClient{ctrl: ctrl}
	mock.recorder = &MockFimdClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFimdClient) EXPECT() *MockFimdClientMockRecorder {
	return m.recorder
}

// CreateWatch mocks base method
func (m *MockFimdClient) CreateWatch(arg0 context.Context, arg1 *pb.FimdConfig, arg2 ...grpc.CallOption) (*pb.FimdHandle, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateWatch", varargs...)
	ret0, _ := ret[0].(*pb.FimdHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWatch indicates an expected call of CreateWatch
func (mr *MockFimdClientMockRecorder) CreateWatch(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWatch", reflect.TypeOf((*MockFimdClient)(nil).CreateWatch), varargs...)
}

// DestroyWatch mocks base method
func (m *MockFimdClient) DestroyWatch(arg0 context.Context, arg1 *pb.FimdConfig, arg2 ...grpc.CallOption) (*pb.Empty, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DestroyWatch", varargs...)
	ret0, _ := ret[0].(*pb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DestroyWatch indicates an expected call of DestroyWatch
func (mr *MockFimdClientMockRecorder) DestroyWatch(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyWatch", reflect.TypeOf((*MockFimdClient)(nil).DestroyWatch), varargs...)
}

// GetWatchState mocks base method
func (m *MockFimdClient) GetWatchState(arg0 context.Context, arg1 *pb.Empty, arg2 ...grpc.CallOption) (pb.Fimd_GetWatchStateClient, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWatchState", varargs...)
	ret0, _ := ret[0].(pb.Fimd_GetWatchStateClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWatchState indicates an expected call of GetWatchState
func (mr *MockFimdClientMockRecorder) GetWatchState(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWatchState", reflect.TypeOf((*MockFimdClient)(nil).GetWatchState), varargs...)
}

// MockFimd_GetWatchStateClient is a mock of Fimd_GetWatchStateClient interface
type MockFimd_GetWatchStateClient struct {
	ctrl     *gomock.Controller
	recorder *MockFimd_GetWatchStateClientMockRecorder
}

// MockFimd_GetWatchStateClientMockRecorder is the mock recorder for MockFimd_GetWatchStateClient
type MockFimd_GetWatchStateClientMockRecorder struct {
	mock *MockFimd_GetWatchStateClient
}

// NewMockFimd_GetWatchStateClient creates a new mock instance
func NewMockFimd_GetWatchStateClient(ctrl *gomock.Controller) *MockFimd_GetWatchStateClient {
	mock := &MockFimd_GetWatchStateClient{ctrl: ctrl}
	mock.recorder = &MockFimd_GetWatchStateClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFimd_GetWatchStateClient) EXPECT() *MockFimd_GetWatchStateClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockFimd_GetWatchStateClient) CloseSend() error {
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockFimd_GetWatchStateClientMockRecorder) CloseSend() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockFimd_GetWatchStateClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockFimd_GetWatchStateClient) Context() context.Context {
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockFimd_GetWatchStateClientMockRecorder) Context() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockFimd_GetWatchStateClient)(nil).Context))
}

// Header mocks base method
func (m *MockFimd_GetWatchStateClient) Header() (metadata.MD, error) {
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockFimd_GetWatchStateClientMockRecorder) Header() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockFimd_GetWatchStateClient)(nil).Header))
}

// Recv mocks base method
func (m *MockFimd_GetWatchStateClient) Recv() (*pb.FimdHandle, error) {
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*pb.FimdHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockFimd_GetWatchStateClientMockRecorder) Recv() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockFimd_GetWatchStateClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockFimd_GetWatchStateClient) RecvMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockFimd_GetWatchStateClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockFimd_GetWatchStateClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *MockFimd_GetWatchStateClient) SendMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockFimd_GetWatchStateClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockFimd_GetWatchStateClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockFimd_GetWatchStateClient) Trailer() metadata.MD {
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockFimd_GetWatchStateClientMockRecorder) Trailer() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockFimd_GetWatchStateClient)(nil).Trailer))
}

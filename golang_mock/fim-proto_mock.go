// Source: github.com/clustergarage/fim-proto/golang (interfaces: FimdClient)
package fim_mock

import (
	context "context"
	golang "github.com/clustergarage/fim-proto/golang"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
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
func (m *MockFimdClient) CreateWatch(arg0 context.Context, arg1 *golang.FimdConfig, arg2 ...grpc.CallOption) (*golang.FimdHandle, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateWatch", varargs...)
	ret0, _ := ret[0].(*golang.FimdHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWatch indicates an expected call of CreateWatch
func (mr *MockFimdClientMockRecorder) CreateWatch(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWatch", reflect.TypeOf((*MockFimdClient)(nil).CreateWatch), varargs...)
}

// DestroyWatch mocks base method
func (m *MockFimdClient) DestroyWatch(arg0 context.Context, arg1 *golang.FimdConfig, arg2 ...grpc.CallOption) (*golang.Empty, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DestroyWatch", varargs...)
	ret0, _ := ret[0].(*golang.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DestroyWatch indicates an expected call of DestroyWatch
func (mr *MockFimdClientMockRecorder) DestroyWatch(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyWatch", reflect.TypeOf((*MockFimdClient)(nil).DestroyWatch), varargs...)
}

// GetWatchState mocks base method
func (m *MockFimdClient) GetWatchState(arg0 context.Context, arg1 *golang.Empty, arg2 ...grpc.CallOption) (golang.Fimd_GetWatchStateClient, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWatchState", varargs...)
	ret0, _ := ret[0].(golang.Fimd_GetWatchStateClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWatchState indicates an expected call of GetWatchState
func (mr *MockFimdClientMockRecorder) GetWatchState(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWatchState", reflect.TypeOf((*MockFimdClient)(nil).GetWatchState), varargs...)
}

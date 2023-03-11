// Code generated by MockGen. DO NOT EDIT.
// Source: service/catalog/interfaces/ticket.go

// Package mocks is a generated GoMock package.
package mocks

import (
	dto "projects/music_microservices/StorageService/domain/dto"
	"projects/music_microservices/StorageService/service/catalog/interfaces"
	reflect "reflect"

	consts "github.com/Yuno-obsessed/music_microservices/ProjectLibrary/consts"
	gomock "github.com/golang/mock/gomock"
)

// MockTicketInterface is a mock of TicketInterface interface.
type MockTicketInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTicketInterfaceMockRecorder
}

// MockTicketInterfaceMockRecorder is the mock recorder for MockTicketInterface.
type MockTicketInterfaceMockRecorder struct {
	mock *MockTicketInterface
}

// NewMockTicketInterface creates a new mock instance.
func NewMockTicketInterface(ctrl *gomock.Controller) *MockTicketInterface {
	mock := &MockTicketInterface{ctrl: ctrl}
	mock.recorder = &MockTicketInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTicketInterface) EXPECT() *MockTicketInterfaceMockRecorder {
	return m.recorder
}

// GetEntity mocks base method.
func (m *MockTicketInterface) GetEntity(id int) (dto.TicketOut, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntity", id)
	ret0, _ := ret[0].(dto.TicketOut)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntity indicates an expected call of GetEntity.
func (mr *MockTicketInterfaceMockRecorder) GetEntity(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntity", reflect.TypeOf((*MockTicketInterface)(nil).GetEntity), id)
}

// GetSumAndAverage mocks base method.
func (m *MockTicketInterface) GetSumAndAverage(id int) (dto.TicketBrief, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSumAndAverage", id)
	ret0, _ := ret[0].(dto.TicketBrief)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSumAndAverage indicates an expected call of GetSumAndAverage.
func (mr *MockTicketInterfaceMockRecorder) GetSumAndAverage(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSumAndAverage", reflect.TypeOf((*MockTicketInterface)(nil).GetSumAndAverage), id)
}

// Subtruct mocks base method.
func (m *MockTicketInterface) Subtruct(id int, ttype consts.TicketType, number int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subtruct", id, ttype, number)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subtruct indicates an expected call of Subtruct.
func (mr *MockTicketInterfaceMockRecorder) Subtruct(id, ttype, number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subtruct", reflect.TypeOf((*MockTicketInterface)(nil).Subtruct), id, ttype, number)
}
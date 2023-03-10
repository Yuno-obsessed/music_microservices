// Code generated by MockGen. DO NOT EDIT.
// Source: service/upload/interfaces/fileupload.go

// Package mocks is a generated GoMock package.
package mocks

import (
	multipart "mime/multipart"
	reflect "reflect"

	consts "github.com/Yuno-obsessed/music_microservices/ProjectLibrary/consts"
	gomock "github.com/golang/mock/gomock"
)

// MockUploadFile is a mock of UploadFile interface.
type MockUploadFile struct {
	ctrl     *gomock.Controller
	recorder *MockUploadFileMockRecorder
}

// MockUploadFileMockRecorder is the mock recorder for MockUploadFile.
type MockUploadFileMockRecorder struct {
	mock *MockUploadFile
}

// NewMockUploadFile creates a new mock instance.
func NewMockUploadFile(ctrl *gomock.Controller) *MockUploadFile {
	mock := &MockUploadFile{ctrl: ctrl}
	mock.recorder = &MockUploadFileMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUploadFile) EXPECT() *MockUploadFileMockRecorder {
	return m.recorder
}

// DeleteFile mocks base method.
func (m *MockUploadFile) DeleteFile(file string, bucket consts.BucketName) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFile", file, bucket)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFile indicates an expected call of DeleteFile.
func (mr *MockUploadFileMockRecorder) DeleteFile(file, bucket interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockUploadFile)(nil).DeleteFile), file, bucket)
}

// ReplaceFile mocks base method.
func (m *MockUploadFile) ReplaceFile(file string, newFile *multipart.FileHeader, bucket consts.BucketName) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplaceFile", file, newFile, bucket)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReplaceFile indicates an expected call of ReplaceFile.
func (mr *MockUploadFileMockRecorder) ReplaceFile(file, newFile, bucket interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplaceFile", reflect.TypeOf((*MockUploadFile)(nil).ReplaceFile), file, newFile, bucket)
}

// UploadFile mocks base method.
func (m *MockUploadFile) UploadFile(file *multipart.FileHeader, name string, bucket consts.BucketName) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", file, name, bucket)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockUploadFileMockRecorder) UploadFile(file, name, bucket interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockUploadFile)(nil).UploadFile), file, name, bucket)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: git/middleware/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	git "github.com/libgit2/git2go/v31"
	middleware "github.com/neel1996/gitconvex/git/middleware"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateBranch mocks base method.
func (m *MockRepository) CreateBranch(arg0 string, arg1 *git.Commit, arg2 bool) (*git.Branch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBranch", arg0, arg1, arg2)
	ret0, _ := ret[0].(*git.Branch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBranch indicates an expected call of CreateBranch.
func (mr *MockRepositoryMockRecorder) CreateBranch(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBranch", reflect.TypeOf((*MockRepository)(nil).CreateBranch), arg0, arg1, arg2)
}

// CreateCommit mocks base method.
func (m *MockRepository) CreateCommit(s string, signature, signature2 *git.Signature, message string, tree *git.Tree, parents ...*git.Commit) (*git.Oid, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{s, signature, signature2, message, tree}
	for _, a := range parents {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateCommit", varargs...)
	ret0, _ := ret[0].(*git.Oid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCommit indicates an expected call of CreateCommit.
func (mr *MockRepositoryMockRecorder) CreateCommit(s, signature, signature2, message, tree interface{}, parents ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{s, signature, signature2, message, tree}, parents...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCommit", reflect.TypeOf((*MockRepository)(nil).CreateCommit), varargs...)
}

// DefaultSignature mocks base method.
func (m *MockRepository) DefaultSignature() (*git.Signature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultSignature")
	ret0, _ := ret[0].(*git.Signature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DefaultSignature indicates an expected call of DefaultSignature.
func (mr *MockRepositoryMockRecorder) DefaultSignature() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultSignature", reflect.TypeOf((*MockRepository)(nil).DefaultSignature))
}

// DiffTreeToTree mocks base method.
func (m *MockRepository) DiffTreeToTree(tree, tree2 *git.Tree, options *git.DiffOptions) (*git.Diff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiffTreeToTree", tree, tree2, options)
	ret0, _ := ret[0].(*git.Diff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiffTreeToTree indicates an expected call of DiffTreeToTree.
func (mr *MockRepositoryMockRecorder) DiffTreeToTree(tree, tree2, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiffTreeToTree", reflect.TypeOf((*MockRepository)(nil).DiffTreeToTree), tree, tree2, options)
}

// Head mocks base method.
func (m *MockRepository) Head() (middleware.Reference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Head")
	ret0, _ := ret[0].(middleware.Reference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Head indicates an expected call of Head.
func (mr *MockRepositoryMockRecorder) Head() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Head", reflect.TypeOf((*MockRepository)(nil).Head))
}

// Index mocks base method.
func (m *MockRepository) Index() (middleware.Index, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Index")
	ret0, _ := ret[0].(middleware.Index)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Index indicates an expected call of Index.
func (mr *MockRepositoryMockRecorder) Index() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Index", reflect.TypeOf((*MockRepository)(nil).Index))
}

// LookupCommit mocks base method.
func (m *MockRepository) LookupCommit(oid *git.Oid) (*git.Commit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookupCommit", oid)
	ret0, _ := ret[0].(*git.Commit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LookupCommit indicates an expected call of LookupCommit.
func (mr *MockRepositoryMockRecorder) LookupCommit(oid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookupCommit", reflect.TypeOf((*MockRepository)(nil).LookupCommit), oid)
}

// LookupTree mocks base method.
func (m *MockRepository) LookupTree(id *git.Oid) (*git.Tree, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookupTree", id)
	ret0, _ := ret[0].(*git.Tree)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LookupTree indicates an expected call of LookupTree.
func (mr *MockRepositoryMockRecorder) LookupTree(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookupTree", reflect.TypeOf((*MockRepository)(nil).LookupTree), id)
}

// Remotes mocks base method.
func (m *MockRepository) Remotes() middleware.Remotes {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remotes")
	ret0, _ := ret[0].(middleware.Remotes)
	return ret0
}

// Remotes indicates an expected call of Remotes.
func (mr *MockRepositoryMockRecorder) Remotes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remotes", reflect.TypeOf((*MockRepository)(nil).Remotes))
}

// Walk mocks base method.
func (m *MockRepository) Walk() (middleware.RevWalk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Walk")
	ret0, _ := ret[0].(middleware.RevWalk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Walk indicates an expected call of Walk.
func (mr *MockRepositoryMockRecorder) Walk() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Walk", reflect.TypeOf((*MockRepository)(nil).Walk))
}

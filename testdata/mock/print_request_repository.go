package mock

import (
	"threedee/entity"

	"github.com/stretchr/testify/mock"
)

type MockPrintRequestRepository struct {
	mock.Mock
}

func (mr *MockPrintRequestRepository) GetAll() ([]*entity.PrintRequest, error) {
	args := mr.Called()
	return args.Get(0).([]*entity.PrintRequest), args.Error(1)
}

func (mr *MockPrintRequestRepository) GetById(id int) (*entity.PrintRequest, error) {
	args := mr.Called(id)
	return args.Get(0).(*entity.PrintRequest), args.Error(1)
}

func (mr *MockPrintRequestRepository) Insert(model *entity.PrintRequest) (int, error) {
	args := mr.Called(model)
	return args.Get(0).(int), args.Error(1)
}

func (mr *MockPrintRequestRepository) Update(model *entity.PrintRequest) (bool, error) {
	args := mr.Called(model)
	return args.Get(0).(bool), args.Error(1)
}

func (mr *MockPrintRequestRepository) Delete(id int) (bool, error) {
	args := mr.Called(id)
	return args.Get(0).(bool), args.Error(1)
}

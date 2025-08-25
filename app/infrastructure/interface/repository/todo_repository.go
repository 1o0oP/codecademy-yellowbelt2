package repository

import (
	"codecademy-yellowbelt2/core/domain/entity"

	"github.com/stretchr/testify/mock"
)

type ITodoRepository interface {
	Create(todo *entity.Todo) error
	GetByID(id string) (*entity.Todo, error)
	GetAll() ([]*entity.Todo, error)
	Update(todo *entity.Todo) error
	Delete(id string) error
}

type MockTodoRepository struct {
	mock.Mock
}

func (m *MockTodoRepository) Create(todo *entity.Todo) error {
	args := m.Called(todo)
	return args.Error(0)
}

func (m *MockTodoRepository) GetByID(id string) (*entity.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Todo), args.Error(1)
}

func (m *MockTodoRepository) GetAll() ([]*entity.Todo, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Todo), args.Error(1)
}

func (m *MockTodoRepository) Update(todo *entity.Todo) error {
	args := m.Called(todo)
	return args.Error(0)
}

func (m *MockTodoRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

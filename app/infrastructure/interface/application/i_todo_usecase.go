package application

import (
	"codecademy-yellowbelt2/core/domain/entity"

	"github.com/stretchr/testify/mock"
)

type ITodoUseCase interface {
	CreateTodo(title, description string) (*entity.Todo, error)
	GetTodoByID(id string) (*entity.Todo, error)
	GetAllTodos() ([]*entity.Todo, error)
	UpdateTodo(id, title, description string) (*entity.Todo, error)
	CompleteTodo(id string) (*entity.Todo, error)
	DeleteTodo(id string) error
}

type MockTodoUseCase struct {
	mock.Mock
}

func (m *MockTodoUseCase) CreateTodo(title, description string) (*entity.Todo, error) {
	args := m.Called(title, description)
	todo, _ := args.Get(0).(*entity.Todo)
	return todo, args.Error(1)
}

func (m *MockTodoUseCase) GetTodoByID(id string) (*entity.Todo, error) {
	args := m.Called(id)
	todo, _ := args.Get(0).(*entity.Todo)
	return todo, args.Error(1)
}

func (m *MockTodoUseCase) GetAllTodos() ([]*entity.Todo, error) {
	args := m.Called()
	todos, _ := args.Get(0).([]*entity.Todo)
	return todos, args.Error(1)
}

func (m *MockTodoUseCase) UpdateTodo(id, title, description string) (*entity.Todo, error) {
	args := m.Called(id, title, description)
	todo, _ := args.Get(0).(*entity.Todo)
	return todo, args.Error(1)
}

func (m *MockTodoUseCase) CompleteTodo(id string) (*entity.Todo, error) {
	args := m.Called(id)
	todo, _ := args.Get(0).(*entity.Todo)
	return todo, args.Error(1)
}

func (m *MockTodoUseCase) DeleteTodo(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

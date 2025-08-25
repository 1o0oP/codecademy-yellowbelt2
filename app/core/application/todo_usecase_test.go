package application

import (
	"codecademy-yellowbelt2/core/domain/entity"
	repoMock "codecademy-yellowbelt2/infrastructure/interface/repository"
	"codecademy-yellowbelt2/infrastructure/repository"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTodoUseCase_CreateTodo(t *testing.T) {
	// Arrange
	repo := repository.NewInMemoryTodoRepository()
	useCase := NewTodoUseCase(repo)

	// Act
	todo, err := useCase.CreateTodo("Test Todo", "Test Description")

	// Assert
	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, "Test Todo", todo.Title, "Expected todo to be created with correct title")
}

func TestTodoUseCase_GetAllTodos(t *testing.T) {
	// Arrange
	repo := repository.NewInMemoryTodoRepository()
	useCase := NewTodoUseCase(repo)
	useCase.CreateTodo("Todo 1", "Description 1")
	useCase.CreateTodo("Todo 2", "Description 2")

	// Act
	todos, err := useCase.GetAllTodos()

	// Assert
	assert.NoError(t, err, "Expected no error")
	assert.Len(t, todos, 2, "Expected 2 todos")
}

func TestTodoUseCase_UpdateTodo(t *testing.T) {
	// Arrange
	repo := repository.NewInMemoryTodoRepository()
	useCase := NewTodoUseCase(repo)
	todo, _ := useCase.CreateTodo("Original", "Original Description")

	// Act
	updated, err := useCase.UpdateTodo(todo.ID, "Updated", "Updated Description")

	// Assert
	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, "Updated", updated.Title, "Expected todo to be updated")
}

func TestTodoUseCase_CompleteTodo(t *testing.T) {
	// Arrange
	repo := repository.NewInMemoryTodoRepository()
	useCase := NewTodoUseCase(repo)
	todo, _ := useCase.CreateTodo("Test", "Test Description")

	// Act
	completed, err := useCase.CompleteTodo(todo.ID)

	// Assert
	assert.NoError(t, err, "Expected no error")
	assert.True(t, completed.Completed, "Expected todo to be completed")
}

func TestTodoUseCase_DeleteTodo(t *testing.T) {
	// Arrange
	repo := repository.NewInMemoryTodoRepository()
	useCase := NewTodoUseCase(repo)
	todo, _ := useCase.CreateTodo("Test", "Test Description")

	// Act
	err := useCase.DeleteTodo(todo.ID)
	_, getErr := useCase.GetTodoByID(todo.ID)

	// Assert
	assert.NoError(t, err, "Expected no error")
	assert.Error(t, getErr, "Expected error when getting deleted todo")
}

func TestShouldReturnErrorWhenCreateTodoFails(t *testing.T) {
	// Arrange
	mockRepo := new(repoMock.MockTodoRepository)
	useCase := NewTodoUseCase(mockRepo)
	expectedErr := errors.New("some error")
	mockRepo.On("Create", mock.AnythingOfType("*entity.Todo")).Return(expectedErr)

	// Act
	todo, err := useCase.CreateTodo("Title", "Description")

	// Assert
	assert.Nil(t, todo, "Expected todo to be nil when creation fails")
	assert.Error(t, err, "Expected error when repository Create fails")
	mockRepo.AssertExpectations(t)
}

func TestShouldReturnErrorWhenUpdateTodoGetByIDFails(t *testing.T) {
	// Arrange
	mockRepo := new(repoMock.MockTodoRepository)
	useCase := NewTodoUseCase(mockRepo)
	expectedErr := errors.New("get by id error")
	mockRepo.On("GetByID", "some-id").Return(&entity.Todo{}, expectedErr)

	// Act
	todo, err := useCase.UpdateTodo("some-id", "New Title", "New Description")

	// Assert
	assert.Nil(t, todo, "Expected todo to be nil when GetByID fails")
	assert.Error(t, err, "Expected error when repository GetByID fails")
	mockRepo.AssertExpectations(t)
}

func TestShouldReturnErrorWhenUpdateTodoUpdateFails(t *testing.T) {
	// Arrange
	mockRepo := new(repoMock.MockTodoRepository)
	useCase := NewTodoUseCase(mockRepo)
	existingTodo := &entity.Todo{ID: "some-id", Title: "Old", Description: "Old"}
	expectedErr := errors.New("update error")
	mockRepo.On("GetByID", "some-id").Return(existingTodo, nil)
	mockRepo.On("Update", mock.AnythingOfType("*entity.Todo")).Return(expectedErr)

	// Act
	todo, err := useCase.UpdateTodo("some-id", "New Title", "New Description")

	// Assert
	assert.Nil(t, todo, "Expected todo to be nil when Update fails")
	assert.Error(t, err, "Expected error when repository Update fails")
	mockRepo.AssertExpectations(t)
}

func TestShouldReturnErrorWhenCompleteTodoGetByIDFails(t *testing.T) {
	// Arrange
	mockRepo := new(repoMock.MockTodoRepository)
	useCase := NewTodoUseCase(mockRepo)
	expectedErr := errors.New("get by id error")
	mockRepo.On("GetByID", "some-id").Return(&entity.Todo{}, expectedErr)

	// Act
	todo, err := useCase.CompleteTodo("some-id")

	// Assert
	assert.Nil(t, todo, "Expected todo to be nil when GetByID fails")
	assert.Error(t, err, "Expected error when repository GetByID fails")
	mockRepo.AssertExpectations(t)
}

func TestShouldReturnErrorWhenCompleteTodoUpdateFails(t *testing.T) {
	// Arrange
	mockRepo := new(repoMock.MockTodoRepository)
	useCase := NewTodoUseCase(mockRepo)
	existingTodo := &entity.Todo{ID: "some-id", Title: "Old", Description: "Old"}
	expectedErr := errors.New("update error")
	mockRepo.On("GetByID", "some-id").Return(existingTodo, nil)
	mockRepo.On("Update", mock.AnythingOfType("*entity.Todo")).Return(expectedErr)

	// Act
	todo, err := useCase.CompleteTodo("some-id")

	// Assert
	assert.Nil(t, todo, "Expected todo to be nil when Update fails")
	assert.Error(t, err, "Expected error when repository Update fails")
	mockRepo.AssertExpectations(t)
}

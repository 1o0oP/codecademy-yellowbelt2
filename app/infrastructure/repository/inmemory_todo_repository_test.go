package repository

import (
	"codecademy-yellowbelt2/core/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateTodoForInMemory(t *testing.T) {
	// Arrange
	repo := NewInMemoryTodoRepository()
	todo := entity.NewTodo("Test Todo", "Test Description")

	// Act
	err := repo.Create(todo)
	retrieved, getErr := repo.GetByID(todo.ID)

	// Assert
	assert.NoError(t, err)
	assert.NoError(t, getErr)
	assert.Equal(t, todo.ID, retrieved.ID)
}

func TestShouldGetAllTodosForInMemory(t *testing.T) {
	// Arrange
	repo := NewInMemoryTodoRepository()
	todo1 := entity.NewTodo("Todo 1", "Description 1")
	todo2 := entity.NewTodo("Todo 2", "Description 2")
	repo.Create(todo1)
	repo.Create(todo2)

	// Act
	todos, err := repo.GetAll()

	// Assert
	assert.NoError(t, err)
	assert.Len(t, todos, 2)
}

func TestShouldUpdateTodoForInMemory(t *testing.T) {
	// Arrange
	repo := NewInMemoryTodoRepository()
	todo := entity.NewTodo("Original", "Original Description")
	repo.Create(todo)
	todo.Update("Updated", "Updated Description")

	// Act
	err := repo.Update(todo)
	retrieved, getErr := repo.GetByID(todo.ID)

	// Assert
	assert.NoError(t, err)
	assert.NoError(t, getErr)
	assert.Equal(t, "Updated", retrieved.Title)
	assert.Equal(t, "Updated Description", retrieved.Description)
}

func TestShouldDeleteTodoForInMemory(t *testing.T) {
	// Arrange
	repo := NewInMemoryTodoRepository()
	todo := entity.NewTodo("Test", "Test Description")
	repo.Create(todo)

	// Act
	err := repo.Delete(todo.ID)
	retrieved, getErr := repo.GetByID(todo.ID)

	// Assert
	assert.NoError(t, err)
	assert.Error(t, getErr)
	assert.Nil(t, retrieved)
}

func TestShouldReturnErrorWhenUpdatingNonExistentTodo(t *testing.T) {
	// Arrange
	repo := NewInMemoryTodoRepository()
	nonExistentTodo := entity.NewTodo("Non-existent", "Does not exist")

	// Act
	err := repo.Update(nonExistentTodo)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "todo not found", err.Error())
}

func TestShouldReturnErrorWhenDeletingNonExistentTodo(t *testing.T) {
	// Arrange
	repo := NewInMemoryTodoRepository()
	nonExistentID := "non-existent-id"

	// Act
	err := repo.Delete(nonExistentID)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "todo not found", err.Error())
}

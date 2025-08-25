package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateNewTodo(t *testing.T) {
	// Arrange
	title := "Test Todo"
	description := "Test Description"

	// Act
	todo := NewTodo(title, description)

	// Assert
	assert.NotEmpty(t, todo.ID, "Expected ID to be generated")
	assert.Equal(t, title, todo.Title, "Expected title to match")
	assert.Equal(t, description, todo.Description, "Expected description to match")
	assert.False(t, todo.Completed, "Expected new todo to be incomplete")
}

func TestShouldMarkTodoAsCompleted(t *testing.T) {
	// Arrange
	todo := NewTodo("Test", "Description")
	originalTime := todo.UpdatedAt

	// Act
	time.Sleep(1 * time.Millisecond)
	todo.MarkAsCompleted()

	// Assert
	assert.True(t, todo.Completed, "Expected todo to be completed")
	assert.True(t, todo.UpdatedAt.After(originalTime), "Expected UpdatedAt to be updated")
}

func TestShouldUpdateTodo(t *testing.T) {
	// Arrange
	todo := NewTodo("Original", "Original Description")
	originalTime := todo.UpdatedAt

	// Act
	time.Sleep(1 * time.Millisecond)
	todo.Update("Updated", "Updated Description")

	// Assert
	assert.Equal(t, "Updated", todo.Title, "Expected title to be updated")
	assert.Equal(t, "Updated Description", todo.Description, "Expected description to be updated")
	assert.True(t, todo.UpdatedAt.After(originalTime), "Expected UpdatedAt to be updated")
}

func TestShouldMarkTodoAsIncomplete(t *testing.T) {
	// Arrange
	todo := NewTodo("Test", "Description")
	todo.MarkAsCompleted()
	originalTime := todo.UpdatedAt

	// Act
	time.Sleep(1 * time.Millisecond)
	todo.MarkAsIncomplete()

	// Assert
	assert.False(t, todo.Completed, "Expected todo to be incomplete")
	assert.True(t, todo.UpdatedAt.After(originalTime), "Expected UpdatedAt to be updated")
}

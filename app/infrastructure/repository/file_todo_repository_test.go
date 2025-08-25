package repository

import (
	"codecademy-yellowbelt2/core/domain/entity"
	"encoding/json"
	"errors"
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func createTempRepo(t *testing.T) (*FileTodoRepository, func()) {
	tmpfile, err := os.CreateTemp("", "todos_*.json")
	assert.NoError(t, err)
	repo := NewFileTodoRepository(tmpfile.Name()).(*FileTodoRepository)
	cleanup := func() {
		os.Remove(tmpfile.Name())
	}
	return repo, cleanup
}

func TestShouldCreateTodo(t *testing.T) {
	// Arrange
	repo, cleanup := createTempRepo(t)
	defer cleanup()
	todo := &entity.Todo{ID: "1", Title: "Test", Completed: false}

	// Act
	err := repo.Create(todo)
	got, getErr := repo.GetByID("1")

	// Assert
	assert.NoError(t, err)
	assert.NoError(t, getErr)
	assert.Equal(t, "Test", got.Title)
	assert.False(t, got.Completed)
}

func TestShouldReturnErrorWhenTodoNotFoundByID(t *testing.T) {
	// Arrange
	repo, cleanup := createTempRepo(t)
	defer cleanup()

	// Act
	_, err := repo.GetByID("notfound")

	// Assert
	assert.Error(t, err)
}

func TestShouldGetAllTodos(t *testing.T) {
	// Arrange
	repo, cleanup := createTempRepo(t)
	defer cleanup()
	todo1 := &entity.Todo{ID: "1", Title: "A", Completed: false}
	todo2 := &entity.Todo{ID: "2", Title: "B", Completed: true}
	repo.Create(todo1)
	repo.Create(todo2)

	// Act
	todos, err := repo.GetAll()

	// Assert
	assert.NoError(t, err)
	assert.Len(t, todos, 2)
}

func TestShouldUpdateTodo(t *testing.T) {
	// Arrange
	repo, cleanup := createTempRepo(t)
	defer cleanup()
	todo := &entity.Todo{ID: "1", Title: "Old", Completed: false}
	repo.Create(todo)
	updated := &entity.Todo{ID: "1", Title: "New", Completed: true}

	// Act
	err := repo.Update(updated)
	got, getErr := repo.GetByID("1")

	// Assert
	assert.NoError(t, err)
	assert.NoError(t, getErr)
	assert.Equal(t, "New", got.Title)
	assert.True(t, got.Completed)
}

func TestShouldReturnErrorWhenUpdateTodoNotFound(t *testing.T) {
	// Arrange
	repo, cleanup := createTempRepo(t)
	defer cleanup()
	todo := &entity.Todo{ID: "notfound", Title: "X", Completed: false}

	// Act
	err := repo.Update(todo)

	// Assert
	assert.Error(t, err)
}

func TestShouldDeleteTodo(t *testing.T) {
	// Arrange
	repo, cleanup := createTempRepo(t)
	defer cleanup()
	todo := &entity.Todo{ID: "1", Title: "Delete", Completed: false}
	repo.Create(todo)

	// Act
	err := repo.Delete("1")
	_, getErr := repo.GetByID("1")

	// Assert
	assert.NoError(t, err)
	assert.Error(t, getErr)
}

func TestShouldReturnErrorWhenDeleteTodoNotFound(t *testing.T) {
	// Arrange
	repo, cleanup := createTempRepo(t)
	defer cleanup()

	// Act
	err := repo.Delete("notfound")

	// Assert
	assert.Error(t, err)
}

func TestShouldReturnErrorOnLoadWhenReadFileFails(t *testing.T) {
	// Arrange
	repo, cleanup := createTempRepo(t)
	defer cleanup()
	patch := monkey.Patch(os.ReadFile, func(string) ([]byte, error) {
		return nil, errors.New("read error")
	})
	defer patch.Unpatch()

	// Act
	todos, err := repo.load()

	// Assert
	assert.Nil(t, todos)
	assert.Error(t, err)
	assert.Equal(t, "read error", err.Error())
}

func TestShouldReturnErrorOnLoadWhenUnmarshalFails(t *testing.T) {
	// Arrange
	repo, cleanup := createTempRepo(t)
	defer cleanup()
	patch := monkey.Patch(os.ReadFile, func(string) ([]byte, error) {
		return []byte("{invalid json"), nil
	})
	defer patch.Unpatch()

	// Act
	todos, err := repo.load()

	// Assert
	assert.Nil(t, todos)
	assert.Error(t, err)
	_, ok := err.(*json.SyntaxError)
	assert.True(t, ok)
}

func TestShouldReturnEmptyMapOnLoadWhenFileNotExist(t *testing.T) {
	// Arrange
	repo, cleanup := createTempRepo(t)
	defer cleanup()
	patch := monkey.Patch(os.ReadFile, func(string) ([]byte, error) {
		return nil, os.ErrNotExist
	})
	defer patch.Unpatch()

	// Act
	todos, err := repo.load()

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, todos)
	assert.Len(t, todos, 0)
}

func TestShouldReturnErrorOnSaveWhenMarshalFails(t *testing.T) {
	// Arrange
	repo, cleanup := createTempRepo(t)
	defer cleanup()
	todos := map[string]*entity.Todo{
		"invalid": nil,
	}
	patch := monkey.Patch(json.MarshalIndent, func(v interface{}, prefix, indent string) ([]byte, error) {
		return nil, errors.New("marshal error")
	})
	defer patch.Unpatch()

	// Act
	err := repo.save(todos)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "marshal error", err.Error())
}

func TestShouldReturnErrorOnSaveWhenWriteFileFails(t *testing.T) {
	// Arrange
	repo, cleanup := createTempRepo(t)
	defer cleanup()
	todos := map[string]*entity.Todo{
		"1": {ID: "1", Title: "Test", Completed: false},
	}
	patch := monkey.Patch(os.WriteFile, func(string, []byte, os.FileMode) error {
		return errors.New("write error")
	})
	defer patch.Unpatch()

	// Act
	err := repo.save(todos)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "write error", err.Error())
}

func TestShouldReturnErrorOnCreateWhenLoadFails(t *testing.T) {
	// Arrange
	repo, cleanup := createTempRepo(t)
	defer cleanup()
	patch := monkey.Patch(os.ReadFile, func(string) ([]byte, error) {
		return nil, errors.New("read error")
	})
	defer patch.Unpatch()
	todo := &entity.Todo{ID: "1", Title: "Test", Completed: false}

	// Act
	err := repo.Create(todo)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "read error")
}

func TestShouldReturnErrorOnGetByIDWhenLoadFails(t *testing.T) {
	// Arrange
	repo, cleanup := createTempRepo(t)
	defer cleanup()
	patch := monkey.Patch(os.ReadFile, func(string) ([]byte, error) {
		return nil, errors.New("read error")
	})
	defer patch.Unpatch()

	// Act
	todo, err := repo.GetByID("1")

	// Assert
	assert.Nil(t, todo)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "read error")
}

func TestShouldReturnErrorOnGetAllWhenLoadFails(t *testing.T) {
	// Arrange
	repo, cleanup := createTempRepo(t)
	defer cleanup()
	patch := monkey.Patch(os.ReadFile, func(string) ([]byte, error) {
		return nil, errors.New("read error")
	})
	defer patch.Unpatch()

	// Act
	todos, err := repo.GetAll()

	// Assert
	assert.Nil(t, todos)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "read error")
}

func TestShouldReturnErrorOnUpdateWhenLoadFails(t *testing.T) {
	// Arrange
	repo, cleanup := createTempRepo(t)
	defer cleanup()
	patch := monkey.Patch(os.ReadFile, func(string) ([]byte, error) {
		return nil, errors.New("read error")
	})
	defer patch.Unpatch()
	todo := &entity.Todo{ID: "1", Title: "Test", Completed: false}

	// Act
	err := repo.Update(todo)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "read error")
}

func TestShouldReturnErrorOnDeleteWhenLoadFails(t *testing.T) {
	// Arrange
	repo, cleanup := createTempRepo(t)
	defer cleanup()
	patch := monkey.Patch(os.ReadFile, func(string) ([]byte, error) {
		return nil, errors.New("read error")
	})
	defer patch.Unpatch()

	// Act
	err := repo.Delete("1")

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "read error")
}

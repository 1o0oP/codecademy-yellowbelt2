package cli

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"codecademy-yellowbelt2/core/domain/entity"
	"codecademy-yellowbelt2/infrastructure/interface/application"

	"github.com/stretchr/testify/assert"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = stdout
	buf.ReadFrom(r)
	return buf.String()
}

func TestShouldCreateTodoSuccessfully(t *testing.T) {
	// Arrange
	mockUseCase := new(application.MockTodoUseCase)
	cli := NewTodoCLI(mockUseCase)
	expectedTodo := &entity.Todo{
		ID:          "1",
		Title:       "Test",
		Description: "Desc",
	}
	mockUseCase.On("CreateTodo", "Test", "Desc").Return(expectedTodo, nil)

	cmd := cli.createCommand()
	cmd.SetArgs([]string{"Test", "Desc"})

	// Act
	output := captureOutput(func() {
		cmd.Execute()
	})

	// Assert
	assert.Contains(t, output, "✅ Tarefa criada com sucesso!")
	assert.Contains(t, output, "ID: 1")
	assert.Contains(t, output, "Título: Test")
	assert.Contains(t, output, "Descrição: Desc")
	mockUseCase.AssertExpectations(t)
}

func TestShouldCreateTodoWithError(t *testing.T) {
	// Arrange
	mockUseCase := new(application.MockTodoUseCase)
	cli := NewTodoCLI(mockUseCase)
	mockUseCase.On("CreateTodo", "Test", "").Return(nil, errors.New("fail"))

	cmd := cli.createCommand()
	cmd.SetArgs([]string{"Test"})

	// Act
	output := captureOutput(func() {
		cmd.Execute()
	})

	// Assert
	assert.Contains(t, output, "Erro ao criar tarefa: fail")
	mockUseCase.AssertExpectations(t)
}

func TestShouldListTodosSuccessfully(t *testing.T) {
	// Arrange
	mockUseCase := new(application.MockTodoUseCase)
	cli := NewTodoCLI(mockUseCase)
	todos := []*entity.Todo{
		{ID: "1", Title: "A", Description: "D", Completed: false},
		{ID: "2", Title: "B", Description: "", Completed: true},
	}
	mockUseCase.On("GetAllTodos").Return(todos, nil)

	cmd := cli.listCommand()
	cmd.SetArgs([]string{})

	// Act
	output := captureOutput(func() {
		cmd.Execute()
	})

	// Assert
	assert.Contains(t, output, "📋 Total de tarefas: 2")
	assert.Contains(t, output, "1. ⏳ A")
	assert.Contains(t, output, "   📄 D")
	assert.Contains(t, output, "   🆔 ID: 1")
	assert.Contains(t, output, "2. ✅ B")
	assert.Contains(t, output, "   🆔 ID: 2")
	mockUseCase.AssertExpectations(t)
}

func TestShouldListTodosWithError(t *testing.T) {
	// Arrange
	mockUseCase := new(application.MockTodoUseCase)
	cli := NewTodoCLI(mockUseCase)
	mockUseCase.On("GetAllTodos").Return(nil, errors.New("fail"))

	cmd := cli.listCommand()
	cmd.SetArgs([]string{})

	// Act
	output := captureOutput(func() {
		cmd.Execute()
	})

	// Assert
	assert.Contains(t, output, "Erro ao listar tarefas: fail")
	mockUseCase.AssertExpectations(t)
}

func TestShouldListTodosEmpty(t *testing.T) {
	// Arrange
	mockUseCase := new(application.MockTodoUseCase)
	cli := NewTodoCLI(mockUseCase)
	mockUseCase.On("GetAllTodos").Return([]*entity.Todo{}, nil)

	cmd := cli.listCommand()
	cmd.SetArgs([]string{})

	// Act
	output := captureOutput(func() {
		cmd.Execute()
	})

	// Assert
	assert.Contains(t, output, "📝 Nenhuma tarefa encontrada!")
	mockUseCase.AssertExpectations(t)
}

func TestShouldShowTodoSuccessfully(t *testing.T) {
	// Arrange
	mockUseCase := new(application.MockTodoUseCase)
	cli := NewTodoCLI(mockUseCase)
	now := time.Now()
	todo := &entity.Todo{
		ID:          "1",
		Title:       "Test",
		Description: "Desc",
		Completed:   true,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	mockUseCase.On("GetTodoByID", "1").Return(todo, nil)

	cmd := cli.showCommand()
	cmd.SetArgs([]string{"1"})

	// Act
	output := captureOutput(func() {
		cmd.Execute()
	})

	// Assert
	assert.Contains(t, output, "🆔 ID: 1")
	assert.Contains(t, output, "📝 Título: Test")
	assert.Contains(t, output, "📄 Descrição: Desc")
	assert.Contains(t, output, "📊 Status: ✅ Concluída")
	assert.Contains(t, output, fmt.Sprintf("📅 Criada em: %s", now.Format("02/01/2006 15:04")))
	assert.Contains(t, output, fmt.Sprintf("🔄 Atualizada em: %s", now.Format("02/01/2006 15:04")))
	mockUseCase.AssertExpectations(t)
}

func TestShouldShowTodoNotFound(t *testing.T) {
	// Arrange
	mockUseCase := new(application.MockTodoUseCase)
	cli := NewTodoCLI(mockUseCase)
	mockUseCase.On("GetTodoByID", "1").Return(nil, errors.New("not found"))

	cmd := cli.showCommand()
	cmd.SetArgs([]string{"1"})

	// Act
	output := captureOutput(func() {
		cmd.Execute()
	})

	// Assert
	assert.Contains(t, output, "❌ Tarefa não encontrada: not found")
	mockUseCase.AssertExpectations(t)
}

func TestShouldUpdateTodoSuccessfully(t *testing.T) {
	// Arrange
	mockUseCase := new(application.MockTodoUseCase)
	cli := NewTodoCLI(mockUseCase)
	todo := &entity.Todo{
		ID:          "1",
		Title:       "Updated",
		Description: "Desc",
	}
	mockUseCase.On("UpdateTodo", "1", "Updated", "Desc").Return(todo, nil)

	cmd := cli.updateCommand()
	cmd.SetArgs([]string{"1", "Updated", "Desc"})

	// Act
	output := captureOutput(func() {
		cmd.Execute()
	})

	// Assert
	assert.Contains(t, output, "✅ Tarefa atualizada com sucesso!")
	assert.Contains(t, output, "📝 Título: Updated")
	assert.Contains(t, output, "📄 Descrição: Desc")
	mockUseCase.AssertExpectations(t)
}

func TestShouldUpdateTodoWithError(t *testing.T) {
	// Arrange
	mockUseCase := new(application.MockTodoUseCase)
	cli := NewTodoCLI(mockUseCase)
	mockUseCase.On("UpdateTodo", "1", "Updated", "").Return(nil, errors.New("fail"))

	cmd := cli.updateCommand()
	cmd.SetArgs([]string{"1", "Updated"})

	// Act
	output := captureOutput(func() {
		cmd.Execute()
	})

	// Assert
	assert.Contains(t, output, "❌ Erro ao atualizar tarefa: fail")
	mockUseCase.AssertExpectations(t)
}

func TestShouldCompleteTodoSuccessfully(t *testing.T) {
	// Arrange
	mockUseCase := new(application.MockTodoUseCase)
	cli := NewTodoCLI(mockUseCase)
	todo := &entity.Todo{
		ID:    "1",
		Title: "Test",
	}
	mockUseCase.On("CompleteTodo", "1").Return(todo, nil)

	cmd := cli.completeCommand()
	cmd.SetArgs([]string{"1"})

	// Act
	output := captureOutput(func() {
		cmd.Execute()
	})

	// Assert
	assert.Contains(t, output, "✅ Tarefa 'Test' marcada como concluída!")
	mockUseCase.AssertExpectations(t)
}

func TestShouldCompleteTodoWithError(t *testing.T) {
	// Arrange
	mockUseCase := new(application.MockTodoUseCase)
	cli := NewTodoCLI(mockUseCase)
	mockUseCase.On("CompleteTodo", "1").Return(nil, errors.New("fail"))

	cmd := cli.completeCommand()
	cmd.SetArgs([]string{"1"})

	// Act
	output := captureOutput(func() {
		cmd.Execute()
	})

	// Assert
	assert.Contains(t, output, "❌ Erro ao completar tarefa: fail")
	mockUseCase.AssertExpectations(t)
}

func TestShouldDeleteTodoSuccessfully(t *testing.T) {
	// Arrange
	mockUseCase := new(application.MockTodoUseCase)
	cli := NewTodoCLI(mockUseCase)
	mockUseCase.On("DeleteTodo", "1").Return(nil)

	cmd := cli.deleteCommand()
	cmd.SetArgs([]string{"1"})

	// Act
	output := captureOutput(func() {
		cmd.Execute()
	})

	// Assert
	assert.Contains(t, output, "🗑️  Tarefa deletada com sucesso!")
	mockUseCase.AssertExpectations(t)
}

func TestShouldDeleteTodoWithError(t *testing.T) {
	// Arrange
	mockUseCase := new(application.MockTodoUseCase)
	cli := NewTodoCLI(mockUseCase)
	mockUseCase.On("DeleteTodo", "1").Return(errors.New("fail"))

	cmd := cli.deleteCommand()
	cmd.SetArgs([]string{"1"})

	// Act
	output := captureOutput(func() {
		cmd.Execute()
	})

	// Assert
	assert.Contains(t, output, "❌ Erro ao deletar tarefa: fail")
	mockUseCase.AssertExpectations(t)
}

func TestShouldGetRootCommandWithAllSubcommands(t *testing.T) {
	// Arrange
	mockUseCase := new(application.MockTodoUseCase)
	cli := NewTodoCLI(mockUseCase)

	// Act
	rootCmd := cli.GetRootCommand()

	// Assert
	assert.Equal(t, "todo", rootCmd.Use)
	subcommands := []string{"create", "list", "show", "update", "complete", "delete"}
	for _, sub := range subcommands {
		found := false
		for _, c := range rootCmd.Commands() {
			if strings.HasPrefix(c.Use, sub) {
				found = true
				break
			}
		}
		assert.True(t, found, "Subcommand %s not found", sub)
	}
}

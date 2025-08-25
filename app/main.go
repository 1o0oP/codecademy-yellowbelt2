package main

import (
	"codecademy-yellowbelt2/core/application"
	"codecademy-yellowbelt2/infrastructure/interface/cli"
	"codecademy-yellowbelt2/infrastructure/interface/repository"
	fileRepo "codecademy-yellowbelt2/infrastructure/repository"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Definir arquivo de dados
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Erro ao obter diretório home:", err)
	}

	dataFile := filepath.Join(homeDir, ".todo-cli", "todos.json")

	// Criar diretório se não existir
	if err := os.MkdirAll(filepath.Dir(dataFile), 0755); err != nil {
		log.Fatal("Erro ao criar diretório de dados:", err)
	}

	// Inicializar repository com arquivo JSON
	var todoRepo repository.ITodoRepository = fileRepo.NewFileTodoRepository(dataFile)

	// Inicializar use case
	todoUseCase := application.NewTodoUseCase(todoRepo)

	// Inicializar CLI
	todoCLI := cli.NewTodoCLI(todoUseCase)

	// Executar comando raiz
	rootCmd := todoCLI.GetRootCommand()
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

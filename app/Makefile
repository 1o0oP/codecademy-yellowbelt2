.PHONY: build run test clean help

# Variáveis
APP_NAME=todo
BUILD_DIR=bin

# Comandos principais
build: ## Compilar o projeto
	@echo "🔨 Compilando projeto..."
	go build -o $(BUILD_DIR)/$(APP_NAME) .

run: build ## Executar a aplicação
	@echo "🚀 Executando aplicação..."
	./$(BUILD_DIR)/$(APP_NAME)

install: ## Instalar dependências
	@echo "📦 Instalando dependências..."
	go mod tidy
	go mod download

test: ## Executar testes com coverage
	@echo "📊 Executando testes com coverage..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

clean: ## Limpar arquivos de build
	@echo "🧹 Limpando arquivos de build..."
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

# Comandos da aplicação
create: build ## Criar uma nova tarefa - uso: make create TITLE="título" DESC="descrição"
	./$(BUILD_DIR)/$(APP_NAME) create "$(TITLE)" "$(DESC)"

list: build ## Listar todas as tarefas
	./$(BUILD_DIR)/$(APP_NAME) list

show: build ## Mostrar tarefa específica - uso: make show ID="id-da-tarefa"
	./$(BUILD_DIR)/$(APP_NAME) show "$(ID)"

complete: build ## Completar tarefa - uso: make complete ID="id-da-tarefa"
	./$(BUILD_DIR)/$(APP_NAME) complete "$(ID)"

delete: build ## Deletar tarefa - uso: make delete ID="id-da-tarefa"
	./$(BUILD_DIR)/$(APP_NAME) delete "$(ID)"

help: ## Mostrar esta ajuda
	@echo "📋 Comandos disponíveis:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

# ⚡ Guia de Início Rápido

Este guia permite que você tenha o Todo List CLI funcionando em menos de 5 minutos.

## 🚀 Instalação Rápida

### Pré-requisitos
- **Go 1.23.5+** ([Download aqui](https://golang.org/dl/))
- **Git** (para clonar o repositório)
- **Make** (opcional, mas recomendado)

### Verificação do Ambiente
```bash
# Verifique se Go está instalado
go version
# Saída esperada: go version go1.23.5 linux/amd64

# Verifique se Make está disponível
make --version
# Se não tiver make, você pode usar comandos go diretamente
```

## 📥 Download e Configuração

### 1. Clone o Repositório
```bash
git clone https://github.com/1o0oP/codecademy-yellowbelt2.git
cd codecademy-yellowbelt2/app
```

### 2. Instale Dependências
```bash
# Usando Make (recomendado)
make install

# Ou manualmente
go mod tidy && go mod download
```

### 3. Compile a Aplicação
```bash
# Usando Make
make build

# Ou manualmente
go build -o bin/todo .
```

### 4. Verifique a Instalação
```bash
# Teste se está funcionando
./bin/todo --help

# Saída esperada:
# Uma ferramenta de linha de comando para gerenciar sua lista de tarefas
# 
# Usage:
#   todo [command]
# 
# Available Commands:
#   complete    Marcar uma tarefa como concluída
#   create      Criar uma nova tarefa
#   delete      Deletar uma tarefa
#   help        Help about any command
#   list        Listar todas as tarefas
#   show        Mostrar detalhes de uma tarefa específica
#   update      Atualizar uma tarefa existente
```

## 🎯 Primeiros Passos

### Workflow Básico de 2 Minutos

#### 1. Crie sua primeira tarefa
```bash
# Via Make (mais fácil)
make create TITLE="Aprender Go" DESC="Estudar Clean Architecture"

# Via CLI direta
./bin/todo create "Aprender Go" "Estudar Clean Architecture"
```

**Saída esperada:**
```
✅ Tarefa criada com sucesso!
ID: 550e8400-e29b-41d4-a716-446655440000
Título: Aprender Go
Descrição: Estudar Clean Architecture
```

#### 2. Liste suas tarefas
```bash
# Via Make
make list

# Via CLI
./bin/todo list
```

**Saída esperada:**
```
📋 Total de tarefas: 1

1. ⏳ Aprender Go
   📄 Estudar Clean Architecture
   🆔 ID: 550e8400-e29b-41d4-a716-446655440000
```

#### 3. Marque como concluída
```bash
# Copie o ID da etapa anterior
make complete ID="550e8400-e29b-41d4-a716-446655440000"

# Ou via CLI
./bin/todo complete "550e8400-e29b-41d4-a716-446655440000"
```

#### 4. Veja os detalhes
```bash
# Veja informações detalhadas
./bin/todo show "550e8400-e29b-41d4-a716-446655440000"
```

**Saída esperada:**
```
🆔 ID: 550e8400-e29b-41d4-a716-446655440000
📝 Título: Aprender Go
📄 Descrição: Estudar Clean Architecture
📊 Status: ✅ Concluída
📅 Criada em: 25/08/2025 14:30
🔄 Atualizada em: 25/08/2025 14:35
```

## 🛠 Comandos Essenciais

### Comandos via Make (Recomendado)

| Comando | Uso | Exemplo |
|---------|-----|---------|
| `make create` | `TITLE="..." DESC="..."` | `make create TITLE="Estudar" DESC="Go lang"` |
| `make list` | | `make list` |
| `make show` | `ID="..."` | `make show ID="abc123"` |
| `make complete` | `ID="..."` | `make complete ID="abc123"` |
| `make update` | `ID="..." TITLE="..." DESC="..."` | `make update ID="abc123" TITLE="Novo título"` |
| `make delete` | `ID="..."` | `make delete ID="abc123"` |

### Comandos Diretos da CLI

```bash
# Criar tarefa
./bin/todo create "Título" "Descrição opcional"

# Listar todas
./bin/todo list

# Ver detalhes
./bin/todo show <id>

# Atualizar
./bin/todo update <id> "Novo título" "Nova descrição"

# Marcar como concluída
./bin/todo complete <id>

# Deletar
./bin/todo delete <id>
```

## 📂 Onde os Dados São Armazenados

As tarefas são salvas em:
```
~/.todo-cli/todos.json
```

**Exemplo do arquivo:**
```json
{
  "550e8400-e29b-41d4-a716-446655440000": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "title": "Aprender Go",
    "description": "Estudar Clean Architecture",
    "completed": true,
    "created_at": "2025-08-25T14:30:00Z",
    "updated_at": "2025-08-25T14:35:00Z"
  }
}
```

## 🧪 Executar Testes

```bash
# Todos os testes
make test

# Com relatório de cobertura
make test-coverage

# Abrir relatório HTML (Linux/Mac)
xdg-open coverage.html
```

## 🎯 Cenários de Uso Comuns

### Gerenciamento Diário de Tarefas
```bash
# Manhã: criar tarefas do dia
make create TITLE="Revisar emails" DESC="Responder emails importantes"
make create TITLE="Reunião 10h" DESC="Discussão sobre projeto X"
make create TITLE="Código feature Y" DESC="Implementar nova funcionalidade"

# Durante o dia: completar tarefas
make list  # ver o que precisa fazer
make complete ID="..."  # marcar como feito

# Final do dia: revisar progresso
make list  # ver o que foi concluído
```

### Workflow de Desenvolvimento
```bash
# Planejamento
make create TITLE="Setup projeto" DESC="Configurar repositório Git"
make create TITLE="Implementar tests" DESC="TDD para feature X"
make create TITLE="Code review" DESC="Revisar PRs pendentes"

# Execução
make complete ID="..."  # conforme progride

# Revisão
make list  # status geral
```

## 🔧 Comandos de Desenvolvimento

### Build e Limpeza
```bash
# Build completo
make build

# Limpeza
make clean

# Reinstalar dependências
make install
```

### Makefile - Todos os Comandos
```bash
# Ver todos os comandos disponíveis
make help

# Saída:
# 📋 Comandos disponíveis:
#   build          Compilar o projeto
#   clean          Limpar arquivos de build
#   complete       Completar tarefa - uso: make complete ID="id-da-tarefa"
#   create         Criar uma nova tarefa - uso: make create TITLE="título" DESC="descrição"
#   delete         Deletar tarefa - uso: make delete ID="id-da-tarefa"
#   help           Mostrar esta ajuda
#   install        Instalar dependências
#   list           Listar todas as tarefas
#   run            Executar a aplicação
#   show           Mostrar tarefa específica - uso: make show ID="id-da-tarefa"
#   test           Executar testes com coverage
```

## ❓ Solução de Problemas Comuns

### Erro: "command not found: todo"
```bash
# Certifique-se de estar no diretório correto
cd codecademy-yellowbelt2/app

# E use o caminho completo
./bin/todo --help
```

### Erro: "no such file or directory: bin/todo"
```bash
# Compile novamente
make build

# Ou manualmente
go build -o bin/todo .
```

### Erro: Go não encontrado
```bash
# Instale Go 1.23.5+
# Ubuntu/Debian:
sudo apt update && sudo apt install golang-go

# macOS:
brew install go

# Ou baixe de: https://golang.org/dl/
```

### Problemas de Permissão
```bash
# Torne o executável executável
chmod +x ./bin/todo
```

## 🎉 Próximos Passos

Agora que você tem tudo funcionando:

1. 📖 [**Leia o Guia de Uso Completo**](usage-guide.md) - Exemplos avançados
2. 🏗 [**Entenda a Arquitetura**](architecture.md) - Como o código está organizado  
3. 🧪 [**Explore os Testes**](testing.md) - Como garantimos qualidade
4. 🤖 [**IA no Desenvolvimento**](ai-development.md) - Como a IA foi utilizada

## 💡 Dicas Finais

- **Use IDs curtos**: Você pode usar apenas os primeiros 8 caracteres do ID
- **Atalhos Make**: Os comandos make são mais fáceis que CLI direta
- **Backup**: O arquivo `~/.todo-cli/todos.json` contém todos seus dados
- **Performance**: A aplicação é rápida, mesmo com centenas de tarefas

---

**🚀 Você está pronto para ser mais produtivo com o Todo List CLI!**

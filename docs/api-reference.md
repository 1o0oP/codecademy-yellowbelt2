# 🔧 API Reference

Documentação técnica completa de todas as interfaces, tipos e métodos do Todo List CLI.

## 📦 Estrutura de Módulos

```
codecademy-yellowbelt2/
├── core/
│   ├── domain/entity/           # Entidades de domínio
│   └── application/            # Casos de uso
├── infrastructure/
│   ├── interface/              # Contratos/Interfaces
│   │   ├── repository/         # Repository interfaces
│   │   └── application/        # Use case interfaces
│   ├── repository/             # Implementações de persistência
│   └── interface/cli/          # Interface de linha de comando
└── main.go                     # Entry point
```

## 🏛 Domain Layer

### `entity.Todo`

Representa uma tarefa no sistema.

#### Definição
```go
type Todo struct {
    ID          string    `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Completed   bool      `json:"completed"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

#### Campos

| Campo | Tipo | Descrição | Validação |
|-------|------|-----------|-----------|
| `ID` | `string` | Identificador único UUID | Gerado automaticamente |
| `Title` | `string` | Título da tarefa | Obrigatório |
| `Description` | `string` | Descrição opcional | Opcional |
| `Completed` | `bool` | Status de conclusão | Default: `false` |
| `CreatedAt` | `time.Time` | Timestamp de criação | Automático |
| `UpdatedAt` | `time.Time` | Timestamp de atualização | Atualizado automaticamente |

#### Construtores

##### `NewTodo`
```go
func NewTodo(title, description string) *Todo
```

**Parâmetros:**
- `title` (`string`): Título da tarefa
- `description` (`string`): Descrição opcional

**Retorno:**
- `*Todo`: Nova instância de Todo

**Exemplo:**
```go
todo := entity.NewTodo("Estudar Go", "Focar em interfaces")
fmt.Println(todo.ID)        // "550e8400-e29b-41d4-a716-446655440000"
fmt.Println(todo.Title)     // "Estudar Go"
fmt.Println(todo.Completed) // false
```

#### Métodos de Negócio

##### `MarkAsCompleted`
```go
func (t *Todo) MarkAsCompleted()
```

Marca a tarefa como concluída e atualiza o timestamp.

**Side Effects:**
- `t.Completed = true`
- `t.UpdatedAt = time.Now()`

**Exemplo:**
```go
todo := entity.NewTodo("Task", "Description")
todo.MarkAsCompleted()
fmt.Println(todo.Completed) // true
```

##### `MarkAsIncomplete`
```go
func (t *Todo) MarkAsIncomplete()
```

Marca a tarefa como não concluída e atualiza o timestamp.

**Side Effects:**
- `t.Completed = false`
- `t.UpdatedAt = time.Now()`

##### `Update`
```go
func (t *Todo) Update(title, description string)
```

Atualiza título e/ou descrição da tarefa.

**Parâmetros:**
- `title` (`string`): Novo título (vazio = não altera)
- `description` (`string`): Nova descrição (vazio = não altera)

**Side Effects:**
- Atualiza campos se não vazios
- `t.UpdatedAt = time.Now()`

**Exemplo:**
```go
todo.Update("Novo Título", "Nova Descrição")
todo.Update("Só Título", "")  // Só altera título
todo.Update("", "Só Desc")    // Só altera descrição
```

## 🔌 Repository Interface

### `repository.ITodoRepository`

Interface para persistência de todos.

#### Definição
```go
type ITodoRepository interface {
    Create(todo *entity.Todo) error
    GetByID(id string) (*entity.Todo, error)
    GetAll() ([]*entity.Todo, error)
    Update(todo *entity.Todo) error
    Delete(id string) error
}
```

#### Métodos

##### `Create`
```go
Create(todo *entity.Todo) error
```

**Parâmetros:**
- `todo` (`*entity.Todo`): Instância da tarefa a criar

**Retorno:**
- `error`: `nil` em sucesso, erro em falha

**Possíveis Erros:**
- Erro de I/O (FileRepository)
- Violação de constraint (implementação específica)

##### `GetByID`
```go
GetByID(id string) (*entity.Todo, error)
```

**Parâmetros:**
- `id` (`string`): ID único da tarefa

**Retorno:**
- `*entity.Todo`: Instância encontrada
- `error`: `nil` em sucesso, `"todo not found"` se não existe

##### `GetAll`
```go
GetAll() ([]*entity.Todo, error)
```

**Retorno:**
- `[]*entity.Todo`: Lista de todas as tarefas
- `error`: `nil` em sucesso, erro em falha de I/O

##### `Update`
```go
Update(todo *entity.Todo) error
```

**Parâmetros:**
- `todo` (`*entity.Todo`): Instância atualizada

**Retorno:**
- `error`: `nil` em sucesso, `"todo not found"` se não existe

**Pré-condição:**
- Todo deve existir no repositório

##### `Delete`
```go
Delete(id string) error
```

**Parâmetros:**
- `id` (`string`): ID da tarefa a deletar

**Retorno:**
- `error`: `nil` em sucesso, `"todo not found"` se não existe

## 💾 Repository Implementations

### `InMemoryTodoRepository`

Implementação em memória para testes e desenvolvimento.

#### Definição
```go
type InMemoryTodoRepository struct {
    todos map[string]*entity.Todo
    mutex sync.RWMutex
}
```

#### Construtor
```go
func NewInMemoryTodoRepository() repository.ITodoRepository
```

#### Características
- ✅ **Thread-Safe**: Usa `sync.RWMutex`
- ✅ **Performance**: Acesso O(1) por ID
- ❌ **Persistência**: Dados perdidos ao reiniciar
- ✅ **Ideal para**: Testes unitários

### `FileTodoRepository`

Implementação com persistência em arquivo JSON.

#### Definição
```go
type FileTodoRepository struct {
    filename string
    mutex    sync.RWMutex
}
```

#### Construtor
```go
func NewFileTodoRepository(filename string) repository.ITodoRepository
```

**Parâmetros:**
- `filename` (`string`): Caminho para arquivo JSON

#### Características
- ✅ **Persistente**: Dados mantidos entre execuções
- ✅ **Thread-Safe**: Usa `sync.RWMutex`
- ✅ **Auto-Create**: Cria arquivo/diretório se não existir
- ✅ **JSON Format**: Formato legível e editável
- ⚠️ **Performance**: I/O em cada operação

#### Formato do Arquivo
```json
{
  "550e8400-e29b-41d4-a716-446655440000": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "title": "Estudar Go",
    "description": "Clean Architecture",
    "completed": false,
    "created_at": "2025-08-25T14:30:00Z",
    "updated_at": "2025-08-25T14:30:00Z"
  }
}
```

#### Métodos Internos

##### `load`
```go
func (r *FileTodoRepository) load() (map[string]*entity.Todo, error)
```

Carrega todos do arquivo JSON.

**Comportamento:**
- Retorna mapa vazio se arquivo não existir
- Trata arquivo vazio como válido
- Retorna erro em JSON inválido

##### `save`
```go
func (r *FileTodoRepository) save(todos map[string]*entity.Todo) error
```

Salva todos no arquivo JSON com indentação.

## 🎯 Application Layer

### `application.ITodoUseCase`

Interface para casos de uso de negócio.

#### Definição
```go
type ITodoUseCase interface {
    CreateTodo(title, description string) (*entity.Todo, error)
    GetTodoByID(id string) (*entity.Todo, error)
    GetAllTodos() ([]*entity.Todo, error)
    UpdateTodo(id, title, description string) (*entity.Todo, error)
    CompleteTodo(id string) (*entity.Todo, error)
    DeleteTodo(id string) error
}
```

#### Métodos

##### `CreateTodo`
```go
CreateTodo(title, description string) (*entity.Todo, error)
```

**Fluxo:**
1. Cria nova instância `entity.Todo`
2. Persiste via repository
3. Retorna instância criada

**Exemplo:**
```go
todo, err := useCase.CreateTodo("Estudar Go", "Clean Architecture")
if err != nil {
    log.Fatal(err)
}
fmt.Println(todo.ID) // UUID gerado
```

##### `GetTodoByID`
```go
GetTodoByID(id string) (*entity.Todo, error)
```

**Delega diretamente** para `repository.GetByID()`.

##### `GetAllTodos`
```go
GetAllTodos() ([]*entity.Todo, error)
```

**Delega diretamente** para `repository.GetAll()`.

##### `UpdateTodo`
```go
UpdateTodo(id, title, description string) (*entity.Todo, error)
```

**Fluxo:**
1. Busca todo existente por ID
2. Aplica método `Update()` da entidade
3. Persiste alterações via repository
4. Retorna instância atualizada

##### `CompleteTodo`
```go
CompleteTodo(id string) (*entity.Todo, error)
```

**Fluxo:**
1. Busca todo existente por ID
2. Aplica método `MarkAsCompleted()` da entidade
3. Persiste alterações via repository
4. Retorna instância atualizada

##### `DeleteTodo`
```go
DeleteTodo(id string) error
```

**Delega diretamente** para `repository.Delete()`.

### `TodoUseCase`

Implementação concreta dos casos de uso.

#### Definição
```go
type TodoUseCase struct {
    todoRepo repository.ITodoRepository
}
```

#### Construtor
```go
func NewTodoUseCase(todoRepo repository.ITodoRepository) application.ITodoUseCase
```

**Parâmetros:**
- `todoRepo` (`repository.ITodoRepository`): Implementação de persistência

**Retorno:**
- `application.ITodoUseCase`: Interface implementada

## 🖥 CLI Interface

### `cli.TodoCLI`

Interface de linha de comando usando Cobra framework.

#### Definição
```go
type TodoCLI struct {
    todoUseCase app_interfaces.ITodoUseCase
}
```

#### Construtor
```go
func NewTodoCLI(todoUseCase app_interfaces.ITodoUseCase) *TodoCLI
```

#### Comandos

##### Root Command
```go
func (cli *TodoCLI) GetRootCommand() *cobra.Command
```

Retorna comando raiz com todos os subcomandos.

**Uso:**
```bash
todo [command]
```

**Subcomandos:**
- `create` - Criar tarefa
- `list` - Listar tarefas
- `show` - Mostrar detalhes
- `update` - Atualizar tarefa
- `complete` - Marcar como concluída
- `delete` - Deletar tarefa

##### Create Command
```bash
todo create [title] [description]
```

**Argumentos:**
- `title` (obrigatório): Título da tarefa
- `description` (opcional): Descrição da tarefa

**Saída de Sucesso:**
```
✅ Tarefa criada com sucesso!
ID: 550e8400-e29b-41d4-a716-446655440000
Título: Estudar Go
Descrição: Clean Architecture
```

##### List Command
```bash
todo list
```

**Saída de Exemplo:**
```
📋 Total de tarefas: 2

1. ⏳ Estudar Go
   📄 Clean Architecture
   🆔 ID: 550e8400-e29b-41d4-a716-446655440000

2. ✅ Fazer exercícios
   🆔 ID: 661f9511-f3ac-52e5-b827-557766551111
```

##### Show Command
```bash
todo show [id]
```

**Argumentos:**
- `id` (obrigatório): ID da tarefa

**Saída de Exemplo:**
```
🆔 ID: 550e8400-e29b-41d4-a716-446655440000
📝 Título: Estudar Go
📄 Descrição: Clean Architecture
📊 Status: ⏳ Pendente
📅 Criada em: 25/08/2025 14:30
🔄 Atualizada em: 25/08/2025 14:30
```

##### Update Command
```bash
todo update [id] [title] [description]
```

**Argumentos:**
- `id` (obrigatório): ID da tarefa
- `title` (obrigatório): Novo título
- `description` (opcional): Nova descrição

##### Complete Command
```bash
todo complete [id]
```

**Argumentos:**
- `id` (obrigatório): ID da tarefa

**Saída de Exemplo:**
```
✅ Tarefa 'Estudar Go' marcada como concluída!
```

##### Delete Command
```bash
todo delete [id]
```

**Argumentos:**
- `id` (obrigatório): ID da tarefa

**Saída de Exemplo:**
```
🗑️  Tarefa deletada com sucesso!
```

## 🔧 Main Entry Point

### `main.main`

Ponto de entrada da aplicação com configuração de dependências.

#### Fluxo de Execução
```go
func main() {
    // 1. Configuração do arquivo de dados
    homeDir, err := os.UserHomeDir()
    dataFile := filepath.Join(homeDir, ".todo-cli", "todos.json")
    
    // 2. Criação de diretório
    os.MkdirAll(filepath.Dir(dataFile), 0755)
    
    // 3. Injeção de dependências
    var todoRepo repository.ITodoRepository = fileRepo.NewFileTodoRepository(dataFile)
    todoUseCase := application.NewTodoUseCase(todoRepo)
    todoCLI := cli.NewTodoCLI(todoUseCase)
    
    // 4. Execução do comando
    rootCmd := todoCLI.GetRootCommand()
    rootCmd.Execute()
}
```

#### Configuração de Dados
- **Diretório**: `~/.todo-cli/`
- **Arquivo**: `todos.json`
- **Permissões**: `0755` para diretório, `0644` para arquivo

## 🧪 Test Doubles

### Mocks

#### `MockTodoRepository`
```go
type MockTodoRepository struct {
    mock.Mock
}

func (m *MockTodoRepository) Create(todo *entity.Todo) error {
    args := m.Called(todo)
    return args.Error(0)
}
```

**Uso em Testes:**
```go
mockRepo := new(MockTodoRepository)
mockRepo.On("Create", mock.AnythingOfType("*entity.Todo")).Return(nil)
mockRepo.AssertExpectations(t)
```

#### `MockTodoUseCase`
```go
type MockTodoUseCase struct {
    mock.Mock
}

func (m *MockTodoUseCase) CreateTodo(title, description string) (*entity.Todo, error) {
    args := m.Called(title, description)
    todo, _ := args.Get(0).(*entity.Todo)
    return todo, args.Error(1)
}
```

## 🔐 Thread Safety

### Concurrent Access Patterns

#### Repository Level
```go
// Leitura (múltiplas goroutines OK)
r.mutex.RLock()
defer r.mutex.RUnlock()
// operação de leitura

// Escrita (exclusiva)
r.mutex.Lock()
defer r.mutex.Unlock()
// operação de escrita
```

#### Application Level
- **Use Cases**: Stateless, thread-safe por design
- **Entity Methods**: Operam em instâncias isoladas
- **CLI**: Single-threaded por natureza

## 📊 Performance Characteristics

### Repository Performance

| Operação | InMemory | File |
|----------|----------|------|
| `Create` | O(1) | O(n) |
| `GetByID` | O(1) | O(n) |
| `GetAll` | O(n) | O(n) |
| `Update` | O(1) | O(n) |
| `Delete` | O(1) | O(n) |

### Memory Usage
- **InMemory**: ~100 bytes por todo + overhead do map
- **File**: ~50 bytes por todo (apenas durante operação)

### I/O Patterns
- **File Repository**: Load/Save completo a cada operação
- **Otimização Futura**: Caching layer, lazy loading

## 🚨 Error Handling

### Error Types

#### Repository Errors
```go
var (
    ErrTodoNotFound = errors.New("todo not found")
    ErrFileAccess   = errors.New("file access error")
    ErrInvalidJSON  = errors.New("invalid JSON format")
)
```

#### Use Case Errors
- **Validation**: Títulos vazios, IDs inválidos
- **Business Logic**: Regras de negócio violadas
- **Persistence**: Erros propagados do repository

#### CLI Errors
- **Argument Parsing**: Cobra validation
- **Command Execution**: Erros de use case propagados
- **User Display**: Mensagens amigáveis com emojis

### Error Propagation Chain

```
CLI Command → Use Case → Repository → File System
     ↑            ↑           ↑           ↓
User Message ← Business ← Persistence ← I/O Error
```

## 📚 Dependencies

### External Libraries
```go
// go.mod
require (
    github.com/google/uuid v1.6.0        // UUID generation
    github.com/spf13/cobra v1.8.0        // CLI framework
    github.com/stretchr/testify v1.11.0  // Testing utilities
    bou.ke/monkey v1.0.2                 // Monkey patching (tests only)
)
```

### Standard Library Usage
- `encoding/json`: Serialização
- `os`: File system operations
- `sync`: Concurrent access control
- `time`: Timestamps
- `path/filepath`: Path manipulation

---

## 🎯 Quick Reference

### Key Types
- `entity.Todo`: Core business entity
- `repository.ITodoRepository`: Persistence interface
- `application.ITodoUseCase`: Business logic interface
- `cli.TodoCLI`: Command-line interface

### Key Functions
- `entity.NewTodo()`: Create new todo
- `fileRepo.NewFileTodoRepository()`: Create file-based repo
- `application.NewTodoUseCase()`: Create use case instance
- `cli.NewTodoCLI()`: Create CLI instance

### Common Patterns
- **Dependency Injection**: Interfaces injected into constructors
- **Error Handling**: Explicit error returns, user-friendly messages
- **Thread Safety**: RWMutex for concurrent access
- **Testing**: Mocks for external dependencies

---

**📖 Esta documentação serve como referência completa para desenvolvedores que querem entender, modificar ou estender o Todo List CLI.**

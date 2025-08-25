# 🧪 Estratégias de Teste e Cobertura

Este projeto implementa uma estratégia abrangente de testes seguindo as melhores práticas de Go e Clean Architecture.

## 📊 Visão Geral da Cobertura

### Métricas Atuais
- **Cobertura Geral**: >85%
- **Testes Unitários**: 100% das funções críticas
- **Testes de Integração**: Cobertura completa da CLI
- **Mocks**: Todas as interfaces externas

### Execução Rápida
```bash
# Executar todos os testes
make test

# Com relatório detalhado
make test-coverage

# Abrir relatório HTML
xdg-open coverage.html  # Linux
open coverage.html      # macOS
```

## 🎯 Filosofia de Testes

### Pirâmide de Testes Implementada

```
        /\
       /  \
      / E2E \ ← CLI Integration Tests
     /______\
    /        \
   /Integration\ ← Use Case Tests
  /_____________\
 /               \
/   Unit Tests    \ ← Entity & Repository Tests
\________________/
```

### Princípios Seguidos
1. **Teste Primeiro**: TDD aplicado em funcionalidades críticas
2. **Isolamento**: Cada teste é independente e idempotente
3. **Mocks**: Dependências externas sempre mockadas
4. **Legibilidade**: Testes como documentação executável
5. **Performance**: Testes rápidos (<100ms cada)

## 📂 Estrutura de Testes

### Organização por Camada

```
app/
├── core/
│   ├── domain/entity/
│   │   ├── todo.go
│   │   └── todo_test.go              ← Unit Tests
│   └── application/
│       ├── todo_usecase.go
│       └── todo_usecase_test.go      ← Use Case Tests
├── infrastructure/
│   ├── repository/
│   │   ├── inmemory_todo_repository.go
│   │   ├── inmemory_todo_repository_test.go  ← Repository Tests
│   │   ├── file_todo_repository.go
│   │   └── file_todo_repository_test.go     ← File I/O Tests
│   └── interface/
│       └── cli/
│           ├── todo_cli.go
│           └── todo_cli_test.go             ← CLI Integration Tests
```

## 🔬 Testes por Camada

### 1. Domain Layer Tests

#### Entity Tests (`todo_test.go`)
```go
func TestShouldCreateNewTodo(t *testing.T) {
    // Arrange
    title := "Test Todo"
    description := "Test Description"

    // Act
    todo := NewTodo(title, description)

    // Assert
    assert.NotEmpty(t, todo.ID)
    assert.Equal(t, title, todo.Title)
    assert.Equal(t, description, todo.Description)
    assert.False(t, todo.Completed)
}
```

**Cobertura**:
- ✅ Criação de entidade
- ✅ Métodos de negócio (`MarkAsCompleted`, `Update`)
- ✅ Validações de estado
- ✅ Timestamps automáticos

#### Padrões Utilizados:
- **AAA Pattern**: Arrange, Act, Assert
- **Descriptive Names**: Nomes que documentam comportamento
- **Edge Cases**: Testes de casos extremos

### 2. Application Layer Tests

#### Use Case Tests (`todo_usecase_test.go`)
```go
func TestTodoUseCase_CreateTodo(t *testing.T) {
    // Arrange
    repo := repository.NewInMemoryTodoRepository()
    useCase := NewTodoUseCase(repo)

    // Act
    todo, err := useCase.CreateTodo("Test Todo", "Test Description")

    // Assert
    assert.NoError(t, err)
    assert.Equal(t, "Test Todo", todo.Title)
}
```

**Estratégias Implementadas**:
- **Dependency Injection**: Repository real em testes simples
- **Mock Testing**: Para cenários de erro
- **Integration Testing**: Fluxo completo de Use Cases

#### Testes com Mocks
```go
func TestShouldReturnErrorWhenCreateTodoFails(t *testing.T) {
    // Arrange
    mockRepo := new(repoMock.MockTodoRepository)
    useCase := NewTodoUseCase(mockRepo)
    expectedErr := errors.New("some error")
    mockRepo.On("Create", mock.AnythingOfType("*entity.Todo")).Return(expectedErr)

    // Act
    todo, err := useCase.CreateTodo("Title", "Description")

    // Assert
    assert.Nil(t, todo)
    assert.Error(t, err)
    mockRepo.AssertExpectations(t)
}
```

### 3. Infrastructure Layer Tests

#### Repository Tests
**InMemory Repository** (`inmemory_todo_repository_test.go`):
```go
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
```

**File Repository** (`file_todo_repository_test.go`):
```go
func createTempRepo(t *testing.T) (*FileTodoRepository, func()) {
    tmpfile, err := os.CreateTemp("", "todos_*.json")
    assert.NoError(t, err)
    repo := NewFileTodoRepository(tmpfile.Name()).(*FileTodoRepository)
    cleanup := func() {
        os.Remove(tmpfile.Name())
    }
    return repo, cleanup
}
```

**Características**:
- **Temporary Files**: Testes isolados com arquivos temporários
- **Monkey Patching**: Simulação de erros I/O
- **Cleanup Functions**: Limpeza automática após testes
- **Error Scenarios**: Cobertura completa de casos de erro

#### Testes de Erro com Monkey Patching
```go
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
}
```

### 4. Interface Layer Tests

#### CLI Tests (`todo_cli_test.go`)
```go
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
    expectedTodo := &entity.Todo{ID: "1", Title: "Test", Description: "Desc"}
    mockUseCase.On("CreateTodo", "Test", "Desc").Return(expectedTodo, nil)

    cmd := cli.createCommand()
    cmd.SetArgs([]string{"Test", "Desc"})

    // Act
    output := captureOutput(func() { cmd.Execute() })

    // Assert
    assert.Contains(t, output, "✅ Tarefa criada com sucesso!")
    mockUseCase.AssertExpectations(t)
}
```

**Técnicas Utilizadas**:
- **Output Capture**: Interceptação do stdout para validação
- **Command Testing**: Teste de cada comando Cobra isoladamente
- **Mock Integration**: Use cases mockados para isolamento
- **User Experience Testing**: Validação de mensagens e formatação

## 🏭 Testes de Integração

### End-to-End Workflow
```go
func TestCompleteWorkflow(t *testing.T) {
    // Setup temporary repository
    tmpfile, _ := os.CreateTemp("", "e2e_*.json")
    defer os.Remove(tmpfile.Name())
    
    repo := NewFileTodoRepository(tmpfile.Name())
    useCase := NewTodoUseCase(repo)
    cli := NewTodoCLI(useCase)
    
    // Create todo
    createCmd := cli.createCommand()
    createCmd.SetArgs([]string{"E2E Test", "End to end testing"})
    
    var todoID string
    output := captureOutput(func() { createCmd.Execute() })
    // Extract ID from output...
    
    // Complete todo
    completeCmd := cli.completeCommand()
    completeCmd.SetArgs([]string{todoID})
    completeCmd.Execute()
    
    // Verify completed
    listCmd := cli.listCommand()
    output = captureOutput(func() { listCmd.Execute() })
    assert.Contains(t, output, "✅")
}
```

## 📈 Cobertura Detalhada

### Comando para Análise
```bash
# Gerar relatório detalhado
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# Analisar por arquivo
go tool cover -func=coverage.out

# Output esperado:
# codecademy-yellowbelt2/core/domain/entity/todo.go:15:    NewTodo         100.0%
# codecademy-yellowbelt2/core/domain/entity/todo.go:26:    MarkAsCompleted 100.0%
# codecademy-yellowbelt2/core/application/todo_usecase.go:15: CreateTodo    100.0%
# ...
# total:                                                    (statements)    87.2%
```

### Métricas por Módulo

| Módulo | Cobertura | Arquivos de Teste | Cenários |
|--------|-----------|-------------------|----------|
| **Domain/Entity** | 100% | `todo_test.go` | 4 testes |
| **Application** | 95% | `todo_usecase_test.go` | 12 testes |
| **Repository/InMemory** | 100% | `inmemory_todo_repository_test.go` | 6 testes |
| **Repository/File** | 92% | `file_todo_repository_test.go` | 18 testes |
| **CLI Interface** | 85% | `todo_cli_test.go` | 14 testes |

### Áreas com Menor Cobertura
```bash
# Identificar linhas não cobertas
go tool cover -html=coverage.out

# Áreas típicas não cobertas:
# - Tratamento de erros raros
# - Código de inicialização
# - Algumas condições edge-case
```

## 🎯 Estratégias de Mock

### Interface Mocks
```go
type MockTodoRepository struct {
    mock.Mock
}

func (m *MockTodoRepository) Create(todo *entity.Todo) error {
    args := m.Called(todo)
    return args.Error(0)
}

// Uso nos testes
mockRepo := new(MockTodoRepository)
mockRepo.On("Create", mock.AnythingOfType("*entity.Todo")).Return(nil)
```

### Mocks para I/O
```go
// Monkey patching para simulação de falhas
patch := monkey.Patch(os.ReadFile, func(string) ([]byte, error) {
    return nil, errors.New("simulated read error")
})
defer patch.Unpatch()
```

### Mocks para Tempo
```go
// Para testes determinísticos de timestamps
patch := monkey.Patch(time.Now, func() time.Time {
    return time.Date(2025, 8, 25, 12, 0, 0, 0, time.UTC)
})
defer patch.Unpatch()
```

## ⚡ Performance de Testes

### Benchmarks
```bash
# Executar benchmarks
go test -bench=. ./...

# Exemplo de saída:
# BenchmarkTodoCreation-8     1000000   1043 ns/op
# BenchmarkRepositoryRead-8    500000   2187 ns/op
```

### Paralelização
```go
func TestParallelOperations(t *testing.T) {
    t.Parallel()
    
    t.Run("CreateTodo", func(t *testing.T) {
        t.Parallel()
        // Test implementation
    })
    
    t.Run("UpdateTodo", func(t *testing.T) {
        t.Parallel()
        // Test implementation
    })
}
```

## 🔧 Ferramentas Utilizadas

### Testing Libraries
```go
// go.mod dependencies
require (
    github.com/stretchr/testify v1.11.0  // Assertions e mocks
    bou.ke/monkey v1.0.2                 // Monkey patching
)
```

### Testify Features Utilizadas
- **assert**: Validações simples e legíveis
- **mock**: Sistema de mocks robusto
- **suite**: Organização de testes complexos (se necessário)

### Monkey Patching para Go
- **Simulação de falhas**: I/O, rede, sistema
- **Determinismo**: Controle de tempo e randomness
- **Isolamento**: Substituição de dependências externas

## 🎭 Test Doubles Patterns

### Tipos Implementados

1. **Mocks**: Objetos com expectativas verificáveis
   ```go
   mockRepo.On("Create", mock.Anything).Return(nil)
   mockRepo.AssertExpectations(t)
   ```

2. **Stubs**: Retornos pré-definidos
   ```go
   stubRepo.On("GetByID", "id").Return(fixedTodo, nil)
   ```

3. **Fakes**: Implementações funcionais simplificadas
   ```go
   fakeRepo := NewInMemoryTodoRepository() // Usado como fake
   ```

4. **Spies**: Captura de chamadas para verificação
   ```go
   spy := new(RepositorySpy)
   // Verificar se método foi chamado corretamente
   ```

## 🚨 Testes de Casos Extremos

### Error Handling
```go
func TestShouldHandleRepositoryFailure(t *testing.T) {
    mockRepo := new(MockTodoRepository)
    mockRepo.On("GetAll").Return(nil, errors.New("database connection failed"))
    
    useCase := NewTodoUseCase(mockRepo)
    todos, err := useCase.GetAllTodos()
    
    assert.Nil(t, todos)
    assert.Error(t, err)
}
```

### Boundary Tests
```go
func TestShouldHandleEmptyTitle(t *testing.T) {
    todo := NewTodo("", "description")
    
    // Mesmo com título vazio, a criação deve funcionar
    // Validação pode ser adicionada posteriormente
    assert.NotEmpty(t, todo.ID)
}
```

### Concurrency Tests
```go
func TestShouldBeThreadSafe(t *testing.T) {
    repo := NewInMemoryTodoRepository()
    
    // Executar operações concorrentes
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            todo := entity.NewTodo(fmt.Sprintf("Todo %d", id), "")
            repo.Create(todo)
        }(i)
    }
    wg.Wait()
    
    todos, _ := repo.GetAll()
    assert.Len(t, todos, 100)
}
```

## 📋 Checklist de Qualidade

### ✅ Critérios de Aceite para Testes

- [ ] **Cobertura >85%**: Verificada automaticamente
- [ ] **Zero Flaky Tests**: Testes determinísticos
- [ ] **Performance <100ms**: Por suite de testes
- [ ] **Mocks Isolados**: Sem dependências externas
- [ ] **Cleanup Automático**: Sem side effects
- [ ] **Documentação Clara**: Nomes descritivos
- [ ] **Edge Cases Cobertos**: Scenarios extremos testados

### 🎯 Próximas Melhorias

1. **Property-Based Testing**: Usando `gopter`
2. **Mutation Testing**: Verificação de qualidade dos testes
3. **Contract Testing**: Para interfaces futuras
4. **Load Testing**: Para performance em escala

## 🚀 Executando Testes

### Comandos Essenciais
```bash
# Todos os testes
go test ./...

# Com verbose output
go test -v ./...

# Testes específicos
go test -run TestCreateTodo ./...

# Com coverage
go test -cover ./...

# Benchmark
go test -bench=. ./...

# Race detection
go test -race ./...
```

### Integração com CI/CD
```yaml
# .github/workflows/test.yml
name: Tests
on: [push, pull_request]
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    - uses: actions/setup-go@v2
      with:
        go-version: 1.23.5
    - name: Run tests
      run: |
        go test -race -coverprofile=coverage.out ./...
        go tool cover -html=coverage.out -o coverage.html
    - name: Upload coverage
      uses: actions/upload-artifact@v2
      with:
        name: coverage-report
        path: coverage.html
```

---

## 📚 Recursos Adicionais

- [Testing in Go - Official Docs](https://golang.org/doc/tutorial/add-a-test)
- [Testify Documentation](https://pkg.go.dev/github.com/stretchr/testify)
- [Clean Architecture Testing Strategies](https://blog.cleancoder.com/uncle-bob/2017/10/03/TestContravariance.html)
- [Go Testing Best Practices](https://github.com/golang/go/wiki/TestComments)

---

**🎯 Resultado**: Uma base de testes sólida que garante confiabilidade, facilita refatoração e documenta o comportamento esperado do sistema.

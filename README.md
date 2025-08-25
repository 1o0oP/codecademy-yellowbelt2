# Todo List CLI - Codecademy Yellow Belt 2

## 🎯 Sobre o Projeto

Este é um projeto desenvolvido para o curso **Codecademy Yellow Belt 2**, especificamente para o exercício "Create a Program Using Generative AI". O objetivo é demonstrar o uso responsável e efetivo de IA Generativa no desenvolvimento de software.

### Aplicação Desenvolvida
**Todo List CLI** - Uma ferramenta de linha de comando para gerenciamento de tarefas pessoais com operações CRUD completas.

## 📋 Tarefas do Curso

### ✅ 1. Come Up With the Application
**Aplicação Escolhida**: Sistema de gerenciamento de tarefas (Todo List) via CLI

**Justificativa**:
- Complexidade suficiente com múltiplos componentes (CRUD, persistência, interface)
- Múltiplas etapas de desenvolvimento (entidades → repositórios → casos de uso → interface)
- Valor prático real para uso diário
- Estrutura que permite desenvolvimento incremental ideal para IA

### ✅ 2. Consider Risks and Limitations
**Riscos Identificados**:
- **Código Incorreto**: IA pode gerar lógica com bugs
- **Over-reliance**: Dependência excessiva da IA
- **Falta de Contexto**: IA não conhece todo o projeto

**Mitigações Aplicadas**:
- Testes unitários para validar código gerado
- Revisão manual de todo código sugerido pela IA
- Desenvolvimento incremental com validação constante
- Documentação clara do contexto para a IA

### 🔄 3. Provide Context
**Contexto Fornecido à IA**:
- **Linguagem**: Go 1.23.5
- **Arquitetura**: Clean Architecture
- **Tipo**: Aplicação CLI (não web/API)
- **Persistência**: In-memory para simplicidade
- **Estrutura**: Separação clara de responsabilidades
- **Testes**: Testes unitários obrigatórios ao lado do código
- **Framework CLI**: Cobra para interface de linha de comando

### 🔄 4. Choose A Tool
**Ferramenta Escolhida**: GitHub Copilot

**Justificativa**:
- Integração nativa com VS Code
- Sugestões contextuais durante desenvolvimento
- Boa para geração de código Go
- Suporte excelente para testes unitários

### 🔄 5. Consider Technology Suggestions
**Consulta à IA**: "Preciso criar uma CLI para gerenciamento de tarefas em Go. Que bibliotecas recomendam?"

**Sugestões Recebidas e Adotadas**:
- **Cobra**: Framework para CLI (adotado)
- **UUID**: Geração de IDs únicos (adotado)
- **Clean Architecture**: Estrutura de projeto (adotado)
- **Table-driven tests**: Padrão de testes Go (aplicado)

### 🔄 6. Start Small
**Primeira Implementação**: Entidade Todo básica

```go
type Todo struct {
    ID          string
    Title       string
    Description string
    Completed   bool
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

**Teste Imediato**: Validação da criação de todos e métodos básicos.

### 🔄 7. Continue Development
**Desenvolvimento Incremental**:
1. **Entidade Todo** → Testada ✅
2. **Interface Repository** → Definida ✅
3. **Implementação InMemory** → Testada ✅
4. **Casos de Uso** → Implementados e testados ✅
5. **Interface CLI** → Desenvolvida com Cobra ✅
6. **Integração** → Main.go conectando tudo ✅

Cada etapa foi testada individualmente antes de prosseguir.

### 🔄 8. Debug
**Problemas Encontrados e Soluções**:
- **Thread Safety**: Repository inicial não era seguro → Adicionado `sync.RWMutex`
- **Imports Incorretos**: IA sugeriu imports errados → Corrigidos manualmente
- **Testes Falhando**: Problemas de timing em testes → Ajustados com `time.Sleep`

### 🔄 9. Finish the Project
**Atividades de Finalização**:
- ✅ **Documentação**: README completo, arquitetura documentada
- ✅ **Makefile**: Automação de builds e comandos
- ✅ **Testes Completos**: Cobertura em todas as camadas
- ✅ **Exemplos de Uso**: Guia prático na pasta samples
- ✅ **Estrutura Clara**: Organização seguindo convenções Go
- ✅ **Estado Compartilhável**: Projeto pronto para uso

## 🏗 Arquitetura Implementada

### Clean Architecture
```
┌─────────────────────────────────────┐
│            Interface Layer           │  ← CLI (Cobra)
├─────────────────────────────────────┤
│           Application Layer          │  ← Use Cases
├─────────────────────────────────────┤
│            Domain Layer              │  ← Entities
├─────────────────────────────────────┤
│         Infrastructure Layer         │  ← Repository Implementation
└─────────────────────────────────────┘
```

### Estrutura de Pastas
```
app/
├── main.go                           # Ponto de entrada
├── core/                            # Camada de negócio
│   ├── domain/entity/               # Entidades
│   └── application/                 # Casos de uso
└── infrastructure/                  # Implementações
    ├── repository/                  # Persistência
    └── interface/
        ├── repository/              # Contratos
        └── cli/                     # Interface CLI
```

## 🚀 Como Usar

### Instalação
```bash
git clone <repo>
cd codecademy-yellowbelt2
make install    # Instala dependências
make build      # Compila a aplicação
```

### Comandos Principais
```bash
# Gerenciamento via Makefile
make create TITLE="Estudar Go" DESC="Aprender interfaces"
make list
make complete ID="<todo-id>"
make delete ID="<todo-id>"

# Uso direto da CLI
./bin/todo create "Nova tarefa" "Descrição opcional"
./bin/todo list
./bin/todo complete <id>
```

### Testes
```bash
make test           # Executar testes
make test-coverage  # Coverage report
```

## 🎓 Principais Aprendizados

### Sobre IA Generativa
- **Contexto é Crucial**: Prompts detalhados geram código melhor
- **Validação Necessária**: Sempre testar código gerado
- **Desenvolvimento Incremental**: Pequenos passos são mais efetivos
- **IA como Ferramenta**: Acelera desenvolvimento, mas não substitui conhecimento

### Sobre Clean Architecture
- **Separação de Responsabilidades**: Facilita testes e manutenção
- **Inversão de Dependências**: Core independente de detalhes técnicos
- **Testabilidade**: Cada camada testável isoladamente

### Sobre Go
- **Simplicidade**: Linguagem clara e objetiva
- **Interfaces**: Contratos bem definidos
- **Testes**: Convenções simples e efetivas
- **Concorrência**: Goroutines e channels para thread safety

## 📊 Métricas do Projeto
- **Linhas de Código**: ~500 linhas
- **Cobertura de Testes**: >85%
- **Tempo de Desenvolvimento**: ~6 horas
- **Uso de IA**: ~60% do código inicial gerado/sugerido por IA
- **Refatorações**: 3 grandes refatorações baseadas em testes

---

**Projeto desenvolvido como exercício prático do curso Codecademy Yellow Belt 2, demonstrando uso efetivo de IA Generativa no desenvolvimento de software.**

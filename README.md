# Todo List CLI

> 🎯 **Uma ferramenta de linha de comando para gerenciamento de tarefas pessoais desenvolvida em Go**

[![Go Version](https://img.shields.io/badge/Go-1.23.5-blue.svg)](https://golang.org)
[![Architecture](https://img.shields.io/badge/Architecture-Clean%20Architecture-green.svg)](#arquitetura)
[![Tests](https://img.shields.io/badge/Coverage-%3E85%25-success.svg)](#testes)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

## 🚀 Quick Start

```bash
# Clone e entre no diretório
git clone https://github.com/1o0oP/codecademy-yellowbelt2.git
cd codecademy-yellowbelt2/app

# Instale dependências e compile
make install && make build

# Crie sua primeira tarefa
make create TITLE="Estudar Go" DESC="Aprender Clean Architecture"

# Liste suas tarefas
make list
```

## Sobre o Projeto

Este projeto é uma **aplicação CLI (Command Line Interface)** para gerenciamento de tarefas pessoais, desenvolvida como parte do curso **Codecademy Yellow Belt 2**. O foco principal é demonstrar o **uso responsável e efetivo de IA Generativa** no desenvolvimento de software.

### ✨ Características Principais

- **🏗 Clean Architecture**: Separação clara de responsabilidades
- **🔒 Thread-Safe**: Operações concorrentes seguras
- **💾 Persistência**: Armazenamento em arquivo JSON local
- **🧪 Testado**: Cobertura de testes >85%
- **⚡ Performance**: Interface CLI rápida e responsiva
- **🔧 Extensível**: Arquitetura preparada para novas funcionalidades

## Funcionalidades

| Comando    | Descrição                     | Exemplo                                 |
|------------|-------------------------------|-----------------------------------------|
| `create`   | Criar nova tarefa             | `todo create "Estudar" "Revisar conceitos"` |
| `list`     | Listar todas as tarefas       | `todo list`                             |
| `show`     | Exibir detalhes de uma tarefa | `todo show <id>`                        |
| `update`   | Atualizar tarefa existente    | `todo update <id> "Novo título"`        |
| `complete` | Marcar como concluída         | `todo complete <id>`                    |
| `delete`   | Remover tarefa                | `todo delete <id>`                      |

## 🏗 Arquitetura

O projeto implementa **Clean Architecture** com as seguintes camadas:

```
┌─────────────────────────────┐
│       Interface Layer       │  ← CLI (Cobra Framework)
├─────────────────────────────┤
│      Application Layer      │  ← Use Cases & Business Logic
├─────────────────────────────┤
│        Domain Layer         │  ← Entities & Business Rules
├─────────────────────────────┤
│     Infrastructure Layer    │  ← Repository & External Deps
└─────────────────────────────┘
```

[📖 **Documentação Detalhada da Arquitetura**](docs/architecture.md)

## 📚 Documentação

- [🏗 **Arquitetura & Design**](docs/architecture.md) - Decisões arquiteturais e padrões
- [⚡ **Guia de Início Rápido**](docs/getting-started.md) - Instalação e primeiros passos
- [🎯 **Guia de Uso**](docs/usage-guide.md) - Exemplos práticos de todos os comandos
- [🧪 **Testes**](docs/testing.md) - Estratégias de teste e cobertura
- [🤖 **IA no Desenvolvimento**](docs/ai-development.md) - Como a IA foi utilizada
- [🔧 **API Reference**](docs/api-reference.md) - Documentação técnica das interfaces
- [🚀 **Deploy & Build**](docs/deployment.md) - Build, distribuição e configuração

## 🧪 Testes

```bash
# Executar todos os testes
make test

# Gerar relatório de cobertura HTML
make test-coverage

# Abrir relatório no navegador
open coverage.html
```

**Métricas de Qualidade:**
- ✅ Cobertura de testes: **>85%**
- ✅ Testes unitários em todas as camadas
- ✅ Mocks para isolamento de dependências
- ✅ Testes de integração da CLI

## 🤖 IA Generativa no Desenvolvimento

Este projeto demonstra o **uso estratégico de IA Generativa** (GitHub Copilot) no desenvolvimento:

- **~60% do código inicial** gerado/sugerido por IA
- **Validação rigorosa** através de testes automatizados
- **Desenvolvimento incremental** com feedback constante
- **Documentação** do processo e lições aprendidas

[📖 **Leia mais sobre o uso de IA no projeto**](docs/ai-development.md)

## 📊 Estatísticas do Projeto

| Métrica                | Valor        |
|------------------------|-------------|
| **Linhas de Código**   | ~800 linhas |
| **Cobertura de Testes**| >85%        |
| **Tempo de Desenvolvimento** | ~8 horas |
| **Contribuição da IA** | ~60%        |
| **Arquivos de Teste**  | 6 arquivos  |
| **Dependências**       | 6 bibliotecas |

## 🤝 Contribuindo

Este é um projeto educacional, mas contribuições são bem-vindas! Consulte as diretrizes de contribuição e a documentação técnica.

## 📄 Licença

Este projeto está licenciado sob a [MIT License](LICENSE).

---

<p align="center">
    <strong>Desenvolvido como exercício prático do curso Codecademy Yellow Belt 2</strong><br>
    <em>Demonstrando uso efetivo de IA Generativa no desenvolvimento de software</em>
</p>

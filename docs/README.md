# 📚 Índice da Documentação

Navegação completa pela documentação do Todo List CLI.

## 🎯 Para Usuários

### ⚡ [Getting Started](getting-started.md)
**Comece aqui!** Guia de instalação e primeiros passos para usar o Todo List CLI em menos de 5 minutos.

**Conteúdo:**
- Instalação rápida e verificação do ambiente
- Primeiro workflow de uso (criar → listar → completar → ver detalhes)
- Comandos essenciais via Make e CLI direta
- Localização dos dados e estrutura de arquivos
- Solução de problemas comuns
- Próximos passos

**Ideal para:** Usuários que querem começar a usar imediatamente

---

### 🎯 [Usage Guide](usage-guide.md) 
**Guia completo de uso** com todos os comandos, exemplos práticos e cenários reais de aplicação.

**Conteúdo:**
- Documentação detalhada de cada comando (`create`, `list`, `show`, `update`, `complete`, `delete`)
- Cenários práticos (planejamento diário, gerenciamento pessoal, desenvolvimento de software)
- Dicas e truques para maior produtividade
- Workflow com pipeline e análise de produtividade
- Comandos avançados (backup, restore, exportação)

**Ideal para:** Usuários que querem dominar todas as funcionalidades

---

## 🔧 Para Desenvolvedores

### 🏗 [Architecture](architecture.md)
**Documentação técnica da arquitetura** - Clean Architecture em Go com todos os detalhes de design e implementação.

**Conteúdo:**
- Estrutura em camadas (Domain, Application, Infrastructure, Interface)
- Fluxo de dados e dependências
- Princípios SOLID aplicados
- Padrões utilizados (Repository, Dependency Injection, etc.)
- Testabilidade e extensibilidade
- Possíveis extensões futuras

**Ideal para:** Desenvolvedores que querem entender o design do sistema

---

### 🧪 [Testing](testing.md)
**Estratégias de teste e cobertura** - Como garantimos qualidade através de testes abrangentes.

**Conteúdo:**
- Pirâmide de testes implementada
- Testes por camada (Entity, Use Case, Repository, CLI)
- Estratégias de Mock e Test Doubles
- Cobertura detalhada (>85%)
- Testes de concorrência e casos extremos
- Performance e benchmarks

**Ideal para:** Desenvolvedores interessados em qualidade e TDD

---

### 🔧 [API Reference](api-reference.md)
**Referência técnica completa** - Documentação de todas as interfaces, tipos e métodos.

**Conteúdo:**
- Documentação completa de tipos (`entity.Todo`, interfaces)
- Métodos e construtores com exemplos
- Repository implementations (InMemory e File)
- Use Cases e CLI interface
- Error handling e thread safety
- Performance characteristics

**Ideal para:** Desenvolvedores que querem modificar ou estender o código

---

### 🚀 [Deployment](deployment.md)
**Build, distribuição e configuração** - Como compilar, empacotar e distribuir o Todo List CLI.

**Conteúdo:**
- Build process (local, cross-platform, otimizações)
- Containerização com Docker
- Instalação em diferentes sistemas (Unix/Linux/Windows)
- Package managers (Homebrew, Debian packages)
- CI/CD pipelines (GitHub Actions, GitLab CI)
- Configuração de runtime e monitoramento

**Ideal para:** DevOps e deployment engineers

---

## 🤖 Para Estudantes de IA

### 🤖 [AI Development](ai-development.md)
**IA Generativa no desenvolvimento** - Como GitHub Copilot foi usado estrategicamente seguindo Codecademy Yellow Belt 2.

**Conteúdo:**
- Processo incremental de desenvolvimento com IA
- Métricas de contribuição (~60% código inicial gerado por IA)
- Estratégias de prompt engineering
- Debugging colaborativo (human + AI)
- Lições aprendidas sobre limitações e eficácia
- Best practices para uso responsável de IA

**Ideal para:** Estudantes e desenvolvedores interessados em IA Generativa

---

## 🎓 Rotas de Aprendizado

### 🚀 **Usuário Rápido** (5-10 minutos)
```
1. Getting Started → Uso básico
2. Usage Guide (seção Quick Reference)
```

### 💼 **Usuário Avançado** (30-45 minutos)
```
1. Getting Started → Instalação
2. Usage Guide → Todos os comandos
3. Usage Guide → Cenários práticos
```

### 👨‍💻 **Desenvolvedor Iniciante** (1-2 horas)
```
1. Architecture → Visão geral
2. Getting Started → Setup de desenvolvimento
3. Testing → Estratégias básicas
4. API Reference → Interfaces principais
```

### 🏗 **Desenvolvedor Arquiteto** (2-3 horas)
```
1. Architecture → Documentação completa
2. API Reference → Detalhes técnicos
3. Testing → Estratégias avançadas
4. Deployment → Build e CI/CD
```

### 🤖 **Estudante de IA** (1-1.5 horas)
```
1. AI Development → Processo completo
2. Architecture → Para entender o contexto
3. Testing → Para ver validação de código IA
```

### 🚀 **DevOps Engineer** (1 hora)
```
1. Deployment → Build e distribuição
2. Architecture → Para entender dependências
3. API Reference → Para configuração
```

---

## 📊 Navegação por Complexidade

### 🟢 **Iniciante**
- [Getting Started](getting-started.md) - Uso básico
- [Usage Guide](usage-guide.md) - Comandos essenciais
- [AI Development](ai-development.md) - Seções sobre lições aprendidas

### 🟡 **Intermediário**  
- [Architecture](architecture.md) - Design patterns
- [Testing](testing.md) - Estratégias de teste
- [Deployment](deployment.md) - Build e instalação

### 🔴 **Avançado**
- [API Reference](api-reference.md) - Detalhes técnicos
- [Architecture](architecture.md) - Extensibilidade
- [Testing](testing.md) - Concorrência e performance
- [Deployment](deployment.md) - CI/CD e produção

---

## 🔍 Navegação por Tema

### **🏗 Arquitetura de Software**
- [Architecture](architecture.md) - Clean Architecture completa
- [API Reference](api-reference.md) - Interfaces e contratos
- [Testing](testing.md) - Testabilidade por design

### **🧪 Qualidade de Código**
- [Testing](testing.md) - Estratégias de teste
- [AI Development](ai-development.md) - Validação de código IA
- [Deployment](deployment.md) - Quality gates

### **🚀 DevOps e Deployment**
- [Deployment](deployment.md) - Build, CI/CD, containers
- [Getting Started](getting-started.md) - Setup de ambiente
- [API Reference](api-reference.md) - Configuração

### **🤖 Inteligência Artificial**
- [AI Development](ai-development.md) - Uso estratégico de IA
- [Testing](testing.md) - Validação de código gerado
- [Architecture](architecture.md) - Decisões humanas vs IA

### **💻 Desenvolvimento Go**
- [API Reference](api-reference.md) - Padrões Go
- [Testing](testing.md) - Testing em Go
- [Architecture](architecture.md) - Interfaces e packages

---

## 🎯 Quick Links

| Quero... | Vá para |
|----------|---------|
| **Usar o todo CLI agora** | [Getting Started](getting-started.md) |
| **Aprender todos os comandos** | [Usage Guide](usage-guide.md) |
| **Entender a arquitetura** | [Architecture](architecture.md) |
| **Ver como os testes funcionam** | [Testing](testing.md) |
| **Modificar o código** | [API Reference](api-reference.md) |
| **Fazer deploy** | [Deployment](deployment.md) |
| **Entender o uso de IA** | [AI Development](ai-development.md) |

---

## 📱 Suporte e Contribuição

### 🐛 **Encontrou um Bug?**
1. Verifique a seção de troubleshooting em [Getting Started](getting-started.md)
2. Consulte issues conhecidos em [Deployment](deployment.md)
3. Abra uma issue no repositório

### 💡 **Quer Contribuir?**
1. Leia [Architecture](architecture.md) para entender o design
2. Veja [Testing](testing.md) para entender os padrões de teste
3. Consulte [API Reference](api-reference.md) para detalhes técnicos
4. Siga as práticas de [AI Development](ai-development.md) se usar IA

### 📚 **Quer Aprender Mais?**
1. **Clean Architecture**: [Architecture](architecture.md)
2. **Testing em Go**: [Testing](testing.md)  
3. **IA no Desenvolvimento**: [AI Development](ai-development.md)
4. **DevOps com Go**: [Deployment](deployment.md)

---

**🎯 Esta documentação foi criada para ser um recurso completo tanto para usuários quanto desenvolvedores. Comece pelo documento que mais se adequa ao seu objetivo atual!**

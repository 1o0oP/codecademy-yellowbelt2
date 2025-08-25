# 🎯 Guia de Uso Completo

Este guia apresenta todos os comandos disponíveis com exemplos práticos e cenários reais de uso.

## 📋 Visão Geral dos Comandos

| Comando | Propósito | Parâmetros Obrigatórios | Parâmetros Opcionais |
|---------|-----------|------------------------|----------------------|
| `create` | Criar nova tarefa | `title` | `description` |
| `list` | Listar todas as tarefas | - | - |
| `show` | Exibir detalhes de uma tarefa | `id` | - |
| `update` | Atualizar tarefa existente | `id`, `title` | `description` |
| `complete` | Marcar como concluída | `id` | - |
| `delete` | Remover tarefa | `id` | - |

## 🔧 Comandos Detalhados

### 1. `create` - Criar Nova Tarefa

Cria uma nova tarefa com título obrigatório e descrição opcional.

#### Sintaxe
```bash
# Via Make
make create TITLE="título" DESC="descrição"

# Via CLI
./bin/todo create "título" ["descrição"]
```

#### Exemplos Práticos

**Tarefa simples (sem descrição):**
```bash
make create TITLE="Comprar café"

# Saída:
# ✅ Tarefa criada com sucesso!
# ID: a1b2c3d4-e5f6-7g8h-9i0j-k1l2m3n4o5p6
# Título: Comprar café
```

**Tarefa completa (com descrição):**
```bash
make create TITLE="Estudar Go" DESC="Focar em interfaces e channels"

# Saída:
# ✅ Tarefa criada com sucesso!
# ID: b2c3d4e5-f6g7-8h9i-0j1k-l2m3n4o5p6q7
# Título: Estudar Go
# Descrição: Focar em interfaces e channels
```

**Via CLI direta:**
```bash
./bin/todo create "Preparar apresentação"
./bin/todo create "Revisar código" "Pull Request #42 precisa de review"
```

#### Comportamento Esperado
- ✅ Gera ID único automaticamente (UUID)
- ✅ Define status como "pendente" (não concluída)
- ✅ Registra timestamp de criação
- ✅ Salva automaticamente no arquivo JSON
- ⚠️ Se título estiver vazio, comando falhará

---

### 2. `list` - Listar Todas as Tarefas

Exibe todas as tarefas com status visual e informações básicas.

#### Sintaxe
```bash
# Via Make
make list

# Via CLI
./bin/todo list
```

#### Exemplo de Saída

**Lista com várias tarefas:**
```
📋 Total de tarefas: 4

1. ⏳ Comprar café
   🆔 ID: a1b2c3d4-e5f6-7g8h-9i0j-k1l2m3n4o5p6

2. ✅ Estudar Go
   📄 Focar em interfaces e channels
   🆔 ID: b2c3d4e5-f6g7-8h9i-0j1k-l2m3n4o5p6q7

3. ⏳ Preparar apresentação
   📄 Slides sobre Clean Architecture
   🆔 ID: c3d4e5f6-g7h8-9i0j-1k2l-m3n4o5p6q7r8

4. ⏳ Fazer exercícios
   📄 30 minutos de corrida
   🆔 ID: d4e5f6g7-h8i9-0j1k-2l3m-n4o5p6q7r8s9
```

**Lista vazia:**
```
📝 Nenhuma tarefa encontrada!
```

#### Símbolos de Status
- ⏳ **Pendente**: Tarefa não concluída
- ✅ **Concluída**: Tarefa finalizada
- 📄 **Descrição**: Descrição opcional da tarefa
- 🆔 **ID**: Identificador único da tarefa

---

### 3. `show` - Exibir Detalhes Específicos

Mostra informações completas de uma tarefa específica.

#### Sintaxe
```bash
# Via Make
make show ID="id-da-tarefa"

# Via CLI
./bin/todo show "id-da-tarefa"
```

#### Exemplos

**Tarefa pendente:**
```bash
make show ID="a1b2c3d4-e5f6-7g8h-9i0j-k1l2m3n4o5p6"

# Saída:
# 🆔 ID: a1b2c3d4-e5f6-7g8h-9i0j-k1l2m3n4o5p6
# 📝 Título: Comprar café
# 📊 Status: ⏳ Pendente
# 📅 Criada em: 25/08/2025 14:30
# 🔄 Atualizada em: 25/08/2025 14:30
```

**Tarefa concluída com descrição:**
```bash
make show ID="b2c3d4e5-f6g7-8h9i-0j1k-l2m3n4o5p6q7"

# Saída:
# 🆔 ID: b2c3d4e5-f6g7-8h9i-0j1k-l2m3n4o5p6q7
# 📝 Título: Estudar Go
# 📄 Descrição: Focar em interfaces e channels
# 📊 Status: ✅ Concluída
# 📅 Criada em: 25/08/2025 09:15
# 🔄 Atualizada em: 25/08/2025 16:45
```

**Tarefa não encontrada:**
```bash
make show ID="inexistente"

# Saída:
# ❌ Tarefa não encontrada: todo not found
```

---

### 4. `update` - Atualizar Tarefa Existente

Permite modificar título e/ou descrição de uma tarefa.

#### Sintaxe
```bash
# Via Make (com descrição)
make update ID="id" TITLE="novo título" DESC="nova descrição"

# Via CLI
./bin/todo update "id" "novo título" ["nova descrição"]
```

#### Exemplos

**Atualizar apenas o título:**
```bash
./bin/todo update "a1b2c3d4-e5f6-7g8h-9i0j-k1l2m3n4o5p6" "Comprar café premium"

# Saída:
# ✅ Tarefa atualizada com sucesso!
# 📝 Título: Comprar café premium
```

**Atualizar título e descrição:**
```bash
./bin/todo update "b2c3d4e5-f6g7-8h9i-0j1k-l2m3n4o5p6q7" "Dominar Go" "Estudar concorrência e performance"

# Saída:
# ✅ Tarefa atualizada com sucesso!
# 📝 Título: Dominar Go
# 📄 Descrição: Estudar concorrência e performance
```

#### Comportamento
- ✅ Mantém o status (concluída/pendente)
- ✅ Atualiza timestamp de modificação
- ✅ Preserva ID original
- ⚠️ Título não pode ser vazio

---

### 5. `complete` - Marcar como Concluída

Marca uma tarefa como finalizada.

#### Sintaxe
```bash
# Via Make
make complete ID="id-da-tarefa"

# Via CLI
./bin/todo complete "id-da-tarefa"
```

#### Exemplos

**Marcar tarefa como concluída:**
```bash
make complete ID="a1b2c3d4-e5f6-7g8h-9i0j-k1l2m3n4o5p6"

# Saída:
# ✅ Tarefa 'Comprar café premium' marcada como concluída!
```

**Verificar mudança:**
```bash
make list

# A tarefa agora aparece com ✅:
# 1. ✅ Comprar café premium
#    🆔 ID: a1b2c3d4-e5f6-7g8h-9i0j-k1l2m3n4o5p6
```

#### Comportamento
- ✅ Altera status para "concluída"
- ✅ Atualiza timestamp de modificação
- ℹ️ Aplicar em tarefa já concluída não causa erro

---

### 6. `delete` - Remover Tarefa

Remove permanentemente uma tarefa do sistema.

#### Sintaxe
```bash
# Via Make
make delete ID="id-da-tarefa"

# Via CLI
./bin/todo delete "id-da-tarefa"
```

#### Exemplos

**Deletar tarefa:**
```bash
make delete ID="d4e5f6g7-h8i9-0j1k-2l3m-n4o5p6q7r8s9"

# Saída:
# 🗑️  Tarefa deletada com sucesso!
```

**Verificar remoção:**
```bash
make list

# A tarefa não aparece mais na lista
```

**Tentar deletar tarefa inexistente:**
```bash
make delete ID="inexistente"

# Saída:
# ❌ Erro ao deletar tarefa: todo not found
```

#### ⚠️ **AVISO**: Operação irreversível!
- A tarefa é removida permanentemente
- Não há funcionalidade de "lixeira" ou "desfazer"
- Faça backup do arquivo `~/.todo-cli/todos.json` se necessário

---

## 🎯 Cenários de Uso Práticos

### 📅 Workflow de Planejamento Diário

**1. Manhã - Planejamento do Dia**
```bash
# Criar tarefas importantes
make create TITLE="Email matinal" DESC="Responder emails urgentes"
make create TITLE="Reunião 10h" DESC="Sprint planning com a equipe"
make create TITLE="Feature X" DESC="Implementar login social"
make create TITLE="Code review" DESC="Revisar PR #123 e #124"

# Ver lista do dia
make list
```

**2. Durante o Dia - Execução**
```bash
# Completar tarefas conforme progride
make complete ID="email-id"
make complete ID="reuniao-id"

# Atualizar se necessário
make update ID="feature-id" TITLE="Feature X - OAuth" DESC="Implementar Google OAuth"
```

**3. Final do Dia - Revisão**
```bash
# Ver progresso
make list

# Criar tarefas para amanhã
make create TITLE="Continuar Feature X" DESC="Implementar Facebook OAuth"
```

### 🏠 Gerenciamento Pessoal

**Tarefas domésticas:**
```bash
make create TITLE="Supermercado" DESC="Leite, pão, frutas"
make create TITLE="Academia" DESC="Treino de pernas - 1h"
make create TITLE="Dentista" DESC="Consulta às 16h - Dr. Silva"
```

**Acompanhamento de hábitos:**
```bash
make create TITLE="Leitura" DESC="20 páginas do livro de Go"
make create TITLE="Água" DESC="Beber 2L de água hoje"
make create TITLE="Meditação" DESC="10 minutos de mindfulness"
```

### 💻 Desenvolvimento de Software

**Sprint planning:**
```bash
make create TITLE="Setup CI/CD" DESC="Configurar GitHub Actions"
make create TITLE="Testes unitários" DESC="Cobertura >90% no módulo auth"
make create TITLE="Documentação" DESC="Atualizar README e API docs"
make create TITLE="Refatoração" DESC="Extrair interface UserService"
```

**Bug tracking:**
```bash
make create TITLE="Bug #456" DESC="Crash no login com email vazio"
make create TITLE="Performance" DESC="Query lenta na listagem de users"
```

---

## 🔍 Dicas e Truques

### ✂️ Usando IDs Curtos
Você não precisa do ID completo, os primeiros 8 caracteres são suficientes:

```bash
# ID completo
a1b2c3d4-e5f6-7g8h-9i0j-k1l2m3n4o5p6

# ID curto (funciona igual)
a1b2c3d4

# Exemplo
make complete ID="a1b2c3d4"
```

### 📋 Copiando IDs Facilmente
```bash
# Listar tarefas e copiar ID
make list | grep "🆔 ID" | head -1
# Saída: 🆔 ID: a1b2c3d4-e5f6-7g8h-9i0j-k1l2m3n4o5p6
```

### 🔄 Workflow com Pipeline
```bash
# Criar várias tarefas em sequência
make create TITLE="Task 1" DESC="Primeira tarefa" && \
make create TITLE="Task 2" DESC="Segunda tarefa" && \
make create TITLE="Task 3" DESC="Terceira tarefa"
```

### 📊 Análise de Produtividade
```bash
# Ver todas as tarefas
make list

# Contar tarefas concluídas
make list | grep "✅" | wc -l

# Contar tarefas pendentes
make list | grep "⏳" | wc -l
```

## ❌ Mensagens de Erro Comuns

### Tarefa Não Encontrada
```
❌ Tarefa não encontrada: todo not found
```
**Solução**: Verifique o ID usando `make list`

### Argumentos Inválidos
```
Error: accepts between 1 and 2 arg(s), received 0
```
**Solução**: Forneça os argumentos obrigatórios

### Arquivo de Dados Corrompido
```
❌ Erro ao listar tarefas: invalid character '}' after object key:value pair
```
**Solução**: Restaure backup ou delete `~/.todo-cli/todos.json`

---

## 🚀 Comandos Avançados

### Backup e Restore
```bash
# Fazer backup
cp ~/.todo-cli/todos.json ~/backup-todos-$(date +%Y%m%d).json

# Restaurar backup
cp ~/backup-todos-20250825.json ~/.todo-cli/todos.json
```

### Exportar Tarefas
```bash
# Ver arquivo JSON raw
cat ~/.todo-cli/todos.json | jq '.'

# Exportar apenas títulos
make list | grep -E "^\d+\." | sed 's/^\d+\. [⏳✅] //'
```

### Resetar Sistema
```bash
# ⚠️ CUIDADO: Remove todas as tarefas
rm ~/.todo-cli/todos.json
```

---

## 📈 Próximos Passos

Agora que você domina todos os comandos:

1. 🏗 [**Entenda a Arquitetura**](architecture.md) - Como tudo funciona
2. 🧪 [**Explore os Testes**](testing.md) - Garantia de qualidade
3. 🔧 [**API Reference**](api-reference.md) - Detalhes técnicos
4. 🤖 [**IA no Desenvolvimento**](ai-development.md) - Como foi construído

---

**💡 Dica Final**: Use o Todo List CLI para gerenciar suas próprias tarefas enquanto explora o código! É uma ótima forma de entender o valor da ferramenta que você está estudando.

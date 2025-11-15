# CLI - Tasks Management

Sistema de gerenciamento de tarefas composto por uma API REST em Go e um cliente CLI para interação via terminal.

## 🏗️ Arquitetura

- **API REST**: Containerizada em Docker, responsável pelo gerenciamento de tarefas
- **CLI**: Cliente em Go para consumo da API via linha de comando
- **PostgreSQL**: Banco de dados para persistência das tarefas

## 📋 Pré-requisitos

- Go 1.25+
- Docker
- PostgreSQL

## 🚀 Instalação e Execução

### 1. Clone o projeto

```bash
git clone https://github.com/outrojao/cli_tasks
cd cli_tasks
```

### 2. Instale as dependências e execute o CLI

```bash
# Instale as dependências
go mod tidy

# Execute o CLI
make run
```

## 📝 Funcionalidades

- ✅ Criar tarefas
- ✅ Listar tarefas
- ✅ Atualizar tarefas
- ✅ Excluir tarefas
- ✅ Marcar tarefas como concluídas

### Estrutura do Projeto

```
.
├── cmd
│   └── main.go
├── internal
   ├── app
   │   └── task
   │       ├── task.go
   ├── cli
   │   └── menu.go
   └── client
       └── client.go
```

### Comandos Make disponíveis

```bash
make run      # Executar o CLI
make build    # Build do projeto
make test     # Executar testes
```

## 🌐 API Endpoints

- `POST /create` - Criar nova tarefa
- `PUT /do/:task_name` - Atualizar tarefa
- `DELETE /remove/:task_name` - Excluir tarefa
- `GET /list` - Listar todas as tarefas

package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	app_interfaces "codecademy-yellowbelt2/infrastructure/interface/application"
)

type TodoCLI struct {
	todoUseCase app_interfaces.ITodoUseCase
}

func NewTodoCLI(todoUseCase app_interfaces.ITodoUseCase) *TodoCLI {
	return &TodoCLI{
		todoUseCase: todoUseCase,
	}
}

func (cli *TodoCLI) GetRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "todo",
		Short: "Todo List CLI - Gerenciador de tarefas",
		Long:  "Uma ferramenta de linha de comando para gerenciar sua lista de tarefas",
	}

	rootCmd.AddCommand(cli.createCommand())
	rootCmd.AddCommand(cli.listCommand())
	rootCmd.AddCommand(cli.showCommand())
	rootCmd.AddCommand(cli.updateCommand())
	rootCmd.AddCommand(cli.completeCommand())
	rootCmd.AddCommand(cli.deleteCommand())

	return rootCmd
}

func (cli *TodoCLI) createCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "create [title] [description]",
		Short: "Criar uma nova tarefa",
		Args:  cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			title := args[0]
			description := ""
			if len(args) > 1 {
				description = args[1]
			}

			todo, err := cli.todoUseCase.CreateTodo(title, description)
			if err != nil {
				fmt.Printf("Erro ao criar tarefa: %v\n", err)
				return
			}

			fmt.Printf("✅ Tarefa criada com sucesso!\n")
			fmt.Printf("ID: %s\n", todo.ID)
			fmt.Printf("Título: %s\n", todo.Title)
			if todo.Description != "" {
				fmt.Printf("Descrição: %s\n", todo.Description)
			}
		},
	}
}

func (cli *TodoCLI) listCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Listar todas as tarefas",
		Run: func(cmd *cobra.Command, args []string) {
			todos, err := cli.todoUseCase.GetAllTodos()
			if err != nil {
				fmt.Printf("Erro ao listar tarefas: %v\n", err)
				return
			}

			if len(todos) == 0 {
				fmt.Println("📝 Nenhuma tarefa encontrada!")
				return
			}

			fmt.Printf("📋 Total de tarefas: %d\n\n", len(todos))
			for i, todo := range todos {
				status := "⏳"
				if todo.Completed {
					status = "✅"
				}
				fmt.Printf("%d. %s %s\n", i+1, status, todo.Title)
				if todo.Description != "" {
					fmt.Printf("   📄 %s\n", todo.Description)
				}
				fmt.Printf("   🆔 ID: %s\n", todo.ID)
				fmt.Println()
			}
		},
	}
}

func (cli *TodoCLI) showCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "show [id]",
		Short: "Mostrar detalhes de uma tarefa específica",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]

			todo, err := cli.todoUseCase.GetTodoByID(id)
			if err != nil {
				fmt.Printf("❌ Tarefa não encontrada: %v\n", err)
				return
			}

			status := "⏳ Pendente"
			if todo.Completed {
				status = "✅ Concluída"
			}

			fmt.Printf("🆔 ID: %s\n", todo.ID)
			fmt.Printf("📝 Título: %s\n", todo.Title)
			if todo.Description != "" {
				fmt.Printf("📄 Descrição: %s\n", todo.Description)
			}
			fmt.Printf("📊 Status: %s\n", status)
			fmt.Printf("📅 Criada em: %s\n", todo.CreatedAt.Format("02/01/2006 15:04"))
			fmt.Printf("🔄 Atualizada em: %s\n", todo.UpdatedAt.Format("02/01/2006 15:04"))
		},
	}
}

func (cli *TodoCLI) updateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "update [id] [title] [description]",
		Short: "Atualizar uma tarefa existente",
		Args:  cobra.RangeArgs(2, 3),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			title := args[1]
			description := ""
			if len(args) > 2 {
				description = args[2]
			}

			todo, err := cli.todoUseCase.UpdateTodo(id, title, description)
			if err != nil {
				fmt.Printf("❌ Erro ao atualizar tarefa: %v\n", err)
				return
			}

			fmt.Printf("✅ Tarefa atualizada com sucesso!\n")
			fmt.Printf("📝 Título: %s\n", todo.Title)
			if todo.Description != "" {
				fmt.Printf("📄 Descrição: %s\n", todo.Description)
			}
		},
	}
}

func (cli *TodoCLI) completeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "complete [id]",
		Short: "Marcar uma tarefa como concluída",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]

			todo, err := cli.todoUseCase.CompleteTodo(id)
			if err != nil {
				fmt.Printf("❌ Erro ao completar tarefa: %v\n", err)
				return
			}

			fmt.Printf("✅ Tarefa '%s' marcada como concluída!\n", todo.Title)
		},
	}
}

func (cli *TodoCLI) deleteCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "delete [id]",
		Short: "Deletar uma tarefa",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]

			err := cli.todoUseCase.DeleteTodo(id)
			if err != nil {
				fmt.Printf("❌ Erro ao deletar tarefa: %v\n", err)
				return
			}

			fmt.Println("🗑️  Tarefa deletada com sucesso!")
		},
	}
}

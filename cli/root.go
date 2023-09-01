package cli

import (
	"todos/actions/add"
	"todos/actions/del"
	"todos/actions/list"
	"todos/actions/markdone"
	unmarkdone "todos/actions/unmarkDone"
	"todos/services"

	"github.com/spf13/cobra"
)

func Start(todoService services.TodoService) error {
	rootCmd := &cobra.Command{
		Use: "todos",
		Short: "add todos from the command line",
		Long: `
		Welcome to the Todos CLI!
		You can manage your todos directly from the command line using this app!
		`,
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(add.MakeCommand(add.NewAddingService(&todoService)))
	rootCmd.AddCommand(list.MakeCommand(list.NewListService(&todoService)))
	rootCmd.AddCommand(markdone.NewCommand(markdone.NewMarkDoneService(&todoService)))
	rootCmd.AddCommand(unmarkdone.NewCommand(unmarkdone.NewMarkUndoneService(&todoService)))
	rootCmd.AddCommand(del.NewCommand(del.NewDeleteService(&todoService)))

	return rootCmd.Execute()
}
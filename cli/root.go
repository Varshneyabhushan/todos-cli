package cli

import (
	"os"
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
		Use:   "todos",
		Short: "add todos from the command line",
		Long: `
		Welcome to the Todos CLI!
		You can manage your todos directly from the command line using this app!
		`,
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(add.MakeCommand(add.NewAddingService(&todoService), os.Stdout))
	rootCmd.AddCommand(list.MakeCommand(list.NewListService(&todoService), os.Stdout))
	rootCmd.AddCommand(markdone.NewCommand(markdone.NewMarkDoneService(&todoService), os.Stdout))
	rootCmd.AddCommand(unmarkdone.NewCommand(unmarkdone.NewunmarkDoneService(&todoService), os.Stdout))
	rootCmd.AddCommand(del.NewCommand(del.NewDeleteService(&todoService), os.Stdout))

	return rootCmd.Execute()
}

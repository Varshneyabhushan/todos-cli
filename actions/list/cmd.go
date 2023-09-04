package list

import (
	"errors"
	"fmt"
	"io"
	"todos/services"

	"github.com/spf13/cobra"
)

type itemFilter struct {
	onlyPending bool
}

func shouldShow(item services.TodoItem, filter itemFilter) bool {
	if filter.onlyPending && item.IsDone {
		return false
	}

	return true
}

func itemToText(item services.TodoItem, index int) string {
	status := " "
	if item.IsDone {
		status = "x"
	}

	return fmt.Sprintf("%d [%s] %s", index, status, item.Text)
}

func MakeCommand(s ListService, w io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use: "list",
		Short: "list the current todos",
		Long: `list the total todos those are added. 
		Even the todos those are marked done are shown here. 
		To skip them, simply delete them`,
		RunE: func(cmd *cobra.Command, args []string) error {
			onlyPending, _ := cmd.Flags().GetBool("pending")

			filter := itemFilter {
				onlyPending: onlyPending,
			}

			items, err := s()
			if err != nil {
				return errors.New("error while getting list : " + err.Error())
			}

			if len(items) == 0 {
				fmt.Fprintln(w, "the list is empty!")
				return nil
			}

			for index, item := range items {
				if shouldShow(item, filter) {
					fmt.Fprintln(w, itemToText(item, index))
				}
			}

			return nil
		},
	}

	cmd.Flags().BoolP("pending", "p", false, "Show only pending todos")
	return cmd
}
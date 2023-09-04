package list

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

func MakeCommand(s ListService) *cobra.Command {
	return &cobra.Command{
		Use: "list",
		Short: "list the current todos",
		Long: `list the total todos those are added. 
		Even the todos those are marked done are shown here. 
		To skip them, simply delete them`,
		RunE: func(cmd *cobra.Command, args []string) error {
			items, err := s()
			if err != nil {
				return errors.New("error while getting list : " + err.Error())
			}

			if len(items) == 0 {
				fmt.Println("the list is empty!")
				return nil
			}

			for index, item := range items {
				status := " "
				if item.IsDone {
					status = "x"
				}

				fmt.Println(fmt.Sprintf("%d [%s] %s", index, status, item.Text))
			}

			return nil
		},
	}
}
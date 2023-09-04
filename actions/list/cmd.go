package list

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

func MakeCommand(s ListService) *cobra.Command {
	cmd := &cobra.Command{
		Use: "list",
		Short: "list the current todos",
		Long: `list the total todos those are added. 
		Even the todos those are marked done are shown here. 
		To skip them, simply delete them`,
		RunE: func(cmd *cobra.Command, args []string) error {
			showPending, _ := cmd.Flags().GetBool("pending")

			items, err := s()
			if err != nil {
				return errors.New("error while getting list : " + err.Error())
			}

			if len(items) == 0 {
				fmt.Println("the list is empty!")
				return nil
			}

			for index, item := range items {
				if showPending && item.IsDone {
					continue
				}
				
				status := " "
				if item.IsDone {
					status = "x"
				}

				fmt.Println(fmt.Sprintf("%d [%s] %s", index, status, item.Text))
			}

			return nil
		},
	}

	cmd.Flags().BoolP("pending", "p", false, "Show only pending todos")
	return cmd
}
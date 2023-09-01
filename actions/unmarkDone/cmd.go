package unmarkdone

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func NewCommand(s MarkUndoneService) *cobra.Command {
	return &cobra.Command{
		Use:   "unmark",
		Short: "unmark a done todo",
		Long:  `When a todo item is already complete, mark it as undone`,
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return errors.New("error while converting id to integer : " + err.Error())
			}

			if err = s(id); err != nil {
				return errors.New("error while unmarking an item done : " + err.Error())
			}

			fmt.Println("task marked as pending..")
			return nil
		},
	}
}

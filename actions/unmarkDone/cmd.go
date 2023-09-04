package unmarkdone

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/spf13/cobra"
)

func NewCommand(s MarkUndoneService, w io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "unmark",
		Short: "unmark a done todo",
		Long:  `When a todo item is already complete, mark it as undone`,
		RunE: func(cmd *cobra.Command, args []string) error {
			index, err := strconv.Atoi(args[0])
			if err != nil {
				return errors.New("error while converting index to integer : " + err.Error())
			}

			if err = s(index); err != nil {
				return errors.New("error while unmarking an item done : " + err.Error())
			}

			fmt.Fprintln(w, "task marked as pending..")
			return nil
		},
	}
}

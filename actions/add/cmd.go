package add

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
)

func MakeCommand(add AddingService, w io.Writer) *cobra.Command {
	return &cobra.Command{
		Use: "add",
		Short: "add a todo item to the list",
		Long: "this method let's you add a todo item to the list",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("no args sent")
			}

			todoText := strings.Join(args, " ")
			err := add(todoText)
			if err != nil {
				return errors.New("error while adding a todo : " + err.Error())
			}

			fmt.Fprintln(w, "task added successfully!")
			return nil
		},
	}
}
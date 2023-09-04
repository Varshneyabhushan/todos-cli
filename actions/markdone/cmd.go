package markdone

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/spf13/cobra"
)

func NewCommand(s MarkDoneService, w io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "mark",
		Short: "mark a todo as done",
		Long: `When a todo item is complete, mark it as done`,
		RunE: func(cmd *cobra.Command, args []string) error {
			index, err := strconv.Atoi(args[0])
			if err != nil {
				return errors.New("error while converting index to integer : " + err.Error())
			}

			if err = s(index); err != nil {
				return errors.New("error while marking an item done : " + err.Error())
			}

			fmt.Fprintln(w, "task marked as complete!")
			return nil
		},
	}
}
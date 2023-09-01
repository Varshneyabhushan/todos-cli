package markdone

import (
	"errors"
	"strconv"

	"github.com/spf13/cobra"
)

func NewCommand(s MarkDoneService) *cobra.Command {
	return &cobra.Command{
		Use:   "mark",
		Short: "mark a todo as done",
		Long: `When a todo item is complete, mark it as done`,
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return errors.New("error while converting id to integer : " + err.Error())
			}

			if err = s(id); err != nil {
				return errors.New("error while marking an item done : " + err.Error())
			}

			return nil
		},
	}
}
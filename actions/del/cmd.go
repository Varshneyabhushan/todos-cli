package del

import (
	"errors"
	"strconv"

	"github.com/spf13/cobra"
)

func NewCommand(s DeleteService) *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "delete an item",
		Long:  `Delete an item regardless of whether it's done or not`,
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

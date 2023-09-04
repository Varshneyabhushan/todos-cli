package del

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func NewCommand(s DeleteService) *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "delete an item",
		Long:  `Delete an item regardless of whether it's done or not`,
		RunE: func(cmd *cobra.Command, args []string) error {
			index, err := strconv.Atoi(args[0])
			if err != nil {
				return errors.New("error while converting index to integer : " + err.Error())
			}

			if err = s(index); err != nil {
				return errors.New("error while marking an item done : " + err.Error())
			}

			fmt.Println("task deleted!")
			return nil
		},
	}
}

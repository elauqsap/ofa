package ad

import "github.com/spf13/cobra"

func NewAddCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "add operations for active directory",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	return addCmd
}

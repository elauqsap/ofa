package ad

import "github.com/spf13/cobra"

func NewADCmd() *cobra.Command {
	adCmd := &cobra.Command{
		Use:   "ad",
		Short: "utilities to peform common active directory tasks",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	adCmd.AddCommand(NewAddCmd())
	adCmd.AddCommand(NewDeleteCmd())
	adCmd.AddCommand(NewGetCmd())

	return adCmd
}

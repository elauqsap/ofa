package ad

import "github.com/spf13/cobra"

func NewGetCmd() *cobra.Command {
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "get operations for active directory",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	getCmd.AddCommand(NewGetGroupCmd())
	getCmd.AddCommand(NewGetUserCmd())

	return getCmd
}

func NewGetUserCmd() *cobra.Command {
	userCmd := &cobra.Command{
		Use:   "user",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	return userCmd
}

func NewGetGroupCmd() *cobra.Command {
	groupCmd := &cobra.Command{
		Use:   "group",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	return groupCmd
}

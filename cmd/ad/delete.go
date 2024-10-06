package ad

import "github.com/spf13/cobra"

func NewDeleteCmd() *cobra.Command {
	delCmd := &cobra.Command{
		Use:     "delete",
		Short:   "delete operations for active directory",
		Aliases: []string{"del"},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	delCmd.AddCommand(NewDeleteGroupMemberCmd())

	return delCmd
}

func NewDeleteGroupMemberCmd() *cobra.Command {
	groupMemberCmd := &cobra.Command{
		Use:     "group-member",
		Short:   "remove a member from an active directory group",
		Aliases: []string{"gm"},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	return groupMemberCmd
}

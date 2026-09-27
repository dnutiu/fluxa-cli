package cmd

import "github.com/spf13/cobra"

func (a *app) accountsCommand() *cobra.Command {
	parent := &cobra.Command{Use: "accounts", Short: "Manage accounts in an entity"}
	parent.AddCommand(a.accountListCommand(), a.accountShowCommand(), a.accountAddCommand(), a.accountEditCommand(), a.accountArchiveCommand(), a.accountRestoreCommand())
	return parent
}

package cmd

import (
	"github.com/dnutiu/fluxa-cli/internal/application"
	"github.com/dnutiu/fluxa-cli/internal/domain"
	"github.com/dnutiu/fluxa-cli/internal/presentation"
	"github.com/spf13/cobra"
)

func (a *app) transactionsCommand() *cobra.Command {
	parent := &cobra.Command{Use: "transactions", Short: "Manage income and expenses"}
	parent.AddCommand(a.transactionListCommand(), a.transactionShowCommand(), a.transactionAddCommand(), a.transactionEditCommand(), a.transactionDeleteCommand())
	return parent
}

func (a *app) transactionShowCommand() *cobra.Command {
	return &cobra.Command{
		Use: "show ID", Short: "Show a transaction", Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			entityID, err := a.entityID()
			if err != nil {
				return err
			}
			transactionID, err := domain.ParseID(args[0])
			if err != nil {
				return err
			}
			repo, err := a.repository()
			if err != nil {
				return err
			}
			result, err := application.ShowTransaction(cmd.Context(), repo, entityID, transactionID)
			if err != nil {
				return err
			}
			return a.print(cmd, result, presentation.Transactions)
		},
	}
}

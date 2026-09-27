package cmd

import (
	"github.com/dnutiu/fluxa-cli/internal/application"
	"github.com/dnutiu/fluxa-cli/internal/domain"
	"github.com/dnutiu/fluxa-cli/internal/presentation"
	"github.com/spf13/cobra"
)

func (a *app) accountListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list", Short: "List accounts", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			entityID, err := a.entityID()
			if err != nil {
				return err
			}
			repo, err := a.repository()
			if err != nil {
				return err
			}
			archived, _ := cmd.Flags().GetBool("include-archived")
			result, err := application.ListAccounts(cmd.Context(), repo, entityID, domain.AccountFilter{IncludeArchived: archived})
			if err != nil {
				return err
			}
			return a.print(cmd, result, presentation.Accounts)
		},
	}
	cmd.Flags().Bool("include-archived", false, "include archived accounts")
	return cmd
}

func (a *app) accountShowCommand() *cobra.Command {
	return &cobra.Command{
		Use: "show ID", Short: "Show an account", Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			entityID, err := a.entityID()
			if err != nil {
				return err
			}
			accountID, err := domain.ParseID(args[0])
			if err != nil {
				return err
			}
			repo, err := a.repository()
			if err != nil {
				return err
			}
			result, err := application.ShowAccount(cmd.Context(), repo, entityID, accountID)
			if err != nil {
				return err
			}
			return a.print(cmd, result, presentation.Accounts)
		},
	}
}

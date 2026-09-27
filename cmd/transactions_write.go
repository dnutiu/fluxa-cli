package cmd

import (
	"fmt"

	"github.com/dnutiu/fluxa-cli/internal/application"
	"github.com/dnutiu/fluxa-cli/internal/domain"
	"github.com/dnutiu/fluxa-cli/internal/infrastructure/fluxa"
	"github.com/dnutiu/fluxa-cli/internal/presentation"
	"github.com/spf13/cobra"
)

func (a *app) transactionAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "add", Short: "Add a transaction from a JSON file", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			entityID, err := a.entityID()
			if err != nil {
				return err
			}
			body, err := readBody(cmd)
			if err != nil {
				return err
			}
			key, _ := cmd.Flags().GetString("idempotency-key")
			if key == "" {
				key, err = fluxa.NewIdempotencyKey()
				if err != nil {
					return err
				}
				fmt.Fprintf(cmd.ErrOrStderr(), "Idempotency-Key: %s\n", key)
			}
			repo, err := a.repository()
			if err != nil {
				return err
			}
			result, err := application.AddTransaction(cmd.Context(), repo, entityID, body, key)
			if err != nil {
				return err
			}
			return a.print(cmd, result, presentation.Transactions)
		},
	}
	addFileFlag(cmd)
	cmd.Flags().String("idempotency-key", "", "reuse a key when retrying the exact same create request")
	return cmd
}

func (a *app) transactionEditCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "edit ID", Short: "Edit a transaction from a JSON file", Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			entityID, err := a.entityID()
			if err != nil {
				return err
			}
			transactionID, err := domain.ParseID(args[0])
			if err != nil {
				return err
			}
			body, err := readBody(cmd)
			if err != nil {
				return err
			}
			repo, err := a.repository()
			if err != nil {
				return err
			}
			result, err := application.EditTransaction(cmd.Context(), repo, entityID, transactionID, body)
			if err != nil {
				return err
			}
			return a.print(cmd, result, presentation.Transactions)
		},
	}
	addFileFlag(cmd)
	return cmd
}

func (a *app) transactionDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete ID", Short: "Soft-delete a transaction", Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := requireYes(cmd); err != nil {
				return err
			}
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
			if err := application.DeleteTransaction(cmd.Context(), repo, entityID, transactionID); err != nil {
				return err
			}
			return a.print(cmd, nil, nil)
		},
	}
	addYesFlag(cmd)
	return cmd
}

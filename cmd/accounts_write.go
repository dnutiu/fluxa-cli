package cmd

import (
	"github.com/dnutiu/fluxa-cli/internal/application"
	"github.com/dnutiu/fluxa-cli/internal/domain"
	"github.com/dnutiu/fluxa-cli/internal/presentation"
	"github.com/spf13/cobra"
)

func (a *app) accountAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "add", Short: "Create an account with flags", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			entityID, err := a.entityID()
			if err != nil {
				return err
			}
			input, err := accountInput(cmd, false)
			if err != nil {
				return err
			}
			if err := input.ValidateCreate(); err != nil {
				return err
			}
			repo, err := a.repository()
			if err != nil {
				return err
			}
			result, err := application.AddAccount(cmd.Context(), repo, entityID, input)
			if err != nil {
				return err
			}
			return a.print(cmd, result, presentation.Accounts)
		},
	}
	addAccountFlags(cmd, false)
	return cmd
}

func (a *app) accountEditCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "edit ID", Short: "Edit account fields with flags", Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			entityID, err := a.entityID()
			if err != nil {
				return err
			}
			accountID, err := domain.ParseID(args[0])
			if err != nil {
				return err
			}
			input, err := accountInput(cmd, true)
			if err != nil {
				return err
			}
			if err := input.ValidateUpdate(); err != nil {
				return err
			}
			repo, err := a.repository()
			if err != nil {
				return err
			}
			result, err := application.EditAccount(cmd.Context(), repo, entityID, accountID, input)
			if err != nil {
				return err
			}
			return a.print(cmd, result, presentation.Accounts)
		},
	}
	addAccountFlags(cmd, true)
	return cmd
}

func (a *app) accountArchiveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "archive ID", Short: "Archive an account", Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := requireYes(cmd); err != nil {
				return err
			}
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
			if err := application.ArchiveAccount(cmd.Context(), repo, entityID, accountID); err != nil {
				return err
			}
			return a.print(cmd, nil, nil)
		},
	}
	addYesFlag(cmd)
	return cmd
}

func (a *app) accountRestoreCommand() *cobra.Command {
	return &cobra.Command{
		Use: "restore ID", Short: "Restore an archived account", Args: cobra.ExactArgs(1),
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
			result, err := application.RestoreAccount(cmd.Context(), repo, entityID, accountID)
			if err != nil {
				return err
			}
			return a.print(cmd, result, presentation.Accounts)
		},
	}
}

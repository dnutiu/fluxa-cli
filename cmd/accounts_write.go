package cmd

import (
	"github.com/dnutiu/fluxa-cli/internal/application"
	"github.com/dnutiu/fluxa-cli/internal/domain"
	"github.com/dnutiu/fluxa-cli/internal/presentation"
	"github.com/spf13/cobra"
)

func (a *app) accountAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "add", Short: "Add an account from a JSON file", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			entityID, err := a.entityID()
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
			result, err := application.AddAccount(cmd.Context(), repo, entityID, body)
			if err != nil {
				return err
			}
			return a.print(cmd, result, presentation.Accounts)
		},
	}
	addFileFlag(cmd)
	return cmd
}

func (a *app) accountEditCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "edit ID", Short: "Edit an account from a JSON file", Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			entityID, err := a.entityID()
			if err != nil {
				return err
			}
			accountID, err := domain.ParseID(args[0])
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
			result, err := application.EditAccount(cmd.Context(), repo, entityID, accountID, body)
			if err != nil {
				return err
			}
			return a.print(cmd, result, presentation.Accounts)
		},
	}
	addFileFlag(cmd)
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

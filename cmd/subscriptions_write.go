package cmd

import (
	"github.com/dnutiu/fluxa-cli/internal/application"
	"github.com/dnutiu/fluxa-cli/internal/domain"
	"github.com/dnutiu/fluxa-cli/internal/presentation"
	"github.com/spf13/cobra"
)

func (a *app) subscriptionAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "add", Short: "Create a subscription with flags", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			entityID, err := a.entityID()
			if err != nil {
				return err
			}
			input, err := subscriptionInput(cmd, false)
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
			result, err := application.AddSubscription(cmd.Context(), repo, entityID, input)
			if err != nil {
				return err
			}
			return a.print(cmd, result, presentation.Subscriptions)
		},
	}
	addSubscriptionFlags(cmd, false)
	return cmd
}

func (a *app) subscriptionEditCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "edit ID", Short: "Edit subscription fields with flags", Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			entityID, err := a.entityID()
			if err != nil {
				return err
			}
			subscriptionID, err := domain.ParseID(args[0])
			if err != nil {
				return err
			}
			input, err := subscriptionInput(cmd, true)
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
			result, err := application.EditSubscription(cmd.Context(), repo, entityID, subscriptionID, input)
			if err != nil {
				return err
			}
			return a.print(cmd, result, presentation.Subscriptions)
		},
	}
	addSubscriptionFlags(cmd, true)
	return cmd
}

func (a *app) subscriptionDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete ID", Short: "Delete a subscription", Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := requireYes(cmd); err != nil {
				return err
			}
			entityID, err := a.entityID()
			if err != nil {
				return err
			}
			subscriptionID, err := domain.ParseID(args[0])
			if err != nil {
				return err
			}
			repo, err := a.repository()
			if err != nil {
				return err
			}
			if err := application.DeleteSubscription(cmd.Context(), repo, entityID, subscriptionID); err != nil {
				return err
			}
			return a.print(cmd, nil, nil)
		},
	}
	addYesFlag(cmd)
	return cmd
}

func (a *app) subscriptionActivityCommand(name string, active bool) *cobra.Command {
	return &cobra.Command{
		Use: name + " ID", Short: name + " a subscription", Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			entityID, err := a.entityID()
			if err != nil {
				return err
			}
			subscriptionID, err := domain.ParseID(args[0])
			if err != nil {
				return err
			}
			repo, err := a.repository()
			if err != nil {
				return err
			}
			result, err := application.SetSubscriptionActive(cmd.Context(), repo, entityID, subscriptionID, active)
			if err != nil {
				return err
			}
			return a.print(cmd, result, presentation.Subscriptions)
		},
	}
}

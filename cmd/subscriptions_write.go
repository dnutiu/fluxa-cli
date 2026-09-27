package cmd

import (
	"github.com/dnutiu/fluxa-cli/internal/application"
	"github.com/dnutiu/fluxa-cli/internal/domain"
	"github.com/dnutiu/fluxa-cli/internal/presentation"
	"github.com/spf13/cobra"
)

func (a *app) subscriptionAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "add", Short: "Add a subscription from a JSON file", Args: cobra.NoArgs,
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
			result, err := application.AddSubscription(cmd.Context(), repo, entityID, body)
			if err != nil {
				return err
			}
			return a.print(cmd, result, presentation.Subscriptions)
		},
	}
	addFileFlag(cmd)
	return cmd
}

func (a *app) subscriptionEditCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "edit ID", Short: "Edit a subscription from a JSON file", Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			entityID, err := a.entityID()
			if err != nil {
				return err
			}
			subscriptionID, err := domain.ParseID(args[0])
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
			result, err := application.EditSubscription(cmd.Context(), repo, entityID, subscriptionID, body)
			if err != nil {
				return err
			}
			return a.print(cmd, result, presentation.Subscriptions)
		},
	}
	addFileFlag(cmd)
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

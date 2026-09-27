package cmd

import (
	"github.com/dnutiu/fluxa-cli/internal/application"
	"github.com/dnutiu/fluxa-cli/internal/domain"
	"github.com/dnutiu/fluxa-cli/internal/presentation"
	"github.com/spf13/cobra"
)

func (a *app) subscriptionsCommand() *cobra.Command {
	parent := &cobra.Command{Use: "subscriptions", Short: "Manage subscriptions"}
	parent.AddCommand(a.subscriptionListCommand(), a.subscriptionShowCommand(), a.subscriptionAddCommand(), a.subscriptionEditCommand(), a.subscriptionDeleteCommand(), a.subscriptionActivityCommand("pause", false), a.subscriptionActivityCommand("resume", true))
	return parent
}

func (a *app) subscriptionListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list", Short: "List subscriptions", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			entityID, err := a.entityID()
			if err != nil {
				return err
			}
			page, _ := cmd.Flags().GetInt("page")
			perPage, _ := cmd.Flags().GetInt("per-page")
			updatedSince, _ := cmd.Flags().GetString("updated-since")
			filter := domain.SubscriptionFilter{Pagination: domain.Pagination{Page: page, PerPage: perPage}, UpdatedSince: updatedSince}
			if cmd.Flags().Changed("active") {
				active, _ := cmd.Flags().GetBool("active")
				filter.Active = &active
			}
			repo, err := a.repository()
			if err != nil {
				return err
			}
			result, err := application.ListSubscriptions(cmd.Context(), repo, entityID, filter)
			if err != nil {
				return err
			}
			return a.print(cmd, result, presentation.Subscriptions)
		},
	}
	cmd.Flags().Int("page", 1, "page number")
	cmd.Flags().Int("per-page", 25, "results per page (maximum 100)")
	cmd.Flags().Bool("active", false, "filter active or paused subscriptions; use --active=false for paused")
	cmd.Flags().String("updated-since", "", "RFC 3339 timestamp")
	return cmd
}

func (a *app) subscriptionShowCommand() *cobra.Command {
	return &cobra.Command{
		Use: "show ID", Short: "Show a subscription", Args: cobra.ExactArgs(1),
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
			result, err := application.ShowSubscription(cmd.Context(), repo, entityID, subscriptionID)
			if err != nil {
				return err
			}
			return a.print(cmd, result, presentation.Subscriptions)
		},
	}
}

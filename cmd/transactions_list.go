package cmd

import (
	"fluxa-cli/internal/application"
	"fluxa-cli/internal/domain"
	"fluxa-cli/internal/presentation"
	"github.com/spf13/cobra"
)

func (a *app) transactionListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list", Short: "List transactions", Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			entityID, err := a.entityID()
			if err != nil {
				return err
			}
			page, _ := cmd.Flags().GetInt("page")
			perPage, _ := cmd.Flags().GetInt("per-page")
			categoryID, _ := cmd.Flags().GetInt("category-id")
			accountID, _ := cmd.Flags().GetInt("account-id")
			kind, _ := cmd.Flags().GetString("kind")
			dateFrom, _ := cmd.Flags().GetString("date-from")
			dateTo, _ := cmd.Flags().GetString("date-to")
			updatedSince, _ := cmd.Flags().GetString("updated-since")
			includeDeleted, _ := cmd.Flags().GetBool("include-deleted")
			filter := domain.TransactionFilter{
				Pagination: domain.Pagination{Page: page, PerPage: perPage},
				CategoryID: domain.ID(categoryID), AccountID: domain.ID(accountID), Kind: kind,
				DateFrom: dateFrom, DateTo: dateTo, UpdatedSince: updatedSince,
				IncludeDeleted: includeDeleted,
			}
			repo, err := a.repository()
			if err != nil {
				return err
			}
			result, err := application.ListTransactions(cmd.Context(), repo, entityID, filter)
			if err != nil {
				return err
			}
			return a.print(cmd, result, presentation.Transactions)
		},
	}
	cmd.Flags().Int("page", 1, "page number")
	cmd.Flags().Int("per-page", 25, "results per page (maximum 100)")
	cmd.Flags().Int("category-id", 0, "filter by category ID")
	cmd.Flags().Int("account-id", 0, "filter by account ID")
	cmd.Flags().String("kind", "", "income or expense")
	cmd.Flags().String("date-from", "", "inclusive start date (YYYY-MM-DD)")
	cmd.Flags().String("date-to", "", "inclusive end date (YYYY-MM-DD)")
	cmd.Flags().String("updated-since", "", "RFC 3339 timestamp")
	cmd.Flags().Bool("include-deleted", false, "include soft-deleted records")
	return cmd
}

package cmd

import (
	"github.com/dnutiu/fluxa-cli/internal/application"
	"github.com/dnutiu/fluxa-cli/internal/domain"
	"github.com/dnutiu/fluxa-cli/internal/presentation"
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
	cmd.Flags().IntP("page", "p", 1, "page number")
	cmd.Flags().IntP("per-page", "l", 25, "results per page (maximum 100)")
	cmd.Flags().IntP("category-id", "g", 0, "filter by category ID")
	cmd.Flags().IntP("account-id", "a", 0, "filter by account ID")
	cmd.Flags().StringP("kind", "k", "", "income or expense")
	cmd.Flags().StringP("date-from", "f", "", "inclusive start date (YYYY-MM-DD)")
	cmd.Flags().StringP("date-to", "t", "", "inclusive end date (YYYY-MM-DD)")
	cmd.Flags().StringP("updated-since", "s", "", "RFC 3339 timestamp")
	cmd.Flags().BoolP("include-deleted", "D", false, "include soft-deleted records")
	return cmd
}

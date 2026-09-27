package cmd

import (
	"github.com/dnutiu/fluxa-cli/internal/domain"
	"github.com/spf13/cobra"
)

func addTransactionFlags(cmd *cobra.Command, editing bool) {
	f := cmd.Flags()
	f.Int64P("category-id", "g", 0, "expense or income category ID")
	f.Int64P("account-id", "a", 0, "account ID")
	f.StringP("amount", "m", "", "nonzero decimal amount, such as 125.00")
	f.StringP("currency", "c", "", "RON, EUR, USD, GBP, or CHF")
	f.StringP("exchange-rate", "r", "", "positive decimal exchange rate")
	f.StringP("date", "d", "", "transaction date (YYYY-MM-DD)")
	f.StringP("occurred-at", "t", "", "exact time (RFC 3339)")
	f.StringP("description", "n", "", "transaction note")
	f.BoolP("exclude-from-analytics", "x", false, "exclude this transaction from analytics")
	if editing {
		f.BoolP("clear-account", "A", false, "unlink the account")
		f.BoolP("clear-description", "N", false, "remove the description")
	}
}

func transactionInput(cmd *cobra.Command, editing bool) (domain.TransactionInput, error) {
	input := domain.TransactionInput{
		CategoryID:           idField(cmd, "category-id"),
		AccountID:            idField(cmd, "account-id"),
		Amount:               stringField(cmd, "amount"),
		Currency:             stringField(cmd, "currency"),
		ExchangeRate:         stringField(cmd, "exchange-rate"),
		Date:                 stringField(cmd, "date"),
		OccurredAt:           stringField(cmd, "occurred-at"),
		Description:          stringField(cmd, "description"),
		ExcludeFromAnalytics: boolField(cmd, "exclude-from-analytics"),
	}
	if !editing {
		return input, nil
	}
	var err error
	input.AccountID, err = nullableIDField(cmd, "account-id", "clear-account")
	if err != nil {
		return input, err
	}
	input.Description, err = nullableStringField(cmd, "description", "clear-description")
	if err != nil {
		return input, err
	}
	return input, nil
}

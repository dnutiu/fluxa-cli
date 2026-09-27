package cmd

import (
	"github.com/dnutiu/fluxa-cli/internal/domain"
	"github.com/spf13/cobra"
)

func addSubscriptionFlags(cmd *cobra.Command, editing bool) {
	f := cmd.Flags()
	f.StringP("name", "n", "", "subscription name")
	f.Int64P("account-id", "a", 0, "linked account ID")
	f.StringP("amount", "m", "", "positive decimal amount")
	f.StringP("currency", "c", "", "RON, EUR, USD, GBP, or CHF")
	f.StringP("exchange-rate", "r", "", "positive decimal exchange rate")
	f.StringP("start-date", "s", "", "first billing date (YYYY-MM-DD)")
	f.StringP("recurrence-interval", "i", "", "monthly or yearly")
	f.BoolP("active", "A", true, "active status; use --active=false to pause")
	f.BoolP("include-vat", "v", false, "include VAT in calculated amount")
	f.BoolP("informative", "f", false, "mark as informative")
	if editing {
		f.BoolP("clear-account", "G", false, "unlink the account")
	}
}

func subscriptionInput(cmd *cobra.Command, editing bool) (domain.SubscriptionInput, error) {
	input := domain.SubscriptionInput{
		Name:               stringField(cmd, "name"),
		AccountID:          idField(cmd, "account-id"),
		Amount:             stringField(cmd, "amount"),
		Currency:           stringField(cmd, "currency"),
		ExchangeRate:       stringField(cmd, "exchange-rate"),
		StartDate:          stringField(cmd, "start-date"),
		RecurrenceInterval: stringField(cmd, "recurrence-interval"),
		Active:             boolField(cmd, "active"),
		IncludeVAT:         boolField(cmd, "include-vat"),
		Informative:        boolField(cmd, "informative"),
	}
	if editing {
		account, err := nullableIDField(cmd, "account-id", "clear-account")
		if err != nil {
			return input, err
		}
		input.AccountID = account
	}
	return input, nil
}

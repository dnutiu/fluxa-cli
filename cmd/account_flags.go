package cmd

import (
	"github.com/dnutiu/fluxa-cli/internal/domain"
	"github.com/spf13/cobra"
)

func addAccountFlags(cmd *cobra.Command, editing bool) {
	f := cmd.Flags()
	f.StringP("name", "n", "", "account name")
	f.StringP("kind", "k", "", "cash, bank, card, savings, or investment")
	f.StringP("currency", "c", "", "RON, EUR, USD, GBP, or CHF")
	f.StringP("opening-balance", "b", "", "opening balance as a decimal string")
	f.StringP("annual-interest", "i", "", "annual interest percentage")
	f.StringP("interest-tax", "t", "", "interest tax percentage")
	f.Int64P("account-group-id", "g", 0, "account group ID")
	if editing {
		f.BoolP("clear-account-group", "G", false, "remove the account group")
	}
}

func accountInput(cmd *cobra.Command, editing bool) (domain.AccountInput, error) {
	input := domain.AccountInput{
		Name:           stringField(cmd, "name"),
		Kind:           stringField(cmd, "kind"),
		Currency:       stringField(cmd, "currency"),
		OpeningBalance: stringField(cmd, "opening-balance"),
		AnnualInterest: stringField(cmd, "annual-interest"),
		InterestTax:    stringField(cmd, "interest-tax"),
		AccountGroupID: idField(cmd, "account-group-id"),
	}
	if editing {
		group, err := nullableIDField(cmd, "account-group-id", "clear-account-group")
		if err != nil {
			return input, err
		}
		input.AccountGroupID = group
	}
	return input, nil
}

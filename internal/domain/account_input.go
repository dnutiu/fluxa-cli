package domain

import "errors"

type AccountInput struct {
	Name           Field[string]
	Kind           Field[string]
	Currency       Field[string]
	OpeningBalance Field[string]
	AnnualInterest Field[string]
	InterestTax    Field[string]
	AccountGroupID Field[ID]
}

func (in AccountInput) HasChanges() bool {
	return in.Name.Set || in.Kind.Set || in.Currency.Set || in.OpeningBalance.Set ||
		in.AnnualInterest.Set || in.InterestTax.Set || in.AccountGroupID.Set
}

func (in AccountInput) ValidateCreate() error { return in.validate(true) }

func (in AccountInput) ValidateUpdate() error { return in.validate(false) }

func (in AccountInput) validate(creating bool) error {
	if !creating && !in.HasChanges() {
		return errors.New("provide at least one account field to edit")
	}
	if err := requireName(in.Name, creating); err != nil {
		return err
	}
	if err := validateChoice("kind", in.Kind, "cash", "bank", "card", "savings", "investment"); err != nil {
		return err
	}
	if err := validateCurrency(in.Currency); err != nil {
		return err
	}
	for _, field := range []struct {
		name  string
		value Field[string]
	}{
		{"opening-balance", in.OpeningBalance},
		{"annual-interest", in.AnnualInterest},
		{"interest-tax", in.InterestTax},
	} {
		if err := validateDecimal(field.name, field.value); err != nil {
			return err
		}
	}
	return validateID("account-group-id", in.AccountGroupID)
}

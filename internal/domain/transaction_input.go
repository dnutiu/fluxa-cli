package domain

import "errors"

type TransactionInput struct {
	CategoryID           Field[ID]
	AccountID            Field[ID]
	Amount               Field[string]
	Currency             Field[string]
	ExchangeRate         Field[string]
	Date                 Field[string]
	OccurredAt           Field[string]
	Description          Field[string]
	ExcludeFromAnalytics Field[bool]
}

func (in TransactionInput) HasChanges() bool {
	return in.CategoryID.Set || in.AccountID.Set || in.Amount.Set ||
		in.Currency.Set || in.ExchangeRate.Set || in.Date.Set ||
		in.OccurredAt.Set || in.Description.Set || in.ExcludeFromAnalytics.Set
}

func (in TransactionInput) ValidateCreate() error { return in.validate(true) }

func (in TransactionInput) ValidateUpdate() error { return in.validate(false) }

func (in TransactionInput) validate(creating bool) error {
	if !creating && !in.HasChanges() {
		return errors.New("provide at least one transaction field to edit")
	}
	if creating && !in.CategoryID.Set {
		return errors.New("category-id is required")
	}
	if in.CategoryID.Null {
		return errors.New("category-id cannot be cleared")
	}
	if err := validateID("category-id", in.CategoryID); err != nil {
		return err
	}
	if err := validateID("account-id", in.AccountID); err != nil {
		return err
	}
	if err := requireDecimal("amount", in.Amount, creating, false); err != nil {
		return err
	}
	if err := validateCurrency(in.Currency); err != nil {
		return err
	}
	if err := validatePositiveDecimal("exchange-rate", in.ExchangeRate); err != nil {
		return err
	}
	if creating && !in.Date.Set && (!in.OccurredAt.Set || in.OccurredAt.Null) {
		return errors.New("date or occurred-at is required")
	}
	if err := validateDate("date", in.Date); err != nil {
		return err
	}
	return validateTimestamp("occurred-at", in.OccurredAt)
}

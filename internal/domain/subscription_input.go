package domain

import "errors"

type SubscriptionInput struct {
	Name               Field[string]
	AccountID          Field[ID]
	Amount             Field[string]
	Currency           Field[string]
	ExchangeRate       Field[string]
	StartDate          Field[string]
	RecurrenceInterval Field[string]
	Active             Field[bool]
	IncludeVAT         Field[bool]
	Informative        Field[bool]
}

func (in SubscriptionInput) HasChanges() bool {
	return in.Name.Set || in.AccountID.Set || in.Amount.Set || in.Currency.Set ||
		in.ExchangeRate.Set || in.StartDate.Set || in.RecurrenceInterval.Set ||
		in.Active.Set || in.IncludeVAT.Set || in.Informative.Set
}

func (in SubscriptionInput) ValidateCreate() error { return in.validate(true) }

func (in SubscriptionInput) ValidateUpdate() error { return in.validate(false) }

func (in SubscriptionInput) validate(creating bool) error {
	if !creating && !in.HasChanges() {
		return errors.New("provide at least one subscription field to edit")
	}
	if err := requireName(in.Name, creating); err != nil {
		return err
	}
	if err := validateID("account-id", in.AccountID); err != nil {
		return err
	}
	if err := requireDecimal("amount", in.Amount, creating, true); err != nil {
		return err
	}
	if err := validateCurrency(in.Currency); err != nil {
		return err
	}
	if err := validatePositiveDecimal("exchange-rate", in.ExchangeRate); err != nil {
		return err
	}
	if creating && !in.StartDate.Set {
		return errors.New("start-date is required")
	}
	if err := validateDate("start-date", in.StartDate); err != nil {
		return err
	}
	return validateChoice("recurrence-interval", in.RecurrenceInterval, "monthly", "yearly")
}

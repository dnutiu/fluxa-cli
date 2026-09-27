package domain

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"
	"time"
)

var decimalPattern = regexp.MustCompile(`^[+-]?[0-9]+(\.[0-9]+)?$`)

func requireName(field Field[string], creating bool) error {
	if creating && !field.Set {
		return fmt.Errorf("name is required")
	}
	if field.Set && strings.TrimSpace(field.Value) == "" {
		return fmt.Errorf("name cannot be empty")
	}
	return nil
}

func requireDecimal(name string, field Field[string], creating, positive bool) error {
	if creating && !field.Set {
		return fmt.Errorf("%s is required", name)
	}
	if !field.Set {
		return nil
	}
	if field.Null || !decimalPattern.MatchString(field.Value) {
		return fmt.Errorf("%s must be a decimal number", name)
	}
	amount, ok := new(big.Rat).SetString(field.Value)
	if !ok || amount.Sign() == 0 || positive && amount.Sign() < 0 {
		if positive {
			return fmt.Errorf("%s must be greater than zero", name)
		}
		return fmt.Errorf("%s must be nonzero", name)
	}
	return nil
}

func validateDecimal(name string, field Field[string]) error {
	if !field.Set {
		return nil
	}
	if field.Null || !decimalPattern.MatchString(field.Value) {
		return fmt.Errorf("%s must be a decimal number", name)
	}
	return nil
}

func validatePositiveDecimal(name string, field Field[string]) error {
	if err := validateDecimal(name, field); err != nil {
		return err
	}
	if field.Set {
		amount, _ := new(big.Rat).SetString(field.Value)
		if amount.Sign() <= 0 {
			return fmt.Errorf("%s must be greater than zero", name)
		}
	}
	return nil
}

func validateChoice(name string, field Field[string], allowed ...string) error {
	if !field.Set {
		return nil
	}
	for _, option := range allowed {
		if field.Value == option {
			return nil
		}
	}
	return fmt.Errorf("%s must be one of: %s", name, strings.Join(allowed, ", "))
}

func validateCurrency(field Field[string]) error {
	return validateChoice("currency", field, "RON", "EUR", "USD", "GBP", "CHF")
}

func validateDate(name string, field Field[string]) error {
	if !field.Set {
		return nil
	}
	date, err := time.Parse("2006-01-02", field.Value)
	if field.Null || err != nil || date.Format("2006-01-02") != field.Value {
		return fmt.Errorf("%s must use YYYY-MM-DD", name)
	}
	return nil
}

func validateTimestamp(name string, field Field[string]) error {
	if !field.Set || field.Null {
		return nil
	}
	if _, err := time.Parse(time.RFC3339, field.Value); err != nil {
		return fmt.Errorf("%s must be an RFC 3339 timestamp", name)
	}
	return nil
}

func validateID(name string, field Field[ID]) error {
	if field.Set && !field.Null && !field.Value.Valid() {
		return fmt.Errorf("%s must be a positive integer", name)
	}
	return nil
}

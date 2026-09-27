package domain

import "testing"

func TestTypedWriteValidation(t *testing.T) {
	validTransaction := TransactionInput{CategoryID: With(ID(12)), Amount: With("125.00"), Date: With("2026-09-27")}
	for _, tc := range []struct {
		name string
		err  error
	}{
		{"valid transaction", validTransaction.ValidateCreate()},
		{"valid account", (AccountInput{Name: With("Cash"), OpeningBalance: With("0.00")}).ValidateCreate()},
		{"valid subscription", (SubscriptionInput{Name: With("Cloud"), Amount: With("12.50"), StartDate: With("2026-09-27")}).ValidateCreate()},
	} {
		if tc.err != nil {
			t.Errorf("%s: %v", tc.name, tc.err)
		}
	}
	for _, tc := range []struct {
		name string
		err  error
	}{
		{"missing account name", (AccountInput{}).ValidateCreate()},
		{"empty account patch", (AccountInput{}).ValidateUpdate()},
		{"zero transaction amount", (TransactionInput{CategoryID: With(ID(12)), Amount: With("0.00"), Date: With("2026-09-27")}).ValidateCreate()},
		{"fraction transaction amount", (TransactionInput{CategoryID: With(ID(12)), Amount: With("1/2"), Date: With("2026-09-27")}).ValidateCreate()},
		{"negative subscription amount", (SubscriptionInput{Name: With("Cloud"), Amount: With("-1.00"), StartDate: With("2026-09-27")}).ValidateCreate()},
		{"cleared category", (TransactionInput{CategoryID: Cleared[ID](), Amount: With("1.00"), Date: With("2026-09-27")}).ValidateCreate()},
	} {
		if tc.err == nil {
			t.Errorf("%s: expected validation error", tc.name)
		}
	}
}

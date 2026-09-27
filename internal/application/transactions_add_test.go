package application

import (
	"context"
	"testing"

	"github.com/dnutiu/fluxa-cli/internal/domain"
)

type transactionCreatorSpy struct {
	calls int
	input domain.TransactionInput
	key   string
}

func (s *transactionCreatorSpy) CreateTransaction(_ context.Context, _ domain.ID, input domain.TransactionInput, key string) (Record[domain.Transaction], error) {
	s.calls++
	s.input = input
	s.key = key
	return Record[domain.Transaction]{Data: domain.Transaction{ID: 42}}, nil
}

func validTransactionInput() domain.TransactionInput {
	return domain.TransactionInput{
		CategoryID: domain.With(domain.ID(12)),
		Amount:     domain.With("125.00"),
		Date:       domain.With("2026-09-27"),
	}
}

func TestAddTransactionValidatesBeforeCallingRepository(t *testing.T) {
	for _, tc := range []struct {
		name  string
		id    domain.ID
		input domain.TransactionInput
		key   string
	}{
		{"entity", 0, validTransactionInput(), "key"},
		{"category", 7, domain.TransactionInput{Amount: domain.With("12.00"), Date: domain.With("2026-09-27")}, "key"},
		{"amount", 7, domain.TransactionInput{CategoryID: domain.With(domain.ID(12)), Amount: domain.With("oops"), Date: domain.With("2026-09-27")}, "key"},
		{"key", 7, validTransactionInput(), " "},
	} {
		t.Run(tc.name, func(t *testing.T) {
			spy := &transactionCreatorSpy{}
			if _, err := AddTransaction(context.Background(), spy, tc.id, tc.input, tc.key); err == nil {
				t.Fatal("expected validation error")
			}
			if spy.calls != 0 {
				t.Fatalf("repository called %d times", spy.calls)
			}
		})
	}
}

func TestAddTransactionForwardsTypedInputAndKey(t *testing.T) {
	spy := &transactionCreatorSpy{}
	input := validTransactionInput()
	result, err := AddTransaction(context.Background(), spy, 7, input, "attempt-1")
	if err != nil {
		t.Fatal(err)
	}
	if spy.calls != 1 || spy.input.Amount.Value != "125.00" || spy.key != "attempt-1" || result.Data.ID != 42 {
		t.Fatalf("spy = %+v, result = %+v", spy, result)
	}
}

package application

import (
	"context"
	"testing"

	"fluxa-cli/internal/domain"
)

type transactionCreatorSpy struct {
	calls int
	body  string
	key   string
}

func (s *transactionCreatorSpy) CreateTransaction(_ context.Context, _ domain.ID, body []byte, key string) (Record[domain.Transaction], error) {
	s.calls++
	s.body = string(body)
	s.key = key
	return Record[domain.Transaction]{Data: domain.Transaction{ID: 42}}, nil
}

func TestAddTransactionValidatesBeforeCallingRepository(t *testing.T) {
	for _, tc := range []struct {
		name string
		id   domain.ID
		body string
		key  string
	}{
		{"entity", 0, `{"amount":"12.00"}`, "key"},
		{"body", 7, `[{"amount":"12.00"}]`, "key"},
		{"key", 7, `{"amount":"12.00"}`, " "},
	} {
		t.Run(tc.name, func(t *testing.T) {
			spy := &transactionCreatorSpy{}
			if _, err := AddTransaction(context.Background(), spy, tc.id, []byte(tc.body), tc.key); err == nil {
				t.Fatal("expected validation error")
			}
			if spy.calls != 0 {
				t.Fatalf("repository called %d times", spy.calls)
			}
		})
	}
}

func TestAddTransactionForwardsExactBodyAndKey(t *testing.T) {
	spy := &transactionCreatorSpy{}
	body := `{"amount":"125.00"}`
	result, err := AddTransaction(context.Background(), spy, 7, []byte(body), "attempt-1")
	if err != nil {
		t.Fatal(err)
	}
	if spy.calls != 1 || spy.body != body || spy.key != "attempt-1" || result.Data.ID != 42 {
		t.Fatalf("spy = %+v, result = %+v", spy, result)
	}
}

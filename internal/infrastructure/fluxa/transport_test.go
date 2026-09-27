package fluxa

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dnutiu/fluxa-cli/internal/domain"
)

func TestCreateTransactionSendsBearerAndIdempotencyKey(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/api/v1/entities/7/transactions" {
			t.Errorf("request = %s %s", r.Method, r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer secret" || r.Header.Get("Idempotency-Key") != "attempt-1" {
			t.Errorf("missing authentication or idempotency header")
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("content type = %q", r.Header.Get("Content-Type"))
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"data":{"id":42,"amount":"125.00"}}`))
	}))
	defer server.Close()
	repo, err := NewRepository(server.URL, "secret", nil)
	if err != nil {
		t.Fatal(err)
	}
	result, err := repo.CreateTransaction(context.Background(), 7, domain.TransactionInput{
		CategoryID: domain.With(domain.ID(12)),
		Amount:     domain.With("125.00"),
		Date:       domain.With("2026-09-27"),
	}, "attempt-1")
	if err != nil {
		t.Fatal(err)
	}
	if result.Data.ID != 42 || result.Data.Amount != "125.00" {
		t.Fatalf("transaction = %+v", result.Data)
	}
}

func TestListAccountKeepsRONBalance(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("include_archived") != "true" {
			t.Errorf("query = %v", r.URL.Query())
		}
		_, _ = w.Write([]byte(`{"data":[{"id":3,"balance":"10.00","balance_ron":"49.75"}]}`))
	}))
	defer server.Close()
	repo, err := NewRepository(server.URL, "secret", nil)
	if err != nil {
		t.Fatal(err)
	}
	result, err := repo.ListAccounts(context.Background(), 7, domain.AccountFilter{IncludeArchived: true})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Data) != 1 || result.Data[0].BalanceRON != "49.75" {
		t.Fatalf("accounts = %+v", result.Data)
	}
}

func TestAPIErrorDoesNotExposeCredential(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Retry-After", "60")
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`{"error":{"code":"rate_limited","message":"Slow down"}}`))
	}))
	defer server.Close()
	repo, err := NewRepository(server.URL, "secret", nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = repo.ListEntities(context.Background())
	var apiErr *Error
	if !errors.As(err, &apiErr) || apiErr.Status != 429 || apiErr.Code != "rate_limited" {
		t.Fatalf("error = %v", err)
	}
	if !strings.Contains(err.Error(), "retry after 60") || strings.Contains(err.Error(), "secret") {
		t.Fatalf("error text = %q", err)
	}
}

func TestRejectsInsecureRemoteURL(t *testing.T) {
	if _, err := NewRepository("http://example.com", "secret", nil); err == nil {
		t.Fatal("accepted insecure remote URL")
	}
	if _, err := NewRepository("https://user:pass@example.com", "secret", nil); err == nil {
		t.Fatal("accepted URL credentials")
	}
}

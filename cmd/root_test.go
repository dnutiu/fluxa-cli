package cmd

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func runCommand(t *testing.T, args ...string) (string, error) {
	t.Helper()
	command := NewRootCommand()
	var out bytes.Buffer
	command.SetOut(&out)
	command.SetErr(&out)
	command.SetArgs(args)
	err := command.Execute()
	return out.String(), err
}

func TestTransactionCreateKeepsExactAmountAndIdempotencyKey(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/api/v1/entities/7/transactions" {
			t.Errorf("request = %s %s", r.Method, r.URL.Path)
		}
		if got := r.Header.Get("Idempotency-Key"); got != "fixed-key" {
			t.Errorf("idempotency key = %q", got)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer test-key" {
			t.Errorf("authorization = %q", got)
		}
		var body struct {
			CategoryID int    `json:"category_id"`
			Amount     string `json:"amount"`
			Date       string `json:"date"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Error(err)
		}
		if body.CategoryID != 12 || body.Amount != "125.00" || body.Date != "2026-09-27" {
			t.Errorf("body = %+v", body)
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"data":{"id":42,"amount":"125.00"}}`))
	}))
	defer server.Close()
	t.Setenv("FLUXA_API_KEY", "test-key")
	out, err := runCommand(t, "-C", filepath.Join(t.TempDir(), "none.yaml"), "-u", server.URL, "-e", "7", "-o", "json", "transactions", "add", "-g", "12", "-m", "125.00", "-d", "2026-09-27", "-i", "fixed-key")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, `"amount": "125.00"`) {
		t.Fatalf("output = %q", out)
	}
}

func TestTransactionListMapsFilters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		if query.Get("date_from") != "2026-09-01" || query.Get("include_deleted") != "true" || query.Get("per_page") != "50" {
			t.Errorf("query = %v", query)
		}
		_, _ = w.Write([]byte(`{"data":[],"meta":{"page":1,"per_page":50,"total":0,"pages":0}}`))
	}))
	defer server.Close()
	t.Setenv("FLUXA_API_KEY", "test-key")
	_, err := runCommand(t, "-C", filepath.Join(t.TempDir(), "none.yaml"), "-u", server.URL, "-e", "7", "transactions", "list", "-f", "2026-09-01", "-D", "-l", "50")
	if err != nil {
		t.Fatal(err)
	}
}

func TestArchiveNeedsExplicitYes(t *testing.T) {
	output, err := runCommand(t, "accounts", "archive", "4")
	if err == nil || !strings.Contains(err.Error(), "--yes") {
		t.Fatalf("output %q, error %v", output, err)
	}
}

func TestConfigDoesNotPersistAPIKey(t *testing.T) {
	t.Setenv("FLUXA_API_KEY", "private-token")
	path := filepath.Join(t.TempDir(), "config.yaml")
	if _, err := runCommand(t, "--config", path, "config", "set-entity", "7"); err != nil {
		t.Fatal(err)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(content), "private-token") {
		t.Fatal("config contains the API key")
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm() != 0600 {
		t.Fatalf("config permissions = %v", info.Mode().Perm())
	}
}

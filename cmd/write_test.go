package cmd

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestFlagWritesSendOnlySuppliedFields(t *testing.T) {
	for _, tc := range []struct {
		name   string
		method string
		path   string
		args   []string
		body   map[string]any
	}{
		{
			"account add", http.MethodPost, "/api/v1/entities/7/accounts",
			[]string{"accounts", "add", "-n", "Everyday cash", "-k", "cash", "-c", "RON", "-b", "100.00"},
			map[string]any{"name": "Everyday cash", "kind": "cash", "currency": "RON", "opening_balance": "100.00"},
		},
		{
			"account edit clears group", http.MethodPatch, "/api/v1/entities/7/accounts/4",
			[]string{"accounts", "edit", "4", "-G"},
			map[string]any{"account_group_id": nil},
		},
		{
			"transaction edit clears and sets false", http.MethodPatch, "/api/v1/entities/7/transactions/42",
			[]string{"transactions", "edit", "42", "-x=false", "-N", "-a", "4"},
			map[string]any{"exclude_from_analytics": false, "description": nil, "account_id": float64(4)},
		},
		{
			"subscription add", http.MethodPost, "/api/v1/entities/7/subscriptions",
			[]string{"subscriptions", "add", "-n", "Cloud", "-m", "12.50", "-s", "2026-09-27", "-A=false"},
			map[string]any{"name": "Cloud", "amount": "12.50", "start_date": "2026-09-27", "active": false},
		},
		{
			"subscription edit clears and sets false", http.MethodPatch, "/api/v1/entities/7/subscriptions/12",
			[]string{"subscriptions", "edit", "12", "-v=false", "-G"},
			map[string]any{"include_vat": false, "account_id": nil},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != tc.method || r.URL.Path != tc.path {
					t.Errorf("request = %s %s", r.Method, r.URL.Path)
				}
				var body map[string]any
				if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(body, tc.body) {
					t.Errorf("body = %#v, want %#v", body, tc.body)
				}
				_, _ = w.Write([]byte(`{"data":{"id":42}}`))
			}))
			defer server.Close()
			t.Setenv("FLUXA_API_KEY", "test-key")
			args := append([]string{"-C", filepath.Join(t.TempDir(), "none.yaml"), "-u", server.URL, "-e", "7"}, tc.args...)
			if _, err := runCommand(t, args...); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestWriteFlagsValidateWithoutSending(t *testing.T) {
	for _, tc := range []struct {
		name string
		args []string
		want string
	}{
		{"account name", []string{"accounts", "add"}, "name is required"},
		{"empty edit", []string{"accounts", "edit", "4"}, "at least one account field"},
		{"transaction category", []string{"transactions", "add", "-m", "12.00", "-d", "2026-09-27"}, "category-id is required"},
		{"invalid money", []string{"transactions", "add", "-g", "12", "-m", "12,00", "-d", "2026-09-27"}, "decimal number"},
		{"conflicting account flags", []string{"transactions", "edit", "42", "-a", "4", "-A"}, "cannot be used together"},
		{"subscription date", []string{"subscriptions", "add", "-n", "Cloud", "-m", "12.50", "-s", "yesterday"}, "YYYY-MM-DD"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			args := append([]string{"-e", "7"}, tc.args...)
			_, err := runCommand(t, args...)
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("error = %v, want %q", err, tc.want)
			}
		})
	}
}

package domain

import "testing"

func TestTransactionFilterRejectsInvalidValues(t *testing.T) {
	base := TransactionFilter{Pagination: Pagination{Page: 1, PerPage: 25}}
	for _, tc := range []struct {
		name   string
		change func(*TransactionFilter)
	}{
		{"category ID", func(f *TransactionFilter) { f.CategoryID = -1 }},
		{"account ID", func(f *TransactionFilter) { f.AccountID = -1 }},
		{"kind", func(f *TransactionFilter) { f.Kind = "transfer" }},
		{"date order", func(f *TransactionFilter) { f.DateFrom, f.DateTo = "2026-09-30", "2026-09-01" }},
		{"timestamp", func(f *TransactionFilter) { f.UpdatedSince = "yesterday" }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			filter := base
			tc.change(&filter)
			if err := filter.Validate(); err == nil {
				t.Fatal("expected validation error")
			}
		})
	}
}

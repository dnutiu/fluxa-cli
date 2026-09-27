package domain

import (
	"errors"
	"time"
)

type Transaction struct {
	ID                   ID      `json:"id"`
	EntityID             ID      `json:"entity_id"`
	CategoryID           ID      `json:"category_id"`
	AccountID            *ID     `json:"account_id"`
	Kind                 string  `json:"kind"`
	Amount               string  `json:"amount"`
	AmountBase           string  `json:"amount_base"`
	Currency             string  `json:"currency"`
	ExchangeRate         string  `json:"exchange_rate"`
	Date                 string  `json:"date"`
	OccurredAt           *string `json:"occurred_at"`
	Description          *string `json:"description"`
	ExcludeFromAnalytics bool    `json:"exclude_from_analytics"`
	Deleted              bool    `json:"deleted"`
	DeletedAt            *string `json:"deleted_at"`
	CreatedAt            *string `json:"created_at"`
	UpdatedAt            *string `json:"updated_at"`
}

type TransactionFilter struct {
	Pagination     Pagination
	CategoryID     ID
	AccountID      ID
	Kind           string
	DateFrom       string
	DateTo         string
	UpdatedSince   string
	IncludeDeleted bool
}

func (f TransactionFilter) Validate() error {
	if err := f.Pagination.Validate(); err != nil {
		return err
	}
	if f.CategoryID < 0 || f.AccountID < 0 {
		return errors.New("filter IDs must be positive integers")
	}
	if f.Kind != "" && f.Kind != "income" && f.Kind != "expense" {
		return errors.New("kind must be income or expense")
	}
	for _, date := range []string{f.DateFrom, f.DateTo} {
		if date == "" {
			continue
		}
		if parsed, err := time.Parse("2006-01-02", date); err != nil || parsed.Format("2006-01-02") != date {
			return errors.New("dates must use YYYY-MM-DD")
		}
	}
	if f.DateFrom != "" && f.DateTo != "" && f.DateFrom > f.DateTo {
		return errors.New("date-from must not be after date-to")
	}
	if f.UpdatedSince != "" {
		if _, err := time.Parse(time.RFC3339, f.UpdatedSince); err != nil {
			return errors.New("updated-since must be an RFC 3339 timestamp")
		}
	}
	return nil
}

package domain

import (
	"errors"
	"time"
)

type Subscription struct {
	ID                 ID      `json:"id"`
	EntityID           ID      `json:"entity_id"`
	AccountID          *ID     `json:"account_id"`
	Name               string  `json:"name"`
	Amount             string  `json:"amount"`
	AmountBase         string  `json:"amount_base"`
	Currency           string  `json:"currency"`
	ExchangeRate       string  `json:"exchange_rate"`
	RecurrenceInterval string  `json:"recurrence_interval"`
	StartDate          string  `json:"start_date"`
	NextOccurrenceDate string  `json:"next_occurrence_date"`
	LastOccurrenceDate *string `json:"last_occurrence_date"`
	Active             bool    `json:"active"`
	IncludeVAT         bool    `json:"include_vat"`
	Informative        bool    `json:"informative"`
	MonthlyPriceBase   string  `json:"monthly_price_base"`
	YearlyPriceBase    string  `json:"yearly_price_base"`
	CreatedAt          *string `json:"created_at"`
	UpdatedAt          *string `json:"updated_at"`
}

type SubscriptionFilter struct {
	Pagination   Pagination
	Active       *bool
	UpdatedSince string
}

func (f SubscriptionFilter) Validate() error {
	if err := f.Pagination.Validate(); err != nil {
		return err
	}
	if f.UpdatedSince != "" {
		if _, err := time.Parse(time.RFC3339, f.UpdatedSince); err != nil {
			return errors.New("updated-since must be an RFC 3339 timestamp")
		}
	}
	return nil
}

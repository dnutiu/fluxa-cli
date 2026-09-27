package fluxa

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/dnutiu/fluxa-cli/internal/application"
	"github.com/dnutiu/fluxa-cli/internal/domain"
)

func (r *Repository) ListTransactions(ctx context.Context, entityID domain.ID, filter domain.TransactionFilter) (application.Collection[domain.Transaction], error) {
	query := url.Values{
		"page":     {strconv.Itoa(filter.Pagination.Page)},
		"per_page": {strconv.Itoa(filter.Pagination.PerPage)},
	}
	if filter.CategoryID.Valid() {
		query.Set("category_id", strconv.FormatInt(int64(filter.CategoryID), 10))
	}
	if filter.AccountID.Valid() {
		query.Set("account_id", strconv.FormatInt(int64(filter.AccountID), 10))
	}
	if filter.Kind != "" {
		query.Set("kind", filter.Kind)
	}
	if filter.DateFrom != "" {
		query.Set("date_from", filter.DateFrom)
	}
	if filter.DateTo != "" {
		query.Set("date_to", filter.DateTo)
	}
	if filter.UpdatedSince != "" {
		query.Set("updated_since", filter.UpdatedSince)
	}
	if filter.IncludeDeleted {
		query.Set("include_deleted", "true")
	}
	var result application.Collection[domain.Transaction]
	err := r.request(ctx, http.MethodGet, collectionPath(entityID, "transactions"), query, nil, "", &result)
	return result, err
}

func (r *Repository) GetTransaction(ctx context.Context, entityID, transactionID domain.ID) (application.Record[domain.Transaction], error) {
	var result application.Record[domain.Transaction]
	err := r.request(ctx, http.MethodGet, itemPath(entityID, "transactions", transactionID), nil, nil, "", &result)
	return result, err
}

func (r *Repository) CreateTransaction(ctx context.Context, entityID domain.ID, input domain.TransactionInput, key string) (application.Record[domain.Transaction], error) {
	var result application.Record[domain.Transaction]
	payload, err := json.Marshal(transactionPayload(input))
	if err != nil {
		return result, err
	}
	err = r.request(ctx, http.MethodPost, collectionPath(entityID, "transactions"), nil, payload, key, &result)
	return result, err
}

func (r *Repository) UpdateTransaction(ctx context.Context, entityID, transactionID domain.ID, input domain.TransactionInput) (application.Record[domain.Transaction], error) {
	var result application.Record[domain.Transaction]
	payload, err := json.Marshal(transactionPayload(input))
	if err != nil {
		return result, err
	}
	err = r.request(ctx, http.MethodPatch, itemPath(entityID, "transactions", transactionID), nil, payload, "", &result)
	return result, err
}

func transactionPayload(input domain.TransactionInput) map[string]any {
	body := make(map[string]any)
	putField(body, "category_id", input.CategoryID)
	putField(body, "account_id", input.AccountID)
	putField(body, "amount", input.Amount)
	putField(body, "currency", input.Currency)
	putField(body, "exchange_rate", input.ExchangeRate)
	putField(body, "date", input.Date)
	putField(body, "occurred_at", input.OccurredAt)
	putField(body, "description", input.Description)
	putField(body, "exclude_from_analytics", input.ExcludeFromAnalytics)
	return body
}

func (r *Repository) DeleteTransaction(ctx context.Context, entityID, transactionID domain.ID) error {
	return r.request(ctx, http.MethodDelete, itemPath(entityID, "transactions", transactionID), nil, nil, "", nil)
}

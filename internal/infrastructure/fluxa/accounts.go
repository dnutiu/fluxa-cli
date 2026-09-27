package fluxa

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/dnutiu/fluxa-cli/internal/application"
	"github.com/dnutiu/fluxa-cli/internal/domain"
)

func (r *Repository) ListAccounts(ctx context.Context, entityID domain.ID, filter domain.AccountFilter) (application.Collection[domain.Account], error) {
	query := url.Values{}
	if filter.IncludeArchived {
		query.Set("include_archived", "true")
	}
	var result application.Collection[domain.Account]
	err := r.request(ctx, http.MethodGet, collectionPath(entityID, "accounts"), query, nil, "", &result)
	return result, err
}

func (r *Repository) GetAccount(ctx context.Context, entityID, accountID domain.ID) (application.Record[domain.Account], error) {
	var result application.Record[domain.Account]
	err := r.request(ctx, http.MethodGet, itemPath(entityID, "accounts", accountID), nil, nil, "", &result)
	return result, err
}

func (r *Repository) CreateAccount(ctx context.Context, entityID domain.ID, input domain.AccountInput) (application.Record[domain.Account], error) {
	var result application.Record[domain.Account]
	payload, err := json.Marshal(accountPayload(input))
	if err != nil {
		return result, err
	}
	err = r.request(ctx, http.MethodPost, collectionPath(entityID, "accounts"), nil, payload, "", &result)
	return result, err
}

func (r *Repository) UpdateAccount(ctx context.Context, entityID, accountID domain.ID, input domain.AccountInput) (application.Record[domain.Account], error) {
	var result application.Record[domain.Account]
	payload, err := json.Marshal(accountPayload(input))
	if err != nil {
		return result, err
	}
	err = r.request(ctx, http.MethodPatch, itemPath(entityID, "accounts", accountID), nil, payload, "", &result)
	return result, err
}

func accountPayload(input domain.AccountInput) map[string]any {
	body := make(map[string]any)
	putField(body, "name", input.Name)
	putField(body, "kind", input.Kind)
	putField(body, "currency", input.Currency)
	putField(body, "opening_balance", input.OpeningBalance)
	putField(body, "annual_interest", input.AnnualInterest)
	putField(body, "interest_tax", input.InterestTax)
	putField(body, "account_group_id", input.AccountGroupID)
	return body
}

func (r *Repository) ArchiveAccount(ctx context.Context, entityID, accountID domain.ID) error {
	return r.request(ctx, http.MethodDelete, itemPath(entityID, "accounts", accountID), nil, nil, "", nil)
}

func (r *Repository) RestoreAccount(ctx context.Context, entityID, accountID domain.ID) (application.Record[domain.Account], error) {
	var result application.Record[domain.Account]
	err := r.request(ctx, http.MethodPost, itemPath(entityID, "accounts", accountID)+"/restore", nil, nil, "", &result)
	return result, err
}

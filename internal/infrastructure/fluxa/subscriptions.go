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

func (r *Repository) ListSubscriptions(ctx context.Context, entityID domain.ID, filter domain.SubscriptionFilter) (application.Collection[domain.Subscription], error) {
	query := url.Values{
		"page":     {strconv.Itoa(filter.Pagination.Page)},
		"per_page": {strconv.Itoa(filter.Pagination.PerPage)},
	}
	if filter.Active != nil {
		query.Set("active", strconv.FormatBool(*filter.Active))
	}
	if filter.UpdatedSince != "" {
		query.Set("updated_since", filter.UpdatedSince)
	}
	var result application.Collection[domain.Subscription]
	err := r.request(ctx, http.MethodGet, collectionPath(entityID, "subscriptions"), query, nil, "", &result)
	return result, err
}

func (r *Repository) GetSubscription(ctx context.Context, entityID, subscriptionID domain.ID) (application.Record[domain.Subscription], error) {
	var result application.Record[domain.Subscription]
	err := r.request(ctx, http.MethodGet, itemPath(entityID, "subscriptions", subscriptionID), nil, nil, "", &result)
	return result, err
}

func (r *Repository) CreateSubscription(ctx context.Context, entityID domain.ID, payload []byte) (application.Record[domain.Subscription], error) {
	var result application.Record[domain.Subscription]
	err := r.request(ctx, http.MethodPost, collectionPath(entityID, "subscriptions"), nil, payload, "", &result)
	return result, err
}

func (r *Repository) UpdateSubscription(ctx context.Context, entityID, subscriptionID domain.ID, payload []byte) (application.Record[domain.Subscription], error) {
	var result application.Record[domain.Subscription]
	err := r.request(ctx, http.MethodPatch, itemPath(entityID, "subscriptions", subscriptionID), nil, payload, "", &result)
	return result, err
}

func (r *Repository) DeleteSubscription(ctx context.Context, entityID, subscriptionID domain.ID) error {
	return r.request(ctx, http.MethodDelete, itemPath(entityID, "subscriptions", subscriptionID), nil, nil, "", nil)
}

func (r *Repository) SetSubscriptionActive(ctx context.Context, entityID, subscriptionID domain.ID, active bool) (application.Record[domain.Subscription], error) {
	body, _ := json.Marshal(struct {
		Active bool `json:"active"`
	}{Active: active})
	return r.UpdateSubscription(ctx, entityID, subscriptionID, body)
}

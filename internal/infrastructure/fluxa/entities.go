package fluxa

import (
	"context"
	"net/http"

	"fluxa-cli/internal/application"
	"fluxa-cli/internal/domain"
)

func (r *Repository) ListEntities(ctx context.Context) (application.Collection[domain.Entity], error) {
	var result application.Collection[domain.Entity]
	err := r.request(ctx, http.MethodGet, "/api/v1/entities", nil, nil, "", &result)
	return result, err
}

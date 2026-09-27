package fluxa

import (
	"context"
	"net/http"

	"fluxa-cli/internal/application"
	"fluxa-cli/internal/domain"
)

func (r *Repository) ListCategories(ctx context.Context, entityID domain.ID) (application.Collection[domain.Category], error) {
	var result application.Collection[domain.Category]
	err := r.request(ctx, http.MethodGet, collectionPath(entityID, "categories"), nil, nil, "", &result)
	return result, err
}

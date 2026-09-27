package fluxa

import (
	"context"
	"net/http"

	"github.com/dnutiu/fluxa-cli/internal/application"
	"github.com/dnutiu/fluxa-cli/internal/domain"
)

func (r *Repository) ListCategories(ctx context.Context, entityID domain.ID) (application.Collection[domain.Category], error) {
	var result application.Collection[domain.Category]
	err := r.request(ctx, http.MethodGet, collectionPath(entityID, "categories"), nil, nil, "", &result)
	return result, err
}

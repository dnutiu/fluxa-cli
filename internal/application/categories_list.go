package application

import (
	"context"
	"github.com/dnutiu/fluxa-cli/internal/domain"
)

type CategoryReader interface {
	ListCategories(context.Context, domain.ID) (Collection[domain.Category], error)
}

func ListCategories(ctx context.Context, reader CategoryReader, entityID domain.ID) (Collection[domain.Category], error) {
	if err := requireIDs(entityID); err != nil {
		return Collection[domain.Category]{}, err
	}
	return reader.ListCategories(ctx, entityID)
}

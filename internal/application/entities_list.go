package application

import (
	"context"
	"fluxa-cli/internal/domain"
)

type EntityReader interface {
	ListEntities(context.Context) (Collection[domain.Entity], error)
}

func ListEntities(ctx context.Context, reader EntityReader) (Collection[domain.Entity], error) {
	return reader.ListEntities(ctx)
}

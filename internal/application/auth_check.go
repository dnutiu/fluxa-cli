package application

import (
	"context"
	"fluxa-cli/internal/domain"
)

func CheckAuthentication(ctx context.Context, reader EntityReader) (Collection[domain.Entity], error) {
	return reader.ListEntities(ctx)
}

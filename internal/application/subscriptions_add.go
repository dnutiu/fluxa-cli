package application

import (
	"context"

	"github.com/dnutiu/fluxa-cli/internal/domain"
)

type SubscriptionCreator interface {
	CreateSubscription(context.Context, domain.ID, domain.SubscriptionInput) (Record[domain.Subscription], error)
}

func AddSubscription(ctx context.Context, repo SubscriptionCreator, entityID domain.ID, input domain.SubscriptionInput) (Record[domain.Subscription], error) {
	if err := requireIDs(entityID); err != nil {
		return Record[domain.Subscription]{}, err
	}
	if err := input.ValidateCreate(); err != nil {
		return Record[domain.Subscription]{}, err
	}
	return repo.CreateSubscription(ctx, entityID, input)
}

package application

import (
	"context"
	"github.com/dnutiu/fluxa-cli/internal/domain"
)

type SubscriptionCreator interface {
	CreateSubscription(context.Context, domain.ID, []byte) (Record[domain.Subscription], error)
}

func AddSubscription(ctx context.Context, repo SubscriptionCreator, entityID domain.ID, payload []byte) (Record[domain.Subscription], error) {
	if err := requireIDs(entityID); err != nil {
		return Record[domain.Subscription]{}, err
	}
	if err := requireObject(payload); err != nil {
		return Record[domain.Subscription]{}, err
	}
	return repo.CreateSubscription(ctx, entityID, payload)
}

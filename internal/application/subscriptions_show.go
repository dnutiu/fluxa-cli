package application

import (
	"context"
	"fluxa-cli/internal/domain"
)

type SubscriptionGetter interface {
	GetSubscription(context.Context, domain.ID, domain.ID) (Record[domain.Subscription], error)
}

func ShowSubscription(ctx context.Context, repo SubscriptionGetter, entityID, subscriptionID domain.ID) (Record[domain.Subscription], error) {
	if err := requireIDs(entityID, subscriptionID); err != nil {
		return Record[domain.Subscription]{}, err
	}
	return repo.GetSubscription(ctx, entityID, subscriptionID)
}

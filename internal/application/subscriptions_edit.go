package application

import (
	"context"

	"github.com/dnutiu/fluxa-cli/internal/domain"
)

type SubscriptionUpdater interface {
	UpdateSubscription(context.Context, domain.ID, domain.ID, domain.SubscriptionInput) (Record[domain.Subscription], error)
}

func EditSubscription(ctx context.Context, repo SubscriptionUpdater, entityID, subscriptionID domain.ID, input domain.SubscriptionInput) (Record[domain.Subscription], error) {
	if err := requireIDs(entityID, subscriptionID); err != nil {
		return Record[domain.Subscription]{}, err
	}
	if err := input.ValidateUpdate(); err != nil {
		return Record[domain.Subscription]{}, err
	}
	return repo.UpdateSubscription(ctx, entityID, subscriptionID, input)
}

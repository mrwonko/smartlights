package pubsubhelper

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

func GetOrCreateTopic(ctx context.Context, cl *pubsub.Client, name string) (*pubsub.Topic, error) {
	topic := cl.Topic(name)
	ok, err := topic.Exists(ctx)
	if err != nil {
		return nil, fmt.Errorf("querying existence of %q: %w", name, err)
	}
	if ok {
		return topic, nil
	}
	topic, err = cl.CreateTopic(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("creating topic %q: %w", name, err)
	}
	return topic, nil
}

func GetOrCreateSubscription(ctx context.Context, cl *pubsub.Client, name, topic string) (*pubsub.Subscription, error) {
	res := cl.Subscription(name)
	ok, err := res.Exists(ctx)
	if err != nil {
		return nil, fmt.Errorf("querying existence of subscription %q: %w", name, err)
	}
	if ok {
		return res, nil
	}
	t, err := GetOrCreateTopic(ctx, cl, topic)
	if err != nil {
		return nil, err
	}
	res, err = cl.CreateSubscription(ctx, name, pubsub.SubscriptionConfig{
		Topic: t,
	})
	if err != nil {
		return nil, fmt.Errorf("creating subscription %q: %w", name, err)
	}
	return res, nil
}

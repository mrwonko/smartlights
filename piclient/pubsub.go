package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/mrwonko/smartlights/config"

	"cloud.google.com/go/pubsub"
)

type pubsubClient struct {
	client              *pubsub.Client
	executeSubscription *pubsub.Subscription
}

func newPubsubClient(ctx context.Context, pi int) (_ *pubsubClient, finalErr error) {
	projectID := os.Getenv("PUBSUB_PROJECT_ID")
	if projectID == "" {
		return nil, errors.New("no PUBSUB_PROJECT_ID set")
	}
	cl, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("creating pubsub client: %s", err)
	}
	defer func() {
		if finalErr != nil {
			if err = cl.Close(); err != nil {
				log.Printf("failed to close new pubsub client: %s", err)
			}
		}
	}()
	res := pubsubClient{
		client: cl,
	}
	name := fmt.Sprintf("execute-%d", pi)
	et := cl.Topic(name)
	ok, err := et.Exists(ctx)
	if err != nil {
		return nil, fmt.Errorf("querying existence of topic %q: %s", name, err)
	}
	if !ok {
		et, err = cl.CreateTopic(ctx, name)
		if err != nil {
			return nil, fmt.Errorf("creating topic %q: %s", name, err)
		}
	}
	res.executeSubscription = cl.Subscription(name)
	ok, err = res.executeSubscription.Exists(ctx)
	if err != nil {
		return nil, fmt.Errorf("querying existence of subscription %q: %s", name, err)
	}
	if !ok {
		res.executeSubscription, err = cl.CreateSubscription(ctx, name, pubsub.SubscriptionConfig{
			Topic: et,
		})
		if err != nil {
			return nil, fmt.Errorf("creating topic %q: %s", name, err)
		}
	}
	return &res, nil
}

func (pc *pubsubClient) ReceiveExecute(ctx context.Context, f func(ctx context.Context, msg *config.ExecuteMessage)) error {
	return pc.executeSubscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		data := config.ExecuteMessage{}
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("unmarshaling %q message: %s", pc.executeSubscription, err)
			return
		}
		f(ctx, &data)
	})
}

func (pc *pubsubClient) Close() error {
	return pc.client.Close()
}

package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mrwonko/smartlights/config"
	"github.com/mrwonko/smartlights/internal/pubsubhelper"

	"cloud.google.com/go/pubsub"
)

const (
	discardMessagesAfter = time.Minute
)

type pubsubClient struct {
	client              *pubsub.Client
	executeSubscription *pubsub.Subscription
	stateTopic          *pubsub.Topic
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
	res.executeSubscription, err = pubsubhelper.GetOrCreateSubscription(ctx, cl, name, name)
	if err != nil {
		return nil, err
	}
	name = "state"
	res.stateTopic, err = pubsubhelper.GetOrCreateTopic(ctx, cl, name)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (pc *pubsubClient) ReceiveExecute(ctx context.Context, f func(ctx context.Context, msg *config.ExecuteMessage)) error {
	return pc.executeSubscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		defer msg.Ack()
		if msg.PublishTime.Add(discardMessagesAfter).Before(time.Now()) {
			log.Printf("skipping message due to age: %v", msg)
			return
		}
		data := config.ExecuteMessage{}
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("unmarshaling %q message: %s", pc.executeSubscription, err)
			return
		}
		f(ctx, &data)
	})
}

func (pc *pubsubClient) State(ctx context.Context) error {
	data, err := json.Marshal(config.StateMessage{
		// TODO
	})
	if err != nil {
		return fmt.Errorf("marshalling message: %s", err)
	}
	_, err = pc.stateTopic.Publish(ctx, &pubsub.Message{
		Data: data,
	}).Get(ctx)
	return err
}

func (pc *pubsubClient) Close() error {
	return pc.client.Close()
}

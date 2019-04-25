package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mrwonko/smartlights/internal/common"
	"github.com/mrwonko/smartlights/internal/config"

	"cloud.google.com/go/pubsub"
)

const (
	discardMessagesAfter = time.Minute
)

type pubsubClient struct {
	client                   *pubsub.Client
	executeSubscription      *pubsub.Subscription
	queryRequestSubscription *pubsub.Subscription
	queryResponseTopic       *pubsub.Topic
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

	execute := fmt.Sprintf("execute-%d", pi)
	res.executeSubscription, err = common.GetOrCreateSubscription(ctx, cl, execute, execute)
	if err != nil {
		return nil, err
	}
	queryRequest := fmt.Sprintf("query-request-%d", pi)
	res.queryRequestSubscription, err = common.GetOrCreateSubscription(ctx, cl, queryRequest, queryRequest)
	if err != nil {
		return nil, err
	}
	res.queryResponseTopic, err = common.GetOrCreateTopic(ctx, cl, "query-response")
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (pc *pubsubClient) ReceiveExecute(ctx context.Context, f func(ctx context.Context, msg *config.ExecuteMessage)) error {
	return pc.executeSubscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		defer msg.Ack()
		if msg.PublishTime.Add(discardMessagesAfter).Before(time.Now()) {
			log.Printf("skipping execute message due to age: %v", msg)
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

func (pc *pubsubClient) ReceiveQueryRequest(ctx context.Context, f func(ctx context.Context, msg *config.QueryRequestMessage)) error {
	return pc.queryRequestSubscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		defer msg.Ack()
		if msg.PublishTime.Add(discardMessagesAfter).Before(time.Now()) {
			log.Printf("skipping queryRequest message due to age: %v", msg)
			return
		}
		data := config.QueryRequestMessage{}
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("unmarshaling %q message: %s", pc.queryRequestSubscription, err)
			return
		}
		f(ctx, &data)
	})
}

func (pc *pubsubClient) Close() error {
	pc.queryResponseTopic.Stop()
	return pc.client.Close()
}

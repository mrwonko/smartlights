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
	"github.com/mrwonko/smartlights/internal/protocol"
	"github.com/mrwonko/smartlights/internal/pubsubhelper"

	"cloud.google.com/go/pubsub"
)

type pubsubClient struct {
	client            *pubsub.Client
	executeTopics     map[int]*pubsub.Topic
	stateSubscription *pubsub.Subscription
}

const (
	discardMessagesAfter = time.Minute
)

func newPubsubClient(ctx context.Context) (_ *pubsubClient, finalErr error) {
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
		client:        cl,
		executeTopics: make(map[int]*pubsub.Topic, len(config.Pis)),
	}
	for pi := range config.Pis {
		name := fmt.Sprintf("execute-%d", pi)
		res.executeTopics[pi], err = pubsubhelper.GetOrCreateTopic(ctx, cl, name)
		if err != nil {
			return nil, err
		}
	}
	name := "state"
	res.stateSubscription, err = pubsubhelper.GetOrCreateSubscription(ctx, cl, name, name)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (pc *pubsubClient) Execute(ctx context.Context, pi int, msg protocol.ExecuteMessage) error {
	topic := pc.executeTopics[pi]
	if topic == nil {
		return fmt.Errorf("invalid pi %d", pi)
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshalling message: %s", err)
	}
	_, err = topic.Publish(ctx, &pubsub.Message{
		Data: data,
	}).Get(ctx)
	return err
}

func (pc *pubsubClient) ReceiveState(ctx context.Context, f func(ctx context.Context, msg *protocol.StateMessage)) error {
	return pc.stateSubscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		defer msg.Ack()
		if msg.PublishTime.Add(discardMessagesAfter).Before(time.Now()) {
			log.Printf("skipping State message due to age: %v", msg)
			return
		}
		data := protocol.StateMessage{}
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("unmarshaling %q message: %s", pc.stateSubscription, err)
			return
		}
		f(ctx, &data)
	})
}

func (pc *pubsubClient) Close() error {
	for _, t := range pc.executeTopics {
		t.Stop()
	}
	return pc.client.Close()
}

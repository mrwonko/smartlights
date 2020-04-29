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
	"github.com/mrwonko/smartlights/internal/common"

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
		res.executeTopics[pi], err = common.GetOrCreateTopic(ctx, cl, name)
		if err != nil {
			return nil, err
		}
	}
	name := "state"
	res.stateSubscription, err = common.GetOrCreateSubscription(ctx, cl, name, name)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// TODO: instead of sending individual commands, send a list of them, so we can do one message per device in an execute intent.
func (pc *pubsubClient) OnOff(ctx context.Context, id config.ID, on bool) error {
	light := config.Lights[id]
	if light == nil {
		return fmt.Errorf("invalid light ID %d", id)
	}
	topic := pc.executeTopics[light.Pi]
	if topic == nil {
		return fmt.Errorf("invalid pi %d", light.Pi)
	}
	data, err := json.Marshal(config.ExecuteMessage{
		GPIO: light.GPIO,
		On:   on,
	})
	if err != nil {
		return fmt.Errorf("marshalling message: %s", err)
	}
	_, err = topic.Publish(ctx, &pubsub.Message{
		Data: data,
	}).Get(ctx)
	return err
}

func (pc *pubsubClient) ReceiveState(ctx context.Context, f func(ctx context.Context, msg *config.StateMessage)) error {
	return pc.stateSubscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		defer msg.Ack()
		if msg.PublishTime.Add(discardMessagesAfter).Before(time.Now()) {
			log.Printf("skipping State message due to age: %v", msg)
			return
		}
		data := config.StateMessage{}
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

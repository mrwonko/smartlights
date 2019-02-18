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
	client        *pubsub.Client
	executeTopics map[int]*pubsub.Topic
}

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
		et := cl.Topic(name)
		ok, err := et.Exists(ctx)
		if err != nil {
			return nil, fmt.Errorf("querying existence of %s: %s", name, err)
		}
		if !ok {
			et, err = cl.CreateTopic(ctx, name)
			if err != nil {
				return nil, fmt.Errorf("creating topic %q: %s", name, err)
			}
		}
		res.executeTopics[pi] = et
	}
	return &res, nil
}

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

func (pc *pubsubClient) Close() error {
	for _, t := range pc.executeTopics {
		t.Stop()
	}
	return pc.client.Close()
}

func execute(ctx context.Context) {

}

package clients

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

type GCPClient struct {
	Subscription *pubsub.Subscription
}

var GCPProjectID = ""
var GCPSubscriptionID = ""

func NewGCPClient(ctx context.Context) (*GCPClient, error) {
	client, err := pubsub.NewClient(ctx, GCPProjectID)
	if err != nil {
		return nil, fmt.Errorf("pubsub.NewClient: %v", err)
	}

	sub := client.Subscription(GCPSubscriptionID)
	exist, err := sub.Exists(ctx)
	if err != nil {
		return nil, fmt.Errorf("sub.Exists: %v", err)
	}
	if !exist {
		return nil, fmt.Errorf("sub not exist: %v", err)
	}

	return &GCPClient{
		Subscription: sub,
	}, nil

}

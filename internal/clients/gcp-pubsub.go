package clients

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/wrandowR/gcp-pubsub-with-redis/internal/entity"
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

	return &GCPClient{
		Subscription: sub,
	}, nil

}

func (g *GCPClient) PullMsgs(ctx context.Context, ch chan entity.Message) error {

	var received int32
	err := g.Subscription.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {

		defer msg.Ack()

		fmt.Println("Got message:", string(msg.Data))
		atomic.AddInt32(&received, 1)

		ch <- entity.Message{
			ID:      fmt.Sprint(received),
			Date:    time.Now().String(),
			Message: string(msg.Data),
		}
	})
	if err != nil {
		return fmt.Errorf("sub.Receive: %v", err)
	}

	fmt.Println("Received messages", received)

	return nil
}

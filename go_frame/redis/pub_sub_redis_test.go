package redis

import (
	"context"
	"testing"
)

func TestPubSub(t *testing.T) {

	ctx := context.Background()
	PubSub(ctx, client)
}

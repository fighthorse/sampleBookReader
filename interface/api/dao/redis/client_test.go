package redis

import (
	"context"
	"testing"
)

func TestGet(t *testing.T) {
	_, _ = LoadOthersDB("base", 0).Client.Get(context.Background(), "test")
}

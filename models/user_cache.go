package models

import (
	"context"
	"mychat/clients"
	"time"
)

// SetUserOnlineInfo /**
func SetUserOnlineInfo(key string, val []byte, timeTTL time.Duration) {
	ctx := context.Background()
	clients.Redis.Set(ctx, key, val, timeTTL)
}

package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// TryLock 尝试获取锁
func TryLock(ctx context.Context, client *redis.Client, key string, expire time.Duration) bool {
	cmd := client.SetNX(ctx, key, "1", expire)
	if err := cmd.Err(); err != nil {
		return false
	} else {
		return cmd.Val()
	}
}

// Unlock 释放锁
func Unlock(ctx context.Context, client *redis.Client, key string) {
	client.Del(ctx, key)
}

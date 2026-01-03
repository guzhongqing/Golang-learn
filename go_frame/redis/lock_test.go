package redis

import (
	"context"
	"testing"
	"time"
)

func TestTryLock(t *testing.T) {

	ctx := context.Background()

	key := "test_lock"
	expire := 10 * time.Second

	// 尝试获取锁
	locked := TryLock(ctx, client, key, expire)
	if locked {
		t.Logf("Acquired lock")
	} else {
		t.Errorf("Failed to acquire lock")
	}
	// 再次尝试获取锁，应该失败
	locked = TryLock(ctx, client, key, expire)
	if locked {
		t.Logf("Acquired lock")
	} else {
		t.Errorf("Failed to acquire lock")
	}

	time.Sleep(expire)

	// 释放锁
	Unlock(ctx, client, key)
}

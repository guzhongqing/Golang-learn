package redis

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

// Publish 发布消息到指定频道
func publish(ctx context.Context, client *redis.Client, channel string, message any) {
	cmd := client.Publish(ctx, channel, message)
	if err := cmd.Err(); err != nil {
		panic(err)
	} else {
		n := cmd.Val() // 发布成功后返回的消息数量
		fmt.Printf("%s发布消息到频道 %s，此时该频道有 %d 个订阅者\n", ctx.Value("publisher_name"), channel, n)
	}
}

// Subscribe 订阅指定频道的消息
func subscribe(ctx context.Context, client *redis.Client, channels []string) {
	ps := client.Subscribe(ctx, channels...)
	defer ps.Close()

	for {

		if msg, err := ps.ReceiveMessage(ctx); err != nil {
			panic(err)
		} else {
			fmt.Printf("%s从频道 %s 收到消息: %s\n", ctx.Value("subscriber_name"), msg.Channel, msg.Payload)
		}
	}
}

// 发布-订阅模式
func PubSub(ctx context.Context, client *redis.Client) {

	// 发布者
	publisher1 := context.WithValue(ctx, "publisher_name", "publisher1")
	publisher2 := context.WithValue(ctx, "publisher_name", "publisher2")

	// 订阅者
	subscriber3 := context.WithValue(ctx, "subscriber_name", "subscriber3")
	subscriber4 := context.WithValue(ctx, "subscriber_name", "subscriber4")

	// 频道
	channel1 := "channel1"
	channel2 := "channel2"

	// 先订阅
	go subscribe(subscriber3, client, []string{channel1})
	go subscribe(subscriber4, client, []string{channel2})

	time.Sleep(1 * time.Second)
	// 发布
	publish(publisher1, client, channel1, "publisher1 向 channel1  发布的消息")
	publish(publisher2, client, channel1, "publisher2 向 channel1  发布的消息")

	time.Sleep(1 * time.Second)
	// 分割线
	fmt.Println(strings.Repeat("*", 50))

	// 第二批订阅者
	go subscribe(subscriber3, client, []string{channel2})

	time.Sleep(1 * time.Second)
	// 发布
	publish(publisher1, client, channel2, "publisher1 向 channel2  发布的消息")

	// 等待20秒，确保所有消息都被处理
	time.Sleep(20 * time.Second)

}

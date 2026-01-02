package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"path"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func InitRedisViper(dir, file, FileType string) *viper.Viper {
	config := viper.New()
	config.AddConfigPath(dir)
	config.SetConfigName(file)
	config.SetConfigType(FileType)

	if err := config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("配置文件 %s 解析失败: %s \n", path.Join(dir, file+"."+FileType), err))
	}

	return config
}

func InitRedis(config *viper.Viper) *redis.Client {
	addr := config.GetString("redis.address")
	username := config.GetString("redis.username")
	password := config.GetString("redis.password")
	db := config.GetInt("redis.db")

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: username,
		Password: password,
		DB:       db,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		slog.Error("redis 连接失败", "err", err)
		return nil
	} else {
		slog.Info("redis 连接成功", "addr", addr)
		return client
	}
}

// 检查错误
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// stirng类型的value
func stringValue(ctx context.Context, client *redis.Client) {
	key := "name"
	value := "张三"

	// 测试之后自动删除
	defer client.Del(ctx, key)

	// 最后一个参数为过期时间，0 表示永不过期
	err := client.Set(ctx, key, value, 0).Err()
	checkError(err)

	// 设置过期时间为 10 秒
	err = client.Expire(ctx, key, time.Second*10).Err()
	checkError(err)

	// 测试获取值，如果key不存在，err返回redis.Nil
	val, err := client.Get(ctx, key).Result()
	checkError(err)
	fmt.Println(val)

	client.Set(ctx, "age", 35, 10*time.Second) // 设置时18为int写入redis时会自动转换为string，过期时间为10秒

	// 测试获取age值,Result返回是string
	age, err := client.Get(ctx, "age").Result()
	checkError(err)
	fmt.Println(age)

	// Int()返回是int
	ageInt, err := client.Get(ctx, "age").Int()
	checkError(err)
	fmt.Println(ageInt)

	// 等待 30 秒，确保过期
	time.Sleep(time.Second * 20)

}

func DeleteKey(ctx context.Context, client *redis.Client) {
	n, err := client.Del(ctx, "not_exist_key").Result()
	checkError(err)
	fmt.Println(n)
}

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func WriteStruct2Redis(ctx context.Context, client *redis.Client) {
	student := Student{
		ID:   1001,
		Name: "张三",
		Age:  18,
	}
	key := fmt.Sprintf("user:info:%d", student.ID)
	value, err := json.Marshal(student)
	checkError(err)
	// []byte go-redis会自动转换为 string，然后存储到 redis 中
	client.Set(ctx, key, value, 0)
	// 测试获取值，返回的是 string 类型
	val, err := client.Get(ctx, key).Result()
	checkError(err)
	fmt.Println(val)
	// 测试获取值，返回的是 []byte  类型
	valByte, err := client.Get(ctx, key).Bytes()
	checkError(err)
	fmt.Println(valByte)
}

// list类型的value
func listValue(ctx context.Context, client *redis.Client) {
	key := "list_key"
	// 测试之后自动删除
	// defer client.Del(ctx, key)
	values := []any{"张三", "李四", "王五", 6}

	// LPush 从列表的头部插入多个值,RPush,存储的时候框架会自动转换为string
	// redis的list只会存储string类型的值
	err := client.RPush(ctx, key, values...).Err()
	checkError(err)

	// 测试获取值,0表示第一个元素，-1表示最后一个元素，[0,-1]闭区间，返回所有元素
	val, err := client.LRange(ctx, key, 0, -1).Result()
	checkError(err)
	fmt.Println(val)

}

// set类型的value
func setValue(ctx context.Context, client *redis.Client) {
	key := "set_key"
	// 测试之后自动删除
	// defer client.Del(ctx, key)
	values := []any{"张三", 6, "李四", "王五", 6}

	// redis的set只会存储string类型的值
	err := client.SAdd(ctx, key, values...).Err()
	checkError(err)

	// 测试获取set中的所有元素
	vals, err := client.SMembers(ctx, key).Result()
	checkError(err)
	fmt.Println(vals)

	// 判断set中是否存在某个元素
	exists, err := client.SIsMember(ctx, key, "张三").Result()
	checkError(err)
	fmt.Println(exists)
}

// Zset类型的value
func zsetValue(ctx context.Context, client *redis.Client) {
	key := "zset_key"
	// 测试之后自动删除
	// defer client.Del(ctx, key)
	values := []redis.Z{
		{Score: 100, Member: "张三"},
		{Score: 90, Member: "李四"},
		{Score: 80, Member: "王五"},
	}

	// redis的zset只会存储string类型的值
	err := client.ZAdd(ctx, key, values...).Err()
	checkError(err)

	// 测试获取zset中的所有元素
	vals, err := client.ZRangeWithScores(ctx, key, 0, -1).Result()
	checkError(err)
	fmt.Println(vals)
}

// hash类型的value
func hashValue(ctx context.Context, client *redis.Client) {
	key := "hash_key"
	// 测试之后自动删除
	// defer client.Del(ctx, key)
	values := map[string]any{
		"name": "张三",
		"age":  18,
	}

	// redis的hash只会存储string类型的值
	err := client.HSet(ctx, key, values).Err()
	checkError(err)

	// 测试获取hash中的所有元素
	vals, err := client.HGetAll(ctx, key).Result()
	checkError(err)
	fmt.Println(vals)
}

// 遍历redis中的所有key
func scanKeys(ctx context.Context, client *redis.Client) {
	// 批量写入redis
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key:%d", i)
		client.Set(ctx, key, "1", 0)
	}

	const COUNT = 10000
	var cursor uint64
	// 正则匹配key以"key:"开头的所有key
	match := "key:*"

	iter := client.Scan(ctx, cursor, match, COUNT).Iterator()
	for iter.Next(ctx) {
		fmt.Println(iter.Val())
	}
}

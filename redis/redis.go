package redis

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

func setAndGet(ctx context.Context, client *redis.Client) {
	key := "name"
	value := "TechAoba"
	err := client.Set(ctx, key, value, 1*time.Minute).Err() // 如expire时间设为0，表示永久生效
	checkError(err)

	v2, err := client.Get(ctx, key).Result()
	checkError(err)
	fmt.Println(v2)

	// 删除插入的key
	// client.Del(ctx, key)
}

func list(ctx context.Context, client *redis.Client) {
	key := "lis"
	values := []interface{}{1, 2, "原神", 66}
	err := client.RPush(ctx, key, values...).Err()
	checkError(err)

	v2, err := client.LRange(ctx, key, 0, -1).Result() // L表示List的意思，后两个参数表示index区间
	checkError(err)
	fmt.Println(v2)

	client.Del(ctx, key)
}

func hashTable(ctx context.Context, client *redis.Client) {
	err := client.HSet(ctx, "用户1", "Name", "张三", "Age", 18, "Height", 173.5).Err()
	checkError(err)

	err = client.HSet(ctx, "用户2", "Name", "李四", "Age", 21, "Height", 178.3).Err()
	checkError(err)

	age, err := client.HGet(ctx, "用户2", "Age").Result() // 指定Field读取，每次只能读取一个值
	checkError(err)
	fmt.Println(age)

	for field, value := range client.HGetAll(ctx, "用户1").Val() {
		fmt.Println(field, value)
	}
	client.Del(ctx, "用户1")
	client.Del(ctx, "用户2")
}

func Rds() {
	client := redis.NewClient(&redis.Options{
		Addr:     "182.92.123.43:39006",
		Password: "Dianzijiaocaiecnu2020.",
		DB:       0,
	})
	ctx := context.TODO()
	// setAndGet(ctx, client)
	// list(ctx, client)
	hashTable(ctx, client)
}

func checkError(err error) {
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key不存在")
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

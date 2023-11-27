package redis

import (
	"PlanetMsg/pkg/signalInfo"
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var (
	client *redis.Client
	ctx    context.Context
)

func Set(key, value string, expire int) (string, error) {
	err := client.Set(ctx, key, value, time.Duration(expire)*time.Minute).Err() // 如expire时间设为0，表示永久生效
	if err != nil {
		return signalInfo.GetMsg(signalInfo.REDIS_FAIL), errors.New("Redis插入失败") // Redis插入失败
	}
	return signalInfo.GetMsg(signalInfo.REDIS_SUCCESS), nil // Redis插入成功
	// 删除插入的key
	// client.Del(ctx, key)
}

// 根据Key值获取Value
func Get(key string) (value, signal string) {
	value, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", signalInfo.GetMsg(signalInfo.REDIS_KEY_NOT_EXIST) // Redis不存在key
	}
	if err != nil {
		return "", signalInfo.GetMsg(signalInfo.REDIS_FAIL) // Redis查询失败
	}
	signal = signalInfo.GetMsg(signalInfo.REDIS_SUCCESS)
	return
}

func List(ctx context.Context, client *redis.Client) {
	key := "lis"
	values := []interface{}{1, 2, "原神", 66}
	err := client.RPush(ctx, key, values...).Err()
	checkError(err)

	v2, err := client.LRange(ctx, key, 0, -1).Result() // L表示List的意思，后两个参数表示index区间
	checkError(err)
	fmt.Println(v2)

	client.Del(ctx, key)
}

func HashTable(ctx context.Context, client *redis.Client) {
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

func init() {
	var HOST, PWD, PORT string
	var DBID int
	viper.SetConfigName("mysql")
	viper.AddConfigPath("./conf") // ./conf or ../../conf
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		panic(fmt.Errorf("读取配置文件失败: %s", err))
	}
	// 读取redis配置文件
	HOST = viper.GetString("redis.HOST")
	PORT = viper.GetString("redis.PORT")
	PWD = viper.GetString("redis.PWD")
	DBID = viper.GetInt("redis.DBID")

	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", HOST, PORT),
		Password: PWD,
		DB:       DBID,
	})
	ctx = context.TODO()
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

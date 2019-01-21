package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Could not connect to redis %v", err)
	}

	return client
}

func ExampleClient(redisClient *redis.Client) {

	err := redisClient.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := redisClient.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := redisClient.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist

}

func main() {

	redisClient := NewRedisClient()
	fmt.Println("client returned", redisClient)
	ExampleClient(redisClient)

	return
}

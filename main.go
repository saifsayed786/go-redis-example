package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("Go Redis Example")
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	err := client.Set("name", "saif", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := client.Get("name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}

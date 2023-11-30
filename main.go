package main

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis"
)

var ErrNotFound = errors.New("not found")

func main() {
	fmt.Println("hello")

	fmt.Println("world")
  
        fmt.Println("by Jacky.")

	helloRedis()
}

func helloRedis() {
	clinet := redis.NewClient(&redis.Options{})
	ret, err := clinet.Set("hello", "hello", 0).Result()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ErrNotFound)

	fmt.Println(ret)
}

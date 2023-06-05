package main

import (
	"fmt"

	"encoding/json"

	"github.com/go-redis/redis"
)

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Creating a Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address and port
		Password: "",               // Redis password (if used)
		DB:       0,                // Redis database index
	})

	// Check to connect to Redis server
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Could not connect to redis server:", err)
		return
	}
	fmt.Println("Connected to Redis server:", pong)

	json, err := json.Marshal(Author{Name: "Ahmet", Age: 25})
	if err != nil {
		fmt.Println(err)
	}
	// Write Data
	err = client.Set("id1234", json, 0).Err()
	if err != nil {
		fmt.Println("Data write error:", err)
	}
	// Read Data
	val, err := client.Get("id1234").Result()
	if err != nil {
		fmt.Println("Data read error:", err)
	}
	fmt.Println(val)

	// Closing the Redis client
	err = client.Close()
	if err != nil {
		fmt.Println("Redis client shutdown error:", err)
		return
	}
}

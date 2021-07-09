package main

import "fmt"
import "github.com/go-redis/redis"
import "time"

func Message() {
	redisdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pubsub := redisdb.Subscribe("mychannel1")

	// Wait for confirmation that subscription is created before publishing anything.
	if msg, err := pubsub.Receive(); err != nil {
		panic(err)
	} else {
		fmt.Printf("subscribe message: %v\n", msg)
	}

	// Go channel which receives messages.
	ch := pubsub.Channel()

	// Publish a message.
	if err := redisdb.Publish("mychannel1", "hello").Err(); err != nil {
		panic(err)
	}
	if err := redisdb.Publish("mychannel1", "world").Err(); err != nil {
		panic(err)
	}

	time.AfterFunc(time.Second, func() {
		// When pubsub is closed channel is closed too.
		_ = pubsub.Close()
	})

	// Consume messages.
	for msg := range ch {
		fmt.Printf("channel: %v, payload: %v\n", msg.Channel, msg.Payload)
	}
}

func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func main() {
	// ExampleNewClient()
	Message()
}

package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/streadway/amqp"
)

var (
	sig chan os.Signal = make(chan os.Signal)
	wg  sync.WaitGroup
)

func main() {

	conn, err := amqp.Dial("amqp://dev:dev@rabbitmq:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	// Read many posts
	err = ch.Publish("rapid", "RequestManyPosts", false, false, amqp.Publishing{
		//Headers:       amqp.Table{"JWT": goodAuth.JwtToken},
		ContentType:   "application/json",
		CorrelationId: "give me all those posts",
		Body:          []byte(string(`{"paging": {"start":"2020-12-10 19:43:19.484 +0000 UTC", "end":50}}`)),
	})
	if err != nil {
		fmt.Println(err)
	}

}

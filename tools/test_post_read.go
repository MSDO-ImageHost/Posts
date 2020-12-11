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

	for i := 0; i < 1; i++ {

		// Read many posts
		err = ch.Publish("rapid", "RequestManyPosts", false, false, amqp.Publishing{
			//Headers:       amqp.Table{"JWT": goodAuth.JwtToken},
			ContentType:   "application/json",
			CorrelationId: "give me all those posts",
			Body:          []byte(string(`{"post_ids": ["5fd25e97ed2bcc5bbb6172b7"], "paging": {"start":45, "end":50}}`)),
		})
		if err != nil {
			fmt.Println(err)
		}

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

	//consumer, err := ch.Consume("Posts", "", true, false, false, false, nil)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//go func() {
	//	for msg := range consumer {
	//
	//		p := string(msg.Body)
	//		msg.Body = []byte("")
	//
	//		log.Println("\n\n")
	//
	//		log.Print(utils.PrettyFormatMap(msg))
	//		fmt.Println(p)
	//		log.Println("\n\n")
	//		msg.Ack(true)
	//	}
	//}()
	//fmt.Println("Now consuming")
	//signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	//<-sig

}

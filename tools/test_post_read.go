package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

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
			//Body:          []byte(string(`{"header":"header!!", "body":"body!!", "image_data": ""}`)),
		})
		if err != nil {
			fmt.Println(err)
		}

	}

	fmt.Println("Now consuming")

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

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

}

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/MSDO-ImageHost/Posts/internal/utils"
	"github.com/streadway/amqp"
)

var sig chan os.Signal = make(chan os.Signal)

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

	// Declare queue
	queue, err := ch.QueueDeclare(
		"test-queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	/** Listen for responses **/
	consume, err := ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for msg := range consume {
			fmt.Println(utils.PrettyFormatMap(msg))
			fmt.Println(string(msg.Body))
			msg.Ack(true)
		}
	}()

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

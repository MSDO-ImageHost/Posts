package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/MSDO-ImageHost/Posts/internal/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/streadway/amqp"
)

var (
	correctSecret   []byte         = []byte("secret")
	incorrectSecret []byte         = []byte("not-this-secret")
	sig             chan os.Signal = make(chan os.Signal)
	wg              sync.WaitGroup
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

	// Declare queue
	pubQueue, err := ch.QueueDeclare("CreateOnePost", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 1000000; i++ {

		goodTokenString, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": strconv.Itoa(i), "role": 0}).SignedString(correctSecret)
		goodAuth, _ := auth.AuthJWT(goodTokenString)

		badTokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": strconv.Itoa(i), "role": 0}).SignedString(incorrectSecret)
		if err != nil {
			fmt.Println(err)
		}
		badAuth, err := auth.AuthJWT(badTokenString)

		// Correct message
		err = ch.Publish("rapid", pubQueue.Name, false, false, amqp.Publishing{
			Headers:       amqp.Table{"JWT": goodAuth.JwtToken},
			ContentType:   "application/json",
			CorrelationId: "this one is good",
			Body:          []byte(string(`{"header":"header!!", "body":"body!!", "image_data": ""}`)),
		})

		// Missing jwt
		err = ch.Publish("rapid", pubQueue.Name, false, false, amqp.Publishing{
			Headers:       amqp.Table{},
			ContentType:   "application/json",
			CorrelationId: "missing jwt header",
			Body:          []byte(string(`{"header":"header!!", "body":"body!!", "image_data": ""}`)),
		})

		// Header is nil
		err = ch.Publish("rapid", pubQueue.Name, false, false, amqp.Publishing{
			Headers:       nil,
			ContentType:   "application/json",
			CorrelationId: "delivery header is nil",
			Body:          []byte(string(`{"header":"header!!", "body":"body!!", "image_data": ""}`)),
		})

		// Invalid jwt
		err = ch.Publish("rapid", pubQueue.Name, false, false, amqp.Publishing{
			Headers:       amqp.Table{"JWT": badAuth.JwtToken},
			ContentType:   "application/json",
			CorrelationId: "invalid jwt",
			Body:          []byte(string(`{"header":"header!!", "body":"body!!", "image_data": ""}`)),
		})

		// Incorrect content type
		err = ch.Publish("rapid", pubQueue.Name, false, false, amqp.Publishing{
			Headers: amqp.Table{"JWT": goodAuth.JwtToken},
			//ContentType:   "application/json",
			CorrelationId: "missing content type",
			Body:          []byte(string(`{"header":"header!!", "body":"body!!", "image_data": ""}`)),
		})

		// Incorrect json structure
		err = ch.Publish("rapid", pubQueue.Name, false, false, amqp.Publishing{
			Headers:       amqp.Table{"JWT": goodAuth.JwtToken},
			ContentType:   "application/json",
			CorrelationId: "invalid json",
			Body:          []byte(string(`{"headerheader!!", "body"`)),
		})

		// Missing header field
		err = ch.Publish("rapid", pubQueue.Name, false, false, amqp.Publishing{
			Headers:       amqp.Table{"JWT": goodAuth.JwtToken},
			ContentType:   "application/json",
			CorrelationId: "missing header field",
			Body:          []byte(string(`{"body":"body!!", "image_data": ""}`)),
		})

		// Missing body field
		err = ch.Publish("rapid", pubQueue.Name, false, false, amqp.Publishing{
			Headers:       amqp.Table{"JWT": goodAuth.JwtToken},
			ContentType:   "application/json",
			CorrelationId: "missing body field",
			Body:          []byte(string(`{"header":"header!!", "image_data": ""}`)),
		})

		// Missing image data field
		err = ch.Publish("rapid", pubQueue.Name, false, false, amqp.Publishing{
			Headers:       amqp.Table{"JWT": goodAuth.JwtToken},
			ContentType:   "application/json",
			CorrelationId: "missing image data field",
			Body:          []byte(string(`{"header":"header!!", "body":"body!!"}`)),
		})

		// Incorrect field types
		err = ch.Publish("rapid", pubQueue.Name, false, false, amqp.Publishing{
			Headers:       amqp.Table{"JWT": goodAuth.JwtToken},
			ContentType:   "application/json",
			CorrelationId: "header is not a string",
			Body:          []byte(string(`{"header":12, "body":"body!!", "image_data": ""}`)),
		})

		// Incorrect field types
		err = ch.Publish("rapid", pubQueue.Name, false, false, amqp.Publishing{
			Headers:       amqp.Table{"JWT": goodAuth.JwtToken},
			ContentType:   "application/json",
			CorrelationId: "body is not a string",
			Body:          []byte(string(`{"header":"header!!", "body":12, "image_data": ""}`)),
		})

		// payload is nil
		err = ch.Publish("rapid", pubQueue.Name, false, false, amqp.Publishing{
			Headers:       amqp.Table{"JWT": goodAuth.JwtToken},
			ContentType:   "application/json",
			CorrelationId: "payload is nil",
			Body:          nil,
		})

		// ReplyTo
		//err = ch.Publish("rapid", pubQueue.Name, false, false, amqp.Publishing{
		//	Headers:       amqp.Table{"JWT": goodAuth.JwtToken},
		//	ContentType:   "application/json",
		//	CorrelationId: "this one is looped",
		//	Body:          []byte(string(`{"header":"header!!", "body":"body!!", "image_data": ""}`)),
		//	ReplyTo:       pubQueue.Name,
		//})
	}

	//subQueue, err := ch.QueueDeclare("ConfirmOnePostCreation", true, false, false, false, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//consumer, err := rabbit.ConsumeChannel.Consume(subQueuef.Name, "", true, false, c.NoLocal, c.NoWait, c.Args)
	//if err != nil {
	//	return err
	//}
	//
	//go func() {
	//	for msg := range consumer {
	//		log.Println(msg)
	//	}
	//}()

	//signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	//<-sig

}

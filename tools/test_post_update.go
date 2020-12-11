package main

import (
	"log"

	"github.com/MSDO-ImageHost/Posts/internal/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/streadway/amqp"
)

var (
	correctSecret []byte = []byte("secret")
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

	goodTokenString, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": "0", "role": 0}).SignedString(correctSecret)
	goodAuth, _ := auth.AuthJWT(goodTokenString)

	// Correct message
	if err := ch.Publish("rapid", "UpdateOnePost", false, false, amqp.Publishing{
		Headers:       amqp.Table{"JWT": goodAuth.JwtToken},
		ContentType:   "application/json",
		CorrelationId: "plz update post 4 me",
		Body:          []byte(string(`{"post_id":"5fd2c61c1e73919ba021d225",  "body":"!!body!!"}`)),
	}); err != nil {
		log.Fatal(err)
	}
}

package broker

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

func NewReqRes(handleConf HandleConfig, handler func(req HandleRequestPayload) (res HandleResponsePayload, err error)) error {

	//Declare subscription queue
	if err := QueueDeclare(handleConf.SubQueueConf); err != nil {
		return err
	}

	// Consumer for subscription messages
	consumer, err := ConsumerDeclare(handleConf)
	if err != nil {
		return err
	}

	// Listen for new messages
	go func() {
		for msg := range consumer {

			log.Println(LOG_TAG, "got msg", msg)

			// Verify integrety of message
			if msg.ContentType != "application/json" {
				msg.Ack(true)
			}

			// Verify user JWT. Get user id and role
			userID := "132-user" //msg.Headers["jwt"].(string)
			userRole := "user"

			// Run business logic handler
			res, err := handler(HandleRequestPayload{
				UserID:  userID,
				Role:    userRole,
				Payload: msg.Body,
			})

			// Don't acknowledge the message if an error happend
			if err != nil {
				msg.Ack(false)
				log.Fatal(err)
			}

			// Publish response
			rabbit.Channel.Publish("", msg.ReplyTo, false, false, amqp.Publishing{
				Headers:       nil,
				ContentType:   "application/json",
				CorrelationId: msg.CorrelationId,
				Timestamp:     time.Now().UTC(),
				Body:          res.Payload,
			})
			msg.Ack(true)
		}
	}()

	return nil
}

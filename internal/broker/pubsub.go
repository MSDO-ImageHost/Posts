package broker

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

func NewSubPub(handleConf HandleConfig, handler func(req HandleRequestPayload) (res HandleResponsePayload, err error)) error {

	//Declare subscription queue
	if err := QueueDeclare(handleConf.SubQueueConf); err != nil {
		return err
	}

	// Declare publishion queue
	if err := QueueDeclare(handleConf.PubQueueConf); err != nil {
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

			// Verify integrety of message
			if msg.ContentType != "application/json" {
				msg.Ack(false)
				continue
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
				continue
			}

			// Set response message headers
			headers := msg.Headers
			headers["status_code"] = res.Status.Code
			headers["status_code_msg"] = res.Status.Message

			// Publish response
			rabbit.Channel.Publish("", handleConf.PubQueueConf.Name, false, false, amqp.Publishing{
				Headers:       headers,
				ContentType:   "application/json",
				CorrelationId: msg.CorrelationId,
				Timestamp:     time.Now().UTC(),
				Body:          res.Payload,
			})

			log.Println(LOG_TAG, "Acknowledged new post request")
			msg.Ack(true)
		}
	}()

	return nil
}

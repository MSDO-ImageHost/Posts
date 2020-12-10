package broker

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/streadway/amqp"
)

func NewSubPub(hc HandleConfig, handler func(req HandleRequestPayload) (res HandleResponsePayload, err error)) error {

	// Declare 'rapid' exchange
	if err := ExchangeDeclare(hc.ExchangeConf); err != nil {
		return err
	}

	//Declare subscription queue
	if err := QueueDeclare(hc.SubQueueConf); err != nil {
		return err
	}
	if err := hc.SubQueueConf.Bind(hc.ExchangeConf); err != nil {
		return err
	}

	// Declare publishion queues
	for _, q := range hc.PubQueueConfs {
		if err := QueueDeclare(q); err != nil {
			return err
		}
		if err := q.Bind(hc.ExchangeConf); err != nil {
			return err
		}
	}

	// Consumer for subscription messages
	c := hc.ConsumerConf
	consumer, err := rabbit.ConsumeChannel.Consume(hc.SubQueueConf.Name, "", c.AutoAck, c.Exclusive, c.NoLocal, c.NoWait, c.Args)
	if err != nil {
		return err
	}

	// Listen for new messages
	go func() {
		for reqMsg := range consumer {
			log.Printf("%s '%s': Received %s request from sender %s\n", _LOG_TAG, reqMsg.CorrelationId, reqMsg.RoutingKey, reqMsg.ConsumerTag)
			start := time.Now()

			// Start building response message
			resMsg := Response{
				Headers:       amqp.Table{},
				ContentType:   "application/json",
				CorrelationId: reqMsg.CorrelationId,
				AppId:         "posts-v1",
			}

			// Verify delivery integrity
			if err := AssertDelivery(reqMsg); err != nil {
				resMsg.Headers["status_code"] = http.StatusBadRequest
				resMsg.Headers["status_code_reqMsg"] = http.StatusText(http.StatusBadRequest)
				log.Printf("%s '%s': Rejected request: %s\t\n", _LOG_TAG, reqMsg.CorrelationId, err)
				if err := resMsg.Publish(hc, reqMsg); err != nil {
					log.Fatalf("%s '%s': Failed to reply back: %s ", _LOG_TAG, reqMsg.CorrelationId, err)
				}
				continue
			}

			// Run business logic handler
			res, err := handler(HandleRequestPayload{
				Headers: (map[string]interface{})(reqMsg.Headers),
				Payload: reqMsg.Body,
			})
			resMsg.Headers["status_code"] = res.Status.Code
			resMsg.Headers["status_code_reqMsg"] = res.Status.Message

			// Error or rejections from business logic
			if err != nil {
				resMsg.Headers["status_code"] = http.StatusInternalServerError
				resMsg.Headers["status_code_reqMsg"] = err.Error()
				log.Printf("%s '%s': Failed to fulfill request: %s ", _LOG_TAG, reqMsg.CorrelationId, err)
				if err := resMsg.Publish(hc, reqMsg); err != nil {
					log.Fatalf("%s '%s': Failed to reply back: %s ", _LOG_TAG, reqMsg.CorrelationId, err)
				}
				continue
			}

			// Publish response message
			resMsg.Headers["processing_time_ns"] = time.Since(start).Nanoseconds()
			resMsg.Timestamp = time.Now().UTC()
			if err := resMsg.Publish(hc, reqMsg); err != nil {
				log.Fatalf("%s '%s': Failed to reply back: %s ", _LOG_TAG, reqMsg.CorrelationId, err)
			}
			log.Printf("%s '%s': Fulfilled request", _LOG_TAG, reqMsg.CorrelationId)
		}
	}()
	log.Printf("%s Registered event %s", _LOG_TAG, hc.SubQueueConf.Name)
	return nil
}

func (r Response) Publish(hc HandleConfig, reqMsg amqp.Delivery) error {
	resMsg := (amqp.Publishing)(r)

	// Build list of response queues
	pubQueues := append(hc.PubQueueConfs, QueueConfig{Name: reqMsg.ReplyTo})

	// Publish response message
	for _, pubQueue := range pubQueues {
		if err := rabbit.PublishChannel.Publish(hc.ExchangeConf.Name, pubQueue.Name, false, false, resMsg); err != nil {
			return fmt.Errorf("Failed to reply back %s", err)
		}
		if err := rabbit.PublishChannel.Publish(hc.ExchangeConf.Name, pubQueue.Name, false, false, resMsg); err != nil {
			return fmt.Errorf("Failed to reply back: %s", err)
		}
	}
	return reqMsg.Ack(true)
}

func AssertDelivery(msg amqp.Delivery) error {

	if msg.ContentType == "" {
		return fmt.Errorf("ContentType is not specified")
	}
	if msg.ContentType != "application/json" {
		return fmt.Errorf("ContentType is not specified as 'application/json'")
	}

	if !json.Valid(msg.Body) {
		return fmt.Errorf("Body payload contains invalid JSON")
	}

	return nil
}

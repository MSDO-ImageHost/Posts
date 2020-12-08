package broker

import (
	"encoding/json"
	"log"
	"net/http"
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

	// Declare 'rapid' exchange
	if err := ExchangeDeclare(handleConf.QueueBindConf.Exchange); err != nil {
		return err
	}

	// Bind queue to rapid with routing key
	handleConf.QueueBindConf.Name = handleConf.SubQueueConf.Name
	if err := QueueBind(handleConf.QueueBindConf); err != nil {
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
			msg.Ack(true)
			log.Println(_LOG_TAG, "Received request with correlation id", msg.CorrelationId)
			start := time.Now()

			headers := msg.Headers
			headers["status_code"] = http.StatusProcessing
			headers["status_code_msg"] = http.StatusText(http.StatusProcessing)

			// Verify content type and verify json body
			if msg.ContentType != "application/json" || !json.Valid(msg.Body) {
				headers["status_code"] = http.StatusBadRequest
				headers["status_code_msg"] = http.StatusText(http.StatusBadRequest)
				log.Println(_LOG_TAG, "Rejected request with correlation id. Incorrect ContentType or JSON payload", msg.CorrelationId)
				if err := PublicateResponse(handleConf, msg, headers, nil, false, start); err != nil {
					log.Fatal(_LOG_TAG, "Failed process response to", msg.CorrelationId, err)
				}
				continue
			}

			// Run business logic handler
			res, err := handler(HandleRequestPayload{
				Headers: (map[string]interface{})(headers),
				Payload: msg.Body,
			})

			// Don't acknowledge the message if an error happened in business logic
			if err != nil {
				headers["status_code"] = http.StatusInternalServerError
				headers["status_code_msg"] = err.Error()
				log.Println(_LOG_TAG, "Failed to fulfill request with correlation id.", msg.CorrelationId, err)
				if err := PublicateResponse(handleConf, msg, headers, res.Payload, false, start); err != nil {
					log.Fatal(_LOG_TAG, "Failed process response to", msg.CorrelationId, err)
				}
				continue
			}

			// Publish response message
			headers["status_code"] = res.Status.Code
			headers["status_code_msg"] = res.Status.Message
			if err := PublicateResponse(handleConf, msg, headers, res.Payload, true, start); err != nil {
				log.Fatal(_LOG_TAG, "Failed process response to", msg.CorrelationId, err)
			}
			log.Println(_LOG_TAG, "Fulfilled request with correlation id", msg.CorrelationId)
		}
	}()

	log.Printf("%s Registered subpub handler for %s -> %s with routing key %s on exchange %s\n",
		_LOG_TAG,
		handleConf.SubQueueConf.Name,
		handleConf.PubQueueConf.Name,
		handleConf.QueueBindConf.Name,
		handleConf.QueueBindConf.Exchange.Name,
	)
	return nil
}

func PublicateResponse(
	conf HandleConfig,
	msg amqp.Delivery,
	headers amqp.Table,
	payload []byte,
	ack bool,
	start time.Time) (err error) {

	// Check if the sender wants a reply in a specific queue
	responseQueue := conf.PubQueueConf.Name
	if msg.ReplyTo != "" {
		responseQueue = msg.ReplyTo
	}

	// Populate response message and publish it
	headers["processing_time_ns"] = time.Since(start).Nanoseconds()
	err = rabbit.Channel.Publish("", responseQueue, false, false, amqp.Publishing{
		Headers:       headers,
		ContentType:   "application/json",
		CorrelationId: msg.CorrelationId,
		Timestamp:     time.Now().UTC(),
		Body:          payload,
		AppId:         "posts-v1",
	})
	if err != nil || msg.Ack(ack) != nil {
		return err
	}
	return nil
}

package broker

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/MSDO-ImageHost/Posts/internal/auth"
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
			log.Println(_LOG_TAG, "Received request with correlation id", msg.CorrelationId)
			start := time.Now()

			headers := msg.Headers
			headers["status_code"] = http.StatusProcessing
			headers["status_code_msg"] = http.StatusText(http.StatusProcessing)

			// Verify content type and verify json body
			if msg.ContentType != "application/json" || !json.Valid(msg.Body) {
				headers["status_code"] = http.StatusNotAcceptable
				headers["status_code_msg"] = "Invalid content type" // http.StatusText(http.StatusNotAcceptable)
				log.Println(_LOG_TAG, "Rejected request with correlation id", msg.CorrelationId)
				PublicateResponse(handleConf, msg, headers, nil, false, start)
				continue
			}

			// Verify JWT and decompose
			jwt, err := auth.Decompose("")
			if err != nil {
				headers["status_code"] = http.StatusNonAuthoritativeInfo
				headers["status_code_msg"] = http.StatusText(http.StatusNonAuthoritativeInfo)
				log.Println(_LOG_TAG, "Rejected request with correlation id", msg.CorrelationId)
				PublicateResponse(handleConf, msg, headers, nil, false, start)
				continue
			}

			// Run business logic handler
			res, err := handler(HandleRequestPayload{
				UserID:  jwt.UserID,
				Role:    jwt.Role,
				Payload: msg.Body,
			})

			// Don't acknowledge the message if an error happened in business logic
			if err != nil {
				headers["status_code"] = http.StatusInternalServerError
				headers["status_code_msg"] = err.Error()
				log.Println(_LOG_TAG, "Failed to fulfill request with correlation id", msg.CorrelationId)
				PublicateResponse(handleConf, msg, headers, res.Payload, false, start)
				continue
			}

			// Publish response message
			headers["status_code"] = res.Status.Code
			headers["status_code_msg"] = res.Status.Message
			PublicateResponse(handleConf, msg, headers, res.Payload, true, start)
			log.Println(_LOG_TAG, "Fulfilled request with correlation id", msg.CorrelationId)
		}
	}()

	log.Printf("Registered subpub handler for %s -> %s\n", handleConf.SubQueueConf.Name, handleConf.PubQueueConf.Name)
	return nil
}

func PublicateResponse(conf HandleConfig, msg amqp.Delivery, headers amqp.Table, payload []byte, ack bool, start time.Time) /*(err error)*/ {

	headers["processing_time_ns"] = time.Since(start).Nanoseconds()

	err := rabbit.Channel.Publish("", conf.PubQueueConf.Name, false, false, amqp.Publishing{
		Headers:       headers,
		ContentType:   "application/json",
		CorrelationId: msg.CorrelationId,
		Timestamp:     time.Now().UTC(),
		Body:          payload,
	})
	if err != nil {
		log.Fatal(_LOG_TAG, "Failed to publish response to", msg.CorrelationId, err)
	}

	if err := msg.Ack(ack); err != nil {
		log.Fatal(_LOG_TAG, "Failed to acknowledge message", msg.MessageId, err)
	}

	//return nil
}

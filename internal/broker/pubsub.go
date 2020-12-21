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

	// Declare subscription queue and bind to rapid
	if err := QueueDeclare(hc.SubQueueConf); err != nil {
		return err
	}
	if err := hc.SubQueueConf.Bind(hc.ExchangeConf, hc.SubIntent); err != nil {
		return err
	}

	// Declare publishion queues
	for _, q := range hc.PubQueueConfs {
		if err := QueueDeclare(q); err != nil {
			return err
		}
		if err := q.Bind(hc.ExchangeConf, hc.PubIntent); err != nil {
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
			log.Printf("%s CID '%s': Received %s request from sender %s\n", _LOG_TAG, reqMsg.CorrelationId, reqMsg.RoutingKey, reqMsg.ConsumerTag)
			start := time.Now()

			// Start building response message
			resMsg := Response{Headers: amqp.Table{}, ContentType: "application/json", CorrelationId: reqMsg.CorrelationId, AppId: "posts-v1"}

			// Verify delivery integrity
			if err := AssertDelivery(reqMsg); err != nil {
				log.Printf("%s CID '%s': Rejected request: %s\t\n", _LOG_TAG, reqMsg.CorrelationId, err)
				resMsg.Publish(start, hc, reqMsg, http.StatusBadRequest, err)
				continue
			}

			// Run business logic handler
			res, err := handler(HandleRequestPayload{
				Headers: (map[string]interface{})(reqMsg.Headers),
				Payload: reqMsg.Body,
			})

			// Error or rejections from business logic
			if err != nil {
				log.Printf("%s CID '%s': Failed to fulfill request: %s ", _LOG_TAG, reqMsg.CorrelationId, err)
				resMsg.Publish(start, hc, reqMsg, res.Status.Code, err)
				continue
			}

			// Publish response message
			resMsg.Body = res.Payload
			resMsg.Timestamp = time.Now().UTC()
			resMsg.Publish(start, hc, reqMsg, res.Status.Code, nil)

			// Roger and over
			log.Printf("%s CID '%s': Fulfilled request", _LOG_TAG, reqMsg.CorrelationId)
		}
	}()
	log.Printf("%s Listing for %s events -> (%s)", _LOG_TAG, hc.SubIntent.String(), hc.SubQueueConf.Name)
	return nil
}

// Emit the response
func (r Response) Publish(start time.Time, hc HandleConfig, reqMsg amqp.Delivery, statusCode int, err error) {
	resMsg := (amqp.Publishing)(r)

	var errMsg string = ""
	if err != nil {
		errMsg = err.Error()
	}

	// Set payload headers
	if reqMsg.Headers["jwt"] != nil {
		resMsg.Headers["jwt"] = reqMsg.Headers["jwt"]
	}
	resMsg.Headers["processing_time_ns"] = time.Since(start).Nanoseconds()
	resMsg.Headers["status_code"] = statusCode
	resMsg.Headers["status_msg"] = fmt.Sprintf("%s. %s", http.StatusText(statusCode), errMsg)

	// Publish message
	if err := rabbit.PublishChannel.Publish(hc.ExchangeConf.Name, hc.PubIntent.String(), false, false, resMsg); err != nil {
		log.Fatalf("%s Failed to emit response %s -> %s: %s", _LOG_TAG, hc.PubIntent.String(), hc.ExchangeConf.Name, err)
	}
	log.Printf("%s CID '%s': Emitted response. %s -> %s", _LOG_TAG, reqMsg.CorrelationId, hc.PubIntent.String(), hc.ExchangeConf.Name)

	if err := reqMsg.Ack(true); err != nil {
		log.Fatalf("%s Failed to acknowledge request: %s", _LOG_TAG, err)
	}
}

// Check if basic parameters are ok on the delivery
func AssertDelivery(msg amqp.Delivery) error {
	if msg.ContentType == "" {
		return fmt.Errorf("ContentType is not specified")
	}
	if msg.ContentType != "application/json" {
		return fmt.Errorf("ContentType is not specified as 'application/json'")
	}

	if !json.Valid(msg.Body) {
		return fmt.Errorf("Message payload contains invalid JSON")
	}
	return nil
}

package broker

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/MSDO-ImageHost/Posts/pkg/api"
	"github.com/streadway/amqp"
)

func (b *RabbitBroker) NewAsyncHandler(queueName string, reqHandler func(amqp.Delivery) (interface{}, bool, error)) {

	/** Routing setup **/
	defaultQueueName := fmt.Sprintf("%s-%s", EConfig.Name, queueName)
	queue, err := b.Channel.QueueDeclare(defaultQueueName, QConfig.Durable, QConfig.AutoDelete, QConfig.Exclusive, QConfig.NoWait, QConfig.Args)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Broker: Declared queue %s\n", queueName)

	if err := b.Channel.QueueBind(queue.Name, queueName, QBConfig.Exchange, QBConfig.NoWait, QBConfig.Args); err != nil {
		log.Fatal(err)
	}
	log.Printf("Broker: Bound queue %s to exchange %s with binding/routing key %s\n", queue.Name, QBConfig.Exchange, queue.Name)

	consume, err := b.Channel.Consume(queue.Name, CConfig.Consumer, CConfig.AutoAck, CConfig.Exclusive, CConfig.NoLocal, CConfig.NoWait, CConfig.Args)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Broker: Opened consumer for queue %s on channel\n", queueName)
	/** **/

	// Process any incomming messages on the queue
	go func() {
		for msg := range consume {

			// Measure processing time
			start := time.Now()

			fmt.Printf("Recived a new message on %s. Producer expects a reply on '%s'\n", queueName, msg.ReplyTo)
			if len(msg.Body) == 0 {
				log.Println("No data in message")
			}

			// Process message and respond
			handlerRes, ack, err := reqHandler(msg)
			if err != nil {
				log.Fatal(err)
			}

			response := api.ResponseWrapper{
				Data:             handlerRes,
				ProcessingTimeNs: time.Since(start).Nanoseconds(),
			}

			// Acknowledge message processed
			msg.Ack(ack)

			jsonRes, err := json.Marshal(response)
			if err != nil {
				log.Fatal(err)
			}

			responsePayload := amqp.Publishing{
				ContentType: "application/json",
				//ContentEncoding: nil,
				Body:          jsonRes,
				Timestamp:     time.Now(),
				AppId:         "posts-v01",
				CorrelationId: msg.CorrelationId,
			}

			if err := b.Channel.Publish("", msg.ReplyTo, false, false, responsePayload); err != nil {
				log.Fatal(err)
			}
		}
	}()
	log.Printf("Broker: Registered async handler for queue %s\n", queueName)
	return
}

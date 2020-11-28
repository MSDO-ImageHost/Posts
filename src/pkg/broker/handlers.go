package broker

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func (b *RabbitBroker) NewAsyncHandler(queueName string, reqHandler func(amqp.Delivery) (interface{}, bool, error)) {

	queue, err := b.Channel.QueueDeclare(
		queueName,
		QConfig.Durable,
		QConfig.AutoDelete,
		QConfig.Exclusive,
		QConfig.NoWait,
		QConfig.Args,
	)
	if err != nil {
		log.Fatal(err)
	}

	consume, err := b.Channel.Consume(
		queue.Name,
		CConfig.Consumer,
		CConfig.AutoAck,
		CConfig.Exclusive,
		CConfig.NoLocal,
		CConfig.NoWait,
		CConfig.Args,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Process any incomming messages on the queue
	go func() {
		for msg := range consume {
			fmt.Printf("Recived a new message on %s\n", queueName)
			if msg.Body == nil {
				log.Println("No data in message")
			}

			log.Println(msg.RoutingKey, msg.ReplyTo)

			// Process message and respond
			res, ack, err := reqHandler(msg)
			if err != nil {
				log.Fatal(err)
			}

			// Acknowledge message processed
			msg.Ack(ack)

			fmt.Println("Processed the message")

			jsonRes, err := json.Marshal(res)
			if err != nil {
				log.Fatal(err)
			}

			resPayload := amqp.Publishing{
				ContentType: "application/json",
				Body:        jsonRes,
			}

			if err := b.Channel.Publish("", msg.ReplyTo, false, false, resPayload); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Responded to the producer on queue %s\n", msg.ReplyTo)
			fmt.Println(string(jsonRes))

		}
	}()
	log.Printf("Broker: Registered async handler for queue %s\n", queueName)
	return
}

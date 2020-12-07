package broker

import (
	"log"

	"github.com/streadway/amqp"
)

// Declares a new exchange for
func ExchangeDeclare(c ExchangeConfig) error {
	log.Printf("%s Declaring Exchange: %+v", _LOG_TAG, c)
	return rabbit.Channel.ExchangeDeclare(c.Name, c.Kind, c.Durable, c.AutoDelete, c.Internal, c.NoWait, c.Args)
}

// Declares a new queue
func QueueDeclare(c QueueConfig) error {
	log.Printf("%s Declaring Queue: %+v", _LOG_TAG, c)

	_, err := rabbit.Channel.QueueDeclare(c.Name, c.Durable, c.AutoDelete, c.Exclusive, c.NoWait, c.Args)
	if err != nil {
		return err
	}
	return nil
}

// Binds a queue to an exchange with specified routing key
func QueueBind(c QueueBindConfig) error {
	log.Printf("%s Binding queue: %+v", _LOG_TAG, c)
	return rabbit.Channel.QueueBind(c.Name, c.Key, c.Exchange.Name, c.NoWait, c.Args)
}

// Declare a new consumer for queue
func ConsumerDeclare(hc HandleConfig) (<-chan amqp.Delivery, error) {
	log.Printf("%s Declaring consumer: %+v", _LOG_TAG, hc)
	c := hc.ConsumerConf
	return rabbit.Channel.Consume(hc.SubQueueConf.Name, "", c.AutoAck, c.Exclusive, c.NoLocal, c.NoWait, c.Args)
}

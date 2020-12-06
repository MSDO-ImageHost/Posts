package broker

import (
	"github.com/streadway/amqp"
)

// Declares a new exchange for
func ExchangeDeclare(c ExchangeConfig) error {
	return rabbit.Channel.ExchangeDeclare(c.Name, c.Kind, c.Durable, c.AutoDelete, c.Internal, c.NoWait, c.Args)
}

// Declares a new queue
func QueueDeclare(c QueueConfig) error {
	_, err := rabbit.Channel.QueueDeclare(c.Name, c.Durable, c.AutoDelete, c.Exclusive, c.NoWait, c.Args)
	if err != nil {
		return err
	}
	return nil
}

// Binds a queue to an exchange with specified routing key
func QueueBind(c QueueBindConfig) error {
	return rabbit.Channel.QueueBind(c.Name, c.Key, c.Exchange.Name, c.NoWait, c.Args)
}

// Declare a new consumer for queue
func ConsumerDeclare(hc HandleConfig) (<-chan amqp.Delivery, error) {
	c := hc.ConsumerConf
	return rabbit.Channel.Consume(hc.SubQueueConf.Name, "", c.AutoAck, c.Exclusive, c.NoLocal, c.NoWait, c.Args)
}

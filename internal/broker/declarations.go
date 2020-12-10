package broker

import (
	"github.com/streadway/amqp"
)

// Declares a new exchange for
func ExchangeDeclare(c ExchangeConfig) error {
	return rabbit.ConsumeChannel.ExchangeDeclare(c.Name, c.Kind, c.Durable, c.AutoDelete, c.Internal, c.NoWait, c.Args)
}

// Declares a new queue
func QueueDeclare(c QueueConfig) error {
	_, err := rabbit.ConsumeChannel.QueueDeclare(c.Name, c.Durable, c.AutoDelete, c.Exclusive, c.NoWait, c.Args)
	if err != nil {
		return err
	}
	return nil
}

// Declares a new queue
func QueuesDeclare(cs []QueueConfig) error {
	for _, c := range cs {
		_, err := rabbit.ConsumeChannel.QueueDeclare(c.Name, c.Durable, c.AutoDelete, c.Exclusive, c.NoWait, c.Args)
		if err != nil {
			return err
		}
	}
	return nil
}

// Binds a queue to an exchange with specified routing key
func (qc *QueueConfig) Bind(e ExchangeConfig, i Intent) error {
	return rabbit.ConsumeChannel.QueueBind(qc.Name, i.String(), e.Name, qc.NoWait, qc.Args)
}

// Declare a new consumer for queue
func ConsumerDeclare(hc HandleConfig) (<-chan amqp.Delivery, error) {
	c := hc.ConsumerConf
	return rabbit.ConsumeChannel.Consume(hc.SubQueueConf.Name, "", c.AutoAck, c.Exclusive, c.NoLocal, c.NoWait, c.Args)
}

func (i Intent) String() string {
	return (string)(i)
}

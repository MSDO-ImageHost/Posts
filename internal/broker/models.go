package broker

import (
	"github.com/MSDO-ImageHost/Posts/internal/api"
	"github.com/streadway/amqp"
)

// Instance to hold the RabbitMQ host and channel
type RabbitBroker struct {
	Host    *amqp.Connection
	Channel *amqp.Channel
}

// Data structure that is passed into any business logic handler
type HandleRequestPayload struct {
	Headers map[string]interface{}
	Payload []byte
}

// Data structure returned by the business logic handler
type HandleResponsePayload struct {
	Payload []byte
	Status  api.StatusCode
}

// Configuration structure for a RabbitMQ exchange
type ExchangeConfig struct {
	Name, Kind                            string
	Durable, AutoDelete, Internal, NoWait bool
	Args                                  amqp.Table
}

// BindQueue configuration for binding queues in RabbitMQ
type QueueBindConfig struct {
	Name, Key string
	Exchange  ExchangeConfig
	NoWait    bool
	Args      amqp.Table
}

// Queue configuration for a new queue in RabbitMQ
type QueueConfig struct {
	Name                                   string
	Durable, AutoDelete, Exclusive, NoWait bool
	Args                                   amqp.Table
	Bind                                   *QueueBindConfig
}

// Configuration of the consumer used to consume messages from RabbitMQ
type ConsumerConfig struct {
	AutoAck, Exclusive, NoLocal, NoWait bool
	Args                                amqp.Table
}

// Configuration object that composes other structures into a single
type HandleConfig struct {
	SubQueueConf, PubQueueConf QueueConfig
	ConsumerConf               ConsumerConfig
	//ExchangeConf               ExchangeConfig
	QueueBindConf QueueBindConfig
}

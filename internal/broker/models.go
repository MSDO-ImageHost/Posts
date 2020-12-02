package broker

import (
	"github.com/MSDO-ImageHost/Posts/internal/api"
	"github.com/streadway/amqp"
)

type RabbitBroker struct {
	Host    *amqp.Connection
	Channel *amqp.Channel
}

type HandleRequestPayload struct {
	UserID, Role string
	Payload      []byte
}

type HandleResponsePayload struct {
	Payload []byte
	Status  api.StatusCode
}

type QueueConfig struct {
	Name                                   string
	Durable, AutoDelete, Exclusive, NoWait bool
	Args                                   amqp.Table
}

type ConsumerConfig struct {
	AutoAck, Exclusive, NoLocal, NoWait bool
	Args                                amqp.Table
}

type HandleConfig struct {
	SubQueueConf, PubQueueConf QueueConfig
	ConsumerConf               ConsumerConfig
}

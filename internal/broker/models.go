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
	Headers map[string]interface{}
	Payload []byte
}

type HandleResponsePayload struct {
	Payload []byte
	Status  api.StatusCode
}

type ExchangeConfig struct {
	Name, Kind                            string
	Durable, AutoDelete, Internal, NoWait bool
	Args                                  amqp.Table
}

type QueueBindConfig struct {
	Name, Key string
	Exchange  ExchangeConfig
	NoWait    bool
	Args      amqp.Table
}
type QueueConfig struct {
	Name                                   string
	Durable, AutoDelete, Exclusive, NoWait bool
	Args                                   amqp.Table
	Bind                                   *QueueBindConfig
}

type ConsumerConfig struct {
	AutoAck, Exclusive, NoLocal, NoWait bool
	Args                                amqp.Table
}

type HandleConfig struct {
	SubQueueConf, PubQueueConf QueueConfig
	ConsumerConf               ConsumerConfig
	ExchangeConf               ExchangeConfig
}

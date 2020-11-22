package broker

import (
	"github.com/streadway/amqp"
)

type RabbitBroker struct {
	Host    *amqp.Connection
	Channel *amqp.Channel
}

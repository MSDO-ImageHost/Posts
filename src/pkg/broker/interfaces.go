package broker

import "github.com/streadway/amqp"

type BrokerInterface interface {
	NewAsyncHandler(queueName string, reqHandler func(amqp.Delivery) (interface{}, bool, error))
}

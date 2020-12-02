package broker

import (
	"log"

	"github.com/streadway/amqp"
)

func QueueDeclare(config QueueConfig) error {
	_, err := rabbit.Channel.QueueDeclare(
		config.Name,
		config.Durable,
		config.AutoDelete,
		config.Exclusive,
		config.NoWait,
		config.Args,
	)
	if err != nil {
		return err
	}

	log.Println(LOG_TAG, "Declared queue", config.Name)
	return nil
}

func ConsumerDeclare(config HandleConfig) (<-chan amqp.Delivery, error) {
	return rabbit.Channel.Consume(
		config.SubQueueConf.Name,
		"",
		config.ConsumerConf.AutoAck,
		config.ConsumerConf.Exclusive,
		config.ConsumerConf.NoLocal,
		config.ConsumerConf.NoWait,
		config.ConsumerConf.Args,
	)
}

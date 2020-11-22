package broker

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

var (
	broker RabbitBroker
	Broker BrokerInterface
)

// Initializes a new connection to RabbitMQ broker
func Init() error {
	log.Println("Opening broker connection")

	// Connect to broker
	connection, err := amqp.Dial(os.Getenv("RABBITMQ_CONN_URI"))
	if err != nil {
		return err
	}

	// Create channel
	channel, err := connection.Channel()
	if err != nil {
		return err
	}

	if err = channel.Qos(1, 0, false); err != nil {
		return err
	}

	// Store reference in memory
	broker = RabbitBroker{
		Host:    connection,
		Channel: channel,
	}

	log.Println("Opened broker connection")
	return nil
}

func GetBroker() RabbitBroker {
	return broker
}

func Deinit() error {
	log.Println("Closing broker connection")
	// Close channel in use
	if err := broker.Channel.Close(); err != nil {
		return err
	}

	// Terminate connection to broker
	if err := broker.Host.Close(); err != nil {
		return err
	}
	log.Println("Closed broker connection")
	return nil
}

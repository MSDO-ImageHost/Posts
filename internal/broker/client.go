package broker

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

var (
	rabbit RabbitBroker
)

// Initializes a new connection to RabbitMQ broker
func Init() error {
	log.Println(LOG_TAG, "Setting up")

	// Connect to broker
	log.Println(LOG_TAG, "Opening connection")
	connection, err := amqp.Dial(os.Getenv("RABBITMQ_CONN_URI"))
	if err != nil {
		return err
	}
	log.Println(LOG_TAG, "Opened connection")

	// Create channel
	log.Println(LOG_TAG, "Opening channel")
	channel, err := connection.Channel()
	if err != nil {
		return err
	}

	if err = channel.Qos(1, 0, false); err != nil {
		return err
	}
	log.Println(LOG_TAG, "Channel opened")

	// Store reference in memory
	rabbit = RabbitBroker{
		Host:    connection,
		Channel: channel,
	}

	log.Println(LOG_TAG, "Finished setup")
	return nil
}

func Deinit() error {
	log.Println(LOG_TAG, "Closing connection")
	// Close channel in use
	if err := rabbit.Channel.Close(); err != nil {
		return err
	}

	// Terminate connection to broker
	if err := rabbit.Host.Close(); err != nil {
		return err
	}
	log.Println(LOG_TAG, "Closed connection")
	return nil
}

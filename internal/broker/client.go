package broker

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

var rabbit RabbitBroker

// Initializes a new connection to RabbitMQ broker
func Init() error {
	log.Println(_LOG_TAG, "Setting up client connection")

	//connUri := fmt.Sprint("amqp://", "dev" ,":", "dev" "@", os.Getenv("RABBITMQ_CONN_URI"))

	// Connect to broker
	log.Println(_LOG_TAG, "Opening connection")
	connection, err := amqp.Dial(os.Getenv("RABBITMQ_CONN_URI"))
	if err != nil {
		return err
	}
	log.Println(_LOG_TAG, "Opened connection")

	// Create consume channel
	log.Println(_LOG_TAG, "Opening consume channel")
	consumeChannel, err := connection.Channel()
	if err != nil {
		return err
	}
	if err = consumeChannel.Qos(1, 0, false); err != nil {
		return err
	}
	log.Println(_LOG_TAG, "Consume channel opened")

	// Create publish channel
	log.Println(_LOG_TAG, "Opening publish channel")
	publishChannel, err := connection.Channel()
	if err != nil {
		return err
	}
	if err = publishChannel.Qos(1, 0, false); err != nil {
		return err
	}
	log.Println(_LOG_TAG, "Consume publish opened")

	// Store reference in memory
	rabbit = RabbitBroker{
		Host:           connection,
		ConsumeChannel: consumeChannel,
		PublishChannel: publishChannel,
	}

	log.Println(_LOG_TAG, "Finished client connection setup")
	return nil
}

// Closes the active connection to the RabbitMQ host
func Deinit() error {
	log.Println(_LOG_TAG, "Closing client connection")
	// Close channels in use
	if err := rabbit.ConsumeChannel.Close(); err != nil {
		return err
	}
	if err := rabbit.PublishChannel.Close(); err != nil {
		return err
	}

	// Terminate connection to broker
	if err := rabbit.Host.Close(); err != nil {
		return err
	}
	log.Println(_LOG_TAG, "Closed client connection")
	return nil
}

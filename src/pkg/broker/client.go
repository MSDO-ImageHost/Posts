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
	log.Println("Broker: Setting up")

	// Connect to broker
	log.Println("Broker: Opening connection")
	connection, err := amqp.Dial(os.Getenv("RABBITMQ_CONN_URI"))
	if err != nil {
		return err
	}
	log.Println("Broker: Opened connection")

	// Create channel
	log.Println("Broker: Opening channel")
	channel, err := connection.Channel()
	if err != nil {
		return err
	}

	if err = channel.Qos(1, 0, false); err != nil {
		return err
	}
	log.Println("Broker: Channel opened")

	//log.Println("Broker: Registering exchange")
	//if err := channel.ExchangeDeclare(EConfig.Name, EConfig.Kind, EConfig.Durable, EConfig.AutoDelete, EConfig.Internal, EConfig.NoWait, EConfig.Args); err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("Broker: Exchange configured")

	// Store reference in memory
	broker = RabbitBroker{
		Host:    connection,
		Channel: channel,
	}

	Broker = &broker
	log.Println("Broker: Finished setup")

	return nil
}

func Deinit() error {
	log.Println("Broker: Closing connection")
	// Close channel in use
	if err := broker.Channel.Close(); err != nil {
		return err
	}

	// Terminate connection to broker
	if err := broker.Host.Close(); err != nil {
		return err
	}
	log.Println("Broker: Closed connection")
	return nil
}

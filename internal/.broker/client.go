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
	log.Println("Broker:\tSetting up")

	// Connect to broker
	log.Println("Broker:\tOpening connection")
	connection, err := amqp.Dial(os.Getenv("RABBITMQ_CONN_URI"))
	if err != nil {
		return err
	}
	log.Println("Broker:\tOpened connection")

	// Create channel
	log.Println("Broker:\tOpening channel")
	channel, err := connection.Channel()
	if err != nil {
		return err
	}

	if err = channel.Qos(1, 0, false); err != nil {
		return err
	}
	log.Println("Broker:\tChannel opened")

	log.Println("Broker:\tRegistering exchange")
	if err := channel.ExchangeDeclare(EConfig.Name, EConfig.Kind, EConfig.Durable, EConfig.AutoDelete, EConfig.Internal, EConfig.NoWait, EConfig.Args); err != nil {
		log.Fatal(err)
	}
	log.Println("Broker:\tExchange configured")

	// Store reference in memory
	broker = RabbitBroker{
		Host:    connection,
		Channel: channel,
	}

	Broker = &broker
	log.Println("Broker:\tFinished setup")

	return nil
}

func Deinit() error {
	log.Println("Broker:\tClosing connection")
	// Close channel in use
	if err := broker.Channel.Close(); err != nil {
		return err
	}

	// Terminate connection to broker
	if err := broker.Host.Close(); err != nil {
		return err
	}
	log.Println("Broker:\tClosed connection")
	return nil
}

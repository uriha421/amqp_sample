package main

import (
	"github.com/streadway/amqp"
	"os"
)

func main() {
	// import an Environment Variable for AMQP_URL
	amqpURL := os.Getenv("AMQP_URL")
	if amqpURL == "" {
		amqpURL = "amqp://guest:guest@localhost:5672"
	}

	// create an amqp connection
	connection, err := amqp.Dial(amqpURL)
	if err != nil {
		panic("could not establish AMQP connection: " + err.Error())
	}

	// create a channel on the connection
	channel, err := connection.Channel()
	if err != nil {
		panic("could not open channel: " + err.Error())
	}

	// create an exchange on which publisher publishes messages
	/*
		arguments of ExchangeDeclare():
			1. the exhange name,
			2. the exchange type, "direct", "fanout", or "topic",
			7. additional configuration parameters.
	*/
	err = channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	// prepare a message for publishing
	message := amqp.Publishing{
		Body: []byte("Hello Wolrd"),
	}

	// publish a message on the exchange
	/*
		arguments of Publish():
			1. the exchange name on which you will publish,
			2. the message's routing key,
			5. the message.
	*/
	err = channel.Publish("events", "some-routing-key", false, false, message)
	if err != nil {
		panic("error while publishing a message: " + err.Error())
	}

	defer connection.Close()
}

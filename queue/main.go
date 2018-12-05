package main

import (
	"fmt"
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

	// declare a new Queue
	_, err = channel.QueueDeclare("my_queue", true, false, false, false, nil)
	if err != nil {
		panic("error while declaring the queue: " + err.Error())
	}

	// bind the Queue to ...
	err = channel.QueueBind("my_queue", "#", "events", false, nil)
	if err != nil {
		panic("error while binding the queue " + err.Error())
	}

	// consume the Queue
	/*
		arguments of Consume():
			1. the name of the queue to be consumed,
			2. a unique identifier, but if it is nil, a unique identifier automatically generated.
	*/
	msgs, err := channel.Consume("my_queue", "", false, false, false, false, nil)
	if err != nil {
		panic("error while consuming the queue: " + err.Error())
	}

	// read the msgs channel
	for msg := range msgs {
		fmt.Println("message received: " + string(msg.Body))
		msg.Ack(false)
	}

	defer connection.Close()
}

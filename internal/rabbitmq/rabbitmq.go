package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

// Service - interface describing rabitmq service
type Service interface {
	Connect() error
	Publish(message string)
}

// RabbitMQ - struct holding service details
type RabbitMQ struct {
	Conn *amqp.Connection
	Channel *amqp.Channel
}

// Connect - 
func (rmq *RabbitMQ) Connect() error {
	fmt.Println("Connecting to RabbitMQ...")
	var err error
	rmq.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	fmt.Println("Successfuly Connected to RabbitMQ")
	
	rmq.Channel, err = rmq.Conn.Channel()
	if err != nil {
		return err
	}

	_, err = rmq.Channel.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	return nil
}

// Publish takes string message and publish it into queue
func (r *RabbitMQ) Publish(message string) error {
	err := r.Channel.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(message),
		},
	)

	if err != nil {
		return err
	}

	fmt.Println("Successfully published message to queue")
	return nil
}

// NewRabbitMQService - 
func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}
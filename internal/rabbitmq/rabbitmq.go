package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

// Service - interface describing rabitmq service
type Service interface {
	Connect() error
}

// RabbitMQ - struct holding service details
type RabbitMQ struct {
	Conn *amqp.Connection
}

func (rmq *RabbitMQ) Connect() error {
	fmt.Println("Connecting to RabbitMQ...")
	var err error
	rmq.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	fmt.Println("Successfuly Connected to RabbitMQ") 
	return nil
}

func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}
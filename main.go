package main

import (
	"fmt"
	"github.com/chowanij/go_rabbitmq-crash/internal/rabbitmq"
)

type App struct {
	Rmq *rabbitmq.RabbitMQ
}

// Run - handle instantation of our application
func Run() error {
	fmt.Println("Golang RabbitMQ")

	rmq := rabbitmq.NewRabbitMQService()

	app := App{
		Rmq: rmq,
	}

	err := app.Rmq.Connect()
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer app.Rmq.Conn.Close()

	err = app.Rmq.Publish("Hi")

	if err != nil {
		fmt.Println(err)
		return err
	}

	app.Rmq.Consume()

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("Error setting up our application")
		fmt.Println(err)
	}
}
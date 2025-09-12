package pkg

import (
	"billing-app/internal/services"
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	MqAddr string
	MqPort string
	MqUser string
	MqPass string

	Services *services.Service
}

func NewRabbitMQ(addr, port, user, pass string, s *services.Service) *RabbitMQ {
	return &RabbitMQ{
		MqAddr: addr,
		MqPort: port,
		MqUser: user,
		MqPass: pass,

		Services: s,
	}
}

func (rmq *RabbitMQ) Connect() {
	addr := fmt.Sprintf("amqp://%s:%s@%s:%s/", rmq.MqUser, rmq.MqPass, rmq.MqAddr, rmq.MqPort)
	fmt.Println(addr)

	conn, err := amqp091.Dial(addr)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"billing", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			err = rmq.Services.InsertBilling(d.Body)
			if err != nil {
				log.Printf("Error: %s", err)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

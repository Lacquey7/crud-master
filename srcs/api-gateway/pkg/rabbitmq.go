package pkg

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	MqAddr string
	MqPort string
	MqUser string
	MqPass string
}

func NewRabbitMQ(addr, port, user, pass string) *RabbitMQ {
	return &RabbitMQ{
		MqAddr: addr,
		MqPort: port,
		MqUser: user,
		MqPass: pass,
	}
}

func (rmq *RabbitMQ) SendToQueue(payload []byte) {
	addr := fmt.Sprintf("amqp://%s:%s@%s:%s/", rmq.MqUser, rmq.MqPass, rmq.MqAddr, rmq.MqPort)
	fmt.Println(addr)

	conn, err := amqp091.Dial(addr)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	defer ch.Close()
	defer conn.Close()
	q, err := ch.QueueDeclare(
		"billing", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        payload,
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", payload)
	return
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

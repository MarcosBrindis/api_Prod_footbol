package messaging

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewRabbitMQ(url, queueName string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		conn:    conn,
		channel: ch,
		queue:   q,
	}, nil
}

func (r *RabbitMQ) PublishEvent(ctx context.Context, message string) error {
	log.Printf("Publishing message: %s", message)
	return r.channel.PublishWithContext(ctx,
		"",
		r.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(message),
		})
}

func (r *RabbitMQ) Close() {
	r.channel.Close()
	r.conn.Close()
}

func InitRabbitMQ() (*RabbitMQ, error) {
	return NewRabbitMQ("amqp://MarcosDaniel:123456789a@54.146.109.24:5672/", "sport_events")
}

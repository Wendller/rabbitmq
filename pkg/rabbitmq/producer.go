package rabbitmq

import (
	"context"
	"log"

	"github.com/Wendller/gorabbitmq/pkg"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Publish(ctx context.Context, channel *amqp.Channel, queue *amqp.Queue, body string) {
	err := channel.PublishWithContext(ctx,
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	pkg.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}

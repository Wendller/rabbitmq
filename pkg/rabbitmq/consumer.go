package rabbitmq

import (
	"bytes"
	"log"
	"time"

	"github.com/Wendller/gorabbitmq/pkg"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Consume(channel *amqp.Channel, queue *amqp.Queue) {
	messages, err := channel.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	pkg.FailOnError(err, "Failed to register a consumer")

	forever := make(chan struct{})

	go func() {
		for message := range messages {
			log.Printf("Received a message: %s", message.Body)

			dotCount := bytes.Count(message.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)

			log.Printf("Done")
			message.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

package rabbitmq

import (
	"github.com/Wendller/gorabbitmq/pkg"
	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*amqp.Channel, *amqp.Queue) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	pkg.FailOnError(err, "Failed to connect to RabbitMQ")

	channel, err := conn.Channel()
	pkg.FailOnError(err, "Failed to connect to RabbitMQ")

	queue, err := channel.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	pkg.FailOnError(err, "Failed to declare the queue")

	err = channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	pkg.FailOnError(err, "Failed to set QoS")

	return channel, &queue
}

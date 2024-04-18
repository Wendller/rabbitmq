package main

import (
	"github.com/Wendller/gorabbitmq/pkg/rabbitmq"
)

func main() {
	channel, queue := rabbitmq.OpenChannel()
	defer channel.Close()

	rabbitmq.Consume(channel, queue)
}

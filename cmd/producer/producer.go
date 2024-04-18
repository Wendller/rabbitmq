package main

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/Wendller/gorabbitmq/pkg/rabbitmq"
)

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "Hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

func main() {
	channel, queue := rabbitmq.OpenChannel()
	defer channel.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := bodyFrom(os.Args)

	rabbitmq.Publish(ctx, channel, queue, body)
}

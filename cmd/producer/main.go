package main

import (
	"context"
	"github.com/carloseduribeiro/working-with-events/pkg/rabbitmq"
)

func main() {
	ctx := context.Background()
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	if err = rabbitmq.Publish(ctx, ch, "amq.direct", "Hello World, RabbitMQ!"); err != nil {
		panic(err)
	}
}

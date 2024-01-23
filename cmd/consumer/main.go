package main

import (
	"context"
	"fmt"
	"github.com/carloseduribeiro/working-with-events/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ctx := context.TODO()
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	messages := make(chan amqp.Delivery)
	go rabbitmq.Consume(ctx, ch, messages, "myQueue")
	for message := range messages {
		fmt.Println(string(message.Body))
		message.Ack(false)
	}
}

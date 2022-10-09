package main

import (
	"context"
	"fmt"
	"github.com/gsoultan/kelinci"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	forever := make(chan bool)
	config := kelinci.NewConfigBuilder()
	config.SetHost("localhost")
	config.SetPassword("guest")
	config.SetUserName("guest")
	config.SetPort("5672")
	config.SetVHost("/")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	a := kelinci.NewConnection("test", *config.Build())
	fmt.Println(a.Reconnect())
	b := kelinci.NewPublisher(a)

	c := kelinci.NewConsumer(a)
	message := "hi"
	m := amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	}

	go func() {
		c.StartConsuming("user.registered.queue", d)
	}()

	fmt.Println(b.Publish(ctx, "user.registered", "fanout", m))

	message = "ho"
	ma := amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	}
	fmt.Println(b.Publish(ctx, "user.registered", "fanout", ma))
	<-forever
}

func d(message amqp091.Delivery) {
	fmt.Println(string(message.Body[:]))
	message.Ack(false)
}

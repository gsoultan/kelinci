package main

import (
	"fmt"
	"github.com/gsoultan/kelinci"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	config := kelinci.NewConfigBuilder()
	config.SetHost("localhost")
	config.SetPassword("guest")
	config.SetUserName("guest")
	config.SetPort("5672")
	config.SetVHost("/")

	a := kelinci.NewConnection("test", *config.Build())
	fmt.Println(a.Reconnect())
	b := kelinci.NewPublisher(a)

	c := kelinci.NewConsumer(a)
	message := "hi"
	m := amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	}

	fmt.Println(b.Publish("user.registered", "fanout", false, false, m))

	message = "ho"
	ma := amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	}
	fmt.Println(b.Publish("user.registered", "fanout", false, false, ma))

	c.StartConsuming("user.registered.queue", d)

}

func d(message amqp091.Delivery) {
	fmt.Println(string(message.Body[:]))
	message.Ack(false)
}

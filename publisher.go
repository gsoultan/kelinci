package kelinci

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	"time"
)

type Publisher interface {
	Publish(exchange string, key string, mandatory bool, immediate bool, msg amqp091.Publishing) error
}

type publisher struct {
	connection *Connection
}

func (p *publisher) Publish(exchange string, key string, mandatory bool, immediate bool, msg amqp091.Publishing) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := <-p.connection.err; err != nil {
		p.connection.Reconnect()
	}

	return p.connection.Channel.PublishWithContext(ctx, exchange, key, mandatory, immediate, msg)
}

func NewPublisher(connection *Connection) Publisher {
	a := new(publisher)
	a.connection = connection
	return a
}

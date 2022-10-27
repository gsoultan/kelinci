package kelinci

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
)

type Publisher interface {
	Publish(ctx context.Context, exchange string, key string, msg amqp091.Publishing) error
	Connection() *Connection
	SetConnection(connection *Connection)
	Mandatory() bool
	SetMandatory(mandatory bool)
	Immediate() bool
	SetImmediate(immediate bool)
	Message() amqp091.Publishing
	SetMessage(message amqp091.Publishing)
	Exchange() string
	SetExchange(exchange string)
	Key() string
	SetKey(key string)
}

type publisher struct {
	connection *Connection
	mandatory  bool
	immediate  bool
	message    amqp091.Publishing
	exchange   string
	key        string
}

func (p *publisher) Exchange() string {
	return p.exchange
}

func (p *publisher) SetExchange(exchange string) {
	p.exchange = exchange
}

func (p *publisher) Key() string {
	return p.key
}

func (p *publisher) SetKey(key string) {
	p.key = key
}

func (p *publisher) Connection() *Connection {
	return p.connection
}

func (p *publisher) SetConnection(connection *Connection) {
	p.connection = connection
}

func (p *publisher) Mandatory() bool {
	return p.mandatory
}

func (p *publisher) SetMandatory(mandatory bool) {
	p.mandatory = mandatory
}

func (p *publisher) Immediate() bool {
	return p.immediate
}

func (p *publisher) SetImmediate(immediate bool) {
	p.immediate = immediate
}

func (p *publisher) Message() amqp091.Publishing {
	return p.message
}

func (p *publisher) SetMessage(message amqp091.Publishing) {
	p.message = message
}

func (p *publisher) Publish(ctx context.Context, exchange string, key string, msg amqp091.Publishing) (err error) {
	if err := <-p.connection.err; err != nil {
		p.connection.Reconnect()
	}

	p.exchange = exchange
	p.key = key
	p.message = msg

	var channel *amqp091.Channel
	if channel, err = p.connection.Channel(); err != nil {
		return err
	}
	defer channel.Close()

	err = channel.PublishWithContext(ctx, p.exchange, p.key, p.mandatory, p.immediate, p.message)
	return err
}

func NewPublisher(connection *Connection) Publisher {
	a := new(publisher)
	a.connection = connection
	return a
}

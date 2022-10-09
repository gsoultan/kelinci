package kelinci

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
)

type Consumer interface {
	StartConsuming(queue string, d func(messages amqp091.Delivery))
	Connection() *Connection
	SetConnection(connection *Connection)
	Queue() string
	SetQueue(queue string)
	Consumer() string
	SetConsumer(consumer string)
	AutoAck() bool
	SetAutoAck(autoAck bool)
	Exclusive() bool
	SetExclusive(exclusive bool)
	NoLocal() bool
	SetNoLocal(noLocal bool)
	NoWait() bool
	SetNoWait(noWait bool)
	Args() amqp091.Table
	SetArgs(args amqp091.Table)
}

type consumer struct {
	connection *Connection
	queue      string
	consumer   string
	autoAck    bool
	exclusive  bool
	noLocal    bool
	noWait     bool
	args       amqp091.Table
}

func (c *consumer) Connection() *Connection {
	return c.connection
}

func (c *consumer) SetConnection(connection *Connection) {
	c.connection = connection
}

func (c *consumer) Queue() string {
	return c.queue
}

func (c *consumer) SetQueue(queue string) {
	c.queue = queue
}

func (c *consumer) Consumer() string {
	return c.consumer
}

func (c *consumer) SetConsumer(consumer string) {
	c.consumer = consumer
}

func (c *consumer) AutoAck() bool {
	return c.autoAck
}

func (c *consumer) SetAutoAck(autoAck bool) {
	c.autoAck = autoAck
}

func (c *consumer) Exclusive() bool {
	return c.exclusive
}

func (c *consumer) SetExclusive(exclusive bool) {
	c.exclusive = exclusive
}

func (c *consumer) NoLocal() bool {
	return c.noLocal
}

func (c *consumer) SetNoLocal(noLocal bool) {
	c.noLocal = noLocal
}

func (c *consumer) NoWait() bool {
	return c.noWait
}

func (c *consumer) SetNoWait(noWait bool) {
	c.noWait = noWait
}

func (c *consumer) Args() amqp091.Table {
	return c.args
}

func (c *consumer) SetArgs(args amqp091.Table) {
	c.args = args
}

func (c *consumer) StartConsuming(queue string, d func(messages amqp091.Delivery)) {
	if err := <-c.connection.err; err != nil {
		c.connection.Reconnect()
	}

	c.queue = queue
	messages, err := c.connection.Channel.Consume(c.queue, c.consumer, c.autoAck, c.exclusive, c.noLocal, c.noWait, c.args)
	if err != nil {
		fmt.Println("rabbitmq", "consumer", "queue", queue, "err", err)
	}
	c.handleMessages(messages, d)
}

func (c *consumer) handleMessages(messages <-chan amqp091.Delivery, d func(messages amqp091.Delivery)) {
	for m := range messages {
		go d(m)
	}
}

func NewConsumer(connection *Connection) Consumer {
	a := new(consumer)
	a.connection = connection
	return a
}

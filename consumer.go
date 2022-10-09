package kelinci

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
)

type Consumer interface {
	StartConsuming(queue string, d func(messages amqp091.Delivery))
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

	messages, err := c.connection.Channel.Consume(queue, c.consumer, c.autoAck, c.exclusive, c.noLocal, c.noWait, c.args)
	if err != nil {
		fmt.Println("rabbitmq", "consumer", "queue", queue, "err", err)
	}

	forever := make(chan bool)
	go func() {
		for m := range messages {
			d(m)
		}
	}()
	<-forever
}

func NewConsumer(connection *Connection) Consumer {
	a := new(consumer)
	a.connection = connection
	return a
}

package kelinci

import (
	"errors"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"time"
)

var (
	pool = make(map[string]*Connection)
)

type Connection struct {
	config        Config
	name          string
	conn          *amqp091.Connection
	err           chan error
	NotifyConfirm chan amqp091.Confirmation
	NotifyClose   chan *amqp091.Error
}

func (c *Connection) Channel() (*amqp091.Channel, error) {
	return c.conn.Channel()
}

func NewConnection(name string, config Config) *Connection {
	if c, ok := pool[name]; ok {
		return c
	}
	c := new(Connection)
	c.err = make(chan error)
	c.config = config
	pool[name] = c
	return c
}

func (c *Connection) GetConnection(name string) *Connection {
	return pool[name]
}

func (c *Connection) Connect() error {
	var err error
	if c.conn, err = amqp091.Dial(c.config.GetUri()); err != nil {
		return fmt.Errorf("error in creating rabbitmq connection with %s : %s", c.config.GetUri(), err.Error())
	}

	c.NotifyClose = make(chan *amqp091.Error)
	c.NotifyConfirm = make(chan amqp091.Confirmation)

	go func() {
		<-c.conn.NotifyClose(c.NotifyClose) //Listen to NotifyClose
		c.err <- errors.New("connection closed")
	}()

	return nil
}

func (c *Connection) OnClosedConnection() error {
	return <-c.err
}

func (c *Connection) Reconnect() error {
	c.Close()
	time.Sleep(5 * time.Second)
	if err := c.Connect(); err != nil {
		return err
	}
	return nil
}

func (c *Connection) Close() {
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}
	close(c.err)
}

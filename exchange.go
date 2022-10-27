package kelinci

import "github.com/rabbitmq/amqp091-go"

type Exchange interface {
	Declare(name string, kind string) error
	Conn() *Connection
	SetConn(conn *Connection)
	Name() string
	SetName(name string)
	Kind() string
	SetKind(kind string)
	Durable() bool
	SetDurable(durable bool)
	AutoDelete() bool
	SetAutoDelete(autoDelete bool)
	Internal() bool
	SetInternal(internal bool)
	NoWait() bool
	SetNoWait(noWait bool)
	Args() amqp091.Table
	SetArgs(args amqp091.Table)
}

type exchange struct {
	conn       *Connection
	name       string
	kind       string
	durable    bool
	autoDelete bool
	internal   bool
	noWait     bool
	args       amqp091.Table
}

func (e *exchange) Conn() *Connection {
	return e.conn
}

func (e *exchange) SetConn(conn *Connection) {
	e.conn = conn
}

func (e *exchange) Name() string {
	return e.name
}

func (e *exchange) SetName(name string) {
	e.name = name
}

func (e *exchange) Kind() string {
	return e.kind
}

func (e *exchange) SetKind(kind string) {
	e.kind = kind
}

func (e *exchange) Durable() bool {
	return e.durable
}

func (e *exchange) SetDurable(durable bool) {
	e.durable = durable
}

func (e *exchange) AutoDelete() bool {
	return e.autoDelete
}

func (e *exchange) SetAutoDelete(autoDelete bool) {
	e.autoDelete = autoDelete
}

func (e *exchange) Internal() bool {
	return e.internal
}

func (e *exchange) SetInternal(internal bool) {
	e.internal = internal
}

func (e *exchange) NoWait() bool {
	return e.noWait
}

func (e *exchange) SetNoWait(noWait bool) {
	e.noWait = noWait
}

func (e *exchange) Args() amqp091.Table {
	return e.args
}

func (e *exchange) SetArgs(args amqp091.Table) {
	e.args = args
}

func (e *exchange) Declare(name string, kind string) (err error) {
	e.name = name
	e.kind = kind

	var channel *amqp091.Channel
	if channel, err = e.conn.Channel(); err != nil {
		return err
	}

	return channel.ExchangeDeclare(e.name, e.kind, e.durable, e.autoDelete, e.internal, e.noWait, e.args)
}

func NewExchange(conn *Connection) Exchange {
	a := new(exchange)
	a.conn = conn
	a.durable = true
	a.kind = amqp091.ExchangeDirect
	return a
}

func NewExchangeWithName(conn *Connection, name string) Exchange {
	a := NewExchange(conn)
	a.Declare(name, a.Kind())
	return a
}

func NewExchangeWithParams(conn *Connection, name string, kind string) Exchange {
	a := NewExchange(conn)
	a.Declare(name, kind)
	return a
}

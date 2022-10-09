package kelinci

import "github.com/rabbitmq/amqp091-go"

type Queue interface {
	Bind(name string, key string, exchange string) error
	Connection() *Connection
	SetConnection(con *Connection)
	Name() string
	SetName(name string)
	Key() string
	SetKey(key string)
	Exchange() string
	SetExchange(exchange string)
	NoWait() bool
	SetNoWait(noWait bool)
	Args() amqp091.Table
	SetArgs(args amqp091.Table)
}

type queue struct {
	con      *Connection
	name     string
	key      string
	exchange string
	noWait   bool
	args     amqp091.Table
}

func (q *queue) Connection() *Connection {
	return q.con
}

func (q *queue) SetConnection(con *Connection) {
	q.con = con
}

func (q *queue) Name() string {
	return q.name
}

func (q *queue) SetName(name string) {
	q.name = name
}

func (q *queue) Key() string {
	return q.key
}

func (q *queue) SetKey(key string) {
	q.key = key
}

func (q *queue) Exchange() string {
	return q.exchange
}

func (q *queue) SetExchange(exchange string) {
	q.exchange = exchange
}

func (q *queue) NoWait() bool {
	return q.noWait
}

func (q *queue) SetNoWait(noWait bool) {
	q.noWait = noWait
}

func (q *queue) Args() amqp091.Table {
	return q.args
}

func (q *queue) SetArgs(args amqp091.Table) {
	q.args = args
}

func (q *queue) Bind(name string, key string, exchange string) error {
	q.name = name
	q.key = key
	q.exchange = exchange
	return q.con.Channel.QueueBind(q.name, q.key, q.exchange, q.noWait, q.args)
}

func NewQueueWithParams(connection *Connection, name string, key string, exchange string) Queue {
	a := NewQueue(connection)
	a.Bind(name, key, exchange)
	return a
}

func NewQueue(connection *Connection) Queue {
	a := new(queue)
	a.con = connection
	return a
}

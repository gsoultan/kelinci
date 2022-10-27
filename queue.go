package kelinci

import "github.com/rabbitmq/amqp091-go"

type Queue interface {
	Declare(name string) (amqp091.Queue, error)
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
	Durable() bool
	SetDurable(durable bool)
	AutoDelete() bool
	SetAutoDelete(autoDelete bool)
	Exclusive() bool
	SetExclusive(exclusive bool)
}

type queue struct {
	con        *Connection
	name       string
	key        string
	exchange   string
	noWait     bool
	args       amqp091.Table
	durable    bool
	autoDelete bool
	exclusive  bool
}

func (q *queue) Durable() bool {
	return q.durable
}

func (q *queue) SetDurable(durable bool) {
	q.durable = durable
}

func (q *queue) AutoDelete() bool {
	return q.autoDelete
}

func (q *queue) SetAutoDelete(autoDelete bool) {
	q.autoDelete = autoDelete
}

func (q *queue) Exclusive() bool {
	return q.exclusive
}

func (q *queue) SetExclusive(exclusive bool) {
	q.exclusive = exclusive
}

func (q *queue) Declare(name string) (amqp091.Queue, error) {
	q.name = name
	channel, _ := q.con.Channel()
	return channel.QueueDeclare(q.name, q.durable, q.autoDelete, q.exclusive, q.noWait, q.args)
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

func (q *queue) Bind(name string, key string, exchange string) (err error) {
	q.name = name
	q.key = key
	q.exchange = exchange

	var channel *amqp091.Channel
	if channel, err = q.con.Channel(); err != nil {
		return err
	}

	return channel.QueueBind(q.name, q.key, q.exchange, q.noWait, q.args)
}

func NewQueue(connection *Connection) Queue {
	a := new(queue)
	a.con = connection
	a.durable = true
	return a
}

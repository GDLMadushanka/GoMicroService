
rpackage event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn *amqp.Connection
	queueName string

}

func NewConsumer(conn *amqp.connection) (Consumer,error) {
	consumer := Consumer {
		conn:conn
	}
	err := consumer.setup()
	if err != nil {
		return Consumer{},err
	}
	return consumer, nil
}

func (consumer *Consumer) setup() error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}

	return declareExchange(channel)
}

type Payload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (consumer *Consumer) Listen(topics []string) error {
	ch,err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	q,err := declareRandomQueue(ch)
	if err != nil {
		return err
	}
	for _,s := range topics {
		ch.QueueBind(
			q.Name,
			s,
			"logs_topic",
			false,
			nil
		)
	}
}
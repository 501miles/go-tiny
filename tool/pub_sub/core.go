package pub_sub

import (
	"github.com/501miles/logger"
	jsoniter "github.com/json-iterator/go"
	"github.com/streadway/amqp"
	"go-tiny/tool/mq/rabbit"
)

const (
	ExchangeName = "Go-Sub-Pub-Exchange"
)

func Subscribe(topic string, dataChan chan interface{}) {
	ch := rabbit.GetChan()
	err := ch.ExchangeDeclare(
		ExchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		logger.Error(err)
	}

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)

	if err != nil {
		logger.Error(err)
	}

	err = ch.QueueBind(
		q.Name,
		topic,
		ExchangeName,
		false,
		nil,
	)
	if err != nil {
		logger.Error(err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Error(err)
	}

	logger.Info("aaaa")
	logger.Info(dataChan)
	for d := range msgs {
		logger.Info("收到消息d", d)
		dataChan <- d.Body
		logger.Info("aa")
	}
}

func Publish(topic string, data interface{}) error {
	ch := rabbit.GetChan()
	err := ch.ExchangeDeclare(
		ExchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Error(err)
		return err
	}

	body, _ := jsoniter.Marshal(data)
	err = ch.Publish(
		ExchangeName,
		topic,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
	return err
}

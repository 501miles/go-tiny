package rabbit

import (
	"github.com/streadway/amqp"
)

var rabbitChan *amqp.Channel

func InitRabbitMQ(ip, port, username, password string) error {
	conn, err := amqp.Dial("amqp://" + username + ":" + password + "@" + ip + ":" + port)
	if err != nil {
		return err
	}
	rabbitChan, err = conn.Channel()
	return err
}

func GetChan() *amqp.Channel {
	return rabbitChan
}

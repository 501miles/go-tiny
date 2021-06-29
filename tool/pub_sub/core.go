package pub_sub

import (
	"github.com/501miles/logger"
	jsoniter "github.com/json-iterator/go"
	"github.com/streadway/amqp"
	"github.com/tidwall/gjson"
	"go-tiny/tool/mq/rabbit"
	"reflect"
	"time"
)

const (
	ExchangeName = "Go-Sub-Pub-Exchange"
)

type ReSendMsg struct {
	RetryTimes uint8
	NextRetryTime int64
	Data interface{}
}

func Subscribe(topic string, dataChan chan interface{}, done chan struct{}, args ...interface{}) {
	go func() {
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
			true,
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
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			logger.Error(err)
		}

		for {
			select {
			case d := <- msgs:
				logger.Debug(jsoniter.MarshalToString(d))
				var result interface{}
				if d.Exchange == "" || len(d.Body) == 0 {
					continue
				}
				if len(args) > 0 {
					m := args[0]
					obj := reflect.New(reflect.TypeOf(m)).Interface()
					err = jsoniter.Unmarshal(d.Body, &obj)
					if err != nil {
						logger.Error(err)
					}
					result = reflect.ValueOf(obj).Elem().Interface()
				} else {
					result = d.Body
				}
				dataChan <- result
			case <- done:
				logger.Info("收到结束订阅消息")
				goto END
			}
		}

	END:
		logger.Info("订阅结束")
		close(dataChan)
		close(done)
	}()
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
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         body,
		},
	)
	if err != nil {
		logger.Error(err)
	}
	returnChan := make(chan amqp.Return)
	ch.NotifyReturn(returnChan)

	go func() {
		for {
			select {
			case d := <-returnChan:
				logger.Info("收到return消息")
				logger.Info(d.ReplyCode)
				logger.Info(d.ReplyText)
				logger.Info(d.Exchange)
				logger.Info(d.RoutingKey)
				logger.Info(d.Headers)
				logger.Info(string(d.Body))
				//dataChan <- d.Body
				logger.Info("aa")
			case <-time.After(10 * time.Second):
				//logger.Info("超时10秒")
				//close(returnChan)
				return
			}
		}
	}()
	return err
}

func covertDataToType2(data []byte, m interface{}) interface{} {
	return nil
}

func covertDataToType(data []byte, m interface{}) interface{} {
	mType := reflect.TypeOf(m)
	logger.Warn(mType.Kind())
	switch mType.Kind() {
	//case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64,:

	case reflect.Ptr:
		logger.Debug("是指针类型")
		obj := reflect.New(mType.Elem()).Interface()
		err := jsoniter.Unmarshal(data, &obj)
		if err != nil {
			logger.Error(err)
		}
		return obj
	case reflect.Struct:
		logger.Debug("是结构体类型")
		obj := reflect.New(mType).Interface()
		err := jsoniter.Unmarshal(data, &obj)
		if err != nil {
			logger.Error(err)
		}
		return reflect.ValueOf(obj).Elem().Interface()
	case reflect.Array, reflect.Slice:
		logger.Debug("是数组\\切片类型")
		typ := mType.Elem()
		obj := reflect.New(reflect.SliceOf(typ).Elem()).Interface()
		var resp []interface{}
		arr := gjson.ParseBytes(data).Array()
		for _, it := range arr {
			logger.Info(it.String())
			resp = append(resp, covertDataToType([]byte(it.String()), obj))
		}
		return resp
	case reflect.Map:
		logger.Debug("是字典类型")
		var resp map[string]interface{}
		err := jsoniter.Unmarshal(data, &resp)
		if err != nil {
			logger.Error(err)
		}
		return resp

	default:
		obj := reflect.New(mType).Interface()
		err := jsoniter.Unmarshal(data, &obj)
		if err != nil {
			logger.Error(err)
		}
		return reflect.ValueOf(obj).Elem().Interface()
	}
}

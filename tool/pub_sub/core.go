package pub_sub

import (
	"fmt"
	"github.com/501miles/go-tiny/tool/mq/rabbit"
	"github.com/501miles/logger"
	"github.com/Jeffail/tunny"
	jsoniter "github.com/json-iterator/go"
	"github.com/streadway/amqp"
	"github.com/tidwall/gjson"
	"reflect"
	"runtime"
	"sync"
	"time"
)

const (
	ExchangeName       = "Go-Sub-Pub-Exchange"
	ExchangeNameBackup = "Go-Sub-Pub-Exchange-Backup"
)

//var topicThreadPool sync.Map
//var pool1 *tunny.Pool

var pubPoolDict map[string]*tunny.Pool
var lock sync.RWMutex

func init() {
	//numCPUs := runtime.NumCPU()
	//pool1 = tunny.NewFunc(numCPUs, func(payload interface{}) interface{} {
	//	pubProcess("test", payload)
	//	return nil
	//})
	//pool1.SetSize(5) // 100 goroutines
	pubPoolDict = make(map[string]*tunny.Pool)
	go func() {
		time.Sleep(2 * time.Second)

		ch := rabbit.GetChan()
		err := ch.ExchangeDeclare(
			ExchangeNameBackup,
			"fanout",
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
			"",
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

		for d := range msgs {
			logger.Info(d)
		}
	}()
}

type ReSendMsg struct {
	RetryTimes    uint8
	NextRetryTime int64
	Data          interface{}
}

type OnMessageReceive func(data interface{}) bool

func Subscribe(topic string, onMessageReceive OnMessageReceive, args ...interface{}) <-chan struct{} {
	var done chan struct{}
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
			case d := <-msgs:
				var result interface{}
				if d.Exchange == "" || len(d.Body) == 0 {
					logger.Info("AAAABBBCCCDDDD")
					logger.Debug(jsoniter.MarshalToString(d))
					//if time.Now().Unix() % 7 == 0 {
					//	time.Sleep(10000 * time.Second)
					//}
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
				go func(r2 interface{}, d2 amqp.Delivery) {
					ok := onMessageReceive(r2)
					if ok {
						tryCount := 0
					RetryAck:
						err = d2.Acknowledger.Ack(d.DeliveryTag, false)
						if err != nil {
							logger.Error(err)
							tryCount++
							if tryCount < 3 {
								goto RetryAck
							}
						}
					} else {
						tryCount := 0
					RetryReject:
						err = d2.Acknowledger.Reject(d2.DeliveryTag, false)
						if err != nil {
							logger.Error(err)
							tryCount++
							if tryCount < 3 {
								goto RetryReject
							}
						}
					}
				}(result, d)

			case <-done:
				logger.Info("收到结束订阅消息")
				goto END
			}
		}

	END:
		logger.Info("订阅结束")
		close(done)
		return
	}()
	return done
}

func Publish(topic string, data interface{}) error {
	go func() {
		lock.RLock()
		defer lock.RUnlock()
		pool := getThreadPool(topic)
		logger.Error(fmt.Sprintf("%p", pool))
		pool.Process(data)
	}()
	return nil
}

func getThreadPool(topic string) *tunny.Pool {
	p, ok := pubPoolDict[topic]
	if ok {
		return p
	}
	lock.RUnlock()
	lock.Lock()
	defer func() {
		lock.Unlock()
		lock.RLock()
	}()
	numCPUs := runtime.NumCPU()
	logger.Info(numCPUs)
	pool := tunny.NewFunc(numCPUs, func(payload interface{}) interface{} {
		pubProcess(topic, payload)
		return nil
	})
	pool.SetSize(100)
	pubPoolDict[topic] = pool
	return pool
}

func pubProcess(topic string, data interface{}) {
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
	time.Sleep(1 * time.Second)
	//go func() {
	//	for {
	//		select {
	//		case d := <-returnChan:
	//			logger.Info("收到return消息")
	//			logger.Info(d.ReplyCode)
	//			logger.Info(d.ReplyText)
	//			logger.Info(d.Exchange)
	//			logger.Info(d.RoutingKey)
	//			logger.Info(d.Headers)
	//			logger.Info(string(d.Body))
	//			//dataChan <- d.Body
	//			logger.Info("aa")
	//		case <-time.After(10 * time.Second):
	//			//logger.Info("超时10秒")
	//			//close(returnChan)
	//			return
	//		}
	//	}
	//}()
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

package receiver

import (
	"encoding/json"
	"fmt"

	C "github.com/mt1976/rmg/config"
	E "github.com/mt1976/rmg/errors"
	L "github.com/mt1976/rmg/language"
	M "github.com/mt1976/rmg/model"
	"github.com/streadway/amqp"
)

var config = C.Configuration

func Run() error {
	user := config.MQUser
	password := config.MQPassword
	host := config.MQHost
	port := config.MQPort
	//target := "amqp://%v:%v@%v:%v/"
	targetAddress := fmt.Sprintf(config.MQAddressFormat, user, password, host, port)

	//fmt.Println("Go RabbitMQ Tutorial")
	conn, err := amqp.Dial(targetAddress)
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	defer conn.Close()

	fmt.Println(L.TxtConnectedToMQ)

	// Let's start by opening a channel to our RabbitMQ instance
	// over the connection we have already established
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(E.ErrError, err)
		panic(err.Error())
	}
	defer ch.Close()
	fmt.Println(L.TxtMQChannelOpen)
	fmt.Printf(L.TxtMQConnectToQueue, config.MQQueue)

	//fmt.Printf("config: %v\n", config)

	// declaring consumer with its properties over channel opened
	msgs, err := ch.Consume(
		config.MQQueue, // queue
		"",             // consumer
		true,           // auto ack
		false,          // exclusive
		false,          // no local
		false,          // no wait
		nil,            //args
	)
	if err != nil {
		fmt.Println(E.ErrError, err)
		panic(err)
	}

	// print consumed messages from queue
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			//fmt.Printf("Received Message: %s\n", msg.Body)
			if msg.ContentType != config.MQContentType {
				fmt.Printf(E.ErrInvalidContentType, msg.ContentType)
				return
			}
			rate, err := Unmarshal(msg.Body)
			if err != nil {
				fmt.Printf(E.ErrUnmarshalJSON, err)
				return
			}
			fmt.Printf(L.TxtRateMessageDebug, rate.GetID(), rate.GetBid(), rate.GetAsk())
		}
	}()
	fmt.Println(L.TxtWaitingForMessages)
	<-forever
	return nil
}

// Unmarshal unmarshals a JSON-encoded byte slice into a Rate struct
func Unmarshal(msg []byte) (M.Rate, error) {
	var rateMsg M.Rate
	err := json.Unmarshal(msg, &rateMsg)
	if err != nil {
		return M.Rate{}, err
	}
	return rateMsg, nil
}

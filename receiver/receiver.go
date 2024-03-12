package receiver

import (
	"fmt"

	C "github.com/mt1976/rmg/config"
	L "github.com/mt1976/rmg/language"
	"github.com/streadway/amqp"
)

var config = C.Configuration

func Run() error {
	user := config.MQUser
	password := config.MQPassword
	host := config.MQHost
	port := config.MQPort
	target := "amqp://%v:%v@%v:%v/"
	targetAddress := fmt.Sprintf(target, user, password, host, port)

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
		fmt.Println(err)
	}
	defer ch.Close()
	fmt.Println(L.TxtMQChannelOpen)
	fmt.Printf(L.TxtMQConnectToQueue, config.MQQueue)

	fmt.Printf("config: %v\n", config)

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
		panic(err)
	}

	// print consumed messages from queue
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			fmt.Printf("Received Message: %s\n", msg.Body)
		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever

	return nil
}

package sender

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
	C "github.com/mt1976/rmg/config"
	E "github.com/mt1976/rmg/errors"
	M "github.com/mt1976/rmg/model"
	"github.com/streadway/amqp"
)

var config = C.Configuration
var Types map[int]string

func Run() error {
	user := "guest"
	password := "guest"
	host := "localhost"
	port := 5672
	target := "amqp://%v:%v@%v:%v/"
	targetAddress := fmt.Sprintf(target, user, password, host, port)

	fmt.Println("Go RabbitMQ Tutorial")
	conn, err := amqp.Dial(targetAddress)
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	defer conn.Close()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")

	// Let's start by opening a channel to our RabbitMQ instance
	// over the connection we have already established
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()
	fmt.Printf("\"Channel Open\": %v\n", "Channel Open")

	// with this channel open, we can then start to interact
	// with the instance and declare Queues that we can publish and
	// subscribe to
	q, err := ch.QueueDeclare(
		config.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)

	// We can print out the status of our Queue here
	// this will information like the amount of messages on
	// the queue
	fmt.Println(q)
	// Handle any errors if we were unable to create the queue
	if err != nil {
		fmt.Println(err)
	}

	Types = make(map[int]string)
	Types[0] = "Unknown"
	Types[1] = "Reuters"
	Types[2] = "Bloomberg"
	Types[3] = "360T"

	file := "./data/" + config.SimulatorSourceFile
	f, err := os.Open(file)
	if err != nil {
		panic(E.ErrCannotOpenFile + err.Error())
		//return err
	}
	defer f.Close()

	//
	r := csv.NewReader(f)
	records, _ := r.ReadAll()
	//	fmt.Printf("records: %v\n", records)
	fmt.Printf("len(records): %v\n", len(records))
	// o := NewScanner(strings.NewReader(records))
	// for o.Scan() {
	// 	println(o.Text("Month"), o.Text("Day"))
	// }
	for recNo, rec := range records {
		//fmt.Printf("rec %v: %v %v\n", recNo, rec, len(rec))
		if recNo != 0 {

			process(ch, recNo, rec)
		}
		if recNo >= 20 {
			continue
		}
	}

	return nil

}

func process(ch *amqp.Channel, recNo int, rec []string) {

	asset := rec[M.TYPE] // Asset Class
	if asset != "FX" {
		return
	}
	rateType := rec[M.ASSET_CLASS]           // Rate Type]
	source, _ := strconv.Atoi(rec[M.SOURCE]) // Source
	sourceName := Types[source]
	rateID := rec[M.BASE_CCY] + rec[M.QUOTE_CCY]
	if rateType == "FXFWD" {
		rateID = rateID + rec[M.TENOR]
	}
	rateID = rateID + "=" // Comment]

	// byte to string conversion
	//instStr := fmt.Sprintf("%c", asset)
	// fmt.Printf("instStr: %v\n", asset)
	// fmt.Printf("source: %v\n", source)
	// fmt.Printf("sourceName: %v\n", sourceName)
	// fmt.Printf("rateType: %v\n", rateType)
	// fmt.Printf("rateID: %v\n", rateID)

	var x M.Rt
	x.SetCat(sourceName)
	x.SetSrc(sourceName)
	x.SetID(rateID)
	x.SetBid(rec[M.BID])
	x.SetAsk(rec[M.OFFER])
	x.SetOwn(rec[M.OWNER])
	x.SetRsk(rec[M.RISK_CENTRE])
	x.SetSts("OK")
	x.SetDTme(NowToDateTime(time.Now()))
	//spew.Dump(x)
	// var col M.Coll
	// col.Rt = append(col.Rt, x)
	// var msg M.Msg
	// msg.SetXsiNoNamespaceSchemaLocation("eurobase-rate.xsd")
	// msg.SetXmlnsXsi("http://www.w3.org/2001/XMLSchema-instance")
	// msg.Coll = col
	spew.Dump(x)
	xx := fmt.Sprintf("x: %v\n", x)

	// attempt to publish a message to the queue!
	err := ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(xx),
		},
	)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Published Message to Queue")

	//os.Exit(1)
}

func NowToDateTime(now time.Time) string {
	//2012-11-28T10:10:10
	return now.Format("2006-01-02T15:04:05")
}

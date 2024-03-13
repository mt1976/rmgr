package sender

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	C "github.com/mt1976/rmg/config"
	E "github.com/mt1976/rmg/errors"
	L "github.com/mt1976/rmg/language"
	M "github.com/mt1976/rmg/model"
	S "github.com/mt1976/rmg/support"
	"github.com/streadway/amqp"
)

var config = C.Configuration
var Types map[int]string

func Run() error {
	user := config.MQUser
	password := config.MQPassword
	host := config.MQHost
	port := config.MQPort
	//target := "amqp://%v:%v@%v:%v/"
	targetAddress := fmt.Sprintf(config.MQAddressFormat, user, password, host, port)

	//log.Println("Go RabbitMQ Tutorial")
	conn, err := amqp.Dial(targetAddress)
	if err != nil {
		log.Panicln(E.ErrError, err)
	}
	defer conn.Close()

	log.Println(L.TxtConnectedToMQ)

	// Let's start by opening a channel to our RabbitMQ instance
	// over the connection we have already established
	ch, err := conn.Channel()
	if err != nil {
		log.Panicln(E.ErrError, err)
	}
	defer ch.Close()
	log.Println(L.TxtMQChannelOpen)
	log.Printf(L.TxtMQConnectToQueue, config.MQQueue)

	// with this channel open, we can then start to interact
	// with the instance and declare Queues that we can publish and
	// subscribe to
	q, err := ch.QueueDeclare(
		config.MQQueue,
		false,
		false,
		false,
		false,
		nil,
	)

	// We can print out the status of our Queue here
	// this will information like the amount of messages on
	// the queue
	log.Println(L.TxtMQQueueConnected)
	log.Println(L.TxtMQName, q.Name)
	log.Println(L.TxtMQConsumers, q.Consumers)
	log.Println(L.TxtMQMessages, q.Messages)
	// Handle any errors if we were unable to create the queue
	if err != nil {
		log.Panicln(E.ErrError, err)
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
	//	log.Printf("records: %v\n", records)
	//log.Printf("len(records): %v\n", len(records))
	// o := NewScanner(strings.NewReader(records))
	// for o.Scan() {
	// 	println(o.Text("Month"), o.Text("Day"))
	// }
	for recNo, rec := range records {
		//log.Printf("rec %v: %v %v\n", recNo, rec, len(rec))
		if recNo != 0 {
			publishRateMessage(ch, rec)
		}
		// if recNo >= 20 {
		// 	continue
		// }
	}
	return nil
}

func publishRateMessage(ch *amqp.Channel, rec []string) {

	//asset := rec[M.TYPE] // Asset Class
	//if asset != "FX" {
	//	return
	//}
	rateType := rec[M.ASSET_CLASS] // Rate Type
	rateForm := C.Configuration.InterestType
	if rateType[0:2] == C.Configuration.FXText {
		rateForm = C.Configuration.RateText
	}
	source, _ := strconv.Atoi(rec[M.SOURCE]) // Source
	sourceName := Types[source]
	rateID := rec[M.BASE_CCY] + rec[M.QUOTE_CCY]
	rateID = rateID + rec[M.TENOR]
	rateID = rateID + L.TxtEquals // Comment]

	routingKey := BuildRoutingKey(rec[M.TYPE], rec[M.BASE_CCY], rec[M.QUOTE_CCY], rec[M.TENOR])
	log.Printf(L.TxtRoutingKey, routingKey)
	// byte to string conversion
	//instStr := log.Sprintf("%c", asset)
	// log.Printf("instStr: %v\n", asset)
	// log.Printf("source: %v\n", source)
	// log.Printf("sourceName: %v\n", sourceName)
	// log.Printf("rateType: %v\n", rateType)
	// log.Printf("rateID: %v\n", rateID)

	var x M.Rate
	x.SetCat(rateForm)
	x.SetSrc(sourceName)
	x.SetID(rateID)
	x.SetBid(rec[M.BID])
	x.SetAsk(rec[M.OFFER])
	x.SetOwn(rec[M.OWNER])
	x.SetRsk(rec[M.RISK_CENTRE])
	x.SetSts(config.DefaultStatus)
	x.SetDTme(NowToDateTime(time.Now()))
	x.SetStaleAfter(config.StaleAfterMS)
	x.SetUUID(S.NewUUID())

	// Marshal the struct into a JSON string
	json, err := json.Marshal(x)
	if err != nil {
		log.Println(E.ErrError, err)
	}
	log.Println(x, string(json))
	//spew.Dump(x)
	// var col M.Coll
	// col.Rt = append(col.Rt, x)
	// var msg M.Msg
	// msg.SetXsiNoNamespaceSchemaLocation("eurobase-rate.xsd")
	// msg.SetXmlnsXsi("http://www.w3.org/2001/XMLSchema-instance")
	// msg.Coll = col
	//spew.Dump(x)
	//xx := log.Sprintf("json: %v\n", json)

	// attempt to publish a message to the queue!
	publishErr := ch.Publish(
		config.MQExchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: config.MQContentType,
			Body:        json,
			MessageId:   x.GetUUID(),
		},
	)

	if publishErr != nil {
		log.Println(E.ErrError, err)
	}
	log.Printf(L.TxtMQMessagePublised, NowToDateTime(time.Now()))

	//os.Exit(1)
}

func BuildRoutingKey(assetType, baseCCY, quoteCCY, tenor string) string {
	routingKeyFormat := C.Configuration.RoutingKeyFormat
	routingKey := fmt.Sprintf(routingKeyFormat, assetType, baseCCY, quoteCCY, tenor)
	routingKey = strings.ToLower(routingKey)
	return routingKey
}

// NowToDateTime returns a formatted date and time string
func NowToDateTime(now time.Time) string {
	return now.Format(config.DateTimeFormat)
}

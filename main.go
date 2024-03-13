package main

import (
	"flag"
	"log"
	"strings"

	C "github.com/mt1976/rmg/config"
	E "github.com/mt1976/rmg/errors"
	L "github.com/mt1976/rmg/language"
	R "github.com/mt1976/rmg/receiver"
	S "github.com/mt1976/rmg/sender"
)

var config = C.Configuration

func main() {

	log.Println(config.ApplicationName)
	log.Println(strings.Repeat("=", len(config.ApplicationName)))

	var sender bool
	var receiver bool

	flag.BoolVar(&sender, L.TxtFlagSend, false, L.TxtUsageSendMode)
	flag.BoolVar(&receiver, L.TxtFlagRecv, false, L.TxtUsageRecvMode)

	flag.Parse()

	if sender && receiver {
		log.Printf(E.ErrCannotBeSenderAndReceiver)
		return
	}
	if !sender && !receiver {
		log.Printf(E.ErrCannotBeSenderAndReceiver)
		return
	}

	if sender {
		log.Println(L.TxtMode, L.TxtModeSending)
	} else {
		log.Println(L.TxtMode, L.TxtModeReceiving)
	}
	// Get the configuration

	port := config.Port
	target := config.Target
	log.Println(L.TxtTarget, target)
	log.Println(L.TxtPort, port)

	//targetAddress := log.Sprintf("%s:%d", target, port)
	//log.Println(L.TxtAddress, targetAddress)

	switch {
	case sender:
		//log.Println(L.TxtMode, L.TxtModeSending)
		send()
	case receiver:
		//log.Println(L.TxtMode, L.TxtModeReceiving)
		receive()
	default:
		log.Println(E.ErrInvalidRole)
	}

}

func send() {
	//log.Println(address)
	//log.Println(C.Configuration)
	S.Run()
}

func receive() {
	//log.Println(address)
	//log.Println(C.Configuration)
	R.Run()
}

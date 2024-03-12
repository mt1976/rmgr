package main

import (
	"flag"
	"fmt"
	"strings"

	C "github.com/mt1976/rmg/config"
	E "github.com/mt1976/rmg/errors"
	L "github.com/mt1976/rmg/language"
	R "github.com/mt1976/rmg/receiver"
	S "github.com/mt1976/rmg/sender"
)

var config = C.Configuration

func main() {

	fmt.Println(config.ApplicationName)
	fmt.Println(strings.Repeat("=", len(config.ApplicationName)))

	var sender bool
	var receiver bool

	flag.BoolVar(&sender, L.TxtFlagSend, false, L.TxtUsageSendMode)
	flag.BoolVar(&receiver, L.TxtFlagRecv, false, L.TxtUsageRecvMode)

	flag.Parse()

	if sender && receiver {
		fmt.Printf(E.ErrCannotBeSenderAndReceiver)
		return
	}
	if !sender && !receiver {
		fmt.Printf(E.ErrCannotBeSenderAndReceiver)
		return
	}

	if sender {
		fmt.Println(L.TxtMode, L.TxtModeSending)
	} else {
		fmt.Println(L.TxtMode, L.TxtModeReceiving)
	}
	// Get the configuration

	port := config.Port
	target := config.Target
	fmt.Println(L.TxtTarget, target)
	fmt.Println(L.TxtPort, port)

	//targetAddress := fmt.Sprintf("%s:%d", target, port)
	//fmt.Println(L.TxtAddress, targetAddress)

	switch {
	case sender:
		//fmt.Println(L.TxtMode, L.TxtModeSending)
		send()
	case receiver:
		//fmt.Println(L.TxtMode, L.TxtModeReceiving)
		receive()
	default:
		fmt.Println(E.ErrInvalidRole)
	}

}

func send() {
	//fmt.Println(address)
	//fmt.Println(C.Configuration)
	S.Run()
}

func receive() {
	//fmt.Println(address)
	//fmt.Println(C.Configuration)
	R.Run()
}

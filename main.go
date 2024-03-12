package main

import (
	"fmt"

	C "github.com/mt1976/rmg/config"
	E "github.com/mt1976/rmg/errors"
	L "github.com/mt1976/rmg/language"
	S "github.com/mt1976/rmg/sender"
)

var config = C.Configuration

func main() {

	// Get the configuration

	port := config.Port
	target := config.Target
	fmt.Println(config.ApplicationName)
	fmt.Println(L.TxtTarget, target)
	fmt.Println(L.TxtPort, port)

	targetAddress := fmt.Sprintf("%s:%d", target, port)
	fmt.Println(L.TxtAddress, targetAddress)

	switch config.Role {
	case C.SEND:
		fmt.Println(L.TxtMode, L.TxtModeSending)
		send(targetAddress)
	case C.RECEIVE:
		fmt.Println(L.TxtMode, L.TxtModeReceiving)
		//receive(targetAddress)
	default:
		fmt.Println(E.ErrInvalidRole)
	}

}

func send(address string) {
	fmt.Println(address)
	fmt.Println(C.Configuration)
	S.Run()
}

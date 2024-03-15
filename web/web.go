package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\"Handler\": %v\n", "homeHandler")
	t, _ := template.ParseFiles("./web/home.html")
	fmt.Printf("\"loaded Template\": %v\n", "loaded Template")
	fmt.Printf("t: %v\n", t)
	t.Execute(w, nil)
	fmt.Printf("\"Executed Template\": %v\n", "Executed Template")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\"Handler\": %v\n", "wsEndpoint")

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Connected")

	defer ws.Close()
	ws.WriteMessage(1, []byte("Hello Client!"))

	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// Assuming the incoming message is the new price
		log.Printf("Received message: %s\n", p)

		if err := ws.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

func setupRoutes() {
	fmt.Printf("\"setup routes\": %v\n", "setup routes")
	fmt.Printf("\"home route\": %v\n", "home route")
	http.HandleFunc("/", homeHandler)
	fmt.Printf("\"ws route\": %v\n", "ws route")
	http.HandleFunc("/ws", wsEndpoint)
	fmt.Printf("\"done\": %v\n", "done")
}

func Run() {
	fmt.Printf("\"setupRoutes\": %v\n", "setupRoutes")
	setupRoutes()
	fmt.Printf("\"http.ListenAndServe\": %v\n", "http.ListenAndServe")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"fmt"
)

// our host address's flag
var addr = flag.String("addr", "0.0.0.0:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options


// our echo function
func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()

		// websocket connection error: (close connection on client side)
		if err != nil {
			log.Println("read error:", err)
			break
		}

		// received a message
		log.Printf("received: %s", message)
		// send the message to the client
		err = conn.WriteMessage(mt, message)

		// error when writing message: (
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func getUserRole(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer conn.Close()

	for {
		asker, message, err := conn.ReadMessage()
		if err  != nil {
			log.Println("read error: ", err)
			break
		}

		if message != nil {
			fmt.Printf("user to get role: %s\n", message)
		}

		err = conn.WriteMessage(asker, message)
		if err != nil {
			break
		}
	}

}

func main() {
	flag.Parse()  // ->
	log.SetFlags(0)

	http.HandleFunc("/echo", echo)
	http.HandleFunc("/getUserRole", getUserRole)

	log.Fatal(http.ListenAndServe(*addr, nil))
}
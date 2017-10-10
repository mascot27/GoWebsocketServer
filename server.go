package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"fmt"
	//"WebservicesLogic"
)

// our host address's flag
var addr = flag.String("addr", "0.0.0.0:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options


type MessageFormatFromStandardClient struct {
	LoginMethod string
	UserName string
	Password string
	FunctionName string
	arguments []string
}


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

		// useless at the moment
		parts := ParseMessage(message)
		fmt.Println(parts)

		// send the message to the client
		err = conn.WriteMessage(mt, message)

		// error when writing message: (
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}


func main() {
	fmt.Println("server started....")
	flag.Parse()  // ->
	log.SetFlags(0)

	//lematin := WebservicesLogic.SayHello2()

	//fmt.Print(lematin)
	http.HandleFunc("/echo", echo)

	log.Fatal(http.ListenAndServe(*addr, nil))
}
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

		// received a message
		log.Printf("received: %s", message)
		isValidFormat := validateRequestValidFormat(message)
		if isValidFormat {
			fmt.Println("valid format")
		}

		// send the message to the client
		err = conn.WriteMessage(mt, message)

		// error when writing message: (
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}


func validateRequestValidFormat(message []byte)(bool){
	// check length not nil
	// message get part

	/* message composition:

	[loginMethod][UserName|Password][functionName][argumentsArray]

	arguments array is a string array
	this implie that we need to check format when sending
	thus, a username/functionName/password/argument cannot contain a '[', a ']' or a '|'

 	var my_message MessageFormatFromStandardClient
	 */

	// validate format
	opening_counter := 0
	closing_counter := 0

	for i:= 0; i < len(message) ; i++  {
		// count '[' and ']' --> this need to be the same number to be valid
		if message[i] == '[' {
			opening_counter++
		} else if message[i] == ']'{
			closing_counter++
		}
	}

	if closing_counter != opening_counter {
		return false
	}

	if opening_counter != 4 {
		return false
	}

	return true
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
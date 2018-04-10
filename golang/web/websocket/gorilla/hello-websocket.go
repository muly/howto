// demonstrate websocket program in go, and test using any websocket client
//
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		fmt.Println("websocket upgrade error:", err)
	}
	//TODO: failed at upgrade step (with below error) when ran with http://localhost:8080/ instead of ws://localhost:8080
	// "websocket: the client is not using the websocket protocol"

	//defer conn.Close() //TODO: isn't this required

	err = conn.WriteMessage(websocket.TextMessage, []byte("Hello from Websocket! "+time.Now().String()))
	if err != nil {
		fmt.Println("websocket write error:", err)
	}

}

// to run:
// use any websocket client tool and hit the below url
// 		ws://localhost:8080

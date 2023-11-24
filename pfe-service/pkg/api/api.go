package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pfe-service/config"
	"github.com/pfe-service/pkg/api/handlers"
)

func Init() {
	// http hello world example
	http.HandleFunc("/status", handlers.HandlePingStatus)
	http.HandleFunc("/save", handlers.HandleSavePart)
	http.HandleFunc("/retrieve", handlers.HandleGetSecret)
	configuration := config.GetConfig()

	fmt.Println("Service " + configuration.Name + " started on " + configuration.Host)
	err := http.ListenAndServe(configuration.Host, nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
		return
	}
}

func InitStatusWS() {
	// Dialer for WebSocket connection
	dialer := websocket.Dialer{}

	// Connect to WebSocket server
	conn, _, err := dialer.Dial("ws://localhost:3000/ws/v1/services/status/"+config.GetConfig().Name, nil)
	for err != nil {
		fmt.Println("Error connecting to the manager, retrying in 5 seconds")
		time.Sleep(5 * time.Second)
		conn, _, err = dialer.Dial("ws://localhost:3000/ws/v1/services/status/"+config.GetConfig().Name, nil)
	}

	defer conn.Close()

	for {
		_, _, err = conn.ReadMessage()
		if err != nil {
			break
		}
	}
}

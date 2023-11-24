package websockets

import (
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"github.com/pfe-manager/pkg/servicesManager"
)

// store websocket connections for each service with their name as key
var servicesStatusWS = make(map[string][]*websocket.Conn)

func ServiceStatusWS(c *websocket.Conn) {
	serviceName := c.Params("serviceName")
	fmt.Println("Connection opened for service " + serviceName)
	// register connection
	servicesStatusWS[serviceName] = append(servicesStatusWS[serviceName], c)
	defer func() {
		fmt.Println("Connection closed for service " + serviceName)
		delete(servicesStatusWS, serviceName)
		servicesManager.UpdateServicesStatus()
	}()
	servicesManager.UpdateServicesStatus()

	// wait until connection is closed
	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			break
		}
	}

}

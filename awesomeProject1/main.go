package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
}

func main() {
	server := gin.Default()

	server.GET("/ws", WSController)
	server.GET("/home", home)
	server.GET("/about", about)

	server.Run(":8080")
}

func about(c *gin.Context) {
	c.String(200, "About")
}

func home(c *gin.Context) {
	//homeTemplate.Execute(c.Writer, "ws://"+c.Request.Host+"/home")
	c.JSON(http.StatusOK, "не знаю")
}

func WSController(c *gin.Context) {
	return
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil); if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	counter := 0

	for {
		if counter > 10 {
			break
		}

		conn.WriteJSON(map[string]interface{}{
			"name": "Egor",
		})
		counter++
		time.Sleep(time.Microsecond * 100)
	}

	conn.ReadMessage()
}
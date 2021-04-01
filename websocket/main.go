package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"time"
)

var upgrader = websocket.Upgrader{}

func main() {
	router := gin.Default()

	router.GET("/", home)
	router.GET("/ws1", ws1)
	router.GET("/ws2", ws2)
	router.GET("/ws3", ws3)

	router.Run(":8080")
}

func ws1(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil); if err != nil {
		c.String(500, err.Error())
		return
	}

	go func(conn *websocket.Conn) {
		for {
			mType, msg, err := conn.ReadMessage(); if err != nil {
				fmt.Println(err.Error())
				return
			}

			err = conn.WriteMessage(mType, msg); if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}(conn)
}

func ws2(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil); if err != nil {
		c.String(500, err.Error())
		return
	}

	go func(conn *websocket.Conn) {
		for {
			_, msg, _ := conn.ReadMessage()
			println(string(msg))
		}
	}(conn)
}

func ws3(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil); if err != nil {
		c.String(500, err.Error())
		return
	}

	go func(conn *websocket.Conn) {
		for {
			ch := time.Tick(5 * time.Second)

			var value time.Time

			for range ch {
				value = <- ch

				conn.WriteMessage(websocket.TextMessage, []byte(value.String()))
			}
		}
	}(conn)
}

func home(c *gin.Context) {
	c.File("index.html")
}
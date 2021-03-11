package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	router := router()
	router.Start(":8000")
}

func router() *echo.Echo {
	router := echo.New()

	// Pretty Debug Logger
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "[ECHO] ${time_custom} | ${status} | ${latency_human} | ${remote_ip} | ${method} | ${path} |\n",
		CustomTimeFormat: "2006/01/02 - 15:04:00",
	}))

	router.GET("/test", test)

	return router
}

func test(c echo.Context) error {
	return c.String(http.StatusOK, "Hello!")
}

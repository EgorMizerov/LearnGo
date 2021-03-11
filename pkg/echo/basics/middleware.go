package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	router := router2()
	router.Start(":8000")
}

func router2() *echo.Echo {
	router := echo.New()
	// Root level middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.GET("/middleware", testMiddleware)

	return router
}

func testMiddleware(c echo.Context) error {
	return c.String(http.StatusOK, "Тест для middleware")
}

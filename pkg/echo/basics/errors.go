package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	router := router5()
	router.Start(":8000")
}

func router5() *echo.Echo {
	router := echo.New()

	router.GET("/def", defError)

	return router
}

func defError(c echo.Context) error {
	return echo.NewHTTPError(http.StatusBadRequest, "адресс недоступен")
}

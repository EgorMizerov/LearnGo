package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"os"
)

func main() {
	router := router6()
	router.Start(":8000")
}

func router6() *echo.Echo {
	router := echo.New()

	router.Use(middleware.Logger())

	router.GET("/getPhoto", getPhoto)
	router.GET("/getPhoto2", getPhoto2)
	router.GET("/redirect", redirect2)
	router.GET("/hello", hello)

	return router
}

func getPhoto(c echo.Context) error {
	return c.File("./unnamed.jpg")
}

func getPhoto2(c echo.Context) error {
	f, err := os.Open("./unnamed.jpg")
	if err != nil {
		return err
	}
	return c.Stream(http.StatusOK, "image/jpg", f)
}

func redirect2(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "hello")
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello!")
}

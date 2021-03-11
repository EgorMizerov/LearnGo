package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type User struct {
	ID string `path:"id" query:"id" form:"id" json:"id" xml:"id"`
}

func main() {
	router := router3()
	router.Start(":8000")
}

func router3() *echo.Echo {
	router := echo.New()

	router.GET("/query", bindQuery)
	router.POST("/form", bindForm)
	router.POST("/json", bindJSON)
	router.GET("/custom", bindCustom)

	return router
}

// curl -X GET http://localhost:8000/query?id=5
func bindQuery(c echo.Context) error {
	input := new(User)

	c.Bind(input)

	return c.JSON(http.StatusOK, input)
}

// curl -X POST http://localhost:8000/form -d 'id=5'
func bindForm(c echo.Context) error {
	input := new(User)

	c.Bind(input)

	return c.JSON(http.StatusOK, input)
}

// curl -X POST http://localhost:8000/json -H "Content-Type: application/json" -d '{"id":"5"}'
func bindJSON(c echo.Context) error {
	input := new(User)

	c.Bind(input)

	return c.JSON(http.StatusOK, input)
}

// curl -X GET http://localhost:8000/custom?id=5
func bindCustom(c echo.Context) error {
	input := new(User)

	err := echo.QueryParamsBinder(c).
		String("id", &input.ID).BindError()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, input)
}

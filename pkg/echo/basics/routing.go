package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	router := router()
	router.Start(":8000")
}

func router() *echo.Echo {
	router := echo.New()

	router.GET("/users", getDataFromQuery)
	router.GET("/users/:id", getDataFromPath)
	router.POST("/users", getDataFromJSON)

	return router
}

func getDataFromPath(c echo.Context) error {
	// Path Parameters
	id := c.Param("id")

	return c.String(http.StatusOK, id)
}

func getDataFromQuery(c echo.Context) error {
	// Query Parameters
	name := c.QueryParam("name")
	lastName := c.QueryParam("last_name")

	return c.String(http.StatusOK, name+" "+lastName)
}

func getDataFromJSON(c echo.Context) error {
	// JSON
	input := new(Post)

	if err := c.Bind(input); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, input)
}

type Post struct {
	Title       string `json:"title"       form:"title"       query:"title"`
	Description string `json:"description" form:"description" query:"description"`
}

// curl запрос

/*
curl -X POST http://localhost:8000/users -d 'title=Егор' -d 'description=Обо мне'
curl -X POST http://localhost:8000/users \-H "Content-Type: application/json" -d '{"title":"Егор","description":"Обо мне"}'
*/

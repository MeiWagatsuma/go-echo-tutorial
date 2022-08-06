package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from the web side!")
}

func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	return c.String(
		http.StatusOK,
		fmt.Sprintf(
			"your cat name is: %s\nnand his type is: %s\n",
			catName,
			catType))
}

func main() {
	fmt.Println("Welcome to the server")

	e := echo.New()
	e.GET("/", hello)
	e.GET("/cats", getCats)

	e.Start(":8080")
}

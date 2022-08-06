package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from the web side!")
}

func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(
			http.StatusOK,
			fmt.Sprintf(
				"your cat name is: %s\nnand his type is: %s\n",
				catName,
				catType))
	} else if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": catName,
			"type": catType,
		})
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "you need to lets us know if you want json or string data",
		})
	}

}

func addCat(c echo.Context) error {
	cat := Cat{}
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed to reading the request body for addCats: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &cat)
	if err != nil {
		log.Printf("Failed unmarshaling in addCats: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	log.Printf("this is your cat: %#v", cat)
	return c.String(http.StatusOK, "we got your cat!")
}

func main() {
	fmt.Println("Welcome to the server")

	e := echo.New()
	e.GET("/", hello)
	e.GET("/cats/:data", getCats)

	e.POST("/cats", addCat)

	e.Start(":8080")
}

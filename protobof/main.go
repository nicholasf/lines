package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Event struct {
	Sentence string    `json:"sentence"`
	When     time.Time `json:"when"`
}

func logJson(c echo.Context) error {
	e := new(Event)

	if err := c.Bind(e); err != nil {
		return c.JSON(http.StatusBadRequest, e)
	}

	println(e.Sentence)

	return c.JSON(http.StatusOK, e)
}

func main() {
	e := echo.New()

	e.POST("/json", logJson)

	e.Logger.Fatal(e.Start(":1323"))
}

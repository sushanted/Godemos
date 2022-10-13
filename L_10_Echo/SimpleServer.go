package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	fmt.Println("Hello world")

	e := echo.New()

	e.GET("/health", Health)
	e.GET("/echo", Echo)

	e.Start(":8082")

	//time.Sleep(5 * time.Minute)
}

func Health(c echo.Context) error {
	c.String(http.StatusOK, "ok")
	return nil
}

func Echo(c echo.Context) error {

	c.String(http.StatusOK, fmt.Sprintf("Welcome %s", c.QueryParam("name")))
	return nil
}

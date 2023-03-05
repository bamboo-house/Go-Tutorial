package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/request", func(c echo.Context) error {
		req := c.Request()
		format := `
		Protocol: %s<br>
		Host: %s<br>
		Remote Address: %s<br>
		Method: %s<br>
		Path: %s<br>
		`
		return c.HTML(http.StatusOK, fmt.Sprintf(format, req.Proto, req.Host, req.RemoteAddr, req.Method, req.URL.Path))
	})
	if err := e.StartTLS(":8080", "cert.pem", "key.pem"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
package main

import (
	// "net/http"

	"github.com/bamboo-house/Go-Tutorial/echo-myapp/controller"
	"github.com/bamboo-house/Go-Tutorial/echo-myapp/model"
	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	return c.String(200, "Hellofffffffvvld!")
}

// func connect(c echo.Context) error {
// 	db, _ := model.DB.DB()
// 	defer db.Close()
// 	err := db.Ping()
// 	if err != nil {
// 		return c.String(http.StatusInternalServerError, "DB接続失ff敗しました")
// 	} else {
// 		return c.String(http.StatusOK, "DB接続しました")
// 	}
// }

func main() {
	e := echo.New()
	db, _ := model.DB.DB()
	defer db.Close()

	// 下記コメントアウトしないとダメ
	// e.GET("/", connect)
	e.GET("/", hello)
	e.GET("/users", controller.GetUsers)
	e.GET("/users/:id", controller.GetUser)
	e.POST("/users", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)
	e.Logger.Fatal(e.Start(":8080"))
}

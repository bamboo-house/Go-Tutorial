package main

import (
	"github.com/bamboo-house/Go-Tutorial/echo-myapp/controller"
	"github.com/bamboo-house/Go-Tutorial/echo-myapp/model"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db, _ := model.DB.DB()
	defer db.Close()

	// 下記コメントアウトしないとダメ
	e.GET("/users", controller.GetUsers)
	e.GET("/users/:id", controller.GetUser)
	e.POST("/users", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)
	e.Logger.Fatal(e.Start(":8080"))
}

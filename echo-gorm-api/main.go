package main

import (
	"fmt"
	// "echo-gorm-api/database"
	// "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	// "os"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// func getUsers(c echo.Context) error {
// 	users := []User{}
// 	DB.Find(&users)
// 	return c.JSON(http.StatusOK, users)
// }

func connect(c echo.Context) error {
	db, _ := DB.DB()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, "DB接続成功")
	}
}

var DB *gorm.DB

func main() {
	e := echo.New()
	// database.Connect()

	// user := os.Getenv("MYSQL_USER")
	// password := os.Getenv("MYSQL_PASSWORD")
	// host := os.Getenv("MYSQL_HOST")
	// port := "3306"
	// database_name := os.Getenv("MYSQL_DATABASE")

	// dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4"
	// connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, database_name) // 修正!!

	dsn := "webuser:webpass@tcp(db:3306)/go_mysql8_development?charset=utf8mb4"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// DB, err := gorm.Open("mysql", connection)

	if err != nil {
		fmt.Println("エラーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーー")
		fmt.Println(err.Error())
	}
	sqlDB, _ := DB.DB()
	defer sqlDB.Close()

	e.GET("/users", connect)

	e.Logger.Fatal(e.Start(":8080"))
}

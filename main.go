package main

import (
    "net/http"
    "strconv"
    "github.com/labstack/echo/v4"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "./models"
    "./controllers"
)

var (
    DB *gorm.DB
)

func init() {
    InitDB()
    InitialMigration()
}

type Config struct {
    DB_Username string
    DB_Password string
    DB_Port string
    DB_Host string
    DB_Name string
}

func InitDB() {
    // Konfigurasi database
    config := Config{
        DB_Username: "root",
        DB_Password: "root123",
        DB_Port: "3306",
        DB_Host: "localhost",
        DB_Name: "crud_go",
    }

    // Membangun string koneksi
    connectionString := config.DB_Username + ":" + config.DB_Password +
        "@tcp(" + config.DB_Host + ":" + config.DB_Port + ")/" +
        config.DB_Name + "?charset=utf8&parseTime=True&loc=Local"

    var err error
    DB, err = gorm.Open("mysql", connectionString)
    if err != nil {
        panic(err)
    }
}

func InitialMigration() {
    // Auto-migrate model User
    DB.AutoMigrate(&models.User{})
}

func main() {
    // Membuat instance Echo
    e := echo.New()

    // Routing menggunakan controller
    e.GET("/users", controllers.GetUsersController)
    e.GET("/users/:id", controllers.GetUserController)
    e.POST("/users", controllers.CreateUserController)
    e.PUT("/users/:id", controllers.UpdateUserController)
    e.DELETE("/users/:id", controllers.DeleteUserController)

    // Memulai server
    e.Logger.Fatal(e.Start(":8000"))
}

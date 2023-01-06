package middleware

import (
    "fmt"
    "os"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

type DbConnectionMiddleware struct {
}

func DbConnection() *DbConnectionMiddleware {
    return &DbConnectionMiddleware{
    }
}

func (t *DbConnectionMiddleware) Connect() {
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    database := os.Getenv("DB_DATABASE")
    user := os.Getenv("DB_USERNAME")
    password := os.Getenv("DB_PASSWORD")

    conn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?parseTime=true"
    fmt.Println("connection : " + conn)

    db, err := gorm.Open("mysql", conn)
    if err != nil {
        panic("failed to connect database")
    }

    DB = db
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

    "app/database"
    "app/repositories"
    "app/controllers"
)

func main() {
    g := gin.Default()
    db := database.Connection()

    // -----
    // route
    // -----
    g.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "hello world",
        })
    })

    test := g.Group("/test")
    {
        v1 := test.Group("/v1")
        {
            r := repositories.Test(db)
            c := controllers.Test(r)
            v1.GET("/test", c.Index)
        }
    }

    g.Run(":8000")
}
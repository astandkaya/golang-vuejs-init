package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

    "app/middleware"
    "app/controllers"
)

func main() {
    g := gin.Default()

    db := middleware.DbConnection()
    db.Connect()

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
            tc := controllers.Top()
            v1.GET("/test", tc.Index)
        }
    }

    g.Run(":8000")
}
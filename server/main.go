package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"app/controllers"
)

func main() {
    engine := gin.Default()
    
    engine.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "hello world",
        })
    })

    test := engine.Group("/test")
    {
        v1 := test.Group("/v1")
        {
            v1.GET("/test", controllers.Test)
        }
    }
    engine.Run(":8000")
}
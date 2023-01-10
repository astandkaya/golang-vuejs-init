package main

import (
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
    v1 := g.Group("/v1")
    {
        r := repositories.Test(db)
        c := controllers.Test(r)
        v1.GET("/test", c.Index)
    }

    g.Run(":8000")
}
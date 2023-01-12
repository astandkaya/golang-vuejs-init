package main

import (
    "github.com/gin-gonic/gin"
    jwt "github.com/appleboy/gin-jwt/v2"

    "app/database"
    "app/repositories"
    "app/controllers"
    "app/middleware"
)

func main() {
    g := gin.Default()
    db := database.Connection()

    identityKey := "id"

    authMiddleware := middleware.Auth(identityKey)
    authMiddleware.MiddlewareInit()

    // -----
    // route
    // -----
    v1 := g.Group("/v1")
    {
        // 通常のAPI
        r := repositories.Test(db)
        c := controllers.Test(r)
        v1.GET("/test", c.Index)

        // 認証API
        auth := v1.Group("/auth")
        {
            auth.POST("/login", authMiddleware.LoginHandler)
            auth.GET("/refresh_token", authMiddleware.RefreshHandler)
        }

        // 認証が必要なAPI
        v1.Use(authMiddleware.MiddlewareFunc())
        {
            v1.GET("/test_auth", func(c *gin.Context) {
                claims := jwt.ExtractClaims(c)
                // user, _ := c.Get(identityKey)
                c.JSON(200, gin.H{
                  "userID":   claims[identityKey],
                  // "userName": user.(*User).UserName,
                  "text":     "Hello World.",
                })
            })
        }
    }

    g.Run(":8000")
}
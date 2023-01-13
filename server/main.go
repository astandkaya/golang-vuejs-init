package main

import (
    "github.com/gin-gonic/gin"
    jwt "github.com/appleboy/gin-jwt/v2"

    "app/database"
    "app/repositories"
    "app/controllers"
    "app/middleware"
    "app/models"
)

func main() {
    g := gin.Default()
    db := database.Connection()

    identityKey := "id"

    authMiddleware := middleware.Auth(identityKey, repositories.User(db))
    authMiddleware.MiddlewareInit()

    // -----
    // route
    // -----
    v1 := g.Group("/v1")
    {

        // 通常のAPI
        test := controllers.Test(repositories.Test(db))
        {
            v1.GET("/test", test.Index)
        }

        // 認証系API
        auth := v1.Group("/auth")
        {
            register := controllers.Register(repositories.User(db))
            {
                auth.POST("/register", register.Store)
            }

            auth.POST("/login", authMiddleware.LoginHandler)
            auth.GET("/refresh_token", authMiddleware.RefreshHandler)
        }

        // 認証が必要なAPI
        v1.Use(authMiddleware.MiddlewareFunc())
        {
            v1.GET("/test_auth", func(c *gin.Context) {
                claims := jwt.ExtractClaims(c)
                user, _ := c.Get(identityKey)
                c.JSON(200, gin.H{
                  "userID":   claims[identityKey],
                  "email": user.(*models.UserModel).Email,
                  "text":     "Hello World.",
                })
            })
        }
    }

    g.Run(":8000")
}
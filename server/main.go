package main

import (
    "github.com/gin-gonic/gin"

    "app/database"
    "app/repositories"
    "app/controllers"
    "app/middleware"
    "app/utils"
)

func main() {
    g := gin.Default()
    db := database.Connection()

    utils.ValidatorInit(db)

    authMiddleware := middleware.Auth(repositories.User(db))
    authMiddleware.MiddlewareInit()

    // -----
    // route
    // -----
    v1 := g.Group("/v1")
    {
        v1.Use(middleware.RateLimit())
        v1.Use(middleware.Cors())

        // 通常のAPI
        test := controllers.Test(repositories.Test(db))
        {
            v1.GET("/test", test.Index)
        }

        // 認証系API
        auth := v1.Group("/auth")
        {
            signup := controllers.Signup(repositories.User(db))
            {
                auth.POST("/signup", signup.Store)
            }

            auth.POST("/login", authMiddleware.LoginHandler)
            auth.GET("/refresh_token", authMiddleware.RefreshHandler)
        }

        // 認証が必要なAPI
        member := v1.Group("/", authMiddleware.MiddlewareFunc())
        {
            testAuth := controllers.TestAuth(repositories.User(db))
            {
                member.GET("/test_auth", testAuth.Index)
            }
        }
    }

    g.Run(":8000")
}
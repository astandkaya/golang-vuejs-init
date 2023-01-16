package middleware

import (
    "time"

    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func Cors() gin.HandlerFunc {
    return cors.New(cors.Config{
        AllowOrigins: []string{
            "http://localhost",
        },
        AllowMethods: []string{
            "POST",
            "GET",
            "OPTIONS",
        },
        AllowHeaders: []string{
            "Content-Type",
        },
        AllowCredentials: false,
        MaxAge: 24 * time.Hour,
    })
}
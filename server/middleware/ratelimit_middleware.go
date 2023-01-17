package middleware

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "golang.org/x/time/rate"
)

func RateLimit() gin.HandlerFunc {
    limiter := rate.NewLimiter(rate.Every(time.Second), 20)

    return func(ctx *gin.Context) {
        if limiter.Allow() == false {
            ctx.JSON(http.StatusTooManyRequests, gin.H{
                "status": "ng",
                "message": "too many requests",
            })
            ctx.Abort()
        }
    }
}
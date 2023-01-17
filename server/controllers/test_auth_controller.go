package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    jwt "github.com/appleboy/gin-jwt/v2"

    "app/middleware"
    "app/models"
)

type TestAuthController struct {
    userRepo models.UserRepository
}

func TestAuth(userRepo models.UserRepository) *TestAuthController {
    return &TestAuthController{
        userRepo: userRepo,
    }
}

func (c *TestAuthController) Index(ctx *gin.Context) {
    claims := jwt.ExtractClaims(ctx)
    userId := claims[middleware.IdentityKey]
    ctx.JSON(http.StatusOK, gin.H{
        "userID": userId,
        "text": "Hello World.",
    })
}

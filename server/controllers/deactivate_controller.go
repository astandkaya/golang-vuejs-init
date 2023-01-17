package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    jwt "github.com/appleboy/gin-jwt/v2"

    "app/middleware"
    "app/models"
)

type DeactivateController struct {
    userRepo models.UserRepository
}

func Deactivate(userRepo models.UserRepository) *DeactivateController {
    return &DeactivateController{
        userRepo: userRepo,
    }
}

func (r *DeactivateController) Store(ctx *gin.Context) {
    claims := jwt.ExtractClaims(ctx)
    user := models.UserModel{
        Email: claims[middleware.IdentityKey].(string),
    }

    if !r.userRepo.Exists(user) {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "status": "ng",
            "message": "user not found",
        })
        return
    }

    r.userRepo.Delete(user)

    ctx.JSON(http.StatusOK, gin.H{
        "status": "ok",
    })
}

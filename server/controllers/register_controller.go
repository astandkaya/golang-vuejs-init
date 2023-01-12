package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "app/models"
    "app/services"
)

type RegisterController struct {
    userRepo models.UserRepository
}

func Register(userRepo models.UserRepository) *RegisterController {
    return &RegisterController{
        userRepo: userRepo,
    }
}

func (r *RegisterController) Store(ctx *gin.Context) {
    username := ctx.Query("username")
    password := services.Hash().Make(ctx.Query("password"))

    user := &models.UserModel{
        UserName: username,
        Password: password,
    }

    r.userRepo.Create(user);

    ctx.JSON(http.StatusCreated, gin.H{
        "status": "ok",
    })
}
package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "app/models"
    "app/services"
    "app/utils"
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
    user := models.User()

    if err := ctx.ShouldBind(user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "status": "ng",
            "message": "bind error",
        })
        return
    }

    validate := utils.Validator()
    if err := validate.Struct(user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "status": "ng",
            "message": "valid error",
        })
        return
    }

    user.Password = services.Hash().Make(user.Password)
    r.userRepo.Create(user);

    ctx.JSON(http.StatusCreated, gin.H{
        "status": "ok",
    })
}

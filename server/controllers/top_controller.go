package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "app/middleware"
    "app/models"
)

type TopController struct {
}

func Top() *TopController {
    return &TopController{
    }
}

func (t *TopController) Index(c *gin.Context) {
    tm := models.Test()
    middleware.DB.Delete(tm)

    c.JSON(http.StatusCreated, gin.H{
        "status": "ok",
    })
}

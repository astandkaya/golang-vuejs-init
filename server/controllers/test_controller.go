package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "app/models"
)

type TestController struct {
    testRepo models.TestRepository
}

func Test(testRepo models.TestRepository) *TestController {
    return &TestController{
        testRepo: testRepo,
    }
}

func (c *TestController) Index(ctx *gin.Context) {
    ctx.JSON(http.StatusCreated, gin.H{
        "status": "ok",
    })
}

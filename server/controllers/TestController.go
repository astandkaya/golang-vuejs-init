package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func Test(c *gin.Context) {
    c.JSON(http.StatusCreated, gin.H{
        "status": "ok",
    })
}

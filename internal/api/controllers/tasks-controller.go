package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTasks(c *gin.Context) {
	responseData := gin.H{
        "name":    "a",
        "text":    "b",
        "user_id": "c",
    }
	
	c.JSON(http.StatusOK, responseData)
}

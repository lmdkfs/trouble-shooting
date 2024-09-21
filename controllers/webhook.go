package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Webhook(c *gin.Context) {
	var requestData interface{}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jsonString, err := json.MarshalIndent(requestData, "", "  ")
	if err != nil {
		c.JSON(500, gin.H{"error": "json marshal error"})
		return
	}
	fmt.Printf("request data: %s\n", jsonString)
	c.JSON(http.StatusOK, gin.H{"message": "request webhook success"})
}

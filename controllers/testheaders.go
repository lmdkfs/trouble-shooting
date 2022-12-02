package controllers

import (
	"trouble-shooting/utils"

	"github.com/gin-gonic/gin"
)

func PrintAllHeaders(c *gin.Context) {
	utils.Logger.Debug(c.Request.Header)
	c.JSON(200, gin.H{
		"Headers":    c.Request.Header,
		"RemoteAddr": c.RemoteIP(),
	})
}

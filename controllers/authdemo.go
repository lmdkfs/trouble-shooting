package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var db = make(map[string]string)

func Authorized(c *gin.Context) {
	user := c.Params.ByName("name")
	value, ok := db[user]
	if ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	}
}

func Admin(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)

	var json struct {
		Value string `json:"value" binding:"required"`
	}
	if c.Bind(&json) == nil {
		db[user] = json.Value
		c.JSON(http.StatusOK, gin.H{"status":"ok"})
	}

}

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Healthz(c *gin.Context) {
	testHttpStatus := os.Getenv("TESTHTTP")
	if testHttpStatus == "0" {
		c.String(http.StatusOK, "healthy")
	} else {
		c.String(http.StatusInternalServerError, "Unhealthy")
	}

}

func ChangeTestHTTPStatus(c *gin.Context) {
	code := c.Params.ByName("code")
	os.Setenv("TESTHTTP", code)

}

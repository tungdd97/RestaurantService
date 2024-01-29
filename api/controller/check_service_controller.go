package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CheckServiceController struct{}

func (csc *CheckServiceController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "Response success!!")
}

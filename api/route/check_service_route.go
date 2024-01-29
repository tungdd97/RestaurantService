package route

import (
	"github.com/gin-gonic/gin"
	"restaurant-service/api/controller"
)

func InitCheckServiceRoute(group *gin.RouterGroup) {

	csc := &controller.CheckServiceController{}
	group.GET("/ping", csc.Ping)
}

package route

import (
	"github.com/gin-gonic/gin"
	"restaurant-service/bootstrap"
	"restaurant-service/database/mongo"
	"time"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("api/v1.0/")
	// All Public APIs
	InitCheckServiceRoute(publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewSignupRouter(env, timeout, db, publicRouter)
	NewRestaurantRouter(env, timeout, db, publicRouter)

}

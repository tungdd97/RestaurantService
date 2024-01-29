package route

import (
	"github.com/gin-gonic/gin"
	"restaurant-service/api/controller"
	"restaurant-service/bootstrap"
	"restaurant-service/database/mongo"
	"restaurant-service/entity"
	"restaurant-service/repository"
	"restaurant-service/usecase"
	"time"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, entity.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}

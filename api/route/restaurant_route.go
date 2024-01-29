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

func NewRestaurantRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewRestaurantRepository(db, entity.CollectionRestaurant)
	resC := controller.RestaurantController{
		RestaurantUsecase: usecase.NewRestaurantUsecase(ur, timeout),
		Env:               env,
	}
	group.POST("/restaurants", resC.Create)
	group.POST("/restaurants/action/filter", resC.FetchByCondition)
}

package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"restaurant-service/bootstrap"
	"restaurant-service/entity"
	"restaurant-service/lib"
	"restaurant-service/usecase"
	"time"
)

type restaurantRequestCreate struct {
	Title       string `form:"title" binding:"required"`
	Code        string `form:"code" binding:"required"`
	Description string `form:"description" binding:"required"`
}

type restaurantResponseCreate struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Code        string  `json:"code"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type restaurantRequestFetch struct {
	Keyword string             `form:"keyword"`
	Rating  map[string]float32 `form:"rating"`
}

type RestaurantController struct {
	RestaurantUsecase usecase.RestaurantUsecase
	Env               *bootstrap.Env
}

func (resC *RestaurantController) Create(c *gin.Context) {
	var request restaurantRequestCreate

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = resC.RestaurantUsecase.GetByCode(c, request.Code)
	if err == nil {
		c.JSON(http.StatusConflict, lib.ErrorResponse{Message: "Restaurant already exists with the given code"})
		return
	}
	accountAction, err := lib.ExtractParamFromRequest(c, "id")
	if err != nil {
		c.JSON(http.StatusForbidden, lib.ErrorResponse{Message: "Token illegal!"})
		return
	}
	objectAccountAction, _ := primitive.ObjectIDFromHex(accountAction)

	restaurant := entity.Restaurant{
		ID:          primitive.NewObjectID(),
		Title:       request.Title,
		Code:        request.Code,
		Description: request.Code,
		CreatedBy:   objectAccountAction,
		UpdatedBy:   objectAccountAction,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	err = resC.RestaurantUsecase.Create(c, &restaurant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.ErrorResponse{Message: err.Error()})
		return
	}

	detailRestaurant, err := resC.RestaurantUsecase.GetByCode(c, request.Code)
	result := restaurantResponseCreate{
		ID:          detailRestaurant.ID.Hex(),
		Title:       detailRestaurant.Title,
		Description: detailRestaurant.Description,
		Code:        detailRestaurant.Code,
		Rating:      detailRestaurant.Rating,
		CreatedAt:   detailRestaurant.CreatedAt.Format(time.DateTime),
		UpdatedAt:   detailRestaurant.UpdatedAt.Format(time.DateTime),
	}
	c.JSON(http.StatusOK, result)
}

func (resC *RestaurantController) FetchByCondition(c *gin.Context) {
	var request restaurantRequestFetch

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.ErrorResponse{Message: err.Error()})
		return
	}

	keyword := request.Keyword

	rating := request.Rating
	fromRating := rating["from"]
	toRating := rating["to"]

	filterQuery := bson.M{}
	if keyword != "" {
		filterQuery["$or"] = []bson.M{
			bson.M{"title": bson.M{"$regex": keyword, "$options": "i"}},
			bson.M{"description": bson.M{"$regex": keyword, "$options": "i"}},
			bson.M{"code": bson.M{"$regex": keyword, "$options": "i"}},
		}
	}

	if fromRating != 0.0 && toRating != 0.0 {
		filterQuery["$and"] = []bson.M{
			bson.M{"rating": bson.M{"$gte": fromRating, "$lte": toRating}},
		}
	}

	restaurants, err := resC.RestaurantUsecase.FetchByCondition(c, filterQuery)
	c.JSON(http.StatusOK, restaurants)
}

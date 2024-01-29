package usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"restaurant-service/entity"
	"restaurant-service/repository"
	"time"
)

type restaurantUsecase struct {
	restaurantRepository repository.RestaurantRepository
	contextTimeout       time.Duration
}

func NewRestaurantUsecase(restaurantRepository repository.RestaurantRepository, timeout time.Duration) RestaurantUsecase {
	return &restaurantUsecase{
		restaurantRepository: restaurantRepository,
		contextTimeout:       timeout,
	}
}

func (resU *restaurantUsecase) Create(c context.Context, restaurant *entity.Restaurant) error {
	ctx, cancel := context.WithTimeout(c, resU.contextTimeout)
	defer cancel()
	return resU.restaurantRepository.Create(ctx, restaurant)
}

func (resU *restaurantUsecase) GetByID(c context.Context, id string) (entity.Restaurant, error) {
	ctx, cancel := context.WithTimeout(c, resU.contextTimeout)
	defer cancel()
	return resU.restaurantRepository.GetByID(ctx, id)
}

func (resU *restaurantUsecase) GetByCode(c context.Context, code string) (entity.Restaurant, error) {
	ctx, cancel := context.WithTimeout(c, resU.contextTimeout)
	defer cancel()
	return resU.restaurantRepository.GetByCode(ctx, code)
}

func (resU *restaurantUsecase) Fetch(c context.Context) ([]entity.Restaurant, error) {
	ctx, cancel := context.WithTimeout(c, resU.contextTimeout)
	defer cancel()
	return resU.restaurantRepository.Fetch(ctx)
}

func (resU *restaurantUsecase) FetchByCondition(c context.Context, condition bson.M) ([]entity.Restaurant, error) {
	ctx, cancel := context.WithTimeout(c, resU.contextTimeout)
	defer cancel()
	return resU.restaurantRepository.FetchByCondition(ctx, condition)
}

func (resU *restaurantUsecase) UpdateByID(c context.Context, id string, update interface{}) (int, error) {
	ctx, cancel := context.WithTimeout(c, resU.contextTimeout)
	defer cancel()
	return resU.restaurantRepository.UpdateByID(ctx, id, update)
}

package usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"restaurant-service/entity"
)

type (
	CheckServiceUsecase interface {
		Ping(c context.Context) (message string, err error)
	}

	RestaurantUsecase interface {
		Create(c context.Context, restaurant *entity.Restaurant) error
		GetByID(c context.Context, id string) (entity.Restaurant, error)
		GetByCode(c context.Context, code string) (entity.Restaurant, error)
		Fetch(c context.Context) ([]entity.Restaurant, error)
		FetchByCondition(c context.Context, condition bson.M) ([]entity.Restaurant, error)
		UpdateByID(c context.Context, id string, update interface{}) (int, error)
	}

	SignupUsecase interface {
		Create(c context.Context, user *entity.User) error
		GetUserByEmail(c context.Context, email string) (entity.User, error)
		CreateAccessToken(user *entity.User, secret string, expiry int) (accessToken string, err error)
		CreateRefreshToken(user *entity.User, secret string, expiry int) (refreshToken string, err error)
	}

	LoginUsecase interface {
		GetUserByEmail(c context.Context, email string) (entity.User, error)
		CreateAccessToken(user *entity.User, secret string, expiry int) (accessToken string, err error)
		CreateRefreshToken(user *entity.User, secret string, expiry int) (refreshToken string, err error)
		UpdateAccessTimeLatest(c context.Context, stringId string) (status int, err error)
	}
)

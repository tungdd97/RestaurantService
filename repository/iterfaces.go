package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"restaurant-service/entity"
)

type (
	UserRepository interface {
		Create(c context.Context, user *entity.User) error
		Fetch(c context.Context) ([]entity.User, error)
		GetByEmail(c context.Context, email string) (entity.User, error)
		GetByID(c context.Context, id string) (entity.User, error)
		UpdateByID(c context.Context, id string, update interface{}) (int, error)
	}

	RestaurantRepository interface {
		Create(c context.Context, restaurant *entity.Restaurant) error
		GetByID(c context.Context, id string) (entity.Restaurant, error)
		GetByCode(c context.Context, code string) (entity.Restaurant, error)
		Fetch(c context.Context) ([]entity.Restaurant, error)
		FetchByCondition(c context.Context, condition bson.M) ([]entity.Restaurant, error)
		UpdateByID(c context.Context, id string, update interface{}) (int, error)
	}
)

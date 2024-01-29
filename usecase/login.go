package usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"restaurant-service/entity"
	"restaurant-service/lib"
	"restaurant-service/repository"
	"time"
)

type loginUsecase struct {
	userRepository repository.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository repository.UserRepository, timeout time.Duration) LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) GetUserByEmail(c context.Context, email string) (entity.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByEmail(ctx, email)
}

func (lu *loginUsecase) CreateAccessToken(user *entity.User, secret string, expiry int) (accessToken string, err error) {
	return lib.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *entity.User, secret string, expiry int) (refreshToken string, err error) {
	return lib.CreateRefreshToken(user, secret, expiry)
}
func (lu *loginUsecase) UpdateAccessTimeLatest(c context.Context, stringId string) (int, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()

	return lu.userRepository.UpdateByID(ctx, stringId, bson.M{"$set": bson.M{"access_time_latest": time.Now().UTC()}})
}

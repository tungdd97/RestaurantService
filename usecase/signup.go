package usecase

import (
	"context"
	"restaurant-service/entity"
	"restaurant-service/lib"
	"restaurant-service/repository"
	"time"
)

type signupUsecase struct {
	userRepository repository.UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepository repository.UserRepository, timeout time.Duration) SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *signupUsecase) Create(c context.Context, user *entity.User) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.Create(ctx, user)
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (entity.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByEmail(ctx, email)
}

func (su *signupUsecase) CreateAccessToken(user *entity.User, secret string, expiry int) (accessToken string, err error) {
	return lib.CreateAccessToken(user, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *entity.User, secret string, expiry int) (refreshToken string, err error) {
	return lib.CreateRefreshToken(user, secret, expiry)
}

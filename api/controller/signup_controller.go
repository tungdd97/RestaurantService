package controller

import (
	"net/http"
	"restaurant-service/bootstrap"
	"restaurant-service/entity"
	"restaurant-service/lib"
	"restaurant-service/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type signupRequest struct {
	Name        string `form:"name" binding:"required"`
	Email       string `form:"email" binding:"required,email"`
	Password    string `form:"password" binding:"required"`
	TypeAccount string `form:"type_account" binding:""`
}

type signupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupController struct {
	SignupUsecase usecase.SignupUsecase
	Env           *bootstrap.Env
}

func (sc *SignupController) Signup(c *gin.Context) {
	var request signupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = sc.SignupUsecase.GetUserByEmail(c, request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, lib.ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	password := request.Password

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.ErrorResponse{Message: err.Error()})
		return
	}

	password = string(encryptedPassword)

	typeAccount := request.TypeAccount
	if typeAccount == "" {
		typeAccount = "web"
	}

	user := entity.User{
		ID:               primitive.NewObjectID(),
		Name:             request.Name,
		Email:            &request.Email,
		Password:         password,
		TypeAccount:      typeAccount,
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
		AccessTimeLatest: time.Now().UTC(),
	}

	err = sc.SignupUsecase.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.ErrorResponse{Message: err.Error()})
		return
	}

	signupResponse := signupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}

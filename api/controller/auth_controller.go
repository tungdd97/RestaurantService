package controller

import (
	"net/http"
	"restaurant-service/bootstrap"
	"restaurant-service/lib"
	"restaurant-service/usecase"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type loginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginController struct {
	LoginUsecase usecase.LoginUsecase
	Env          *bootstrap.Env
	Log          lib.Logger
}

func (lc *LoginController) Login(c *gin.Context) {
	var request loginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := lc.LoginUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, lib.ErrorResponse{Message: "User not found with the given email"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, lib.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := loginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	userId := user.ID.Hex()
	_, err = lc.LoginUsecase.UpdateAccessTimeLatest(c, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, loginResponse)
}

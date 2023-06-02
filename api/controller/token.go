package controller

import (
	"fmt"
	"net/http"
	"todo-app-go/bootstrap"
	"todo-app-go/domain"

	"github.com/gin-gonic/gin"
)

type TokenController struct {
	TokenUsecase domain.RefreshTokenUsecase
	Env          *bootstrap.Env
}

func (t *TokenController) RefreshToken(c *gin.Context) {
	var tokenRequest domain.RefreshTokenRequest

	err := c.ShouldBind(&tokenRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	fmt.Println("refresh token request: ", tokenRequest)

	id, err := t.TokenUsecase.ExtractIDFromToken(tokenRequest.RefreshToken, t.Env.RefreshTokenSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User not found"})
		return
	}

	user, err := t.TokenUsecase.GetUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User not found"})
		return
	}

	accessToken, err := t.TokenUsecase.CreateAccessToken(&user, t.Env.AccessTokenSecret, t.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := t.TokenUsecase.CreateRefreshToken(&user, t.Env.RefreshTokenSecret, t.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshTokenResponse := domain.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, refreshTokenResponse)
}

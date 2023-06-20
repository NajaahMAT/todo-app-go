package controller

import (
	"fmt"
	"net/http"
	"todo-app-go/bootstrap"
	"todo-app-go/domain"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}

func (l *LoginController) Login(c *gin.Context) {
	var loginRequest domain.LoginRequest

	err := c.ShouldBind(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	fmt.Println("login request: ", loginRequest)
	user, err := l.LoginUsecase.GetUserByEmail(c, loginRequest.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found with the given email"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	accessToken, err := l.LoginUsecase.CreateAccessToken(&user, l.Env.AccessTokenSecret, l.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := l.LoginUsecase.CreateRefreshToken(&user, l.Env.RefreshTokenSecret, l.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       user.ID.Hex(),
	}

	c.SetCookie("access_token", accessToken, l.Env.AccessTokenExpiryHour, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refreshToken, l.Env.RefreshTokenExpiryHour, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", l.Env.AccessTokenExpiryHour, "/", "localhost", false, false)
	c.SetCookie("user_id", user.ID.Hex(), 0, "", "", false, false)

	c.JSON(http.StatusOK, loginResponse)
}

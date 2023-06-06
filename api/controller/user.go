package controller

import (
	"fmt"
	"net/http"
	"todo-app-go/bootstrap"
	"todo-app-go/domain"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}

func (u *UserController) GetUserProfile(c *gin.Context) {
	email := c.Param("email")

	fmt.Println("email: ", email)

	user, err := u.LoginUsecase.GetUserByEmail(c, email)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found with the given email"})
		return
	}

	c.JSON(http.StatusOK, user)
}

package route

import (
	"fmt"
	"time"
	"todo-app-go/api/controller"
	"todo-app-go/bootstrap"
	"todo-app-go/domain"
	"todo-app-go/mongo"
	"todo-app-go/repository"
	"todo-app-go/usecase"

	"github.com/gin-gonic/gin"
)

func NewTokenRouter(env *bootstrap.Env, t time.Duration, db mongo.Database, g *gin.RouterGroup) {
	fmt.Println("route.signup.go/NewTokenRouter")

	ur := repository.NewUserRepository(db, domain.CollectionUser)

	tc := &controller.TokenController{
		TokenUsecase: usecase.NewRefreshTokenUsecase(ur, t),
		Env:          env,
	}

	g.POST("/refresh-token", tc.RefreshToken)
}

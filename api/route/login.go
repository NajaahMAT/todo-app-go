package route

import (
	"time"
	"todo-app-go/api/controller"
	"todo-app-go/bootstrap"
	"todo-app-go/domain"
	"todo-app-go/mongo"
	"todo-app-go/repository"
	"todo-app-go/usecase"

	"github.com/gin-gonic/gin"
)

func NewLoginRouter(env *bootstrap.Env, t time.Duration, db mongo.Database, g *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, t),
		Env:          env,
	}

	g.POST("/login", lc.Login)
}

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

// func CreateTaskRouter(env *bootstrap.Env, t time.Duration, db mongo.Database, g *gin.RouterGroup) {
// 	ur := repository.NewTaskRepository(db, domain.CollectionTask)

// 	tc := &controller.TaskController{
// 		TaskUsecase: usecase.NewTaskUsecase(ur, t),
// 		Env:         env,
// 	}

// 	g.POST("/task", tc.CreateTask)
// }

func NewTaskRouter(env *bootstrap.Env, t time.Duration, db mongo.Database, g *gin.RouterGroup) {
	ur := repository.NewTaskRepository(db, domain.CollectionTask)

	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(ur, t),
		Env:         env,
	}

	g.POST("/task", tc.CreateTask)
	g.GET("/task/user/:user_id", tc.GetTasksByUserID)
	g.GET("/task/user/:user_id/status/:status", tc.GetTasksByStatus)
	g.GET("/task/:id", tc.GetTaskByID)
	g.GET("/task/user/:user_id/start/:start_date/end/:end_date", tc.GetTasksByDateRange)
	g.PUT("/task", tc.UpdateTask)
	g.PUT("/task/status", tc.UpdateStatus)
	g.DELETE("/task/:id", tc.DeleteTaskByID)
}

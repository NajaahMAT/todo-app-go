package route

import (
	"time"

	// "todo-app-go/api/middleware"
	"todo-app-go/api/middleware"
	"todo-app-go/bootstrap"
	"todo-app-go/mongo"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// publicRouter.Use(CORS())

	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// protectedRouter.Use(CORS())

	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewTokenRouter(env, timeout, db, publicRouter)

	// All Private APIs
	// CreateTaskRouter(env, timeout, db, protectedRouter)
	// GetTasksRouter(env, timeout, db, protectedRouter)
	NewTaskRouter(env, timeout, db, publicRouter)
	NewUserRouter(env, timeout, db, publicRouter)
	NewSignoutRouter(publicRouter)
}

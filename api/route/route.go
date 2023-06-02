package route

import (
	"time"

	// "todo-app-go/api/middleware"
	"todo-app-go/bootstrap"
	"todo-app-go/mongo"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	// NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	// NewRefreshTokenRouter(env, timeout, db, publicRouter)

	// protectedRouter := gin.Group("")
	// // Middleware to verify AccessToken
	// protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// // All Private APIs
	// NewProfileRouter(env, timeout, db, protectedRouter)
	// NewTaskRouter(env, timeout, db, protectedRouter)
}

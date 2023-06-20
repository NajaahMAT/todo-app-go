package route

import (
	"fmt"
	"todo-app-go/api/controller"

	"github.com/gin-gonic/gin"
)

func NewSignoutRouter(g *gin.RouterGroup) {
	fmt.Println("route.signup.go/NewSignoutRouter")

	lc := &controller.LogoutController{}

	g.POST("/signout", lc.LogoutUser)
}

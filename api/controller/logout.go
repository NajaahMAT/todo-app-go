package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogoutController struct{}

func (l *LogoutController) LogoutUser(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("logged_in", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusAccepted, gin.H{"status": http.StatusOK, "message": "Success"})
}

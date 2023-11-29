package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckIsAdmin(c *gin.Context) {
	password := c.Request.PostFormValue("password")

	c.SetSameSite(http.SameSiteNoneMode)

	if password == "TAMER_GDE_BABKI" {
		c.SetCookie("isAdmin", "true", 3600, "/", "127.0.0.1", true, true)
		c.JSON(http.StatusOK, gin.H{"message": "Admin is successfully confirmed!"})
		return
	}

	c.SetCookie("isAdmin", "false", 3600, "/", "127.0.0.1", true, true)
	c.JSON(http.StatusUnauthorized, gin.H{"message": "Access is denied!"})
}

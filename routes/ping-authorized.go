package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addPingAuthorizedRoutes(routerGroup *gin.RouterGroup) {
	pingAuthorized := routerGroup.Group("/ping-auth", gin.BasicAuth(gin.Accounts{
		"admin": "admin", //user:admin password:admin
		"user":  "123",
	}))

	pingAuthorized.GET("/", func(ctx *gin.Context) {

		user := ctx.MustGet(gin.AuthUserKey).(string)
		fmt.Println("User auth: " + user)

		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})

	})

}

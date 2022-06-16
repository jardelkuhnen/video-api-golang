package routes

import "github.com/gin-gonic/gin"

func addPingRoutes(routerGroup *gin.RouterGroup) {
	ping := routerGroup.Group("/ping")

	ping.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

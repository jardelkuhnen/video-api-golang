package main

import (
	"github.com/jardelkuhnen/video-api/routes"
)

// var (
// 	videoService    service.VideoService       = service.New()
// 	videoController controller.VideoController = controller.New(videoService)
// )

func main() {
	// server := gin.Default()

	// server.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// server.GET("/videos", func(ctx *gin.Context) {
	// 	ctx.JSON(200, videoController.FindAll())
	// })

	// server.POST("/videos", func(ctx *gin.Context) {
	// 	ctx.JSON(200, videoController.Save(ctx))
	// })

	// server.DELETE("/videos/", func(ctx *gin.Context) {
	// 	id := ctx.Query("id")
	// 	if id == "" {
	// 		ctx.Header("Content-Type", "application/json")
	// 		ctx.JSON(http.StatusNotFound, gin.H{"Error: ": "Invalid startingIndex on search filter!"})
	// 		ctx.Abort()
	// 		return
	// 	}

	// 	ctx.JSON(200, videoController.Delete(id))
	// })

	// server.Run(":8080")

	routes.Run()
}

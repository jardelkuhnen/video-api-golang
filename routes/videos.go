package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jardelkuhnen/video-api/controller"
	"github.com/jardelkuhnen/video-api/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func addVideosRoutes(routerGroup *gin.RouterGroup) {
	videos := routerGroup.Group("/videos")

	videos.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	videos.POST("/", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Video inserted with success!"})
		}

	})

	videos.DELETE("/", func(ctx *gin.Context) {
		id := ctx.Query("id")
		if id == "" {
			ctx.Header("Content-Type", "application/json")
			ctx.JSON(http.StatusNotFound, gin.H{"Error: ": "Invalid startingIndex on search filter!"})
			ctx.Abort()
			return
		}

		idInt, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		ctx.JSON(200, videoController.Delete(idInt))
	})
}

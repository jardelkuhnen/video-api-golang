package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jardelkuhnen/video-api/controller"
	"github.com/jardelkuhnen/video-api/entity"
	"github.com/jardelkuhnen/video-api/repository"
	"github.com/jardelkuhnen/video-api/service"
	"github.com/jardelkuhnen/video-api/utils"
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.New(videoRepository)
	videoController controller.VideoController = controller.New(videoService)
	video           entity.Video
	err             error
)

func addVideosRoutes(routerGroup *gin.RouterGroup) {
	videos := routerGroup.Group("/videos")

	videos.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})

	videos.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		utils.BadRequestIfConditionTruly(id == "", "Invalida Id. It must be informed!", ctx)

		idInt, err := strconv.ParseUint(id, 10, 64)
		utils.BadRequestIfError(err, ctx)

		var video entity.Video
		video, err = videoController.FindById(idInt)
		if err != nil {
			ctx.Header("Content-Type", "application/json")
			ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, video)
	})

	videos.POST("/", func(ctx *gin.Context) {
		var video entity.Video
		var err error
		video, err = videoController.Save(ctx)

		utils.BadRequestIfError(err, ctx)

		ctx.JSON(http.StatusOK, video)
	})

	videos.PUT("/", func(ctx *gin.Context) {
		video, err = videoController.Update(ctx)

		utils.BadRequestIfError(err, ctx)

		ctx.JSON(http.StatusOK, video)
	})

	videos.DELETE("/", func(ctx *gin.Context) {
		id := ctx.Query("id")

		utils.BadRequestIfConditionTruly(id == "", "Invalid startingIndex on search filter!", ctx)

		idInt, err := strconv.ParseUint(id, 10, 64)
		utils.BadRequestIfError(err, ctx)

		message, err := videoController.Delete(idInt)

		utils.BadRequestIfError(err, ctx)

		ctx.JSON(http.StatusOK, message)
	})
}

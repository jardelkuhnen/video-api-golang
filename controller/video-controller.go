package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jardelkuhnen/video-api/entity"
	"github.com/jardelkuhnen/video-api/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
	Delete(id string) string
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}

func (c *controller) Delete(id string) string {
	c.service.Delete(id)
	return "Removed with success!"
}

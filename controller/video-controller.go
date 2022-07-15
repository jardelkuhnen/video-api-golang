package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jardelkuhnen/video-api/entity"
	"github.com/jardelkuhnen/video-api/service"
	"github.com/jardelkuhnen/video-api/validators"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) (entity.Video, error)
	Update(ctx *gin.Context) (entity.Video, error)
	FindById(id uint64) (entity.Video, error)
	Delete(id uint64) (string, error)
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

// constructor
func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-bad-language", validators.ValidateBadLanguage)
	return &controller{
		service: service,
	}
}

func (c *controller) FindById(id uint64) (entity.Video, error) {
	return c.service.FindById(id)
}

func (c *controller) Update(ctx *gin.Context) (entity.Video, error) {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return entity.Video{}, err
	}
	err = validate.Struct(video)
	if err != nil {
		return entity.Video{}, err
	}

	c.service.Update(video)
	return video, nil
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) (entity.Video, error) {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return entity.Video{}, err
	}
	err = validate.Struct(video)
	if err != nil {
		return entity.Video{}, err
	}

	video, err = c.service.Save(video)
	if err != nil {
		return entity.Video{}, err
	}
	return video, err
}

func (c *controller) Delete(id uint64) (string, error) {
	err := c.service.Delete(id)
	if err != nil {
		return "", err
	}
	return "Removed with success!", nil
}

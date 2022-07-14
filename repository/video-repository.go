package repository

import "github.com/jardelkuhnen/video-api/entity"

type VideoRepository interface {
	Save(video entity.Video)
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
}

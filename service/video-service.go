package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jardelkuhnen/video-api/entity"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
	Delete(videoId string)
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	id := uuid.New()
	video.Id = id.String()
	service.videos = append(service.videos, video)
	return video
}
func (service *videoService) FindAll() []entity.Video {
	return service.videos

}
func (service *videoService) Delete(videoId string) {
	fmt.Println("Id para remoçao ", videoId)
	for i := 0; i < len(service.videos); i++ {
		if service.videos[i].Id == videoId {
			service.videos = removeElementByIndex(service.videos, i)
		}
	}

	fmt.Println(service.videos)
}

func removeElementByIndex[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

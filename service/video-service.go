package service

import (
	"fmt"

	"github.com/jardelkuhnen/video-api/entity"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
	Delete(videoId uint64)
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}
func (service *videoService) FindAll() []entity.Video {
	return service.videos

}
func (service *videoService) Delete(videoId uint64) {
	fmt.Println("Id para remoçao ", videoId)
	for i := 0; i < len(service.videos); i++ {
		if service.videos[i].ID == videoId {
			service.videos = removeElementByIndex(service.videos, i)
		}
	}

	fmt.Println(service.videos)
}

func removeElementByIndex[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

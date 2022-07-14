package service

import (
	"github.com/jardelkuhnen/video-api/entity"
	"github.com/jardelkuhnen/video-api/repository"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	Update(video entity.Video)
	FindAll() []entity.Video
	FindById(id uint64) entity.Video
	Delete(videoId uint64)
}

type videoService struct {
	// videos []entity.Video
	videoRepository repository.VideoRepository
}

func New(repo repository.VideoRepository) VideoService {
	return &videoService{
		videoRepository: repo,
	}
}

// FindById implements VideoService
func (db *videoService) FindById(id uint64) entity.Video {
	return db.videoRepository.FindById(id)
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videoRepository.Save(video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videoRepository.FindAll()

}

func (service *videoService) Update(video entity.Video) {
	service.videoRepository.Update(video)
}

func (service *videoService) Delete(videoId uint64) {
	service.videoRepository.Delete(service.FindById(videoId))
}

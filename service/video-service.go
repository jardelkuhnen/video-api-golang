package service

import (
	"github.com/jardelkuhnen/video-api/entity"
	"github.com/jardelkuhnen/video-api/repository"
)

type VideoService interface {
	Save(entity.Video) (entity.Video, error)
	Update(video entity.Video) entity.Video
	FindAll() []entity.Video
	FindById(id uint64) (entity.Video, error)
	Delete(videoId uint64) error
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
func (db *videoService) FindById(id uint64) (entity.Video, error) {
	return db.videoRepository.FindById(id)
}

func (service *videoService) Save(video entity.Video) (entity.Video, error) {
	video, err := service.videoRepository.Save(video)
	if err != nil {
		return entity.Video{}, err
	}
	return video, nil
}

func (service *videoService) FindAll() []entity.Video {
	return service.videoRepository.FindAll()

}

func (service *videoService) Update(video entity.Video) entity.Video {
	service.videoRepository.Update(video)
	return video
}

func (service *videoService) Delete(videoId uint64) error {
	video, err := service.FindById(videoId)
	if err != nil {
		return err
	}

	service.videoRepository.Delete(video)
	return nil
}

package repository

import (
	"fmt"

	"github.com/jardelkuhnen/video-api/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Save(video entity.Video) (entity.Video, error)
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
	FindById(id uint64) (entity.Video, error)
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

// constructor
func NewVideoRepository() VideoRepository {
	db, err := gorm.Open(sqlite.Open("video-api.db"))
	if err != nil {
		panic("Error to connect database")
	}

	db.AutoMigrate(&entity.Video{}, &entity.Person{})
	return &database{
		connection: db,
	}
}

// FindById implements VideoRepository
func (db *database) FindById(id uint64) (entity.Video, error) {
	var video = entity.Video{ID: id}
	result := db.connection.Preload("Author").First(&video, id)
	if result.Error != nil {
		return entity.Video{}, result.Error
	}

	return video, nil
}

// CloseDB implements VideoRepository
func (db *database) CloseDB() {
	sqlDB, err := db.connection.DB()
	if err != nil {
		panic("Failed to get DB")
	}

	sqlDB.Close()
	if err != nil {
		panic("Failed to close connection")
	}
}

// Delete implements VideoRepository
func (db *database) Delete(video entity.Video) {
	db.connection.Delete(video)
}

// FindAll implements VideoRepository
func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Preload("Author").Find(&videos)

	return videos
}

// Save implements VideoRepository
func (db *database) Save(video entity.Video) (entity.Video, error) {
	result := db.connection.Create(&video)
	if result.Error != nil {
		return entity.Video{}, result.Error
	}
	fmt.Println(result)
	return video, nil
}

// Update implements VideoRepository
func (db *database) Update(video entity.Video) {
	db.connection.Save(video)
}

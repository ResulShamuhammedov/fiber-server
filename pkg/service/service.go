package service

import (
	"github.com/ResulShamuhammedov/fiber-server/pkg/models"
	"github.com/ResulShamuhammedov/fiber-server/pkg/repository"
)

type Service struct {
	Video
}

type Video interface {
	AddVideo(video models.Video) (int, error)
	GetVideoByID(id int) (models.Video, error)
	GetVideoByTitle(title string) (models.Video, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Video: NewVideoService(repo.Video),
	}
}

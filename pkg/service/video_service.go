package service

import (
	"github.com/ResulShamuhammedov/fiber-server/pkg/models"
	"github.com/ResulShamuhammedov/fiber-server/pkg/repository"
)

type VideoService struct {
	repo repository.Video
}

func NewVideoService(repo repository.Video) *VideoService {
	return &VideoService{repo: repo}
}

func (s *VideoService) AddVideo(video models.Video) (int, error) {
	return s.repo.AddVideo(video)
}

func (s *VideoService) GetVideoByID(id int) (models.Video, error) {
	return s.repo.GetVideoByID(id)
}

func (s *VideoService) GetVideoByTitle(title string) (models.Video, error) {
	return s.repo.GetVideoByTitle(title)
}

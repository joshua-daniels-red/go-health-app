package service

import (
	"go-health-app/internal/repository"
)

type DataService struct {
	repo *repository.MovieRepository
}

func NewDataService(repo *repository.MovieRepository) *DataService {
	return &DataService{repo: repo}
}

func (s *DataService) GetPaginatedData(page, pageSize int) []repository.Movie {
	all := s.repo.GetAll()
	start := (page - 1) * pageSize
	if start > len(all) {
		return []repository.Movie{}
	}
	end := start + pageSize
	if end > len(all) {
		end = len(all)
	}
	return all[start:end]
}

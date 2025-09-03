package repository

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Movie struct {
	ID         int     `json:"id"`
	Title      string  `json:"movie_title"`
	Genre      string  `json:"genre"`
	IMDbRating float64 `json:"imdb"`
}

type MovieRepository struct {
	movies []Movie
}

func NewMovieRepository(jsonPath string) (*MovieRepository, error) {
	absPath, err := filepath.Abs(jsonPath)
	if err != nil {
		return nil, err
	}
	file, err := os.ReadFile(absPath)
	if err != nil {
		return nil, err
	}
	var movies []Movie
	if err := json.Unmarshal(file, &movies); err != nil {
		return nil, err
	}
	return &MovieRepository{movies: movies}, nil
}

func (r *MovieRepository) GetAll() []Movie {
	return r.movies
}

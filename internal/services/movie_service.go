package services

import (
	"lion_parcel/internal/models"
	"lion_parcel/internal/repositories"
)

type MovieService struct {
	MovieRepo *repositories.MovieRepository
}

func NewMovieService(movieRepo *repositories.MovieRepository) *MovieService {
	return &MovieService{MovieRepo: movieRepo}
}

func (s *MovieService) CreateMovie(movie *models.Movie) error {
	return s.MovieRepo.Create(movie)
}

func (s *MovieService) UpdateMovie(movie *models.Movie) error {
	return s.MovieRepo.Update(movie)
}

func (s *MovieService) ListMovies(page, limit int) ([]models.Movie, error) {
	return s.MovieRepo.FindAll(page, limit)
}

func (s *MovieService) SearchMovies(query string) ([]models.Movie, error) {
	return s.MovieRepo.Search(query)
}

func (s *MovieService) TrackView(movieID uint) error {
	return s.MovieRepo.IncrementViewCount(movieID)
}

func (s *MovieService) GetMostViewedMovies() ([]models.Movie, error) {
	return s.MovieRepo.GetMostViewedMovies()
}
func (s *MovieService) GetMostViewedGenres() ([]string, error) {
	return s.MovieRepo.GetMostViewedGenres()
}

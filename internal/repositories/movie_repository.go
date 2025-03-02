package repositories

import (
	"lion_parcel/internal/models"

	"gorm.io/gorm"
)

type MovieRepository struct {
	DB *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{DB: db}
}

func (r *MovieRepository) Create(movie *models.Movie) error {
	return r.DB.Create(movie).Error
}

func (r *MovieRepository) Update(movie *models.Movie) error {
	return r.DB.Save(movie).Error
}

func (r *MovieRepository) FindAll(page, limit int) ([]models.Movie, error) {
	var movies []models.Movie
	offset := (page - 1) * limit
	err := r.DB.Offset(offset).Limit(limit).Find(&movies).Error
	return movies, err
}

func (r *MovieRepository) Search(query string) ([]models.Movie, error) {
	var movies []models.Movie
	err := r.DB.Where("title LIKE ? OR description LIKE ? OR artists LIKE ? OR genres LIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%").Find(&movies).Error
	return movies, err
}

func (r *MovieRepository) IncrementViewCount(id uint) error {
	return r.DB.Model(&models.Movie{}).Where("id = ?", id).Update("view_count", gorm.Expr("view_count + ?", 1)).Error
}

func (r *MovieRepository) GetMostViewedMovies() ([]models.Movie, error) {
	var movies []models.Movie
	err := r.DB.Order("view_count DESC").Limit(10).Find(&movies).Error
	return movies, err
}
func (r *MovieRepository) GetMostViewedGenres() ([]string, error) {
	var genres []string
	err := r.DB.Model(&models.Movie{}).
		Select("genres").
		Group("genres").
		Order("SUM(view_count) DESC").
		Limit(5).
		Pluck("genres", &genres).Error
	return genres, err
}

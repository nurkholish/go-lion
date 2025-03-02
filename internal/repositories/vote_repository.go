package repositories

import (
	"lion_parcel/internal/models"

	"gorm.io/gorm"
)

type VoteRepository struct {
	DB *gorm.DB
}

func NewVoteRepository(db *gorm.DB) *VoteRepository {
	return &VoteRepository{DB: db}
}

func (r *VoteRepository) Create(vote *models.Vote) error {
	return r.DB.Create(vote).Error
}

func (r *VoteRepository) Delete(userID, movieID uint) error {
	return r.DB.Where("user_id = ? AND movie_id = ?", userID, movieID).Delete(&models.Vote{}).Error
}

func (r *VoteRepository) Exists(userID, movieID uint) bool {
	var count int64
	r.DB.Model(&models.Vote{}).Where("user_id = ? AND movie_id = ?", userID, movieID).Count(&count)
	return count > 0
}

func (r *VoteRepository) GetUserVotes(userID uint) ([]models.Vote, error) {
	var votes []models.Vote
	err := r.DB.Where("user_id = ?", userID).Find(&votes).Error
	return votes, err
}

package repositories

import (
	"lion_parcel/internal/models"

	"gorm.io/gorm"
)

type BlacklistedTokenRepository struct {
	DB *gorm.DB
}

func NewBlacklistedTokenRepository(db *gorm.DB) *BlacklistedTokenRepository {
	return &BlacklistedTokenRepository{DB: db}
}

// Add a token to the blacklist
func (r *BlacklistedTokenRepository) AddToBlacklist(token string) error {
	blacklistedToken := models.BlacklistedToken{Token: token}
	return r.DB.Create(&blacklistedToken).Error
}

// Check if a token is blacklisted
func (r *BlacklistedTokenRepository) IsBlacklisted(token string) bool {
	var count int64
	r.DB.Model(&models.BlacklistedToken{}).Where("token = ?", token).Count(&count)
	return count > 0
}

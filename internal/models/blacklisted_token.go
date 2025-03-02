package models

import "time"

type BlacklistedToken struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Token         string    `json:"token" gorm:"unique;not null"`
	BlacklistedAt time.Time `json:"blacklisted_at" gorm:"default:CURRENT_TIMESTAMP"`
}

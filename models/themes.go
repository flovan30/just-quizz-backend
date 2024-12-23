package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Themes struct {
	ID        uuid.UUID   `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string      `json:"name" gorm:"type:varchar;not null"`
	Icon_url  string      `json:"icon_url" gorm:"type:varchar;not null"`
	CreatedAt time.Time   `json:"-"`
	Questions []Questions `json:"questions" gorm:"foreignKey:Theme_id"`
}

func (theme *Themes) BeforeCreate(tx *gorm.DB) (err error) {
	theme.ID = uuid.New()
	return
}

// struct for input
type CreateThemeInput struct {
	Name     string `json:"name" binding:"required"`
	Icon_url string `json:"icon_url" binding:"required"`
}

type UpdateThemeInput struct {
	Name     string `json:"name" binding:"required"`
	Icon_url string `json:"icon_url" binding:"required"`
}

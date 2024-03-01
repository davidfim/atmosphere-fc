package player

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (Player) TableName() string {
	return "player"
}

package player

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlayerService struct {
	db *gorm.DB
}

func NewPlayerService(db *gorm.DB) *PlayerService {
	return &PlayerService{db: db}
}

func (s *PlayerService) CreatePlayer(player *Player) error {
	player.ID = uuid.New()
	return s.db.Create(player).Error
}

func (s *PlayerService) RemovePlayer(id uuid.UUID) error {
	return s.db.Delete(&Player{}, id).Error
}

func (s *PlayerService) GetPlayers() ([]Player, error) {
	var players []Player
	if err := s.db.Find(&players).Error; err != nil {
		return nil, err
	}
	return players, nil
}

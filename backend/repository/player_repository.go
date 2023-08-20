package repository

import (
	"backend/model"

	"gorm.io/gorm"
)

type IPlayerRepository interface {
	GetPlayerByEmail (player *model.Player, email string) error
	CreatePlayer(player *model.Player) error
}

type playerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository (db *gorm.DB) IPlayerRepository {
	return &playerRepository{db}
}

func (pr *playerRepository) GetPlayerByEmail (player *model.Player, email string) error {
	if err := pr.db.Where("email=?", email).First(player).Error; err != nil {
		return err
	}
	return nil
}

func (pr *playerRepository) CreatePlayer(player *model.Player) error{
	if err := pr.db.Create(player).Error; err != nil {
		return err
	}
	return nil
}
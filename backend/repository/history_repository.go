package repository

import (
	"backend/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IHistoryRepository interface {
	GetAllHistory(history *[]model.History, playerId uint) error
	GetWinById(history *model.History, playerId uint, win uint) error
	GetLoseById(history *model.History, playerId uint, lose uint) error
	CreateHistory(history *model.History) error
	UpdateHistoryByWin(history *model.History, playerId uint, historyId uint) error
	UpdateHistoryByLose(history *model.History, playerId uint, historyId uint) error
}

type historyRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) {
	return &historyRepository{db}
}

func (hr *historyRepository) GetAllHistory(history *[]model.History, playerId uint) error {
	if err := hr.db.Joins("Player").Where("player_id=?", playerId).Order("created_at").Find(history).Error; err != nil {
		return err
	}
	return nil
}

func (hr *historyRepository) GetWinById(history *model.History, playerId uint, win uint) error {
	if err := hr.db.Joins("Player").Where("player_id=?", playerId).First(history, win).Error; err != nil {
		return err
	}
	return nil
}

func (hr *historyRepository) GetLoseById(history *model.History, playerId uint, lose uint) error {
	if err := hr.db.Joins("Player").Where("player_id=?", playerId).First(history, lose).Error; err != nil {
		return err
	}
	return nil
}

func (hr *historyRepository) CreateHistory(history *model.History) error {
	if err := hr.db.Create(history).Error; err != nil {
		return err
	}
	return nil
}

func (hr *historyRepository) UpdateHistoryByWin(history *model.History, playerId uint, historyId uint) error {
	result := hr.db.Model(history).Clauses(clause.Returning{}).Where("id=? AND player_id", historyId, playerId).Update("win", history.Win)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exists.")
	}
	return nil
}

func (hr *historyRepository) UpdateHistoryByLose(history *model.History, playerId uint, historyId uint) error {
	result := hr.db.Model(history).Clauses(clause.Returning{}).Where("id=? AND player_id", historyId, playerId).Update("lose", history.Lose)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exists.")
	}
	return nil
}

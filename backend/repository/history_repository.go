package repository

import (
	"backend/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IHistoryRepository interface {
	GetAllHistory(history *[]model.History, playerId uint) error
	GetPlayerById(history *model.History, playerId uint, historyId uint) error
	CreateHistory(history *model.History) error
	UpdateHistoryByWinAndLose(history *model.History, playerId uint, historyId uint) error
	// UpdateHistoryByMoney(history *model.History, playerId uint, historyId uint) error
}

type historyRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) IHistoryRepository {
	return &historyRepository{db}
}

func (hr *historyRepository) GetAllHistory(history *[]model.History, playerId uint) error {
	if err := hr.db.Joins("Player").Where("player_id=?", playerId).Order("created_at").Find(history).Error; err != nil {
		return err
	}
	return nil
}

func (hr *historyRepository) GetPlayerById(history *model.History, playerId uint, historyId uint) error {
	if err := hr.db.Joins("Player").Where("player_id=?", playerId).First(history, historyId).Error; err != nil {
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

func (hr *historyRepository) UpdateHistoryByWinAndLose(history *model.History, playerId uint, historyId uint) error {
	// Update("win", history.Win).
	// Update("lose", history.Lose).
	// Update("money", history.Money)
	resultInWin := hr.db.Model(history).Clauses(clause.Returning{}).Where("id=? AND player_id=?", historyId, playerId).
		Updates(map[string]interface{}{
			"win": history.Win, "lose": history.Lose, "money": history.Money,
		})
	if resultInWin.Error != nil {
		return resultInWin.Error
	}
	if resultInWin.RowsAffected < 1 {
		return fmt.Errorf("object does not exists.")
	}
	return nil
}

// func (hr *historyRepository) UpdateHistoryByMoney(history *model.History, playerId uint, historyid uint) error {
// 	result := hr.db.Model(history).Clauses(clause.Returning{}).Where("id=? AND player_id", historyid, playerId).Update("Money", history.Money)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	if result.RowsAffected < 1 {
// 		return fmt.Errorf("object does not exists.")
// 	}
// 	return nil
// }

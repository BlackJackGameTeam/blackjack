package usecase

import (
	"backend/model"
	"backend/repository"
)

type IHistoryUsecase interface {
	GetAllHistory(playerId uint) ([]model.HistoryResponse, error)
}

type historyUsecase struct {
	hr repository.IHistoryRepository
}

func NewHistoryUsecase(pr repository.IHistoryRepository) IHistoryUsecase {
	return &historyUsecase{pr}
}

func (hu *historyUsecase) GetAllHistory(playerId uint) ([]model.HistoryResponse, error) {
	history := []model.History{}
	if err := hu.hr.GetAllHistory(&history, playerId); err != nil {
		return nil, err
	}
	resHistory := []model.HistoryResponse{}
	for _, v := range history {
		t := model.HistoryResponse{
			ID: v.ID,
			Win: v.Win,
			Lose: v.Lose,
			Money: v.Money,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resHistory = append(resHistory, t)
	}
	return resHistory, nil
}
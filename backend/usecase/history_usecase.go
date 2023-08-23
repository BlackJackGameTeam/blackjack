package usecase

import (
	"backend/model"
	"backend/repository"
	"backend/validator"
)

type IHistoryUsecase interface {
	GetAllHistory(playerId uint) ([]model.HistoryResponse, error)
	GetPlayerById(playerId uint, historyId uint) (model.HistoryResponse, error)
	CreateHistory(history model.History) (model.HistoryResponse, error)
	UpdateHistoryByWinAndLose(history model.History, playerId uint, historyId uint) (model.HistoryResponse, error)
	// UpdateHistoryByMoney(history model.History, playerId uint, historyId uint) (model.HistoryResponse, error)
}

type historyUsecase struct {
	hr repository.IHistoryRepository
	hv validator.IHistoryValidation
}

func NewHistoryUsecase(hr repository.IHistoryRepository, hv validator.IHistoryValidation) IHistoryUsecase {
	return &historyUsecase{hr, hv}
}

func (hu *historyUsecase) GetAllHistory(playerId uint) ([]model.HistoryResponse, error) {
	history := []model.History{}
	if err := hu.hr.GetAllHistory(&history, playerId); err != nil {
		return nil, err
	}
	resHistory := []model.HistoryResponse{}
	for _, v := range history {
		t := model.HistoryResponse{
			ID:        v.ID,
			Win:       v.Win,
			Lose:      v.Lose,
			Money:     v.Money,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resHistory = append(resHistory, t)
	}
	return resHistory, nil
}

func (hu *historyUsecase) GetPlayerById(playerId uint, historyId uint) (model.HistoryResponse, error) {
	history := model.History{}
	if err := hu.hr.GetPlayerById(&history, playerId, historyId); err != nil {
		return model.HistoryResponse{}, err
	}
	resHistory := model.HistoryResponse{
		ID:        history.ID,
		Win:       history.Win,
		Lose:      history.Lose,
		Money:     history.Money,
		CreatedAt: history.CreatedAt,
		UpdatedAt: history.UpdatedAt,
	}
	return resHistory, nil
}

func (hu *historyUsecase) CreateHistory(history model.History) (model.HistoryResponse, error) {
	if err := hu.hv.HistoryValidate(history); err != nil {
		return model.HistoryResponse{}, err
	}
	if err := hu.hr.CreateHistory(&history); err != nil {
		return model.HistoryResponse{}, err
	}
	resHistory := model.HistoryResponse{
		ID:        history.ID,
		Win:       history.Win,
		Lose:      history.Lose,
		Money:     history.Money,
		CreatedAt: history.CreatedAt,
		UpdatedAt: history.UpdatedAt,
	}
	return resHistory, nil
}

func (hu *historyUsecase) UpdateHistoryByWinAndLose(history model.History, playerId uint, historyId uint) (model.HistoryResponse, error) {
	if err := hu.hv.HistoryValidate(history); err != nil {
		return model.HistoryResponse{}, err
	}
	if err := hu.hr.UpdateHistoryByWinAndLose(&history, playerId, historyId); err != nil {
		return model.HistoryResponse{}, err
	}

	resHistory := model.HistoryResponse{
		ID:        history.ID,
		Win:       history.Win,
		Lose:      history.Lose,
		Money:     history.Money,
		CreatedAt: history.CreatedAt,
		UpdatedAt: history.UpdatedAt,
	}
	return resHistory, nil
}

// func (hu *historyUsecase) UpdateHistoryByMoney(history model.History, playerId uint, historyId uint) (model.HistoryResponse, error) {
// 	if err := hu.hr.UpdateHistoryByMoney(&history, playerId, historyId); err != nil {
// 		return model.HistoryResponse{}, err
// 	}

// 	resHistory := model.HistoryResponse{
// 		ID:        history.ID,
// 		Win:       history.Win,
// 		Lose:      history.Lose,
// 		Money:     history.Money,
// 		CreatedAt: history.CreatedAt,
// 		UpdatedAt: history.UpdatedAt,
// 	}
// 	return resHistory, nil
// }

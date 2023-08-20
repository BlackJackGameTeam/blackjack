package model

import "time"

type History struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Win uint `json:"win"`
	Lose uint `json:"lose"`
	Money uint `json:"money"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Player Player `json:"player" gorm:"foreignKey:PlayerId; constraint:OnDelete:CASCADE"`
	PlayerId uint `json:"player_id" gorm:"not null"`
}

type HistoryResponse struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Win uint `json:"win"`
	Lose uint `json:"lose"`
	Money uint `json:"money"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`	
}
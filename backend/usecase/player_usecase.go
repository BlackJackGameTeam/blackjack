package usecase

import (
	"backend/model"
	"backend/repository"
	"backend/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IPlayerUsecase interface {
	SignUp(player model.Player) (model.PlayerResponse, error)
	Login(player model.Player) (string, error)
}

type playerUsecase struct {
	pr repository.IPlayerRepository
	pv validator.IPlayerValidation
}

func NewPlayerUsecase(pr repository.IPlayerRepository, pv validator.IPlayerValidation) IPlayerUsecase {
	return &playerUsecase{pr, pv}
}

func (pu *playerUsecase) SignUp(player model.Player) (model.PlayerResponse, error) {
	if err := pu.pv.PlayerValidate(player); err != nil {
		return model.PlayerResponse{}, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(player.Password), 10)
	if err != nil {
		return model.PlayerResponse{}, err
	}
	newPlayer := model.Player{Email: player.Email, Password: string(hash)}
	if err := pu.pr.CreatePlayer(&newPlayer); err != nil {
		return model.PlayerResponse{}, err
	}
	resPlayer := model.PlayerResponse{
		ID:    newPlayer.ID,
		Email: newPlayer.Email,
	}
	return resPlayer, nil
}

func (pu *playerUsecase) Login(player model.Player) (string, error) {
	if err := pu.pv.PlayerValidate(player); err != nil {
		return "", err
	}
	storedUser := model.Player{}
	if err := pu.pr.GetPlayerByEmail(&storedUser, player.Email); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(player.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"player_id": storedUser.ID,
		"exp":       time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

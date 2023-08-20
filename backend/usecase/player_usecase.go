package usecase

import (
	"backend/model"
	"backend/repository"
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
}

func NewPlayerUsecase(pr repository.IPlayerRepository) IPlayerUsecase {
	return &playerUsecase{pr}
}

func (pu *playerUsecase) SignUp(player model.Player) (model.PlayerResponse, error) {
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
	storedPlayer := model.Player{}
	if err := pu.pr.GetPlayerByEmail(&storedPlayer, player.Email); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedPlayer.Password), []byte(player.Password))
	if err != nil {
		return "", err
	}
	// jwtトークンの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"player_id": storedPlayer.ID,
		"exp":       time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, err
}

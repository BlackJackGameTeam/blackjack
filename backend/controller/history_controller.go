package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IHistoryController interface {
	GetAllHistory(c echo.Context) error
	GetPlayerById(c echo.Context) error
	CreateHistory(c echo.Context) error
	UpdateHistoryByWinAndLose(c echo.Context) error
	// UpdateHistoryByMoney(c echo.Context) error
}

type historyController struct {
	hu usecase.IHistoryUsecase
}

func NewHistoryController(hu usecase.IHistoryUsecase) IHistoryController {
	return &historyController{hu}
}

func (hc *historyController) GetAllHistory(c echo.Context) error {
	// jwtTokenに型アサーション
	player := c.Get("user").(*jwt.Token)
	claims := player.Claims.(jwt.MapClaims)
	playerId := claims["player_id"]

	historyRes, err := hc.hu.GetAllHistory(uint(playerId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, historyRes)
}

func (hc *historyController) GetPlayerById(c echo.Context) error {
	player := c.Get("user").(*jwt.Token)
	claims := player.Claims.(jwt.MapClaims)
	playerId := claims["player_id"]
	id := c.Param("historyId")
	historyId, _ := strconv.Atoi(id)
	historyRes, err := hc.hu.GetPlayerById(uint(playerId.(float64)), uint(historyId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, historyRes)
}

func (hc *historyController) CreateHistory(c echo.Context) error {
	player := c.Get("user").(*jwt.Token)
	claims := player.Claims.(jwt.MapClaims)
	playerId := claims["player_id"]

	history := model.History{}
	if err := c.Bind(&history); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	history.PlayerId = uint(playerId.(float64))
	historyRes, err := hc.hu.CreateHistory(history)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, historyRes)
}

func (hc *historyController) UpdateHistoryByWinAndLose(c echo.Context) error {
	player := c.Get("user").(*jwt.Token)
	claims := player.Claims.(jwt.MapClaims)
	playerId := claims["player_id"]
	id := c.Param("historyId")
	historyId, _ := strconv.Atoi(id)

	history := model.History{}
	if err := c.Bind(&history); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	historyRes, err := hc.hu.UpdateHistoryByWinAndLose(history, uint(playerId.(float64)), uint(historyId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, historyRes)
}

// func (hc *historyController) UpdateHistoryByMoney(c echo.Context) error {
// 	player := c.Get("user").(*jwt.Token)
// 	claims := player.Claims.(jwt.MapClaims)
// 	playerId := claims["player_id"]
// 	id := c.Param("historyId")
// 	historyId, _ := strconv.Atoi(id)

// 	history := model.History{}
// 	if err := c.Bind(&history); err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}
// 	historyRes, err := hc.hu.UpdateHistoryByMoney(history, uint(playerId.(float64)), uint(historyId))
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, historyRes)
// }

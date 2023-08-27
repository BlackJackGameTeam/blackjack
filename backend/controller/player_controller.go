package controller

import (
	"backend/model"
	"backend/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type IPlayerController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type playerController struct {
	pu usecase.IPlayerUsecase
}

func NewPlayerController(pu usecase.IPlayerUsecase) IPlayerController {
	return &playerController{pu}
}

func (pc *playerController) SignUp(c echo.Context) error {
	player := model.Player{}
	if err := c.Bind(&player); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	playerRes, err := pc.pu.SignUp(player)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, playerRes)
}

func (pc *playerController) Login(c echo.Context) error {
	player := model.Player{}
	if err := c.Bind(&player); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := pc.pu.Login(player)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie) // HTTPResponseに含める
	return c.NoContent(http.StatusOK)
}

func (pc *playerController) Logout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (pc *playerController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}

package router

import (
	"backend/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IPlayerController, hc controller.IHistoryController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken}, // header経由でXCSRFTOKENを取得できる
		AllowMethods:     []string{"GET", "PUT", "POST"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		// CookieSameSite: http.SameSiteDefaultMode,
		CookieMaxAge:   60,
	}))
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
	e.GET("/csrf", uc.CsrfToken)

	h := e.Group("/history")
	// middlewareを追加
	h.Use(echojwt.WithConfig(echojwt.Config{
		// jwtを生成した時と同じSECRETkeyを指定
		SigningKey: []byte(os.Getenv("SECRET")),
		// clientから送られてくるjwtTokenがどこに格納されているかを示す
		TokenLookup: "cookie:token",
	}))

	h.GET("", hc.GetAllHistory)
	h.GET("/:historyId", hc.GetPlayerById)
	h.POST("", hc.CreateHistory)
	h.PUT("/:historyId", hc.UpdateHistoryByWinAndLose)
	// h.PUT("/:playerId", hc.UpdateHistoryByMoney)
	return e
}

package main

import (
	"backend/controller"
	"backend/db"
	"backend/repository"
	"backend/router"
	"backend/usecase"
	"backend/validator"
)

func main() {
	db := db.NewDB()
	playerValidator := validator.NewPlayerValidator()
	historyValidator := validator.NewHistoryValidator()
	playerRepository := repository.NewPlayerRepository(db)
	historyRepository := repository.NewHistoryRepository(db)
	playerUsecase := usecase.NewPlayerUsecase(playerRepository, playerValidator)
	historyUsecase := usecase.NewHistoryUsecase(historyRepository, historyValidator)
	playerController := controller.NewPlayerController(playerUsecase)
	historyController := controller.NewHistoryController(historyUsecase)
	e := router.NewRouter(playerController, historyController)
	e.Logger.Fatal(e.Start(":8080"))
}

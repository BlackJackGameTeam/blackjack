package main

import (
	"backend/controller"
	"backend/db"
	"backend/repository"
	"backend/router"
	"backend/usecase"
)

func main() {
	db := db.NewDB()
	playerRepository := repository.NewPlayerRepository(db)
	playerUsecase := usecase.NewPlayerUsecase(playerRepository)
	playerController := controller.NewPlayerController(playerUsecase)
	e := router.NewRouter(playerController)
	e.Logger.Fatal(e.Start(":8080"))
}

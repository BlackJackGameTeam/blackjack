package main

import (
	"backend/db"
	"backend/model"
	"fmt"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrate")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.Player{}, &model.History{})
}

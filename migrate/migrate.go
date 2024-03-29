package main

import (
	"fmt"
	"kitashiruAPI/db"
	"kitashiruAPI/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Profile{}, &model.Admin{}, &model.Article{}, &model.Favorite{})

}

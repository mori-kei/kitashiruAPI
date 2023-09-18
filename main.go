package main

import (
	"kitashiruAPI/controller"
	"kitashiruAPI/db"
	"kitashiruAPI/repository"
	"kitashiruAPI/router"
	"kitashiruAPI/usecase"
)

func main() {
	db := db.NewDB()
	useRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(useRepository)
	usserController := controller.NewUserController(userUsecase)
	e := router.NewRouter(usserController)
	e.Logger.Fatal(e.Start(":8080"))
}
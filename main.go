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
	profileRepository := repository.NewProfileRepository(db)
	userUsecase := usecase.NewUserUsecase(useRepository)
	profileUsecase := usecase.NewProfileUsecase(profileRepository)
	userController := controller.NewUserController(userUsecase)
	profileController := controller.NewProfileController(profileUsecase)
	e := router.NewRouter(userController,profileController)
	e.Logger.Fatal(e.Start(":8080"))
}
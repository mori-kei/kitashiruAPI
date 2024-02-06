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
	//repository
	useRepository := repository.NewUserRepository(db)
	profileRepository := repository.NewProfileRepository(db)
	adminRepository := repository.NewAdminRepository(db)
	authRepository := repository.NewAuthRepository(db)
	articleRepository := repository.NewArticleRepository(db)
	favoriteRepository := repository.NewFavoriteRepository(db)
	//usecase
	userUsecase := usecase.NewUserUsecase(useRepository)
	profileUsecase := usecase.NewProfileUsecase(profileRepository)
	adminUsecase := usecase.NewAdminUsecase(adminRepository)
	authUsecase := usecase.NewAuthUsecase(authRepository)
	articleUsecase := usecase.NewArticleUsecase(articleRepository, profileRepository)
	favoriteUsecase := usecase.NewFavoriteUsecase(favoriteRepository)
	//controller
	userController := controller.NewUserController(userUsecase)
	profileController := controller.NewProfileController(profileUsecase)
	adminController := controller.NewAdminController(adminUsecase)
	authContoroller := controller.NewAuthController(authUsecase)
	articleController := controller.NewArticleController(articleUsecase)
	favoriteController := controller.NewFavoriteController(favoriteUsecase)
	e := router.NewRouter(userController, profileController, adminController, authContoroller, articleController, favoriteController)
	e.Logger.Fatal(e.Start(":8080"))
}

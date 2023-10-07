package router

import (
	"kitashiruAPI/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController, pc controller.IProfileController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	p := e.Group("/profile")
	p.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	p.POST("", pc.CreateProfile)
	p.GET("", pc.GetProfileByUserId)
	p.PUT("", pc.UpdateProfile)
	p.DELETE("", pc.DeleteProfile)
	return e
}

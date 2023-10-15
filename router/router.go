package router

import (
	"fmt"
	"kitashiruAPI/controller"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, pc controller.IProfileController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,

		//Postmanで動作確認をする場合はsecure属性をfalseにする必要がある
		//SameSiteをnonemodeにしてしまうと自動的にsecrureがonになるためPostmanで動作確認する時はsamasiteをDefaultModeに設定する]
		//↓【通信用】フロントとの通信の際にはコメントアウトを消しPostmanで確認する際はコメントアウトする
		// CookieSameSite: http.SameSiteNoneMode,
		//↓【API開発用】Postmanで確認する際はコメントアウトを消しフロントとの通信の際にはコメントアウトする
		CookieSameSite: http.SameSiteDefaultMode,
		//CookieMaxAgeは秒単位で有効期限を指定できる
		//CookieMaxAge:   60,
	}))
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)
	e.GET("/me", func(c echo.Context) error {
		cookie, err := c.Cookie("token") // "jwt_cookie_name" は実際のCookieの名前に置き換えてください
		if err != nil {
			return err
		}

		tokenString := cookie.Value
		fmt.Println(tokenString)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 適切な署名キーを返すためのコードをここに記述する
			return []byte(os.Getenv("SECRET")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// ペイロードの内容をJSON形式でレスポンスする
			response := map[string]interface{}{
				"user_id": claims["user_id"],
			}
			return c.JSON(http.StatusOK, response)
		} else {
			return echo.ErrUnauthorized
		}
	})
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

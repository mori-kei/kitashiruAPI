package router

import (
	// "fmt"
	"kitashiruAPI/controller"
	"net/http"
	"os"

	// "github.com/dgrijalva/jwt-go"

	"github.com/golang-jwt/jwt"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//maybe-later:責務と関心を分離させる
func AdminOnlyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token not found")
		}

		tokenString := cookie.Value
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 適切な署名キーを返すためのコードをここに記述する
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userType, ok := claims["user_type"].(string)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid user type")
			}

			// userTypeがadminでない場合はエラーを返す
			if userType != "admin" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Admin access only")
			}

			// userTypeがadminの場合は次のハンドラを呼び出す
			return next(c)
		}

		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
	}
}

func NewRouter(uc controller.IUserController, pc controller.IProfileController, ac controller.IAdminController, auc controller.IAuthController, arc controller.IArticleController) *echo.Echo {
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

		// Skipper: func(c echo.Context) bool {
		// 	// 特定のルートを除外する場合などに使う、スキップする条件を設定

		// 	return c.Path() == "/login" || c.Path() == "/signup"
		// },
		//Postmanで動作確認をする場合はsecure属性をfalseにする必要がある
		//SameSiteをnonemodeにしてしまうと自動的にsecrureがonになるためPostmanで動作確認する時はsamasiteをDefaultModeに設定する]
		//↓【通信用】フロントとの通信の際にはコメントアウトを消しPostmanで確認する際はコメントアウトする
		CookieSameSite: http.SameSiteNoneMode,
		CookieSecure:   false,
		//↓【API開発用】Postmanで確認する際はコメントアウトを消しフロントとの通信の際にはコメントアウトする
		// CookieSameSite: http.SameSiteDefaultMode,
		//CookieMaxAgeは秒単位で有効指定できる
		//CookieMaxAge:   60,
	}))

	//admin
	a := e.Group("/admin")
	a.POST("/signup", ac.SignUp)
	a.POST("/login", ac.LogIn)
	a.POST("/logout", ac.Logout)
	// a.POST("/article", arc.CreateArticle, AdminOnlyMiddleware)
	a.POST("/article", arc.CreateArticle)
	//user
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)
	e.GET("/me", uc.GetUserByJwt)
	e.GET("/auth", auc.GetAuthByJwt)

	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Admin Access Granted")
	}, AdminOnlyMiddleware)
	//articlegroup
	ar := e.Group("/articles")
	ar.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	ar.GET("", arc.GetMatchArticles)
	ar.GET("/:articleId", arc.GetArticle)
	ar.GET("/random", arc.GetAllArticleRandom)
	//profilegroup
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

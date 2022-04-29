package server

import (
	"example/anon/trust/deal"
	"example/anon/trust/sinput"
	"example/anon/trust/test"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key          = []byte("super-secret-key")
	sessionStore = sessions.NewCookieStore(key)
)

func Server(archives *deal.Archives) {

	port := os.Getenv("PORT")
	CorsOrigin := "https://trading-frontend-seven.vercel.app"

	if port == "" {
		log.Fatal("$PORT must be set")
		port = "1323"
		CorsOrigin = "http://localhost:3000"
	}

	e := echo.New()

	// todo: review all these
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{CorsOrigin},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	sessionStore.Options = &sessions.Options{
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   3600 * 8, // 8 hours
		HttpOnly: true,
	}

	//e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(session.Middleware(sessionStore))

	//e.Use(Header)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.GET("/api/users", test.GetUsers)
	e.GET("/api/users/:userName", func(c echo.Context) error { return getUser(c, archives) })
	e.POST("/api/users/:userName/edit", func(c echo.Context) error { return editUser(c, archives) })

	e.GET("/api/offers", func(c echo.Context) error { return GetOffers(c, archives) })
	e.POST("/api/offers", func(c echo.Context) error { return GetOffersPost(c, archives) })
	e.GET("/api/offers/:id", func(c echo.Context) error { return GetOffer(c, archives) })

	e.POST("/api/offers/:id/alter", func(c echo.Context) error { return ModOffer(c, archives) })

	e.POST("/offers/new", func(c echo.Context) error {

		if id, ok := getUserIdFromSession(c); ok {
			return sinput.CreateOffer(c, *archives, id.(int64))
		}
		println("POST new offer: no id found in session")
		return c.NoContent(http.StatusForbidden)
	})

	e.POST("/login", func(c echo.Context) error { return loginX(c, archives) })

	//e.GET("/test/login", loginTest)
	//e.GET("/test/logged", loggedTest)
	e.GET("/test/logout", logoutTest)

	e.Logger.Fatal(e.Start(":" + port))

}

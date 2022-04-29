package goserver

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key          = []byte("super-secret-key")
	sessionStore = sessions.NewCookieStore(key)
)

func goserv() {

	sessionStore.Options = &sessions.Options{
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   3600 * 8, // 8 hours
		HttpOnly: true,
	}
	http.HandleFunc("/log2", gosecret)
	http.HandleFunc("/log", gologin)
	http.ListenAndServe(":1323", nil)
}

func gologin(w http.ResponseWriter, r *http.Request) {

	println("GOLOGIN")

	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// cookie will get expired after 1 year
	expires := time.Now().AddDate(1, 0, 0)

	ck := http.Cookie{
		Name:    "JSESSION_ID",
		Value:   "someval",
		Domain:  "localhost",
		Path:    "/",
		Expires: expires,
	}

	// value of cookie
	ck.Value = "value of this awesome cookie"

	// write the cookie to response
	http.SetCookie(w, &ck)

	session, _ := sessionStore.Get(r, "cookie-name")

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Values["dad"] = "son"
	session.Save(r, w)

	w.WriteHeader(200)

}

func gosecret(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	session, _ := sessionStore.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		println("GOSECRET FAIL")
		return
	}

	println("GOSECRET SUCCECSS:", session.Values["dad"])
	// Print secret message
	fmt.Fprintln(w, "GOSECRET: The cake is a lie!")
}

// Header middleware adds a `Server` header to the response.
func Header(next echo.HandlerFunc) echo.HandlerFunc {

	/*
		These must exists:

		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	*/

	return func(c echo.Context) error {
		//c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		c.Response().Header().Set(echo.HeaderAccessControlAllowCredentials, "true")
		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "localhost:3000")
		c.Response().Header().Set(echo.HeaderAccessControlAllowHeaders, "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Origin")
		/*
			c.Response().Header().Set(echo.HeaderAccessControlAllowHeaders, "Content-Type")
			c.Response().Header().Set(echo.HeaderAccessControlAllowHeaders, "Content-Length")
			c.Response().Header().Set(echo.HeaderAccessControlAllowHeaders, "Accept-Encoding")
			c.Response().Header().Set(echo.HeaderAccessControlAllowHeaders, "X-CSRF-Token")
			c.Response().Header().Set(echo.HeaderAccessControlAllowHeaders, "Authorization")
			c.Response().Header().Set(echo.HeaderAccessControlAllowHeaders, "Accept")
		*/
		//c.Response().Header().Set(echo.HeaderAuthorization, "true")
		return next(c)
	}
}

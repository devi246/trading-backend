package server

import (
	"encoding/json"
	"example/anon/trust/deal"
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type LoginForm struct {
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}

func tryLogin(email string, password string, archives *deal.Archives) *deal.Account {

	// test if email and password combo is found.
	account := archives.Accounts.FindByEmail(email)
	if account != nil {
		println("tryLogin email found")
		if account.Password == password {
			println("tryLogin password ok")
			return account
		}
		println("tryLogin password fail")
	}
	println("tryLogin not found")
	return nil
}

func loginX(c echo.Context, archives *deal.Archives) error {

	println("LOGIN X")

	l := new(LoginForm)
	if err := c.Bind(l); err != nil {
		println("LOGIN BAD REQUEST")
		return c.String(http.StatusBadRequest, "NOT OK")
	}

	account := tryLogin(l.Email, l.Password, archives)

	println("LOGIN X try:", l.Email, l.Password)

	if account == nil {
		println("LOGIN X forbidden")
		return c.String(http.StatusForbidden, "NOT FOUND")
	}

	/*
		These have already been added earlier
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
	*/

	sess, _ := session.Get("cookie-name", c)
	sess.Values["authenticated"] = true
	sess.Values["userId"] = account.Id

	println("LOGIN X user id:", account.Id)
	println("LOGIN X userId:", sess.Values["userId"].(int64))

	sess.Save(c.Request(), c.Response())

	println("LOGIN X isNew:", sess.IsNew, " Auth:", sess.Values["authenticated"].(bool))
	println("logged in with:", l.Email)

	s, _ := json.Marshal(struct{ Name string }{Name: account.Name})

	//return c.String(http.StatusOK, "OK")
	return c.JSON(http.StatusOK, string(s))
}

func loginTest(c echo.Context) error {

	sess, _ := session.Get("cookie-name", c)

	sess.Values["foo"] = "bar"
	sess.Values["authenticated"] = true
	sess.Save(c.Request(), c.Response())

	//println("LOGING", sess.Values["authenticated"].(bool))
	println("\nLOGIN TEST is new:", sess.IsNew)

	//return c.String(http.StatusOK, "OK")
	return c.String(http.StatusOK, "OK")
}

func loggedTest(c echo.Context) error {

	println("\nSESSION TEST")

	// SESSION TEST
	sess, _ := session.Get("cookie-name", c)

	fmt.Printf("SESS: %+v\n", sess)

	if auth, ok := sess.Values["authenticated"].(bool); !ok || !auth {
		//http.Error(w, "Forbidden", http.StatusForbidden)
		fmt.Printf(">>NOT LOGGED IN %+v\n", sess.Values)
	} else {
		fmt.Printf(">>IS LOGGED IN %+v\n", sess.Values)
	}
	// SESSION TEST

	//return c.String(http.StatusOK, "OK")
	return c.String(http.StatusOK, "OK")
}

func logoutTest(c echo.Context) error {
	sess, _ := session.Get("cookie-name", c)

	if auth, ok := sess.Values["authenticated"].(bool); !ok || !auth {
		//http.Error(w, "Forbidden", http.StatusForbidden)
		fmt.Printf("not logged %+v\n", sess.Values)
		println("\nLOGOUT: already logged out?")
	} else {
		sess.Values["authenticated"] = false
		sess.Save(c.Request(), c.Response())
		println("\nLOGOUT: logged out")
	}

	return c.String(http.StatusOK, "OK")
}

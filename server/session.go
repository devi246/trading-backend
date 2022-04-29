package server

import (
	"example/anon/trust/deal"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func getUserFromSession(sess *sessions.Session, archives *deal.Archives) *deal.Account {
	if userId, ok := sess.Values["userId"]; ok {
		account := archives.Accounts.FindById(userId.(int64))
		return account
	} else {
		println("Could ")
		return nil
	}
}

func getUserIdFromSession(c echo.Context) (interface{}, bool) {
	sess, _ := session.Get("cookie-name", c)
	//fmt.Printf("\nSESSS %+v\n", sess)
	id, ok := sess.Values["userId"]
	return id, ok
}

package server

import (
	"example/anon/trust/deal"
	"example/anon/trust/sinput"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func getUser(c echo.Context, archives *deal.Archives) error {
	userName := c.Param("userName")
	user := archives.Accounts.FindByName(userName)
	if user == nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, user.AsJsonString())
}

func editUser(c echo.Context, archives *deal.Archives) error {

	sess, _ := session.Get("cookie-name", c)
	if auth, ok := sess.Values["authenticated"].(bool); !ok || !auth {
		println("editUser: NOT AUTHENTICATED")
		return c.NoContent(http.StatusForbidden)
	}

	user := getUserFromSession(sess, archives)
	if user != nil {

		userName := c.Param("userName")

		if user.Name == userName {

			return sinput.EditUser(c, user)

		} else {
			println("editUser: tried to edit stranger")
			return c.String(http.StatusForbidden, "can only edit self")
		}

	} else {
		println("ALTER OFFER FAIL: NO USER ID IN SESSION")
		return c.NoContent(http.StatusInternalServerError)
	}
}

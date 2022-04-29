package sinput

import (
	"example/anon/trust/deal"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserEdit struct {
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
}

func EditUser(c echo.Context, user *deal.Account) error {

	alterForm := new(UserEdit)
	if err := c.Bind(alterForm); err != nil {
		fmt.Printf("\n EditUser ERROR: %+v\n", alterForm)
		println("EditUser ERROR:", err.Error())
		return c.String(http.StatusBadRequest, "NOT OK")
	}

	fmt.Printf("\n EditUser got data: %+v\n", alterForm)

	makeUserChanges(user, alterForm)

	return c.String(http.StatusOK, "ok")
}

func makeUserChanges(user *deal.Account, form *UserEdit) {

	if form.Email != "" {
		user.Email = form.Email
	}
	if form.Phone != "" {
		user.Phone = form.Phone
	}

}

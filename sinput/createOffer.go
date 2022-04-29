package sinput

import (
	"errors"
	"example/anon/trust/deal"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type NewOffer struct {
	//Name  string `json:"name" form:"name" query:"name"`
	//Email string `json:"email" form:"email" query:"email"`

	Name        string
	Description string
	Way         string

	Price    string
	Currency string
}

func CreateOffer(c echo.Context, arch deal.Archives, id int64) error {

	oForm := new(NewOffer)
	if err := c.Bind(oForm); err != nil {
		fmt.Printf("\n CREATE ERROR: %+v\n", oForm)
		println("Err:", err.Error())
		return c.String(http.StatusBadRequest, "NOT OK")
	}
	fmt.Printf("\nCREATE OK: %+v\n", oForm)

	fmt.Printf("\nCREATE OFFER got data: %+v\n", oForm)

	account := arch.Accounts.FindById(id)

	var o = deal.NewOffer(account, nil, "offername")

	err := initOffer(*oForm, o)
	if err != nil {
		return c.String(http.StatusConflict, err.Error())
	}

	arch.AddPublicOffer(o)

	oid := strconv.FormatInt(o.Id, 10)

	return c.String(http.StatusCreated, oid)
}

func initOffer(form NewOffer, offer *deal.Offer) error {

	price, err := strconv.ParseFloat(form.Price, 64)
	if err != nil {
		return errors.New("price must be a float number")
	}
	offer.Price = float64(price)

	offer.Name = form.Name
	offer.Content.Description = form.Description

	offer.Currency = form.Currency

	return nil
}

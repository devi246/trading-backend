package test

import (
	"example/anon/trust/deal"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	users = []string{"Joe", "Veer", "Zion"}
)

var (
	offers = []string{"Joe", "Veer", "Zion"}
)

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func GetOffers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func InitData(archives *deal.Archives) {

	joe := archives.Accounts.AddAccount("testjoe", "testjoe@example.com", "pass")
	mia := archives.Accounts.AddAccount("testmia", "testmia@example.com", "pass")
	archives.Accounts.AddAccount("testbob", "testbob@example.com", "pass")
	archives.Accounts.AddAccount("testjim", "testjim@example.com", "pass")

	o := archives.CreateAndAddPublicOffer(joe, nil, "Car of Fortune")
	o.Content = deal.Content{ContentType: deal.CtMoney, Amount: 100, Currency: "USD", Description: "This car brings you good luck!", Image: "classic"}
	o.Price = 10
	o.Currency = "Euro"

	o2 := archives.CreateAndAddPublicOffer(mia, nil, "Car of Speed")
	o2.Content = deal.Content{ContentType: deal.CtMoney, Amount: 100, Currency: "USD", Description: "This car is fast. Vroom!", Image: "dodge"}
	o2.Price = 10
	o2.Currency = "Euro"
}

package server

import (
	"example/anon/trust/deal"
	"example/anon/trust/deal/jeison"
	"example/anon/trust/sinput"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func ModOffer(c echo.Context, archives *deal.Archives) error {

	sess, _ := session.Get("cookie-name", c)
	if auth, ok := sess.Values["authenticated"].(bool); !ok || !auth {
		println("Mod Offer: NOT AUTHENTICATED")
		return c.NoContent(http.StatusForbidden)
	}

	id := c.Param("id")

	println("Alter Offer got offer id:", id)

	offerId, err := strconv.Atoi(id)
	if err != nil {
		println("ModOffer: got improper offerId:", id)
		return c.NoContent(http.StatusBadRequest)
	}

	userId, ok := getUserIdFromSession(c)
	if ok == true {

		account := archives.Accounts.FindById(userId.(int64))
		if account == nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		return sinput.AlterOffer(c, *archives, int64(offerId), account)
	} else {
		println("ALTER OFFER FAIL: NO USER ID IN SESSION")
		return c.NoContent(http.StatusForbidden)
	}
}

func GetOffer(c echo.Context, archives *deal.Archives) error {
	id := c.Param("id")

	offerId, err := strconv.Atoi(id)
	if err != nil {
		println("GetOffer: got improper offerId:", id)
		return c.NoContent(http.StatusBadRequest)
	}

	//println("GetOffer id:", offerId)
	of := jeison.GetOffer(archives, int64(offerId))

	if of == nil {
		println("GetOffer: could not find with id:", offerId)
		return c.JSON(http.StatusNotFound, "")
	} else {
		offer := string(of)
		//println("GetOffer of:", offer)
		return c.JSON(http.StatusOK, offer)
	}
}

func GetOffers(c echo.Context, archives *deal.Archives) error {

	var max int64 = 0
	var term string = ""

	println("GetOffers")

	// SESSION TEST
	sess, _ := session.Get("cookie-name", c)

	fmt.Printf("SESS: %+v\n", sess)

	if auth, ok := sess.Values["authenticated"].(bool); !ok || !auth {
		//http.Error(w, "Forbidden", http.StatusForbidden)
		fmt.Printf("not logged %+v\n", sess.Values)
	} else {
		fmt.Printf("logged %+v\n", sess.Values)
	}
	// SESSION TEST

	err := echo.QueryParamsBinder(c).
		Int64("max", &max).
		String("term", &term).
		BindError() // returns first binding error

	if err != nil {
		println("trust GetOffers error: %v", err)
		return c.JSON(http.StatusOK, []string{"GetOffers: CLIENT did not give a parameter: max"})
	}

	o := string(jeison.GetOffers(archives, term, 10))

	//println("return: ", o)
	println("max: ", max, "term:", term)

	return c.JSON(http.StatusOK, o)
}

func GetOffersPost(c echo.Context, archives *deal.Archives) error {

	return sinput.GetOffersPOST(c, *archives)
}

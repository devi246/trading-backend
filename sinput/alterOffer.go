package sinput

import (
	"errors"
	"example/anon/trust/deal"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AlterOffer(c echo.Context, arch deal.Archives, offerId int64, account *deal.Account) error {

	alterForm := new(deal.AlterOfferJson)
	if err := c.Bind(alterForm); err != nil {
		fmt.Printf("\nALTER ERROR: %+v\n", alterForm)
		println("ALTER ERROR:", err.Error())
		return c.String(http.StatusBadRequest, "NOT OK")
	}

	fmt.Printf("\nALTER OFFER got data: %+v\n", alterForm)

	offer := arch.GetOfferById(offerId)

	if offer == nil {
		println("ALTER OFFER could not find offer with id:", offerId)
		return c.NoContent(http.StatusNotFound)
	}

	if err := makeChanges(alterForm, account, offer); err == nil {
		println("ALTER OFFER made changes. return ok")
		return c.String(http.StatusOK, "ok")
	} else {
		println("ALTER OFFER error. return err:", err.Error())
		return c.String(http.StatusForbidden, err.Error())
	}
}

func makeChanges(alterForm *deal.AlterOfferJson, who *deal.Account, offer *deal.Offer) error {

	// A FINISHED OFFER CANT BE CHANGED ANYMORE

	if offer.State == "finished" {
		return errors.New("Can't edit a finished offer")
	}

	// BUY REQUEST
	if alterForm.Buy == true {
		if offer.Giver == who {
			return errors.New("Can't buy from self")
		}
		offer.Receiver = who
		offer.State = "finished"
		println("makeChanges: BOUGHT IT")
		return nil
	}

	// ALTER REQUEST
	if offer.State == "declared" {
		// make changes immediately

		member := false

		if offer.Giver == who {
			member = true
		}
		if offer.Receiver == who {
			member = true
		}

		if member {
			println("Is a member: making changes")
			makeChangesForReal(alterForm, offer)
			return nil
		} else {
			println("Not a member")
			return errors.New("not a member")
		}

	} else {
		// add it to proposals
		// todo
		return nil
	}

}

func makeChangesForReal(alters *deal.AlterOfferJson, offer *deal.Offer) {

	if alters.Name != "" {
		println("changed name")
		offer.Name = alters.Name
	}
	if alters.Price != "" {

		if price, err := strconv.Atoi(alters.Price); err == nil {
			println("changed price")
			offer.Price = float64(price)
		} else {
			println("could not change price:", err.Error())
		}
	}
}

type AlterOfferJsonEXAMPLE struct {
	Name  string
	Price int64
}

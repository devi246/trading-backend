package sinput

import (
	"encoding/json"
	"example/anon/trust/deal"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetOffersPOST(c echo.Context, arch deal.Archives) error {

	form := new(OffersPOSTForm)
	if err := c.Bind(form); err != nil {
		fmt.Printf("\n getOffersPOST ERROR: %+v\n", form)
		println("getOffersPOST ERROR:", err.Error())
		return c.String(http.StatusBadRequest, "NOT OK")
	}

	fmt.Printf("\n getOffersPOST FORM: %+v\n", form)

	offers := searchOffers(arch, *form)

	if form.Sort == "hiprice" {
		println("sorting hi")

		sort.SliceStable(offers, func(i int, j int) bool {
			if offers[i] == nil {
				return true
			}
			if offers[j] == nil {
				return false
			}
			return offers[i].Price < offers[j].Price
		})

	} else if form.Sort == "loprice" {
		println("sorting lo")

		sort.SliceStable(offers, func(i int, j int) bool {
			if offers[j] == nil {
				return true
			}
			if offers[i] == nil {
				return false
			}
			return offers[i].Price > offers[j].Price
		})

	}

	js, err := json.Marshal(offers)
	if err != nil {
		println("JSON MARSHAL ERROR: ", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, string(js))
}

func searchOffers(arch deal.Archives, form OffersPOSTForm) []*deal.Offer {
	maxCount := 10
	maxLoop := 50

	slice := make([]*deal.Offer, maxCount, maxCount)

	i := 0
	j := 0

	for k, v := range arch.PublicOffers {

		if strings.Contains(k, form.Search) {

			if filter(form, v) {

				slice[j] = v
				j++
				if j >= maxCount {
					break
				}
			}
		}

		i++
		if i >= maxLoop {
			break
		}
	}

	slice = slice[:cap(slice)]

	return slice
}

// { search: "", giver: "", receiver: "", buy: false, sell: false, mine: false, finished: false, max: 10}
type OffersPOSTForm struct {
	Search   string `json:"search,omitempty"`
	Giver    string `json:"giver,omitempty"`
	Receiver string `json:"receiver,omitempty"`
	Way      string `json:"way,omitempty"`
	Max      int64  `json:"max,omitempty"`
	Sort     string `json:"sort,omitempty"`
	Finished bool   `json:"finished,omitempty"`
	User     string `json:"user,omitempty"`
}

func filter(form OffersPOSTForm, offer *deal.Offer) bool {

	if form.Finished == true {

		if offer.State != "finished" {
			return false
		}

		if offer.Giver != nil && offer.Giver.Name == form.User {
			return true
		} else if offer.Receiver != nil && offer.Receiver.Name == form.User {
			return true
		}
		return false

	}

	if form.Way == "give" {
		if offer.Giver != nil && offer.Receiver == nil {

			if form.User != "" {
				if offer.Giver.Name == form.User {
					return true
				}
				return false
			}

			return true
		}
	} else if form.Way == "receive" {
		if offer.Receiver != nil && offer.Giver == nil {

			if form.User != "" {
				if offer.Receiver.Name == form.User {
					return true
				}
				return false
			}

			return true
		}
	}

	return false
}

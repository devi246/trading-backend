package deal

import (
	"strings"
)

type Archives struct {
	Accounts       *AccountArchive
	Trades         TradeArchive
	PublicOffers   map[string]*Offer
	OfferProposals map[string]*AlterOfferJson
}

func InitArchives() Archives {
	a := Archives{}

	a.Accounts = NewAccountArchive()

	a.PublicOffers = make(map[string]*Offer)

	a.Trades = TradeArchive{
		Open_trades:   make(map[string]Trade),
		Closed_trades: make(map[string]Trade),
	}

	return a
}

func (a Archives) AddPublicOffer(o *Offer) {
	a.PublicOffers[o.Name] = o
	o.Id = int64(len(a.PublicOffers))
}

func (a Archives) CreateAndAddPublicOffer(who *Account, recipient *Account, offerName string) *Offer {
	var o = NewOffer(who, recipient, offerName)
	a.PublicOffers[o.Name] = o
	o.Id = int64(len(a.PublicOffers))
	return o
}

func (a Archives) GetOfferById(id int64) *Offer {

	for _, v := range a.PublicOffers {
		if v.Id == id {
			return v
		}
	}

	return nil
}

func (a Archives) GetOffer(term string) *Offer {
	o, found := a.PublicOffers[term]

	//fmt.Printf("GetOffer %+v", o)

	if found {
		return o
	}
	return nil
}

func (a Archives) GetPublicOffers(term string, maxCount int, maxLoop int) []*Offer {
	slice := make([]*Offer, maxCount, maxCount) //

	i := 0
	j := 0

	for k, v := range a.PublicOffers {

		if strings.Contains(k, term) {
			slice[j] = v
			j++
			if j >= maxCount {
				break
			}
		}

		i++
		if i >= maxLoop {
			break
		}
	}

	//println("GetPublicOffers: ", i, "", len(slice), cap(slice))
	slice = slice[:cap(slice)]

	return slice
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

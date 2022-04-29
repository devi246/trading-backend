package deal

import (
	"time"
)

type Offer struct {
	Id   int64
	Name string

	Giver     *Account
	Receiver  *Account
	Deliverer *Account
	Insurer   *Account

	State string // declared, finished

	GiverPromise    bool
	ReceiverPromise bool

	ReceiverCanTrade bool

	When time.Time

	Price    float64
	Currency string

	OpenToBids bool

	Content Content
}

type Content struct {
	ContentType ContentType
	Amount      int64
	Currency    string
	Description string
	Image       string
}

func NewOffer(offerer *Account, recipient *Account, offerName string) *Offer {
	o := Offer{Giver: offerer, Receiver: recipient, Name: offerName, State: "declared"}
	return &o
}

func (o Offer) NewDeliverer(who *Account) {
	o.Deliverer = who
}

func (o Dispute) BeneficiaryReportsCompleted() {

}

type OfferState int

const (
	OfferDeclared OfferState = iota
	OfferStarted
	OfferPaused
	OfferFinished
)

type ContentType int

const (
	CtMoney ContentType = iota
	CtHouseMoney
	CtGoods
)

type AlterOfferJson struct {
	Name     string `json:",omitempty"`
	Price    string `json:",omitempty"`
	Giver    string `json:",omitempty"`
	Receiver string `json:",omitempty"`
	Buy      bool   `json:",omitempty"`
}

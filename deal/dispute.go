package deal

type Dispute struct {
	// a deal involves two people; it can hold promises both ways
	Offer   *Offer
	Starter *Account
}

func NewDispute(deal *Offer) *Dispute {
	d := Dispute{}
	return &d
}

package deal

type Delivery struct {
	offering *Offer

	from        string
	destination string
}

func NewDelivery(name string) *Delivery {

	d := Delivery{}

	return &d
}

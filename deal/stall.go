package deal

type Stall struct {
	Name string

	Offers map[string]*Offer
	Trades map[string]*Trade
}

func NewStall(name string) *Stall {

	s := Stall{Name: name}

	s.Offers = make(map[string]*Offer)
	s.Trades = make(map[string]*Trade)

	return &s
}

package deal

/*
TRADE PATH

	first delivery completed
	second delivery completed

	then trade is completed. both get a tick on "completed trades"

*/

type TradeArchive struct {
	Open_trades   map[string]Trade
	Closed_trades map[string]Trade
}

type Deal struct {
	creator *Account
	other   *Account
}

type Trade struct {
	// this could be a loan or purchase
	Deal

	// open trade: anyone can accept anytime.
	open bool

	First  *Offer
	Second *Offer

	Ordered           bool
	first_is_optional bool

	first_accepted  bool
	second_accepted bool

	first_delivered  int
	second_delivered int

	date_started int
	cancel_time  int

	state TradeState
}

func NewTrade(who *Account) *Trade {
	t := Trade{}
	t.Deal.creator = who
	return &t
}

func (t Trade) AcceptTrade() {}

func (t Trade) CancelTrade() {

}

func (t Trade) ListTrades(who *Account) {

}

func (t Trade) DropTrade(friendly bool) {
	t.state = TradeCompleted
}

func (t Trade) OpenTrade(who *Account) {
	t.state = TradeStarted
	t.second_accepted = true
	t.Deal.other = who
}

func (t Trade) FirstDeliveryCompleted() {
	t.Deal.creator.SendMessage("first delivery completed")
	t.first_delivered = 100
	if t.second_delivered == 100 {
		t.TradeCompleted()
	}
}

func (t Trade) SecondDeliveryCompleted() {
	t.Deal.creator.SendMessage("second delivery completed")
	t.second_delivered = 100
	if t.first_delivered == 100 {
		t.TradeCompleted()
	}
}

func (t Trade) TradeCompleted() {
	t.state = TradeCompleted
	t.Deal.creator.Stats.CompletedTrades += 1
	t.Deal.other.Stats.CompletedTrades += 1
}

type TradeState int

const (
	TradeWaiting TradeState = iota
	TradeStarted
	TradeCompleted
)

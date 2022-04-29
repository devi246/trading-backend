package deal

type Insurer struct {
	promisesInsured       int
	promisesInsurancePaid int
}

type InsurancePool struct {
}

type Insurance struct {
	insurer *Account
	value   int
	kind    InsuranceType
}

func newInsurance(who *Account, value int) *Insurance {
	i := Insurance{insurer: who, value: value}
	return &i
}

func InsurePromise() {
}

func RequestInsurance(kind InsuranceType, promise *Offer) {
}

func OfferInsurance(promise *Offer) {
}

type InsuranceType int

const (
	CommonPool InsuranceType = iota
	Guaranteed
)

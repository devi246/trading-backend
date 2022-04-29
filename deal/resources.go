package deal

type HouseMoney struct {
	Amount   int
	Currency string
}

func NewMoney() *HouseMoney {
	m := HouseMoney{}
	return &m
}

type UserResource struct {
	Name         string
	Amount       int
	ResourceType UserResourceType
}

type UserResourceType int

const (
	UserMoney = iota
	UserGoods
	UserOther
)

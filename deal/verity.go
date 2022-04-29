package deal

type Verity struct {
	value    string
	vtype    VerityType
	verified bool
}

type VerityType int

const (
	Phone VerityType = iota
	Email
	Website
)

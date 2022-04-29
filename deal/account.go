package deal

import (
	"encoding/json"
	"fmt"
	"strings"
)

type AccountArchive struct {
	All map[string]*Account
	// Temporary Accounts:
	Days map[int]*ArchiveDay
}

func NewAccountArchive() *AccountArchive {
	a := AccountArchive{}
	a.All = make(map[string]*Account)
	return &a
}

func (archive AccountArchive) SearchAccounts(name string, maxCount int) []*Account {

	results := make([]*Account, 0) // len(results) = maxCount

	elem, ok := archive.All[name]
	if ok {
		results = append(results, elem)
	}

	count := 0
	for key, element := range archive.All {
		fmt.Println("Key:", key, "=>", "Element:", element)

		if strings.Contains(key, name) {
			if key != name { // we dont want exact match here, because it was already added above
				results = append(results, element)
				count++
				if count == maxCount {
					break
				}
			}
		}
	}

	return results
}

func (archive AccountArchive) FindByName(name string) *Account {

	account, exists := archive.All[name]

	if !exists {
		return nil
	}

	return account
}

func (archive AccountArchive) FindByEmail(email string) *Account {

	for _, element := range archive.All {
		println(element.Email, "vs", email)
		if element.Email == email {
			return element
		}
	}

	return nil
}

func (archive AccountArchive) FindById(id int64) *Account {

	for _, element := range archive.All {
		if element.Id == id {
			return element
		}
	}

	return nil
}

func (a Account) AsJsonString() string {

	js, err := json.Marshal(a)

	if err != nil {
		println("JSON MARSHAL ERROR: ", err)
	}
	return string(js)
}

type Account struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"` // todo: must be hashed

	Active     bool // An active account is used by the user. Once fully deactivated, the user cant access it anymore.
	TimeToLive int64
	Created    int64

	Stats AccountStats `json:"-"`

	Likes int

	Verities   map[VerityType]*Verity `json:"-"`
	Insurances map[string]*Insurance  `json:"-"`

	// Inactive
	StallIn   *Stall            `json:"-"`
	StallsOut map[string]*Stall `json:"-"`

	// Active
	OffersIn  []*Offer `json:"-"`
	OffersOut []*Offer `json:"-"`

	FakeMoney int64
	Moneys    map[string]HouseMoney   `json:"-"`
	Resources map[string]UserResource `json:"-"`
	Deposits  map[string]HouseMoney   `json:"-"`
}

type Deposits struct {
	moneys []HouseMoney
}

type AccountStats struct {
	Delivered_promises int
	Failed_promises    int
	CompletedTrades    int
}

func newAccount(name string) *Account {

	a := Account{Name: name}

	a.StallsOut = make(map[string]*Stall)
	a.StallsOut["main"] = NewStall("main")

	a.StallIn = NewStall("main")

	a.OffersOut = make([]*Offer, 0)

	return &a
}

func (a AccountArchive) AddAccount(name string, email string, password string) *Account {

	// test if name exists in archive
	_, exists := a.All[name]

	if !exists {
		ac := newAccount(name)
		ac.Email = email
		ac.Password = password
		ac.Id = int64(len(a.All))
		a.All[name] = ac
		return ac
	}

	return nil
}

func (a Account) AddVerity(vtype VerityType) {
	v := Verity{}
	a.Verities[vtype] = &v
}

func (a Account) Deactivate() {}

func (a Account) Cleanse() {
	// Remove personal data from account
}

func (a Account) RemoveHistory(since int) {
}

func (a Account) SendMessage(msg string) {
	println("Account: SendMessage: ", a.Name, msg)
}

func (a AccountArchive) AddTemporaryAccount(account *Account, days int) {

	current_day := 0

	day := current_day + days

	_, exists := a.Days[day]

	if !exists {
		new_ad := ArchiveDay{
			Accounts: make(map[string]Account),
		}
		a.Days[day] = &new_ad
	}

	ad := a.Days[day]

	account.TimeToLive = int64(days)

	ad.Accounts[account.Name] = *account

}

type ArchiveDay struct {
	Accounts map[string]Account
}

func (a AccountArchive) removeDay(day int) {
	delete(a.Days, day)
}

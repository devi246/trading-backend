package test

import (
	"example/anon/trust/deal"
	"fmt"
)

var archives *deal.Archives

func All(a *deal.Archives) {

	archives = a

	var joe = archives.Accounts.AddAccount("joe", "joe@example.com", "pass")
	var mia = archives.Accounts.AddAccount("mia", "mia@example.com", "pass")
	var _ = archives.Accounts.AddAccount("bob", "bob@example.com", "pass")

	createOffers(joe, mia)

	Gifts(joe, mia)

	test_basic_trade(joe, mia)
}

func createOffers(joe *deal.Account, mia *deal.Account) {

	archives.CreateAndAddPublicOffer(joe, nil, "o1")
	archives.CreateAndAddPublicOffer(joe, nil, "o2")
	archives.CreateAndAddPublicOffer(joe, nil, "o3")

	archives.CreateAndAddPublicOffer(mia, nil, "o1")

	//o := jeison.GetOffers(archives, 5)

	//fmt.Printf("here is all: %+v\n", o)

	//println("end all")
	//println("all as tring: ", string(o))
}

func Gifts(joe *deal.Account, mia *deal.Account) {

	// jim creates a gift
	var gift = deal.NewOffer(joe, nil, "giftTo?")
	joe.OffersOut = append(joe.OffersOut, gift)

	// jim offers it to mia
	mia.OffersIn = append(joe.OffersOut, gift)
}

func test_account_stuff() {
	// create a profile or two
	var jim = archives.Accounts.AddAccount("jim", "jim@example.com", "pass")

	jim.RemoveHistory(1)
	jim.Cleanse()
	jim.Deactivate()
}

func test_basic_trade(joe *deal.Account, mia *deal.Account) {

	var trade = deal.NewTrade(joe)
	// add a promise
	trade.Second = deal.NewOffer(mia, nil, "")

	// add a price
	trade.First = deal.NewOffer(joe, nil, "")

	trade.Ordered = true // In an ORDERED trade, the "first" must complete the transaction before the "second begins"

	// joe puts it in main stall
	joe.StallsOut["main"].Trades["some"] = trade

	// mia searches accounts
	results := archives.Accounts.SearchAccounts("", 10)

	for _, v := range results {
		fmt.Printf("Results: %v \n", v.Name)
	}

	// mia views joes stall
	//fmt.Printf("%+v\n", joe.Stalls["main"].Trades["some"])

	// mia makes an order
	trade.OpenTrade(mia)

	// mia delivers
	trade.FirstDeliveryCompleted()

	// joe delivers
	trade.SecondDeliveryCompleted()
}

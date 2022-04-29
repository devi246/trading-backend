package jeison

import (
	"encoding/json"
	"example/anon/trust/deal"
	"fmt"
)

type Command struct {
	Name string `json:"command"`
	Args string `json:"args"`
}

type NewPromise struct {
	Promiser bool   `json:"promiser"`
	PType    string `json:"type"`
}

func TestJson() {
	test_json_to_SERVER()
}

func test_json_to_SERVER() {

	println("TEST JSON")

	input := `{"command":"json", "args":"yada"}`

	var command Command
	json.Unmarshal([]byte(input), &command)
	fmt.Printf("Name: %s, Args: %s", command.Name, command.Args)

	var newPromise NewPromise
	input = `{"promiser":true, "type":"money"}`
	json.Unmarshal([]byte(input), &newPromise)
	fmt.Printf("\nPromiser: %t, PType: %s", newPromise.Promiser, newPromise.PType)

	var anotherPromise NewPromise
	input = `{"ghfthrth"}`
	json.Unmarshal([]byte(input), &anotherPromise)
	fmt.Printf("\nPromiser: %t, PType: %s", anotherPromise.Promiser, anotherPromise.PType)

	Cmd_NewPromise(newPromise)
}

func GetOffers(a *deal.Archives, term string, max int) []byte {

	var js []byte

	offers := a.GetPublicOffers(term, max, 20)

	js, err := json.Marshal(offers)

	if err != nil {
		println("JSON MARSHAL ERROR: ", err)
	}
	return js
}

func GetOffer(a *deal.Archives, id int64) []byte {

	o := a.GetOfferById(id)

	js, err := json.Marshal(o)

	if err != nil {
		println("JSON MARSHAL ERROR: ", err)
		return make([]byte, 0)
	}
	return js
}

func Cmd_NewPromise(command NewPromise) {

}

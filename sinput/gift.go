package sinput

type Account struct {
	XType      string `json:"type"`
	Url        string `json:"url"`
	Identifier string `json:"identifier"`
}

type Person struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Phone   string  `json:"phone"`
	Account Account `json:"account"`
}

type NewGift struct {
	Offerer   Person `json:"offerer"`
	Recipient Person `json:"recipient"`

	XType string `json:"type"`
	Value string `json:"value"`

	Timeout int `json:"timeout"`
}

/*
GIFT
	A UNIQUE ID IS CREATED FOR THE GIFT -- when no other contact points exist
		GiftArchive
		if one is a known member: add it to their lists

	RECIPIENT INFO
		if nothing: you get a code back, that you can otherwise give to recipient

		COM : email (phone, or other way to message recipient SECURELY)
		ADDRESS: street, money account nr

	TIMEOUT
		if a reciepient doesnt collect the gift. then the gift is returned to sender.
			if sender also does not accept it, it goes to house.

	CONTENT
		router gift: house takes care of sending bitcoin or other currency. takes care of delivery.
		non-routed: it is entirely between offerer and recipient. House does no tracking.

	DELIVERY
		both can raise an error.


OREF-Common
	a reference to a Gift, Promise or Trade

	OREF-Sender		-- only sender knows this
	OREF-Reciever	-- only receiver knows this

	PUBLIC-COMMON-SECRET
*/

/*
FLOW

	"I know only recipients EMAIL. I dont want to know her ADDRESS or BANK ACCOUNT NR"
		-it will cost you 1$ to send an EMAIL with an OFFERING

		"I want to be ANONYMOUS to HOUSE or RECIPIENT"

	-YOU MUST KNOW SOMETHING ABOUT RECIPIENT
		either direct address: street adress, or bank account nr
		or form of communication: email, etc

	-YOU MUST GIVE US SOMETHING
		a price || an account + email + phone

	VERBAL GIFT
		ie: give instructions to receive a gift.
		house can track it.

	HOUSE MONEY GIFT
		recipient contact or account
		house delivers and tracks

	PHYSICAL GIFT
		physichal delivery: your adress, recipient contact or address
		house sets delivery and tracks


	RECIPIENT NOTIFICATION
		a recipient must be notified in advance of any delivery.
		and given a chance to reject.

*/

/*
SNAIL CARDS
	As a gift card for example.

	GIFT CARD snail and email versions


ACCOUNT CREATION
	by: email, bank account nr,
		street address,

	INFORMATION POINTS
		-> account id


*/

/*

INPUT FIELDS

	RECIPIENT INFO:
		email/phone/pm  OR  account  OR  address
			forewarn? require confirmation?

	CONTENT:
		verbal, money, goods

	DELIVERY:
		via house or self-organized or third-party

	CHANGE OF OWNERSHIP
		COO happens for sure when recipient has ACCEPTED. But he may have Right to "return to sender" for a period of time.
		If a recipient receives but has not accepted, then the recipient has the right to decide ownership. (Ownership is in intermediate state)
			He is not obliged to send back. And he has no other obligations to the sender.
		OOO - Offer Of Ownership
			during OOO, a recipient can ACCEPT, and then the ownership changes.
		immediate, when sent, when delivered .. or .. (timed) .. recipient must accept, and can reject anytime before
			rejection?

	TIMING
		delayed send. delayed msg.

	PRIVACY
		public OR private
			private options: via codes. account required.

	-------------
		optional: multi output
			pot max size. pot rules. pot fill.
			recipient fill: fixed as initiated, over time, timeout
		optional: multi input. (restrictions: who can, amounts, what items, timeout)
		pot options:
			randomize outputs. or spinning wheel or lottery kind of things. or guess a number. pick a number. ordered prize giving.
	-------------

	SENDER INFO:
	account  OR  email/phone/pm  OR  account

>>> ENTER INPUT >>> (sender client to server)

	"Good, now you can deliver at any time, to START the GIFT"

<<< SEND REQUEST <<< (server asks sender to send)

	IF CONTENT.HOUSE
		.MONEY

			>>> SEND MONEY TO HOUSE >>>

			sender sends money to house. house sends money to receiver. =RECORDED
			house send MSG to receiver (if com known)

		.GOODS

			>>> SETUP DELIVERY >>>

			sender sends parcel. house arranges courier. courier tracks and informs house.
			house sends MSG to receiver (if com known)

	ELIF CONTENT.VERBAL


*/

/*
SIGNALLING

	via webpage
	via

	ID/reference tokens


*/

/*
TRADE BARTER
	myoffer	-- counter-offer		<-- highlight changes from past. display changes.
	myoffer	-- counter-offer
	myoffer	-- counter-offer
	myoffer -- dismissed or accepted

	subtrade -- subtrade

NEGOTIATION when used with verbals
*/

/*
LEGO PUZZLE
	Base Components
		Offer Object. Can be multiplexed. Upgraded to promise.
		Trade Object. Multiplexable. Two-way agreement.

*/

/*
PROVIDERS
	delivery, insurance
	snail cards, nice emails

	PROVIDER INTEGRATION
		house initiative, user initiative, provider initiative

USER FLOW
	via house
	via provider (store, deliverer, insurer, )
		common provider flows
*/

/*

NICE GIFT MESSAGE
	visuals, message

NICE GIFT CARDS

MANY RECIPIENTS
	a unique or secret id for each receiver


TRADER INTEGRATION
	a trader may setup the whole thing.

	how could I as a TRADER utlize this? i have a webshop for example. Give it GIFT functionality?
		send a nice gift message.
		TRADER sets up a proposed GIFT


CONTROL SIGNALS and INFO POINTS
	info point: a single piece of info of a sender or recipient.
		usually needed for a subsequent phase to start.

	control signals: signals about state changes in external agents; such as a transition from one phase to another
		phase signal: delivery started, delivery received
		activation signal: start, pause, finished (as agreed originally, as renegotiated, as unagreed)
		warning/notice signals:
		confirmation signals: confirm, reject, pause

A SIGNALLING SYSTEM AT HEART


MULTIPLE OFFERINGS
	POT
		Distribution: Round Robin, Equal, First Server

		Timeout/Rest Distribution
			or return to sender

*/

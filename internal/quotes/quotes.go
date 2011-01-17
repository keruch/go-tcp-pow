package quotes

import "math/rand"

type Quote int

const (
	Quote1 Quote = iota
	Quote2
	Quote3
	Quote4
	Quote5

	quoteNum // special variable for detecting a number of quotes
)

func (q Quote) String() string {
	switch q {
	case Quote1:
		return "I grew up on the crime side, the New York Times side stayin' alive was no jive"
	case Quote2:
		return "Had second hands, Mom's bounced on old man so then we moved to Shaolin land"
	case Quote3:
		return "Only way I begin the G off was drug loot and let's start it like this, son"
	case Quote4:
		return "Cash rules everything around me"
	case Quote5:
		return "Get the money, dollar, dollar bill, y'all"
	default:
		return ""
	}
}

func RandQuote() string {
	return Quote(rand.Int() % int(quoteNum)).String()
}

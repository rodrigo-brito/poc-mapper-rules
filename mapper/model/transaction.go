package model

type Transaction struct {
	Operation string
	Card      Card
	Country   string
	Amount    int
}

type Card struct {
	Brand  string
	Method string
}

package main

// All Ice Creams have a flavor and a price
type IceCream struct {
	flavor string
	price  int
	name   string
}

var iceCreams = []IceCream{
	{"Chocolate", 10, "Chocolate"},
	{"Vanilla", 10, "Vanilla"},
	{"Strawberry", 10, "Strawberry"},
	{"Mint", 10, "Mint"},
	{"Pistachio", 10, "Pistachio"},
	{"Rocky Road", 10, "Rocky Road"},
}

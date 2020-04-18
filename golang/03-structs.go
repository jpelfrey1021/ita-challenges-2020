package main

import "fmt"

type customer struct {
	name    string
	address string
}

func main() {
	jim := customer{
		name:    "Jim",
		address: "123 avenue st. Charlotte, NC 27483",
	}

	jim.updateAddress("123 go st. Raleigh, NC 38387")
	fmt.Println(jim)
}

func (c *customer) updateAddress(newAdd string) {
	(*c).address = newAdd
}

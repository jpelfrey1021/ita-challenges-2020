package challenge

type customer struct {
	name    string
	address string
}

func (c *customer) updateAddress(a string) {
	c.address = a
}
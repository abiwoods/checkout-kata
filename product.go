
type product struct {
	SKU       string
	UnitPrice int
	Offer     offer
}

type offer struct {
	Count int
	Price int
}
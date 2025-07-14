package pricing

type SpecialOffer struct {
	Quantity int
	Price    int
}

type Rule struct {
	UnitPrice int
	Offer     *SpecialOffer
}

type Rules map[string]Rule

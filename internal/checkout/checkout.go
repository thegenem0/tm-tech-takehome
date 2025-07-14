package checkout

type ICheckout interface {
	// Scan(SKU string) (err error)
	GetTotalPrice() (totalPrice int, err error)
}

type checkout struct {
}

func NewCheckout() ICheckout {
	return &checkout{}
}

func (c *checkout) GetTotalPrice() (totalPrice int, err error) {
	return 0, nil
}

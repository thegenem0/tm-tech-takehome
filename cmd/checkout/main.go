package main

import (
	"fmt"
	"log"

	"github.com/thegenem0/tm-tech-takehome/internal/checkout"
	"github.com/thegenem0/tm-tech-takehome/internal/pricing"
)

func main() {
	pricingRules := pricing.Rules{
		"A": {UnitPrice: 50, Offer: &pricing.SpecialOffer{Quantity: 3, Price: 130}},
		"B": {UnitPrice: 30, Offer: &pricing.SpecialOffer{Quantity: 2, Price: 45}},
		"C": {UnitPrice: 20, Offer: nil},
		"D": {UnitPrice: 15, Offer: nil},
	}

	checkout := checkout.NewCheckout(pricingRules)

	fmt.Println("Scanning items: A, B, A, C, B, A...")

	basket := []string{"A", "B", "A", "C", "B", "A"}

	for _, item := range basket {
		if err := checkout.Scan(item); err != nil {
			log.Fatalf("error scanning item %s: %v", item, err)
		}
	}

	total, err := checkout.GetTotalPrice()
	if err != nil {
		log.Fatalf("error calculating total price: %v", err)

	}

	// Should return 195
	// 3 * A = 130
	// 2 * B = 45
	// 1 * C = 20
	fmt.Printf("Total price for basket: %d pence\n", total)
}

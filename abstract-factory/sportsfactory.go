package main

import "fmt"

type SportsFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}

func GetFactory(brand string) (SportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}
	if brand == "Nike" {
		return &Nike{}, nil
	}
	return nil, fmt.Errorf("factory unknown")
}

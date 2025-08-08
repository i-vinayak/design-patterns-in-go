package main

import "fmt"

func main() {
	adidasFactory, _ := GetFactory("adidas")
	nikeFactory, _ := GetFactory("Nike")

	adShirt := adidasFactory.MakeShirt()
	adShoe := adidasFactory.MakeShoe()

	adShirt.setSize(10)
	adShoe.setSize(10)

	nikeShirt := nikeFactory.MakeShirt()
	nikeShirt.setSize(2)
	nikeShoe := nikeFactory.MakeShoe()
	nikeShoe.setSize(2)

	fmt.Println(nikeShoe, nikeShirt)
	fmt.Println(adShoe, adShirt)

}

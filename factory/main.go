package main

import "fmt"

func main() {

	ak, _ := GetGun("ak47")
	musket, _ := GetGun("musket")

	printDetails(ak)
	printDetails(musket)
}

func printDetails(gun IGun) {
	fmt.Println(gun.GetName(), ": ", gun.GetPower())
}

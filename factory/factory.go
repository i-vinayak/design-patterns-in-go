package main

import "fmt"

func GetGun(gunTpe string) (IGun, error) {
	if gunTpe == "ak47" {
		return NewAk(), nil
	}
	if gunTpe == "musket" {
		return NewMusket(), nil
	}
	return nil, fmt.Errorf("gun type invalid: %s", gunTpe)
}

package main

type musket struct {
	Gun
}

func NewMusket() IGun {
	return &musket{
		Gun{
			power: 100,
			name:  "musket",
		},
	}
}

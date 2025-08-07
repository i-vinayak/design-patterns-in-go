package main

type ak47 struct {
	Gun
}

func NewAk() IGun {
	return &ak47{
		Gun{
			name:  "ak47",
			power: 10,
		},
	}
}

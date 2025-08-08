package main

type Adidas struct{}

func (a *Adidas) MakeShoe() IShoe {
	return &AdidasShoe{
		Shoe{logo: "adidas", size: 2},
	}
}

func (a *Adidas) MakeShirt() IShirt {
	return &AdidasShirt{
		Shirt{logo: "adidas", size: 10},
	}
}

type Nike struct{}

func (a *Nike) MakeShoe() IShoe {
	return &NikeShoe{
		Shoe{logo: "adidas", size: 2},
	}
}

func (a *Nike) MakeShirt() IShirt {
	return &NikeShirt{
		Shirt{logo: "adidas", size: 10},
	}
}

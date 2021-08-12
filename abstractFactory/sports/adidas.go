package main

type adidas struct {}

func (*adidas) getShoe() IShoe {
	return &adidasShoe{
		shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (*adidas) getShirt() IShirt {
	return &adidasShirt{shirt{
		logo: "adidas",
		size: 14,
	}}
}
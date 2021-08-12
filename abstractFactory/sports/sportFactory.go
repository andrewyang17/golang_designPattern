package main

import "errors"

type ISportFactory interface {
	getShoe() IShoe
	getShirt() IShirt
}

func GetSportsFactory(brand string) (ISportFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}

	if brand == "nike" {
		return &nike{}, nil
	}

	return nil, errors.New("wrong brand type passed")
}
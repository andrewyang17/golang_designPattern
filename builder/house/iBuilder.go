package main

import "errors"

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func getBuilder(builderType string) (iBuilder, error) {
	if builderType == "normal" {
		return &normalBuilder{}, nil
	}
	if builderType == "igloo" {
		return &iglooBuilder{}, nil
	}

	return nil, errors.New("incorrect builder type passed")
}



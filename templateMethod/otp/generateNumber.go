package main

import (
	"math/rand"
	"strconv"
	"strings"
)

func generateNumber(length int) string {
	var numbers []string
	for i := 0; i < length; i++ {
		numbers = append(numbers, strconv.Itoa(rand.Intn(9)))
	}
	return strings.Join(numbers, "")
}
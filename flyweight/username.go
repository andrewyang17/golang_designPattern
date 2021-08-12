package main

import (
	"fmt"
	"strings"
)

type User struct {
	FullName string
}

func NewUser(fullName string) *User {
	return &User{FullName: fullName}
}

// Better way
var allNames []string

type User2 struct {
	// Flyweight
	names []uint8
}

func (u *User2) FullName() string {
	var parts []string

	for _, name := range u.names {
		parts = append(parts, allNames[name])
	}

	return strings.Join(parts, " ")
}

func NewUser2(fullName string) *User2 {
	getOrAdd := func(s string) uint8 {
		strings.ToLower(s)
		for i := range allNames {
			if allNames[i] == s {
				return uint8(i)
			}
		}
		allNames = append(allNames, s)
		return uint8(len(allNames)-1)
	}

	u := User2{}
	parts := strings.Split(fullName, " ")

	for _, p := range parts {
		u.names = append(u.names, getOrAdd(p))
	}
	return &u
}

func main() {
	u1 := NewUser("Andrew Yang")
	u2 := NewUser("Catherine Yang")
	u3 := NewUser("Andrew Yong")
	fmt.Println(u1.FullName)
	fmt.Println(u2.FullName)
	fmt.Println(u3.FullName)
	fmt.Printf("Memory taken: %d bytes \n", len([]byte(u1.FullName)) + len([]byte(u2.FullName)) +len([]byte(u3.FullName)))

	// Flyweight way
	u4 := NewUser2("Andrew Yang")
	u5 := NewUser2("Catherine Yang")
	u6 := NewUser2("Andrew Yong")
	fmt.Println(u4.FullName())
	fmt.Println(u5.FullName())
	fmt.Println(u6.FullName())

	totalMemory := 0

	for _, name := range allNames {
		totalMemory += len([]byte(name))
	}
	totalMemory += len(u4.names)
	totalMemory += len(u5.names)
	totalMemory += len(u6.names)
	fmt.Printf("Memory taken: %d bytes \n", totalMemory)
}
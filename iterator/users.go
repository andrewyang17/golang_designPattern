package main

import "fmt"

type user struct {
	name string
	age int
}

type iterator interface {
	hasNext() bool
	getNext() *user
}

type collection interface {
	createIterator() iterator
}

type userIterator struct {
	index int
	users []*user
}

func (u *userIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

func (u *userIterator) getNext() *user {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}

type userCollection struct {
	users []*user
}

func (u *userCollection) createIterator() iterator {
	return &userIterator{
		users: u.users,
	}
}

func main() {
	user1 := &user{
		name: "Andrew",
		age:  25,
	}

	user2 := &user{
		name: "Yang",
		age:  25,
	}

	userCollection := &userCollection{users: []*user{user1, user2}}
	userIterator := userCollection.createIterator()

	for userIterator.hasNext() {
		user := userIterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
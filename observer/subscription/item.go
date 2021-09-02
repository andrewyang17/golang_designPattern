package main

import "fmt"

type item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func (i *item) updateAvailability() {
	fmt.Printf("Item %s is now in stock \n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *item) register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) deregister(o Observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

func newItem(name string) *item {
	return &item{name: name}
}

func removeFromSlice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, Observer := range observerList {
		if observerToRemove.getID() == Observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}



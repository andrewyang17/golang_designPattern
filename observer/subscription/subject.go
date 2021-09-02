package main

type Publisher interface {
	register(Observer)
	deregister(Observer)
	notifyAll()
}

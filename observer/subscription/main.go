package main

func main() {
	apple := newItem("Iphone13")

	observer1 := newCustomer("yangzhisiang@gmail.com")
	observer2 := newCustomer("yangzhisiangg@gmail.com")

	apple.register(observer1)
	apple.register(observer2)

	apple.updateAvailability()
}
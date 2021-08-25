package main

import "fmt"

type computer interface {
	lightningPort()
}

type client struct {}

func (*client) insertLightningConnectorIntoComputer(com computer) {
	fmt.Println("Client inserts lightning connector into computer")
	com.lightningPort()
}

type mac struct {}

func (*mac) lightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine")
}

// adaptee
type windows struct {}

func (*windows) USBPort() {
	fmt.Println("USB connector is plugged into window machine")
}

type windowsAdapter struct {
	windowMachine *windows
}

func (w *windowsAdapter) lightningPort() {
	fmt.Println("Window adapter converts lightning signal to USB")
	w.windowMachine.USBPort()
}

func main() {
	client := &client{}
	mac := &mac{}
	windows := &windows{}

	windowAdapter := &windowsAdapter{windowMachine: windows}

	client.insertLightningConnectorIntoComputer(mac)
	client.insertLightningConnectorIntoComputer(windowAdapter)
}
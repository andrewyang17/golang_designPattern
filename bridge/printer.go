package main

import "fmt"

type printer interface {
	printFile()
}

type epson struct {}

func (*epson) printFile() {
	fmt.Println("Printing by a EPSON Printer")
}

type hp struct {}

func (*hp) printFile() {
	fmt.Println("Printing by a HP Printer")
}

type computer interface {
	print()
	setPrinter(printer)
}

type mac struct {
	printer printer
}

func (m *mac) print() {
	fmt.Println("Print request from mac")
	m.printer.printFile()
}

func (m *mac) setPrinter(p printer) {
	m.printer = p
}

type windows struct {
	printer printer
}

func (w *windows) print() {
	fmt.Println("Print request from windows")
	w.printer.printFile()
}

func (w *windows) setPrinter(p printer) {
	w.printer = p
}

func main() {
	// Implementation
	hpPrinter := &hp{}
	epsonPrinter := &epson{}

	// Abstraction
	macComputer := &mac{}
	winComputer := &windows{}

	macComputer.setPrinter(hpPrinter)
	macComputer.print()

	winComputer.setPrinter(epsonPrinter)
	winComputer.print()
}
package main

import "fmt"

type Person struct {
	StreetAddress, Postcode, City string
	CompanyName, Position string
	AnnualIncome int
}

type PersonBuilder struct {
	person *Person
}

func (pb *PersonBuilder) Build() *Person {
	return pb.person
}

func (pb *PersonBuilder) Work() *PersonJobBuilder {
	return &PersonJobBuilder{*pb}
}

func (pb *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*pb}
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (pjb *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	pjb.person.CompanyName = companyName
	return pjb
}

func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}

func (pjb *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	pjb.person.AnnualIncome = annualIncome
	return pjb
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (pab *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	pab.person.StreetAddress = streetAddress
	return pab
}

func (pab *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	pab.person.City = city
	return pab
}

func (pab *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	pab.person.Postcode = postcode
	return pab
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

func main() {
	p := NewPersonBuilder()
	p.Lives().
		At("Kebun Sultan").
		In("Kota Bharu").
		WithPostcode("15300")
	p.Work().
		At("Superceed").
		AsA("Programmer").
		Earning(000)

	fmt.Println(p.Build())
}

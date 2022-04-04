package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string {
	return p.first + p.last
}

type doubleZero struct {
	person
	LicenseToKill bool
	First         string
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Selva", "Mohandoss", 56}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.fullName())

	p3 := doubleZero{
		person: person{
			first: "Selva",
			last:  "Mohandoss",
			age:   25,
		},
		First:         "If looks could kill",
		LicenseToKill: true,
	}
	fmt.Println(p3.First, p3.person.first)

}

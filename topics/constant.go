package main

import "fmt"

func main() {

	// single constant
	const name string = "Kuntiarso"
	const age = 24

	// multiple constants
	const (
		lastName string = "Gilang"
		year     int    = 1997
		isBirth  bool   = true
	)

	fmt.Println(name)
	fmt.Println(age)

	fmt.Println(lastName)
	fmt.Println(year)
	fmt.Println(isBirth)
}

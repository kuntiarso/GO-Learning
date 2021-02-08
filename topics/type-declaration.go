package main

import "fmt"

func main() {
	type name = string

	var myName name = string("Kuntiarso")
	fmt.Printf("My name is { %s }\n", myName)

}

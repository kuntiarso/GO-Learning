package main

import "fmt"

func main() {
	var name string = "Kuntiarso"
	var char byte = name[2]
	// type conversion
	var str string = string(char)

	fmt.Printf("My name is { %s } and my char is { %s }", name, str)

}

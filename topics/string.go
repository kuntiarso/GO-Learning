package main

import "fmt"

func main() {
	fmt.Println("print 'hello world' string with GO.")
	fmt.Println(len("helloworld"))
	fmt.Println("helloworld"[0])
	fmt.Println("helloworld"[len("helloworld")-1])
}

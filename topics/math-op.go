package main

import "fmt"

func main() {
	var res = 10 + 10
	fmt.Printf("res { %d }\n", res)

	res += 10
	fmt.Printf("res { %d }\n", res)
}

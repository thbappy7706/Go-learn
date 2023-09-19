package main

import (
	"fmt"
)

func main() {
	const PI float32 = 3.1416
	const (
		A int = 1
		B     = 3.14
		C     = "Hi!"
	)

	fmt.Println("Value of pi:", PI)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}

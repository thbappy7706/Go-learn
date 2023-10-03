package main

import (
	"fmt"
)

func main() {
	getLengthOfAnArray()
}

func getLengthOfAnArray() {
	arr := [...]int{5, 5, 34, 53, 67, 88, 33, 23}
	fmt.Println(len(arr))
}

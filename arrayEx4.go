// Array Initialization
package main

import (
	"fmt"
)

func main() {
	arr1 := [5]int{}
	arr2 := [5]int{1, 2}
	arr3 := [5]int{1, 2, 3, 4, 5}
	arr4 := [...]string{"1", "5"}
	mixedArray := []interface{}{"foo", 5, 3.50, "hello", nil}

	fmt.Println(arr1, arr2, arr3, arr4, mixedArray)
}

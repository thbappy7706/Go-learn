// Change Elements of an Array
package main

import (
	"fmt"
)

func main() {
	myExpenses := [...]int{100, 200, 350, 400, 700, 670}
	myExpenses[2] += 50
	fmt.Println(myExpenses)
}

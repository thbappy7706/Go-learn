package main

import (
	"fmt"
)

func main() {
	var numberArr = [4]int{5, 6, 8, 9}
	var stringArr = [2]string{"Tanvir Hossen Bappy", "Software Engineer"}
	testNumbers := [5]int{4, 5, 6, 7, 8}

	fmt.Println("The arrays of number:", numberArr)
	fmt.Println("The arrays of string:", stringArr)
	fmt.Println("test:", testNumbers)

	//declares two arrays (arr1 and arr2) with inferred lengths:
	other()
}

func other() {
	//declares two arrays (arr1 and arr2) with inferred lengths:
	var arr1 = [...]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)
}

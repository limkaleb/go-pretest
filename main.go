package main

import "fmt"


func main() {
	test1 := isPalindrome("radar")
	fmt.Printf("Answer 1:\n%t\n", test1)

	fmt.Println()

	arr := []int{3,5,1,9,2}
	test2 := maxValue(arr)
	fmt.Printf("Answer 2:\n%d\n", test2)

	fmt.Println()

	fmt.Println("Answer 3:")
	printTriangle(5)

	fmt.Println()

	test4 := factorial(5)
	fmt.Printf("Answer 4:\n%d\n", test4)
}

package main

import "fmt"

func printTriangle(input int) {
	i := 0
	j := 0
	for i < input {
		for j <= i {
			fmt.Print("*")
			j +=1
		}
		fmt.Println()
		j = 0
		i += 1
	}
}

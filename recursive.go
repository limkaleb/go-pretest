package main

func factorial(input int) int {
	if input == 1 {
		return 1
	}

	return input * factorial(input - 1)
}

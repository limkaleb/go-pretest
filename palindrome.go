package main

func isPalindrome(input string) bool {
	for i := range input {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

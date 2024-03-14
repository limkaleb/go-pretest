package main

func maxValue(input []int) int {
	max := 0
	for _, num := range input {
		if num >= max {
			max = num
		}
	}
	return max
}

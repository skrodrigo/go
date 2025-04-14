package main

// Sum Simple
func Sum(a int, b int) int {
	return a + b
} 

// Sum of multiple numbers
func TotalSum(x ...int) int {
	result := 0

	for _, v := range x {
		result += v
	}

	return result
}

// Division of simple return
func Division(a int, b int) (result int) {
	result = a / b 
	return
}
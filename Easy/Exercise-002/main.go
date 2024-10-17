package main

import "fmt"

func factorial(num int) int {
	res := 1

	if num == 0 {
		return 1
	}

	for i := 1; i <= num; i++ {
		res *= i
	}

	return res
}

func main() {
	var num int
	fmt.Print("Enter a number to calculate the factorial: ")
	fmt.Scan(&num)
	fmt.Printf("The factorial of %d is equal %s", num, fmt.Sprint(factorial(num)))
}

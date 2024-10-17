package main

import (
	"errors"
	"fmt"
)

func mapSqr(num int) (map[int]int, error) {
	//num is integral number between 1 and n
	resMap := make(map[int]int)
	if num == 0 {
		return nil, errors.New("zero num is not allowed")
	}
	for i := 1; i <= num; i++ {
		resMap[i] = i * i
	}

	return resMap, nil

}

func main() {
	var num int

	fmt.Print("Enter a number to calculate: ")
	fmt.Scan(&num)

	res, err := mapSqr(num)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The result map for num: %d is %v", num, res)
}

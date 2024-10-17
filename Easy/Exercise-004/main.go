package main

import (
	"fmt"
	"strconv"
	"strings"
)

func concat(str string) ([]int, error) {
	var resSlice []int

	res := strings.Split(str, ", ")
	for _, num := range res {
		toInt, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}
		resSlice = append(resSlice, toInt)
	}
	return resSlice, nil
}

func main() {
	inputStr := "1, 2, 3, 4, 5, 6"
	res, err := concat(inputStr)
	if err != nil {
		return
	}

	fmt.Printf("%v", res)
}

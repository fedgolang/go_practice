package main

import (
	"fmt"
	"strconv"
	"strings"
)

func devisable(lowNum, highNum int) string {
	var result string

	res := []string{}

	for i := lowNum; i <= highNum; i++ {
		if i == 0 {
			fmt.Println("Devide by zero attempt, continuing")
			continue
		}
		if i%7 == 0 && i%5 != 0 {
			res = append(res, strconv.Itoa(i))
		}
	}
	result = strings.Join(res, ", ")

	return result
}

func main() {
	var lowNum, highNum int
	fmt.Println("Enter lowest and highest number in format: low high")
	fmt.Scanf("%d %d", &lowNum, &highNum)
	fmt.Println(devisable(lowNum, highNum))

}

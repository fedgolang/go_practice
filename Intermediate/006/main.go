package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func execFormula(d int) int {
	c, h := 50, 30

	res := int(math.Round(math.Sqrt(float64((2 * c * d) / h))))

	return res
}

func main() {
	var numbers string
	var formulaRes []int
	var res []string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter numbers for formula, separate them by ',': ")
	if scanner.Scan() {
		numbers = scanner.Text()
	} else {
		fmt.Printf("%s", "Input error")
	}

	variables := strings.Split(numbers, ", ")

	for _, num := range variables {
		intNum, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("There is an error with conversation")
			return
		}
		formulaRes = append(formulaRes, execFormula(intNum))
	}

	for _, num := range formulaRes {
		res = append(res, strconv.Itoa(num))
	}

	fmt.Println(res)

}

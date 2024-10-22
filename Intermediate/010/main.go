package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func distinctSortLine(line string) string {
	lineSlice := strings.Split(line, " ")

	slices.Sort(lineSlice)

	for i := 1; i < len(lineSlice); i++ {
		if lineSlice[i] == lineSlice[i-1] {
			lineSlice[i], lineSlice[i-1] = " ", " "
		}
	}

	res := strings.Join(lineSlice, " ")

	return strings.Replace(res, "  ", "", -1)
}

func main() {
	var line string

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		line = distinctSortLine(scanner.Text())
	} else {
		fmt.Println("Input error")
		return
	}

	fmt.Println(line)
}

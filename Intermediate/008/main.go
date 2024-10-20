package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func wordsSorting(words string) string {
	wordsSlice := strings.Split(words, ",")

	slices.Sort(wordsSlice)

	res := strings.Join(wordsSlice, ",")

	return res
}

func main() {
	var text string

	fmt.Print("Enter the words separated by ',': ")
	scan := bufio.NewScanner(os.Stdin)

	if scan.Scan() {
		text = scan.Text()
	} else {
		fmt.Println("Input error")
	}

	println(wordsSorting(text))
}

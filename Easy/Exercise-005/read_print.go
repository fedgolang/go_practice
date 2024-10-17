package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadString() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your sentence: ")
	if scanner.Scan() {
		text := scanner.Text()
		return text
	} else {
		return "Input error"
	}

}

func PrintString(str string) string {
	return strings.ToUpper(str)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func linesToUpper(lines string) string {
	return strings.ToUpper(lines)
}

func main() {
	var lines string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		lines += linesToUpper(scanner.Text()) + "\n"
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Input problem")
		return
	}

	fmt.Println(lines)
}

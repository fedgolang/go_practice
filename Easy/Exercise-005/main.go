package main

import "fmt"

func main() {
	newStr := ReadString()
	fmt.Println(newStr)

	fmt.Println(PrintString(newStr))
}

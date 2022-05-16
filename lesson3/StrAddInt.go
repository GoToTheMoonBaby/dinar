package main

import (
	"fmt"
	"strconv"
)

func StrAddInt(first string, second int) string {
	secondString := strconv.Itoa(second)
	return (first + secondString)
}

func StrAddInt2(first string, second int) string {
	secondString := fmt.Sprint(second)
	return (first + secondString)
}

func main() {
	fmt.Printf(StrAddInt("Go to the Moon ", 3000))
	fmt.Printf(StrAddInt2("\nGo to the Moon ", 3000))
}

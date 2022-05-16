package main

import (
	"fmt"
	"strconv"
)

func StrAddInt(first string, second int) string {
	secondString := strconv.Itoa(second)
	return (first + secondString)
}

func main() {
	fmt.Printf(StrAddInt("Go to the Moon ", 3000))
}

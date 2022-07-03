package main

import (
	"fmt"
)

func StrLen(str string) int {
	return len([]rune(str))
}

func main() {
	fmt.Print(StrLen("abcдцйr"))
}

package main

import "fmt"

func StrConcat(first, second string) string {
	return (first + second)
}

func main() {
	fmt.Printf(StrConcat("Hello ", "world!"))
}

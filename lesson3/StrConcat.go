package lesson3

import "fmt"

func StrConcat(first, second string) string {
	return first + second
}

func tt() {
	fmt.Printf(StrConcat("Hello ", "world!"))
}

package lesson3

import (
	"fmt"
	"strconv"
)

func StrAddInt(first string, second int) string {
	return first + strconv.Itoa(second)
}

func StrAddInt2(first string, second int) string {
	return first + fmt.Sprint(second)
}

package lesson3

import (
	"fmt"
)

func DelSubStr(str, subStr string, n int) string {
	lenSubStr := len(subStr)
	lenStr := len(str)
	if lenSubStr > lenStr || lenSubStr == 0 {
		return str
	}
	inputNumber := 0
	i := 0
	for i < lenStr {
		j := 0
		for j < lenSubStr && str[i+j] == subStr[j] {
			j++
		}
		if j == lenSubStr {
			inputNumber++
			if inputNumber == n || n == -1 {
				str = str[:i] + str[i+j:]
				fmt.Println("DEBUG: ", str)
				lenStr -= lenSubStr
				if n != -1 {
					return str
				}
			} else {
				i++
			}
		} else {
			i++
		}
	}
	return str
}

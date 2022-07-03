package main

import (
	"fmt"
)

func StrTrim(str, substr string, n int) string {
	if len(substr) > len(str) {
		return str
	}
	if len(substr) == 0 {
		return str
	}
	newStr, newStr1, newStr2 := str, str, str

	// отдельная реализация удаления всех совпадений
	if n == -1 {
		for i := len(str) - 1; i >= 0; i-- {
			k := i // индекс совпадения символа
			flag := false
			if str[i] == substr[len(substr)-1] {
				for j := len(substr) - 1; j >= 0; j-- {
					if (str[i-(len(substr)-(j+1))] == substr[j]) && (i+len(substr)-j) > 0 {
						flag = true
					} else {
						flag = false
					}
				}
				if flag == true {
					newStr1 = newStr[:k-len(substr)+1]
					newStr2 = newStr[k+1:]
					newStr = newStr1 + newStr2
					fmt.Println(newStr)
				}
			}
		}
		return newStr
	} else {
		// удаление n-го совпадения
		f := 0 //номер совпадения
		for i := 0; i < len(str); i++ {
			k := i // индекс совпадения символа
			flag := false
			if str[i] == substr[0] {
				for j := 0; j < len(substr); j++ {
					if (i+j < len(str)) && (str[i+j] == substr[j]) {
						flag = true
					} else {
						flag = false
					}
				}
				if flag == true {
					f += 1
					if f == n {
						newStr1 = newStr1[:k]
						newStr2 = newStr2[len(substr)+k:]
						newStr = newStr1 + newStr2
						fmt.Println(newStr)
					}
				}
			}
		}
	}
	return newStr
}

func main() {
	res := StrTrim("1MoonGo2Moon to 3Moonthe 4Moon!!5Moon!", "Moon", -1)
	fmt.Printf(res)
}

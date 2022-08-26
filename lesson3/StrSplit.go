package lesson3

import "fmt"

func StrSplit(str string, sep string) []string {
	var splitStr []string
	k := 0
	for i := 0; i < len(str); i++ {
		if string(str[i]) == sep {
			k += i
			splitStr = append(splitStr, str[:k])
		}
		//	fmt.Printf(string(str[i]))
	}
	return splitStr
}

func lesson3() {
	res := StrSplit("naabc,bas", ",")
	fmt.Println(res)
}

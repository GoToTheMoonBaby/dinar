package main

import (
	"fmt"
)

func StrSub(first, second string) bool {
	var flag = false
	if len(second) > len(first) {
		return false
	}
	if len(second) == 0 {
		return true
	}
	for i := 0; i < len(first); i++ {
		if first[i] == second[0] {
			for j := 0; j < len(second); j++ {
				if (i+j < len(first)) && (first[i+j] == second[j]) {
					flag = true
				} else {
					flag = false
				}
			}
			if flag == true {
				return flag
			}
		}
	}
	return flag
}

func main() {
	res := StrSub("Go to the Moon", "Moon")
	//st := string(res)
	//	fmt.Printf(st)
	fmt.Printf("%v", res)
}

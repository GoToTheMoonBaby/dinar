package lesson3

import "strings"

func StrSub(str, substr string) bool {
	var flag = false
	if len(substr) > len(str) {
		return false
	}
	if len(substr) == 0 {
		return true
	}
	for i := 0; i < len(str); i++ {
		if str[i] == substr[0] {
			for j := 0; j < len(substr); j++ {
				if (i+j < len(str)) && (str[i+j] == substr[j]) {
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

func StrSub2(str, substr string) bool {
	var flag = false
	firstLen := len(str)
	secondLen := len(substr)
	if secondLen > firstLen {
		return false
	}
	if secondLen == 0 {
		return true
	}
	for i := 0; i <= firstLen-secondLen; i++ {
		if str[i] == substr[0] {
			flag = true
			for j := 0; j < secondLen; j++ {
				if str[i+j] != substr[j] {
					flag = false
					break
				}
			}
			if flag {
				return flag
			}
		}
	}
	return flag
}

func StrSub3(str, substr string) bool {
	return strings.Contains(str, substr)
}

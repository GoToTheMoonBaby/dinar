package main

import "testing"

func Equal(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestStrConcat(t *testing.T) {
	first1 := "Go to the "
	second1 := "moon!"
	first2 := "Hello "
	second2 := ""
	first3 := ""
	second3 := "brother"
	first4 := "123"
	second4 := "-456"

	exp1 := "Go to the moon!"
	exp2 := "Hello "
	exp3 := "brother"
	exp4 := "123-456"

	result1 := StrConcat(first1, second1)
	result2 := StrConcat(first2, second2)
	result3 := StrConcat(first3, second3)
	result4 := StrConcat(first4, second4)

	if !(Equal(exp1, result1)) {
		t.Errorf("Incorrect result! exp '%s', got '%s'", exp1, result1)
	}
	if !(Equal(exp2, result2)) {
		t.Errorf("Incorrect result! exp '%s', got '%s'", exp2, result2)
	}
	if !(Equal(exp3, result3)) {
		t.Errorf("Incorrect result! exp  '%s', got '%s'", exp3, result3)
	}
	if !(Equal(exp4, result4)) {
		t.Errorf("Incorrect result! exp  '%s', got '%s'", exp4, result4)
	}

}

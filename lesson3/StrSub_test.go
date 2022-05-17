package main

import "testing"

func TestStrSub(t *testing.T) {
	first1 := "Go to the moon "
	second1 := "moon "
	first2 := "Hello brother"
	second2 := "hero"
	first3 := "Go to the moon "
	second3 := "M"
	first4 := "Go to the moon "
	second4 := ""

	exp1 := true
	exp2 := false
	exp3 := false
	exp4 := true

	result1 := StrSub(first1, second1)
	result2 := StrSub(first2, second2)
	result3 := StrSub(first3, second3)
	result4 := StrSub(first4, second4)

	if !(result1 == exp1) {
		t.Errorf("Incorrect result! exp '%v', got '%v'", exp1, result1)
	}
	if !(result2 == exp2) {
		t.Errorf("Incorrect result! exp '%v', got '%v'", exp2, result2)
	}
	if !(result3 == exp3) {
		t.Errorf("Incorrect result! exp  '%v', got '%v'", exp3, result3)
	}
	if !(result4 == exp4) {
		t.Errorf("Incorrect result! exp  '%v', got '%v'", exp4, result4)
	}
}

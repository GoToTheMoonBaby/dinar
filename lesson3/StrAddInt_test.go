package lesson3

import (
	"testing"
)

func TestStrAddInt(t *testing.T) {
	first1 := "Go to the "
	second1 := 88888
	first2 := "Hello "
	var second2 int
	first3 := ""
	second3 := -12345
	first4 := "123"
	second4 := 00

	exp1 := "Go to the 88888"
	exp2 := "Hello 0"
	exp3 := "-12345"
	exp4 := "1230"

	result1 := StrAddInt(first1, second1)
	result2 := StrAddInt(first2, second2)
	result3 := StrAddInt(first3, second3)
	result4 := StrAddInt(first4, second4)

	if exp1 != result1 {
		t.Errorf("Incorrect result! exp '%s', got '%s'", exp1, result1)
	}
	if exp2 != result2 {
		t.Errorf("Incorrect result! exp '%s', got '%s'", exp2, result2)
	}
	if exp3 != result3 {
		t.Errorf("Incorrect result! exp  '%s', got '%s'", exp3, result3)
	}
	if exp4 != result4 {
		t.Errorf("Incorrect result! exp  '%s', got '%s'", exp4, result4)
	}

}

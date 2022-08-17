package lesson3

import "testing"

func TestDelSubStr(t *testing.T) {
	first1 := "Go to the moon"
	second1 := "moon"
	inputNumber1 := 1
	first2 := "Hello brother"
	second2 := "hwwe"
	inputNumber2 := 1
	first3 := "Go to the moon "
	second3 := "o"
	inputNumber3 := -1
	first4 := "Go to the moon "
	second4 := "t"
	inputNumber4 := 2

	exp1 := "Go to the "
	exp2 := "Hello brother"
	exp3 := "G t the mn "
	exp4 := "Go to he moon "

	result1 := DelSubStr(first1, second1, inputNumber1)
	result2 := DelSubStr(first2, second2, inputNumber2)
	result3 := DelSubStr(first3, second3, inputNumber3)
	result4 := DelSubStr(first4, second4, inputNumber4)

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

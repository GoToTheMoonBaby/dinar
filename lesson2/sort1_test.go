package lesson2

import "testing"

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestBubbleSort(t *testing.T) {
	test1 := []int{}
	test2 := []int{1}
	test3 := []int{3, 0, -1, -2}
	test4 := []int{1, 2, 3, 4}
	test5 := []int{1, 1, 1, 1}
	test6 := []int{4, 3, 2, -1, -2, -3}
	exp1 := []int{}
	exp2 := []int{1}
	exp3 := []int{-2, -1, 0, 3}
	exp4 := []int{1, 2, 3, 4}
	exp5 := []int{1, 1, 1, 1}
	exp6 := []int{-3, -2, -1, 2, 3, 4}

	result1 := BubbleSort(test1)
	result2 := BubbleSort(test2)
	result3 := BubbleSort(test3)
	result4 := BubbleSort(test4)
	result5 := BubbleSort(test5)
	result6 := BubbleSort(test6)

	if !(Equal(result1, exp1)) {
		t.Errorf("Incorrect result!!!exp %d, got %d", exp1, result1)
	}
	if !(Equal(result2, exp2)) {
		t.Errorf("Incorrect result!!!exp %d, got %d", exp2, result2)
	}
	if !(Equal(result3, exp3)) {
		t.Errorf("Incorrect result!!!exp %d, got %d", exp3, result3)
	}
	if !(Equal(result4, exp4)) {
		t.Errorf("Incorrect result!!!exp %d, got %d", exp4, result4)
	}
	if !(Equal(result5, exp5)) {
		t.Errorf("Incorrect result!!!exp %d, got %d", exp5, result5)
	}
	if !(Equal(result6, exp6)) {
		t.Errorf("Incorrect result!!!exp %d, got %d", exp6, result6)
	}
}

package main

import "fmt"

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func FindIndexOfValue(arr []int, value int) (bool, int) {
	left := 0             //левая граница
	right := len(arr) - 1 //правая граница
	for left < right {
		middle := (left + right) / 2
		if value <= arr[middle] {
			right = middle //смещение правой границы
		} else {
			left = middle + 1 //смещение левой границы
		}
	}
	if value == arr[right] {
		return true, right
	} else {
		return false, right
	}
}

func main() {
	data := []int{8, 4, 7, 9, 1, 3, 5, 4, 2}
	findValue := 9

	fmt.Println("Изначальный массив:    ", data)
	sortedData := BubbleSort(data)
	fmt.Println("Отсортированный массив:", sortedData)

	check, index := FindIndexOfValue(sortedData, findValue)
	if check == false {
		fmt.Println("Данного числа нет в массиве")
	} else {
		fmt.Println("Число ", findValue, "находится на позиции: ", index)
	}
}

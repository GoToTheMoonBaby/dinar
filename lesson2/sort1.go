package main

import "fmt"

func BubbleSort(slice []int) []int {
	for i := len(slice); i > 0; i-- {
		for j := 1; j < i; j++ {
			if slice[j-1] > slice[j] {
				slice[j-1], slice[j] = slice[j], slice[j-1] //меняем значения местами
			}
		}
	}
	return (slice)
}

func ShakerSort(slice []int) {
	left := 0               //левая граница
	right := len(slice) - 1 //правая граница

	for left <= right {
		// обход слева-направо
		for i := left; i <= right; i++ {
			if slice[i] > slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i] //меняем значения местами
			}
			right -= 1 //сужаем правую границу
		}
		// обход справа-налево
		for i := right; i >= left; i-- {
			if slice[i-1] > slice[i] {
				slice[i-1], slice[i] = slice[i], slice[i-1] //меняем значения местами
			}
			left += 1 //сужаем левую границу
		}
	}
}

func main() {
	arr := []int{4, 56, 546456, 8, 11, 0, -1} //инициализируем массив (слайс)
	fmt.Println(arr)                          //вывод изначального массива
	BubbleSort(arr)                           //выполнение функции сортировки Пузырьком
	fmt.Println("BubbleSort", arr)            //вывод отсортированного Пузырьком массива
	ShakerSort(arr)                           ////выполнение функции сортировки перемешиванием
	fmt.Println("ShakerSort", arr)            //вывод отсортированного массива методом перемешивания
}

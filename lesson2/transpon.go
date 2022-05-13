package main

import "fmt"

func transpon(slice [][]int) [][]int {
	//создание нового среза с размером элемента входящего среза
	transponSlice := make([][]int, len(slice[0]))
	for i := 0; i < len(slice[0]); i++ {
		//создание вложенного среза с размером входящего среза
		transponSlice[i] = make([]int, len(slice))
		for j := 0; j < len(slice); j++ {
			// присваиваем новому срезу значения изначальгого меняя индексы
			transponSlice[i][j] = slice[j][i]
		}
	}
	return transponSlice // возвращаем новый срез
}

func main() {
	arr := [][]int{{4, 56}, {8, 11}, {-1, 1}, {0, -100}} //инициализируем среза
	fmt.Println(arr)                                     //вывод изначального среза
	fmt.Println(transpon(arr))                           //вывод транспонированного среза
}

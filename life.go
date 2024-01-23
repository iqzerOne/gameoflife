package main

import "fmt"

// Поле
func field(array [][]int) {
	for i := 0; i < len(array); i++ {
		fmt.Print("\n")
		for j := 0; j < len(array); j++ {
			if array[i][j] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
	}
}

// подсчет соседей
func cell_counter(array [][]int, y, x int) int {
	res := 0
	res += predel(array, y, x+1)
	res += predel(array, y, x-1)
	res += predel(array, y+1, x)
	res += predel(array, y-1, x)
	res += predel(array, y-1, x-1)
	res += predel(array, y-1, x+1)
	res += predel(array, y+1, x-1)
	res += predel(array, y+1, x+1)
	return res
}

// Копирования двумерного массива
func duplicate(array [][]int) [][]int {
	duplicate := make([][]int, len(array))
	for i := range array {
		duplicate[i] = make([]int, len(array[i]))
		copy(duplicate[i], array[i])
	}
	return duplicate
}

// Проверка клетки на живую или мертвую
func livedead(array [][]int) [][]int {
	new_array := duplicate(array)
	for y := 0; y < len(array); y++ {
		for x := 0; x < len(array); x++ {
			counter := cell_counter(array, y, x)
			if (counter > 3 || counter < 2) && array[y][x] == 1 {
				new_array[y][x] = 0
			}
			if counter == 3 && array[y][x] == 0 {
				new_array[y][x] = 1
			}
		}
	}
	return new_array
}

// проверка выхода за предел массива
func predel(array [][]int, y, x int) int {
	if (y >= 0 && x >= 0) && (y < len(array) && x < len(array)) {
		if array[y][x] == 1 {
			return 1
		}
	}
	return 0
}
func main() {
	array := [][]int{
		{0, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	}
	field(array)
	// array = livedead(array)
	// field(array)
}

package main

import "fmt"

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
		{0, 0, 1},
		{1, 1, 1},
	}
	field(array)
	fmt.Print("\n", cell_counter(array, 2, 1))
}

// res += array[y][x+1]
// res += array[y][x-1]
// res += array[y+1][x]
// res += array[y-1][x]
// res += array[y-1][x-1]
// res += array[y-1][x+1]
// res += array[y+1][x-1]
// res += array[y+1][x+1]

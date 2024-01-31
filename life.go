package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Поле
func field(array [][]int) {
	for i := 0; i < len(array); i++ {
		fmt.Print("\n")
		for j := 0; j < len(array[i]); j++ {
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

// Считывание длины файла
func lenght(stroka string) int {
	line := strings.Split(stroka, "\n")
	max := 0
	for i := 0; i < len(line); i++ {
		if max < len(line[i]) {
			max = len(line[i])
		}
	}
	return max
}

// Считывание высоты файла
func height(stroka string) int {
	line := strings.Split(stroka, "\n")
	y := len(line)
	return y
}

// Создание 2-мерного массива
func array_2d(y, x int) [][]int {
	array := [][]int{}
	for i := 0; i < y; i++ {
		arrayi := []int{}
		for j := 0; j < x; j++ {
			arrayi = append(arrayi, 0)
		}
		array = append(array, arrayi)
	}
	return array
}

func add_value(array [][]int, file string) [][]int {
	spFile := strings.Split(file, "\n")

	for i := 0; i < len(spFile); i++ {
		lens := spFile[i]
		for j := 0; j < len(lens); j++ {
			if lens[j] == 42 {
				array[i][j] = 1
			}
		}
	}
	return array
}

// Считывание файла
func readfile(file string) [][]int {

	x := lenght(file)
	y := height(file)

	array := array_2d(y, x)

	array = add_value(array, file)

	return array
}

func argument2() string {
	args1 := os.Args[2:]
	file := ""

	for i := 0; i < len(args1); i++ {
		figura, err := os.ReadFile(args1[i])
		if err != nil {
			return err.Error()
		}
		file += string(figura)
	}

	return file
}

func argument1(num int) int {

	flag.IntVar(&num, "num", 0, "num use")

	flag.Parse()

	return num
}

func gLife(array [][]int, cycle_num int) {
	for i := 0; i < cycle_num; i++ {
		field(array)
		array = livedead(array)
	}

}

func main() {
	num := 0
	num = argument1(num)
	figura := argument2()
	array := readfile(figura)
	gLife(array, num)
}

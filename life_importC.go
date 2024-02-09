package main

/*

int predel(char* array,int y,int x,int osx,int osy) {
	if ((y >= 0 && x >= 0) && (y < osy && x < osx)) {
		if (array[y*osx+x] == 1) {
			return 1;
		}
	}
	return 0;
}

int cell_counter(char* array,int y,int x,int osx,int osy) {
	int res = 0;
	res += predel(array, y, x+1, osx, osy);
	res += predel(array, y, x-1, osx, osy);
	res += predel(array, y+1, x, osx, osy);
	res += predel(array, y-1, x, osx, osy);
	res += predel(array, y-1, x-1, osx, osy);
	res += predel(array, y-1, x+1, osx, osy);
	res += predel(array, y+1, x-1, osx, osy);
	res += predel(array, y+1, x+1, osx, osy);
	return res;
}


*/
import "C"
import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Поле
func field(array []byte, y, x int) {
	for i := 0; i < y; i++ {
		fmt.Print("\n")
		for j := 0; j < x; j++ {
			if array[i*x+j] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
	}
}

// подсчет соседей

// Копирования массива
func duplicate(array []byte) []byte {
	duplicate := make([]byte, len(array))
	copy(duplicate, array)
	return duplicate
}

// Проверка клетки на живую или мертвую
func livedead(array []byte, osy, osx int) []byte {
	new_array := duplicate(array)
	for y := 0; y < osy; y++ {
		for x := 0; x < osx; x++ {
			counter := C.cell_counter((*C.char)(C.CBytes(array)), C.int(y), C.int(x), C.int(osx), C.int(osy))
			if (C.int(counter) > 3 || C.int(counter) < 2) && array[y*osx+x] == 1 {
				new_array[y*osx+x] = 0
			}
			if C.int(counter) == 3 && array[y*osx+x] == 0 {
				new_array[y*osx+x] = 1
			}
		}
	}
	return new_array
}

// проверка выхода за предел массива

// Считывание длины файла
func lenght(stroka string) int {
	line := strings.Split(stroka, "\n")
	max := 0
	for i := 0; i < len(line); i++ {
		lineone := len(line[i])
		if max < lineone {
			max = lineone
		}
	}
	//fmt.Printf("lenght = %d", max)
	return max
}

// Считывание высоты файла
func height(stroka string) int {
	line := strings.Split(stroka, "\n")
	y := len(line)
	return y
}

// Создание  массива
func array_1d(y, x int) []byte {
	array := []byte{}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			array = append(array, 0)
		}
	}
	return array
}

// Добавление значений в массив
func add_value(array []byte, file string, x int) []byte {
	spFile := strings.Split(file, "\n")
	for i := 0; i < len(spFile); i++ {
		line := spFile[i]
		for j := 0; j < len(line); j++ {
			if line[j] == 42 {
				array[x*i+j] = 1

			}
		}
	}
	return array
}

// Считывание высоты ширины файла, возвращение массива со значениями
func readfile(file string) []byte {

	x := lenght(file)
	y := height(file)

	array := array_1d(y, x)
	array = add_value(array, file, x)

	return array
}

// Считывание файла с командной строки
func argument2() string {
	args1 := os.Args[2:]
	file := ""

	for i := 0; i < len(args1); i++ {
		figura, err := os.ReadFile(args1[i])
		//fmt.Print(figura)
		if err != nil {
			return err.Error()
		}
		file += string(figura)
	}

	return file
}

// Считывание количество циклов с командной строки
func argument1() int {
	num := 0
	flag.IntVar(&num, "num", 0, "num use")

	flag.Parse()

	return num
}

// Запуск игры Жизнь
func gLife(array []byte, cycle_num, y, x int) {
	for i := 0; i < cycle_num; i++ {
		field(array, y, x)
		array = livedead(array, y, x)
	}

}

func main() {

	num := argument1()
	figura := argument2()
	array := readfile(figura)
	x, y := lenght(figura), height(figura)
	gLife(array, num, y, x)

}

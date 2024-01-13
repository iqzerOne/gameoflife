package main

import "fmt"
   
func field(array [][]int){
    for i := 0; i < len(array);i++{
	fmt.Print("\n")
	for j := 0; j < len(array); j++{
	    if array[i][j] == 0{
		fmt.Print(" ")
	    } else {
	    fmt.Print("*")
	    }
	}
    }
}

func main() {
    array := [][]int{
    {0, 1, 0},
    {0, 0, 1},
    {1, 1, 1},
    }
field(array)
}
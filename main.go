package main

import (
	"fmt"
)

var GLOBAL1 = 100
var GLOBAL2 = -50
var GLOBAL3 = 1

func main() {
	fmt.Println("Введите число:")
	var num int
	fmt.Scan(&num)
	result := f1(f2(f3(num)))

	fmt.Println(result)
}

func f1(i int) int {
	i = i + GLOBAL1 + GLOBAL2
	return i
}

func f2(i int) int {
	i = i + GLOBAL1 + GLOBAL3
	return i
}

func f3(i int) int {
	i = i + GLOBAL2 + GLOBAL3
	return i
}

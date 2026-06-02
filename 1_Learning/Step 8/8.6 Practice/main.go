package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := [5]int{3, 8, 1, 8, 1}
	res := handlerSecretCode(arr)
	fmt.Println(res)
}

func handlerSecretCode(arr [5]int) (sCode string) {
	min, max := searchMinMax(arr) // Поиск мин. значения и макс. значения
	sCode = strconv.Itoa(min)
	for _, v := range arr {
		if v%2 == 0 {
			sCode = sCode + "E" + strconv.Itoa(v) // Префикс четного числа
		} else {
			sCode = sCode + "O" + strconv.Itoa(v) // Префикс нечетного числа
		}
	}
	sCode = sCode + strconv.Itoa(max)
	return
}

func searchMinMax(arr [5]int) (min int, max int) {
	min, max = arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return

}

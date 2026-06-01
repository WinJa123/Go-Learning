// Псевдо случайны числа (взяты и каких-либо данных). Не случайны
package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	random_number := rand.IntN(100) // [0, 100)
	fmt.Println("От 0 до 100. Число:", random_number)

	// Для диапозона
	min := 30
	max := 50
	random_number_2 := rand.IntN(max-min) + min
	// (max - min) ---> 50 - 10 = 40, диапозон 40 чисел (от 0 до 40)
	// + min ---> Получив число от 0 до 40 мы добавляем min и (допустим дисло 5) тогда 5 + 10 = 15
	fmt.Printf("Диапозон. От %v до %v. Число: %v\n", min, max, random_number_2)
}

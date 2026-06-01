package main

import (
	"fmt"
	"os"
)

var (
	weight   float64                                            // Вес
	height   float64                                            // Рост
	bmi      float64 = float64(weight) / float64(height*height) // Формула ИМТ
	category string
)

func main() {
	fmt.Print("Введите вес (кг): ")
	_, err := fmt.Scan(&weight)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Print("Введите рост (см): ")
	_, err = fmt.Scan(&height)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	height /= 100
	bmi = weight / (height * height)

	if bmi < 18.5 {
		category = "Недостаточный вес"
	} else if bmi >= 18.5 && bmi < 25 {
		category = "Нормальный вес"
	} else if bmi >= 25 && bmi < 30 {
		category = "Избыточный вес"
	} else {
		category = "Ожирение"
	}

	fmt.Printf("Ваш ИМТ: %.2f\n", bmi)
	fmt.Printf("Категория: %s\n", category)
}

/*
1. Добавить обработку ошибок для ввода данных, чтобы программа не завершалась при неправильном вводе.
2. Добавить возможность повторного ввода данных, если пользователь ввел некорректные значения.
3. Добавить комментарии к коду для улучшения его читаемости.
4. Разделить код на функции для улучшения структуры и повторного использования.
*/

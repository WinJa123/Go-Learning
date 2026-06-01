package main

import (
	"fmt"
	"time"
)

func testSpeed() {
	start := time.Now()

	count := 0
	for i := 0; i < 1000000000; i++ {
		count++
	}

	duration := time.Since(start)

	fmt.Printf("Результат: %d\n", count)
	fmt.Printf("Время выполнения: %.6f сек\n", duration.Seconds())
}

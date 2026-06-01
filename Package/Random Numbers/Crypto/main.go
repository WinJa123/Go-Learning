// Криптографические числа. Абсолютная случайность
package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

func main() {
	n, err := rand.Int(rand.Reader, big.NewInt(100)) // [0, 100)
	if err != nil {
		log.Fatalf("Ошибка генерации случайного числа: %v\n", err.Error())
	}
	fmt.Println("Случайное число:", n.Int64())
}

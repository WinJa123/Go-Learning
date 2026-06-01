package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	scaner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Введите текст: ")
		if scaner.Scan() {
			echo := scaner.Text()
			if echo == "67" {
				fmt.Println("Пошел нахуй")
				time.Sleep(300 * time.Millisecond)
				break
			}
			fmt.Printf("Ваш текст: %v\n", echo)
		}
		time.Sleep(time.Second)
	}
}

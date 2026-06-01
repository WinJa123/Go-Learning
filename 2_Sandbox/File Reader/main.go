package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	scaner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите путь к файлу...")
	fmt.Println("0 - Завершить")
	for {
		fmt.Print("--> ")
		if scaner.Scan() {
			file_path := scaner.Text()
			if file_path == "0" {
				break
			}
			if err := openFile(file_path); err != nil {
				fmt.Println(err)
			}

		}
	}
}

func openFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	fmt.Println("Content: \n", string(content))
	return nil
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"unicode/utf8"
)

func main() {
	for {
		userInput, err := getInput()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if userInput != "" {
			DisplayResults(CountCharacters(userInput))
		}
	}
}

func getInput() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Введите текст (для очистки консоли нажмите \"Enter\"):  ")
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", fmt.Errorf("\nscan error: %v", err)
		}
		return "", errors.New("\nunable to read input")
	}
	if scanner.Text() != "" {
		return scanner.Text(), nil
	}
	clearCmd()
	return "", nil
}

func CountCharacters(text string) (letters, digits, spaces, punctuation int) {
	letters = utf8.RuneCountInString(text)
	for _, char := range text {
		strDig := string(char)
		if _, err := strconv.Atoi(strDig); err != nil {
			continue
		}
		digits++
		letters--
	}
	for _, char := range text {
		if char == ' ' {
			spaces++
			letters--
		}
	}
	for _, char := range text {
		if char == '.' ||
			char == ',' ||
			char == '?' ||
			char == '!' {
			punctuation++
			letters--
		}
	}
	return
}

func DisplayResults(letters, digits, spaces, punctuation int) {
	fmt.Printf("Количество букв: %d\n", letters)
	fmt.Printf("Количество цифр: %d\n", digits)
	fmt.Printf("Количество пробелов: %d\n", spaces)
	fmt.Printf("Количество знаков препинания: %d\n", punctuation)
}

func clearCmd() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

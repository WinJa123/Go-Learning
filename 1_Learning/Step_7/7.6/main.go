package main

import (
	"bufio"
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
		}
		if userInput != "" {
			DisplayResults(CountCharacters(userInput))
		}
	}
}

func getInput() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Введите текст (для очистки консоли нажмите \"Enter\"):  ")
	if scanner.Scan() {
		if scanner.Text() != "" {
			return scanner.Text(), nil
		}
		clearCmd()
	}
	return "", nil
}

func CountCharacters(text string) (letters, digits, spaces, punctuation int) {
	letters = utf8.RuneCountInString(text)
	digits = 0 // ДОДЕЛАТЬ
	spaces = 0
	punctuation = 0
	for _, searchDigits := range text {
		strDig := string(searchDigits)
		if _, err := strconv.Atoi(strDig); err != nil {
			continue
		}
		digits++
		letters--
	}
	for _, searchSpace := range text {
		if searchSpace == ' ' {
			spaces++
			letters--
		}
	}
	for _, searchPunctuation := range text {
		if searchPunctuation == '.' ||
			searchPunctuation == ',' ||
			searchPunctuation == '?' ||
			searchPunctuation == '!' {
			punctuation++
			letters--
		}
	}
	return
}

func DisplayResults(letters, digits, spaces, punctuation int) {
	fmt.Println("Количество букв:", letters)
	fmt.Println("Количество цифр:", digits)
	fmt.Println("Количество пробелов:", spaces)
	fmt.Println("Количество знаков препинания:", punctuation)
}

func clearCmd() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

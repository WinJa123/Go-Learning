package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет! Ты зашёл на мой сервер на Go!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Сервер запущен на http://localhost:8000")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Ошибка запуска:", err)
	}
}

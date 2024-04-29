package web

import (
	"fmt"
	"log"
	"main.go/logic"
	"net/http"
	"os"
	"strconv"
)

//выводит содержимое файла на ответ HTTP
func render(w http.ResponseWriter, r *http.Request, filePath string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Все фигня давай по новой")
		return
	}
	fmt.Fprintf(w, "%s", file)
}

//обрабатывает запросы и отображает страницу
func serveHandler(w http.ResponseWriter, r *http.Request) {
	render(w, r, "./web/index.html")
}

//брабатывает запросы и обрабатывает введенное пользователем число
func formHandler(w http.ResponseWriter, r *http.Request) {

	render(w, r, "./web/result.html")
	fs := logic.FibonacciService{}
	number, err := strconv.Atoi(r.FormValue("numberValue"))
	if err != nil || number < 0 { // является ли введенное число корректным
		fmt.Fprintf(w, "<h2>Введённое число некорректно!!!</h2>")
	} else {
		fmt.Fprintf(w, "<a>Введённое число: %d\n</a>", number)
		// является ли число Фибоначчи
		if fs.IsFibonacci(number) {
			prev, next := fs.GetAdjacentFibonacci(number)
			fmt.Fprintf(w, "<a>Предыдущее число Фибоначчи: %d</a>", prev)
			fmt.Fprintf(w, "<a>Следующее число Фибоначчи: %d</a>", next)
		} else {
			closest := fs.GetNearestFibonacci(number)
			fmt.Fprintf(w, "<a>Ближайшее число Фибоначчи: %d</a>", closest)
		}
	}

	fmt.Fprintf(w, "<a href=\"/\">Вернуться</a>")
}

//запуск сервера на порту 8080
func StartServer() {
	http.HandleFunc("/", serveHandler)
	http.HandleFunc("/result", formHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
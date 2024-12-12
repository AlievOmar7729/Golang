package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func processSliceGET(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
		<!DOCTYPE html>
		<html>
		<body>
		<form action="/" method="GET">
			<label>Введіть числа через кому: <input type="text" name="numbers"></label><br>
			<input type="submit" value="Обчислити">
		</form>
		`)

	numbersStr := r.URL.Query().Get("numbers")

	if numbersStr != "" {
		numbersSlice := strings.Split(numbersStr, ",")
		var numbers []float64
		sumNegatives := 0.0
		product := 1.0
		valid := true

		for _, numStr := range numbersSlice {
			num, err := strconv.ParseFloat(strings.TrimSpace(numStr), 64)
			if err != nil {
				fmt.Fprintf(w, `<p style="color: red;">Помилка: "%s" не є числом.</p>`, numStr)
				valid = false
				break
			}
			numbers = append(numbers, num)
			if num < 0 {
				sumNegatives += num
			}
			product *= num
		}

		if valid {
			fmt.Fprintf(w, "<h3>Результати:</h3>")
			fmt.Fprintf(w, "<p>Сума від’ємних елементів: %f</p>", sumNegatives)
			fmt.Fprintf(w, "<p>Добуток елементів: %f</p>", product)
		}
	}

	fmt.Fprint(w, "</body></html>")
}

func main() {
	http.HandleFunc("/", processSliceGET)
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func solveQuadraticPOST(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
		<!DOCTYPE html>
		<html>
		<body>
		<form action="/" method="POST">
			<label>a: <input type="text" name="a"></label><br>
			<label>b: <input type="text" name="b"></label><br>
			<label>c: <input type="text" name="c"></label><br>
			<input type="submit" value="Розв'язати">
		</form>
		`)

	if r.Method == "POST" {
		r.ParseForm()
		aStr := r.FormValue("a")
		bStr := r.FormValue("b")
		cStr := r.FormValue("c")

		if aStr != "" && bStr != "" && cStr != "" {
			a, _ := strconv.ParseFloat(aStr, 64)
			b, _ := strconv.ParseFloat(bStr, 64)
			c, _ := strconv.ParseFloat(cStr, 64)

			D := b*b - 4*a*c
			fmt.Fprintf(w, "<h3>Результат:</h3>")
			if D < 0 {
				fmt.Fprint(w, "Немає дійсних коренів.")
			} else if D == 0 {
				x := -b / (2 * a)
				fmt.Fprintf(w, "x = %f", x)
			} else {
				x1 := (-b + math.Sqrt(D)) / (2 * a)
				x2 := (-b - math.Sqrt(D)) / (2 * a)
				fmt.Fprintf(w, "x1 = %f, x2 = %f", x1, x2)
			}
		}
	}

	fmt.Fprint(w, "</body></html>")
}

func main() {
	http.HandleFunc("/", solveQuadraticPOST)
	http.ListenAndServe(":8080", nil)
}

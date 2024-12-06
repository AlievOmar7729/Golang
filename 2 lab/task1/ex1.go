package main

import "fmt"

func main() {
	test1 := findMin(1.5, 3, -10)
	fmt.Println("findMin =", test1)

	test2 := findAverage(1.5, 3, 6)
	fmt.Println("findAverage =", test2)

	test3 := findFirstEquation(3.1, 6)
	fmt.Println("findFirstEquation =", test3)
}

func findMin(x1 float64, x2 float64, x3 float64) float64 {
	var min float64

	min = x1
	if x2 < min {
		min = x2
	}
	if x3 < min {
		min = x3
	}
	return min
}

func findAverage(x1 float64, x2 float64, x3 float64) float64 {
	res := (x1 + x2 + x3) / 3
	return res
}

func findFirstEquation(a float64, b float64) float64 {
	if a == 0 {
		fmt.Printf("Kofiecient {a} ne povinen dorivnyvaty nuly\n")
		return 0 
	}
	res := -b / a
	return res
}

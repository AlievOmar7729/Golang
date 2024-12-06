package main


import (
	"ex2/math"
	"fmt"
)

func main() {
	test1 := math.findMin(1.5, 3, -10)
	fmt.Println("findMin =", test1)

	test2 := math.findAverage(1.5, 3, 6)
	fmt.Println("findAverage =", test2)

	test3 := math.findFirstEquation(3.1, 6)
	fmt.Println("findFirstEquation =", test3)
}
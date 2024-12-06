package main


import (
	"fmt"
	mymath "ex3/math"
	
)

func main() {
	test1 := mymath.findMin(1.5, 3, -10)
	fmt.Println("findMin =", test1)

	test2 := mymath.findAverage(1.5, 3, 6)
	fmt.Println("findAverage =", test2)

	test3 := mymath.findFirstEquation(3.1, 6)
	fmt.Println("findFirstEquation =", test3)
}
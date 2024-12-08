package main

import (
	"fmt"
	"os"
)
func main() {
	var (
		a     int = 1103515245
		c     int = 12345
		x     int
		m     = 1 << 31 // 2^31
		count int = 10000
	)
	fmt.Printf("Введіть перший елемент х (0-99): ")
	fmt.Fscan(os.Stdin, &x)

	for x < 0 || x >= 100 {
		fmt.Printf("Помилка! Значення x введено неправильно! х: ")
		fmt.Fscan(os.Stdin, &x)
	}
	arr := generateRealSequence(a, c, x, m, count)


	fmt.Println("\nПерші 10 значень згенерованої послідовності:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%.4f\n", arr[i])
	}
}


func generateRealSequence(a int, c int, x int, m int, count int) []float64 {
	var arr []float64
	for i := 0; i < count; i++ {
		x = (a*x + c) % m
		randReal := float64(x) / float64(m)
		scaledValue := randReal * 100
		arr = append(arr, scaledValue)
	}
	return arr
}
package main

import (
	"fmt"
	"os"
	"math"
)

func main() {
	var (
		a     int = 1103515245
		c     int = 12345
		x     int
		m     = 1 << 3
		count int
	)

	fmt.Printf("Введіть перший елемент х (0-99): ")
	fmt.Fscan(os.Stdin, &x)
	fmt.Printf("Введіть довжину послідовності: ")
	fmt.Fscan(os.Stdin, &count)

	for x < 0 || x >= 100 {
		fmt.Printf("Помилка! Значення x введено неправильно! х: ")
		fmt.Fscan(os.Stdin, &x)
	}

	for count <= 0 {
		fmt.Printf("Помилка! Значення довжини послідовності введено неправильно! Довжина послідовності: ")
		fmt.Fscan(os.Stdin, &count)
	}

	arr := Congruent(a, c, x, m, count)

	calculateStatistics(arr)
}



func Congruent(a int, c int, x int, m int, count int) []int {
	var arr []int
	for i := 0; i < count; i++ {
		x = (a*x + c) % m
		arr = append(arr, x%100)
	}
	return arr
}

func calculateStatistics(arr []int) {
	var frequency [100]int // масив для частоти появи чисел
	var sum int

	for _, num := range arr {
		frequency[num]++
		sum += num
	}

	mean := float64(sum) / float64(len(arr))

	var variance float64
	for _, num := range arr {
		variance += math.Pow(float64(num)-mean, 2)
	}
	variance /= float64(len(arr))

	stdDev := math.Sqrt(variance)

	fmt.Println("\nЧастота появи чисел та статистична ймовірність:")
	for i, freq := range frequency {
		if freq > 0 {
			probability := float64(freq) / float64(len(arr))
			fmt.Printf("%d: частота = %d, ймовірність = %.4f\n", i, freq, probability)
		}
	}


	fmt.Printf("\nМатематичне сподівання: %.4f\n", mean)
	fmt.Printf("Дисперсія: %.4f\n", variance)
	fmt.Printf("Середньоквадратичне відхилення: %.4f\n", stdDev)
}

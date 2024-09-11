package main

import "fmt"

func main() {
	var x, y, z uint8

	x = 9
	y = 28
	z = x

	fmt.Println("Битовые операции")

	fmt.Printf("^x      = ^(%d)      = ^(%.8b)            = %.8b = %d\n", x, x, ^x, ^x)  
	//оператор ^ копіює біт в результат, якщо біт задани в одному из операндів, но не в обох

	fmt.Printf("x << 2  = (%d << 2)  = (%.8b << 2)        = %.8b = %d\n", x, x, x<<2, x<<2)
	// Зсув вліво на 2 біти

	fmt.Printf("x >> 2  = (%d >> 2)  = (%.8b >> 2)        = %.8b = %d\n", x, x, x>>2, x>>2)
	// Зсув вправо на 2 біти

	fmt.Printf("x & y   = (%d & %d)  = (%.8b & %.8b)  = %.8b = %d\n", x, y, x, y, x&y, x&y)
	// Побітова операція AND

	fmt.Printf("x | y   = (%d | %d)  = (%.8b | %.8b)  = %.8b = %d\n", x, y, x, y, x|y, x|y)
	// Побітова операція OR

	fmt.Printf("x ^ y   = (%d ^ %d)  = (%.8b ^ %.8b)  = %.8b = %d\n", x, y, x, y, x^y, x^y)
	// Побітова операція XOR

	fmt.Printf("x &^ y  = (%d &^ %d) = (%.8b &^ %.8b) = %.8b = %d\n", x, y, x, y, x&^y, x&^y)
	// Побітова операція AND NOT

	fmt.Printf("x %% y   = (%d %% %d)  = (%.8b %% %.8b)  = %.8b = %d\n", x, y, x, y, x%y, x%y)
	// Залишок від ділення змінної x на y


	fmt.Println("\nБитовые операции с присваиванием")

	x = z
	x &= y
	fmt.Printf("x &= y   = (%d &= %d)  = (%.8b &= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)
	// Побітова операція AND з присвоєнням

	x = z
	x |= y
	fmt.Printf("x |= y   = (%d |= %d)  = (%.8b |= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)
	// Побітова операція OR з присвоєнням

	x = z
	x ^= y
	fmt.Printf("x ^= y   = (%d ^= %d)  = (%.8b ^= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)
	// Побітова операція XOR з присвоєнням

	x = z
	x &^= y
	fmt.Printf("x &^= y  = (%d &^= %d) = (%.8b &^= %.8b) = %.8b = %d\n", z, y, z, y, x, x)
	// Побітова операція AND NOT з присвоєнням

	x = z
	x %= y
	fmt.Printf("x %%= y   = (%d %%= %d)  = (%.8b %%= %.8b)  = %.8b = %d\n", z, y, z, y, x, x)
	// Операція залишку від ділення з присвоєнням

}

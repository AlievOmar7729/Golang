package main

import (
	"fmt"
	"reflect" // Импорт пакета reflect
)

func main() {
	fmt.Println("Синонимы целых типов\n")
	fmt.Println("byte    - int8")
	fmt.Println("rune    - int32")
	fmt.Println("int     - int32, или int64, в зависимости от ОС")
	fmt.Println("uint    - uint32, или uint64, в зависимости от ОС")

	// Задание.
	// 1. Определить разрядность ОС
	var x int
	y := reflect.TypeOf(x)
	fmt.Printf("Ваша операційна система: %d-біт\n", y.Size()*8)
}


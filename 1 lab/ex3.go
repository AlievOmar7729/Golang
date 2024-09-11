package main

import "fmt"

func main() {
	//Инициализация переменных
	var userinit8 uint8 = 1
	var userinit16 uint16 = 2
	var userinit64 int64 = -3
	var userautoinit = -4 //Такой вариант инициализации также возможен

	fmt.Println("Values: ", userinit8, userinit16, userinit64, userautoinit, "\n")

	//Краткая запись объявления переменной
	//только для новых переменных
	intVar := 10

	fmt.Printf("Value = %d Type = %T\n", intVar, intVar)

	//Задание.
	
	
	//1. Вывести типы всех переменных
	fmt.Printf("Тип змінної userinit8 = %T\nТип змінної userinit16 =%T\nТип змінної int64 =%T\nТип змінної userautoinit =%T\n", userinit8, userinit16, userinit64, userautoinit)
	//2. Присвоить переменной intVar переменные userinit16 и userautoinit. Результат вывести.
	intVar = userautoinit
	fmt.Printf("Value = %d Type = %T\n", intVar, intVar)
	intVar = int(userinit8)  // якщо не призводити до типу , то виникає помилка - cannot use userinit8 (variable of type uint8) as int value in assignment
	fmt.Printf("Value = %d Type = %T\n", intVar, intVar)
}

package main

import "fmt"

func main() {
	//Все переменные инициализируются значением по умолчанию, т.е. 0
	var defaultInt8 int8
	var defaultInt16 int16
	var defaultInt32 int32
	var defaultInt64 int64
	var defaultInt int

	fmt.Println("Default values (signed): ", defaultInt8, defaultInt16, defaultInt32, defaultInt64, defaultInt)

	var defaultuInt8 uint8
	var defaultuInt16 uint16
	var defaultuInt32 uint32
	var defaultuInt64 uint64
	var defaultuInt uint

	fmt.Println("Default values (unsigned): ", defaultuInt8, defaultuInt16, defaultuInt32, defaultuInt64, defaultuInt)

	//Задание.
	//1. Создать целочисленную переменную (результат не отображать)

	//var x int  - якщо ми не використовуєму змінну , то в go це помилка
	var _y byte
	var _z int32 // при використанні _ ми робимо пусту  змінну , яку go ігнорує

}

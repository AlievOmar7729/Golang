package main

import "fmt"

func main() {
	var first, second bool
	var third bool = true
	fourth := !third
	var fifth = true

	fmt.Println("first  = ", first)       // false  - по замовчуванню
	fmt.Println("second = ", second)      // false - по замовчуванню
	fmt.Println("third  = ", third)       // true - операнд був призначений
	fmt.Println("fourth = ", fourth)      // false - інверсія від third , якому було призначено true
	fmt.Println("fifth  = ", fifth, "\n") // true - операнд був призначений

	fmt.Println("!true  = ", !true)        // false - інверсія
	fmt.Println("!false = ", !false, "\n") // true - інверсія

	fmt.Println("true && true   = ", true && true)         // true - обидва операнда true (&&)
	fmt.Println("true && false  = ", true && false)        // false - один із операндів false (&&)
	fmt.Println("false && false = ", false && false, "\n") // false - обидва операнда false (&&)

	fmt.Println("true || true   = ", true || true)         // true - обидва операнда true (||)
	fmt.Println("true || false  = ", true || false)        // true - один із операндів true (||)
	fmt.Println("false || false = ", false || false, "\n") // false - обидва операнда false (||)
	// (||) - або 
	// (&&) - і

	fmt.Println("2 < 3  = ", 2 < 3)        // true
	fmt.Println("2 > 3  = ", 2 > 3)        // false
	fmt.Println("3 < 3  = ", 3 < 3)        // false
	fmt.Println("3 <= 3 = ", 3 <= 3)       // true
	fmt.Println("3 > 3  = ", 3 > 3)        // false
	fmt.Println("3 >= 3 = ", 3 >= 3)       // true
	fmt.Println("2 == 3 = ", 2 == 3)       // false
	fmt.Println("3 == 3 = ", 3 == 3)       // true
	fmt.Println("2 != 3 = ", 2 != 3)       // true
	fmt.Println("3 != 3 = ", 3 != 3, "\n") // false
	// за законами математики 
	// < - меньше 
	// > - більше 
	// <= - меньше або дорівнює
	// >= - більше або дорівнює 
	// != - не дорівнює 
	// == - дорівнює

}

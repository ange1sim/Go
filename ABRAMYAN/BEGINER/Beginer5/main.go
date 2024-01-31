package main

//Дана длина ребра куба a. Найти объем куба V = (a)3 и площадь его поверхности S = (6·a)2
import "fmt"

func main() {
	var a int
	fmt.Println("Введите длину стороны куба: ")
	fmt.Scan(&a)
	V := a*a*a
	S := (6 * a)*(6*a)
	fmt.Println("Объёь куба равен: ", V)
	fmt.Println("Площадь поверхности куба равен: ", S)
}
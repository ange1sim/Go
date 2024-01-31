package main

import "fmt"
// Необходимо блоки кода разбить на функции. К примеру, блок кода со сложением двух переменных можно вынести в отдельную функцию.
func difference(num1, num2 int) int {
	return num1-num2
}
func main(){
	var num1, num2 int
	fmt.Println("Введите 1-е число: ")
	fmt.Scan(&num1)
	fmt.Println("Введите 2-е число: ")
	fmt.Scan(&num2)
	result :=difference(num1, num2)
	fmt.Println("Разница чисел: ", result)
}

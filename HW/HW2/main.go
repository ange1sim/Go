package main

import "hw/pkg"


import "fmt"
// Необходимо блоки кода разбить на функции. К примеру, блок кода со сложением двух переменных можно вынести в отдельную функцию.

func main(){
	
	var num1, num2 float64
	fmt.Println("Введите 1-е число: ")
	fmt.Scan(&num1)
	fmt.Println("Введите 2-е число: ")
	fmt.Scan(&num2)
	result := pkg.Diff(num1, num2)
	fmt.Println("Разница чисел: ", result)
	result = pkg.Division(num1, num2)
	fmt.Println("Частное чисел: ", result)
	result = pkg.Mult(num1, num2)
	fmt.Println("Произведение чисел: ", result)
	result = pkg.Sum(num1, num2)
	fmt.Println("Сумма чисел: ", result)
}
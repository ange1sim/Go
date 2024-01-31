package main

import "fmt"

//площадь прямоугольника
func main() {
	var a, b int

	fmt.Println("Введите ширину:")
	fmt.Scan(&a)
	fmt.Println("Введите длинну:")
	fmt.Scan(&b)

	fmt.Println("периметр прямоугольника =", a*b)
}

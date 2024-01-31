package main

//Даны два числа a и b. Найти их среднее арифметическое: (a + b)/2.
import "fmt"

func main() {
	var a, b int
	fmt.Println("Введите 1-е число: ")
	fmt.Scan(&a)
	fmt.Println("Введите 2-е число: ")
	fmt.Scan(&b)

	fmt.Println("среднее арифметическое: ", (a+b)/2)
}

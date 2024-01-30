package main

//Дана сторона квадрата a. Найти его периметр P = 4·a
import "fmt"

func main() {
	var a int
	fmt.Println("Введите сторону квадрата: ")
	fmt.Scan(&a)
	fmt.Println("Периметр равен = ", 4*a)
}

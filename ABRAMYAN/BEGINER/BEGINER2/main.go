package main

//Дана сторона квадрата a. Найти его площадь S = (a)2
import "fmt"

func main() {
	var a int
	fmt.Println("Введите сторону квадрата: ")
	fmt.Scan(&a)
	fmt.Println("Площадь квадрата равна: ", a*a)
}

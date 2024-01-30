package main

/* Дан диаметр окружности d. Найти ее длину L = π·d. В качестве
значения π использовать 3.14
*/
import "fmt"

func main() {
	var d float64
	pi := 3.14
	fmt.Println("Введите диаметр окружности: ")
	fmt.Scan(&d)
	L := pi * d
	fmt.Println("Длина окружности равна: ", L)
}

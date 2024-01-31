package main
// Найти длину окружности L и площадь круга S заданного радиуса R: L = 2·π·R, S = (π·R)2. В качестве значения π использовать 3.14
import "fmt"
func main() {
	var R float64
	pi := 3.14
	fmt.Println("Введите радиус окружности: ")
	fmt.Scan(&R)
	
	L := 2*pi*R
	S := (pi*R)*(pi*R)
	fmt.Println("Длинна окружности равна: ", L)
	fmt.Println("Площадь окружности равна: ", S)
}
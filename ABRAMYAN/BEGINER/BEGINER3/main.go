package main
/*Даны стороны прямоугольника a и b. Найти его площадь S = a·b и
периметр P = 2·(a + b).
*/
import "fmt"

func main() {
	var a, b int
	fmt.Println("Введите ширину прямоугольника :")
	fmt.Scan(&a)

	fmt.Println("Введите длинну прямоугольника :")
	fmt.Scan(&b)

	fmt.Println("Площадь прямоугольника: ", a*b)
	fmt.Println("Периметр прямоугольника: ", 2*(a+b))
}
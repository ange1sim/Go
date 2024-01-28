package main

import "fmt"
//периметр квадрата
func main() {
	var a int
	fmt.Println("Ведите сторону квадрата:")
	fmt.Scan(&a)

	fmt.Println("Периметр квадрата равен: ",a*4)	
}

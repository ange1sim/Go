package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	log.Println("Программа запущена")
	var num float64
	fmt.Println("Ведите ваш возраст:")
	fmt.Scan(&num) // 35
	num = num/10
	num2 := math.Ceil(num)
	num2 = (num2*10)-(num*10) 
	fmt.Println("До юбилея вам осталось:", num2)
	log.Println("Программа завершена")
}

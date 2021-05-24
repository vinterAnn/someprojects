package main

import "fmt"

func main() {
	fmt.Println("Введите число")
	var a int
	fmt.Scan(&a)
	fmt.Print("Числовой ряд от 0 до ", a, ":\n")
	for i := 0; i <= a; i++ {
		fmt.Println(i)
	}
}

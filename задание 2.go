package main

import "fmt"

func main() {
	fmt.Println("Введите 2 числа, которые нужно сложить:")
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	sum := 0
	for i := 0; i <= b; i++ {
		sum = a + i
		//Чтобы вывести каждый шаг сложения:
		//fmt.Println(sum)
	}
	fmt.Print("Сумма ", a, " и ", b, " равна ", sum, "\n")
}

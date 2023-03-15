package main

import "fmt"

func main() {
	var num1, num2 int
	for {
		fmt.Println("digite o numero de dias trabalhados:")
		fmt.Scan(&num1)
		if num1 <= 0 {
			fmt.Println("tem que ter trabalhado um dia pelo menos")
			continue
		}
		if num1 > 0 {
			break
		}
	}
	fmt.Println("digite o valor diário:")
	fmt.Scan(&num2)
	salario := num1 * num2
	fmt.Printf("o seu salário é de: %d", salario)
}

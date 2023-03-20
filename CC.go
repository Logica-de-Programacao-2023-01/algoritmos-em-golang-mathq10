package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Println("digite um número:")
	fmt.Scan(&num1)
	fmt.Println("digite o segundo número:")
	fmt.Scan(&num2)
	fmt.Println("digite o terceiro numero:")
	fmt.Scan(&num3)
	if num1 < num2 && num1 < num3 {
		fmt.Printf("%d é menor que todos os outros números acima", num1)
	} else if num2 < num1 && num2 < num3 {
		fmt.Printf("%d é menor que todos os outros números acima", num2)
	} else if num3 < num1 && num3 < num2 {
		fmt.Printf("%d é menor que todos os outros números acima", num3)
	} else {
		fmt.Printf("todos os números são iguais")
	}

}

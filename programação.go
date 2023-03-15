package main

import "fmt"

func main() {
	var num float64
	fmt.Println("digite o preço do produto:")
	fmt.Scan(&num)
	valor_com_desconto := (num * 0.1) + num
	fmt.Printf("o valor com desconto desse produto é de: %.2f", valor_com_desconto)
}

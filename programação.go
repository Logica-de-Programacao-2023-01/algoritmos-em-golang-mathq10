package main

import "fmt"

func main() {
	var num int
	fmt.Println("digite um número inteiro:")
	fmt.Scan(&num)

	dobro := num * 2
	triplo := num * 3
	quadruplo := num * 4
	fmt.Println("o dobro desse número inteiro é", dobro)
	fmt.Println("o triplo desse número inteiro é", triplo)
	fmt.Println("o quadruplo desse numero inteiro é", quadruplo)
}

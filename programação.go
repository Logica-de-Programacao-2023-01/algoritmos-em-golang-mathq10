package main

import "fmt"

func main() {
	var num int
	fmt.Println("digite um número:")
	fmt.Scan(&num)
	antecessor := num - 1
	sucessor := num + 1
	fmt.Printf("o antecessor de %d é %d e o seu sucessor é %d", num, antecessor, sucessor)
}

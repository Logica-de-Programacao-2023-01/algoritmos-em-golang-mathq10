package main

import "fmt"

func main() {
	var peso float64
	fmt.Println("digite seu peso em kg:")
	fmt.Scan(&peso)
	libra := peso / 0.45
	fmt.Printf("seu peso em libras é de: %.2f", libra)
}

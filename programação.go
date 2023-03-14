package main

import "fmt"

func main() {
	var peso float64
	var altura float64
	fmt.Println("digite seu peso em kg:")
	fmt.Scan(&peso)
	fmt.Println("dgite sua altura em metros:")
	fmt.Scan(&altura)
	IMC := peso / (altura * altura)
	if altura <= 0 {
		fmt.Print("a altura tem que ser maior que 0")
		return
	}
	fmt.Printf("o seu IMC é: %.2f\n", IMC)
}

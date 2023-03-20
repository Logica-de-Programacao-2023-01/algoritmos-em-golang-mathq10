package main

import "fmt"

func main() {
	var altura, peso, imc float64
	var sexo string
	fmt.Println("qual sua altura?")
	fmt.Scan(&altura)
	fmt.Println("qual seu sexo?")
	fmt.Scan(&sexo)
	fmt.Println("qual seu peso?")
	fmt.Scan(&peso)
	imc := peso / (altura * altura)
	if sexo == "homem" {
		if imc < 18
			fmt.Printf("você esta abaixo do peso ideal")
	} else if imc >= 18 && imc <= 25 {
		fmt.Printf("você esta no peso ideal, parabens")
	} else {
		fmt.Printf("você esta acima do peso")
	} else if sexo == "mulher" {
		if imc < 18 {
			fmt.Printf("você esta abaixo do peso ideal")
		} else if imc >= 18 && imc <= 24
			fmt.Printf("você esta no peso ideal, parabens")
	} else {
		fmt.Printf("você esta acima do peso ideal")
	}
}



}

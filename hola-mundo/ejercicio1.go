package main

import (
	"fmt"
)

/// cambio!!! /////

func main() {
	fmt.Println("Escriba un numero")
	var numero1 int
	fmt.Scanln(&numero1)
	fmt.Println("Escriba otro numero")
	var numero2 int
	fmt.Scanln(&numero2)
	var contador int
	if numero1%2 != 0 || numero2%2 != 0 {
		for i := numero1; i <= numero2; i++ {
			if i%2 != 0 {
				contador++
				if i == numero1 || i == numero2 {
					fmt.Printf("%d es el numero impar\n", i)
				}
			}
		}
	} else {
		for i := numero1; i <= numero2; i++ {
			if i%2 != 0 {
				contador++
			}
		}
		fmt.Println("no escribiste ningun numero inpar")
	}
	fmt.Printf("Entre %d y %d hay %d numeros impares.", numero1, numero2, contador)

}

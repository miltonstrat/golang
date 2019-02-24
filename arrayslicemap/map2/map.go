package main

import "fmt"

func main() {
	funcESalario := map[string]float64{
		"José João":      11300.25,
		"Gabriela Silva": 12000.10,
		"Pedro Jr":       1000.1,
	}

	funcESalario["Rafael filho"] = 1300.1
	delete(funcESalario, "inexistente")

	for nome, salario := range funcESalario {
		fmt.Println(nome, salario)
	}
}

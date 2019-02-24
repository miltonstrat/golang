package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15480.78,
			"Guga Pereira":   8400.32,
		},
		"J": {
			"Jose Silva": 11220.32,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}

	delete(funcsPorLetra, "P")
	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}

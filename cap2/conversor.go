// EXEMPLO 1: CONVERSOR DE MEDIDAS
//O primeiroexemplo é um simplesconversor de medidas.
//Ele aceita como en- trada uma lista de valores com sua unidade demedida, e produz uma lista de valores convertidos.
//Por questões didáticas, o conversor trabalha apenas com dois tipos de conversões: de graus Celsius para Fahrenheit,
//e de quilômetros pra milhas.

// Exemplo de execução do programa
// go run cap2/conversor.go 32 celsius
// => Saída: 32.00 celsius = 89.60 fahrenheit

// Exemplo com vários valores
// go run cap2/conversor.go 32 27.4 3 0 celsius
// => Saída
//32.00 celsius = 89.60 fahrenheit
//27.40 celsius = 81.32 fahrenheit
//3.00 celsius = 37.40 fahrenheit
//0.00 celsius = 32.00 fahrenheit

// Exemplo de execução do programa
// go run cap2/conversor.go 10 quilometros
// => Saída: 10.00 quilometros = 6.21 milhas

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: conversor <valores> <unidade>")
		os.Exit(1)
	}

	unidadeOrigem := os.Args[len(os.Args)-1]
	valoresOrigem := os.Args[1 : len(os.Args)-1]

	var unidadeDestino string

	if unidadeOrigem == "celsius" {
		unidadeDestino = "fahrenheit"
	} else if unidadeOrigem == "quilometros" {
		unidadeDestino = "milhas"
	} else {
		fmt.Printf("%s não é uma unidade conhecida!", unidadeOrigem)
		os.Exit(1)
	}

	for i, v := range valoresOrigem {
		valorOrigem, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf("O valor %s na posição %d não é um número válido!\n", v, i)
			os.Exit(1)
		}

		var valorDestino float64

		if unidadeOrigem == "celsius" {
			valorDestino = valorOrigem*1.8 + 32
		} else {
			valorDestino = valorOrigem / 1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s\n",
			valorOrigem, unidadeOrigem, valorDestino, unidadeDestino)
	}
}

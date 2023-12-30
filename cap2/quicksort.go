// EXEMPLO 2: QUICKSORT E FUNÇÕES
//O segundo exemplo é uma implementação simples do famoso algoritmo de ordenação quicksort.
//A ideia básica deste algoritmo é:
//1) Eleger um elementoda lista como pivô e removê-lo da lista;
//2) Particionar a lista em duaslistas distintas: uma contendo elementos me- nores que o pivô, e outra
//contendo elementos maiores;
//3) Ordenar as duas listas recursivamente;
//4) Retornar a combinação da lista ordenada de elementos menores, o próprio pivô, e a lista ordenada
//de elementos maiores.

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1:]
	numbers := make([]int, len(input))

	for i, n := range input {
		number, err := strconv.Atoi(n)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		numbers[i] = number

	}

	fmt.Println(quicksort(numbers))
}

func quicksort(numbers []int) []int {
	if len(numbers) < 1 {
		return numbers
	}

	numbersCopy := make([]int, len(numbers))
	copy(numbersCopy, numbers)

	indexPivot := len(numbersCopy) / 2
	pivot := numbersCopy[indexPivot]

	numbersCopy = append(numbersCopy[:indexPivot], numbersCopy[indexPivot+1:]...)

	minors, bigger := split(numbersCopy, pivot)

	return append(append(quicksort(minors), pivot), quicksort(bigger)...)
}

func split(numbers []int, pivot int) (miniors []int, bigger []int) {
	for _, number := range numbers {
		if number <= pivot {
			miniors = append(miniors, number)
		} else {
			bigger = append(bigger, number)
		}
	}

	return miniors, bigger
}

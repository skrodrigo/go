package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	~int | ~int64 | ~float64
}

type MyNumber int

func Generics[T Number](m map[string]T) T {
	var soma T

	for _, v := range m {
		soma += v
	}
	return soma
}

func Comparable[T comparable](num1 T, num2 T) T {
	if num1 == num2 {
		fmt.Println("Igual")
	}
	fmt.Println("Não Igual")
	return num1
}

func Constraints[T constraints.Ordered](num1 T, num2 T) T {
	if num1 > num2 {
		fmt.Println("Maior")
	}
	fmt.Println("Menor")
	return num1
}


package main

import (
	"fmt"
	"math/big"
)

func main() {
	bigIntA := big.NewInt(33554432)
	bigIntB := big.NewInt(16777216)

	fmt.Println("Результат сложения: ", new(big.Int).Add(bigIntA, bigIntB))

	fmt.Println("Результат вычитания: ", new(big.Int).Sub(bigIntA, bigIntB))

	fmt.Println("Результат умножения: ", new(big.Int).Mul(bigIntA, bigIntB))

	fmt.Println("Результат деления: ", new(big.Int).Div(bigIntA, bigIntB))

}

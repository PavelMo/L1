package main

import "fmt"

func main() {
	var (
		startD int64
		i, bit int
	)
	startD = 10
	i = 2
	bit = 0
	fmt.Printf("Начальное число: %b \n", startD)
	switch bit {
	case 1:
		//Используем поразрядное сложение,если нужно установить iый бит 1
		startD = startD | (1 << (i - 1))
	case 0:
		//Сбрасываем бит,если нужно установить iый бит 0
		startD = startD &^ (1 << (i - 1))
	}
	fmt.Printf("Число после преобразования: %b \n", startD)
}

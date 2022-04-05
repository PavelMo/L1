package main

import (
	"fmt"
)

func main() {
	digits := make(chan int)
	doub := make(chan int)

	arr := []int{2, 3, 5, 6, 8, 9, 12}
	//В отедльной горутине пишем значения в канал
	go in(arr, digits)
	//В другой читаем значения из первого канала и во второй отправляем *2 значение
	go double(digits, doub)

	for val := range doub {
		fmt.Println(val)
	}
}
func in(arr []int, digits chan<- int) {
	for _, val := range arr {
		digits <- val
	}
	close(digits)
}
func double(digits <-chan int, double chan<- int) {
	for val := range digits {
		double <- val * 2
	}
	close(double)
}

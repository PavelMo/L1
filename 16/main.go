package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{4, 4, -30, -30, 1, 39, 44, 4, 1, -30}

	fmt.Println(sort(arr))
	//Метод Ints из пакета sort реализует быструю сортировку
	/*
		arr := []int{4, 0, -30, 99, 53, 181, -99}
		fmt.Println("До сортировки:", intSlice) //До сортировки: [4 0 -30 99 53 181 -99]
		sort.Ints(intSlice)
		fmt.Println("После:", intSlice)//После: [-99 -30 0 4 53 99 181]
	*/

}

//Делим массив на 3 части , что бы не обрабатывать вхождения одинаковых чисел
func sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[rand.Intn(len(arr))]

	left, mid, right := []int{}, []int{}, []int{}

	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		} else {
			mid = append(mid, v)
		}
	}
	return append(sort(left), append(mid, sort(right)...)...)
}

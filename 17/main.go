package main

import (
	"fmt"
)

func main() {
	arr := []int{-99, -10, -3, 3, 5, 12, 82, 99, 101, 233}
	i, ok := search(arr, 101)
	fmt.Printf("%s %d", ok, i)
	//SearchInts из пакета sort реализует бинарный поиск
	/*
		i := sort.SearchInts(arr, 12)
		fmt.Printf("Found %d at index %d in %v\n", 12, i, arr)//Found 12 at index 5 in [-99 -10 -3 3 5 12 82 99 101 233]
	*/
}
func search(arr []int, target int) (int, string) {
	var min int
	max := len(arr) - 1

	for min <= max {
		mid := (max + min) / 2

		switch {
		case arr[mid] == target:
			return mid, "Found in position:"
		case arr[mid] > target:
			max = mid - 1
		case arr[mid] < target:
			min = mid + 1
		}
	}
	return -1, "Not found"
}

package main

import "fmt"

func main() {
	var i int
	arr := []int{2, 4, 6, 8, 10}
	i = 1

	arr = append(arr[:i-1], arr[i:]...)

	fmt.Println(arr)

}

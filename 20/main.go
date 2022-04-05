package main

import (
	"fmt"
	"strings"
)

func main() {
	var res string
	st := "snow dog sun"

	arr := strings.Fields(st)

	for i := len(arr) - 1; i >= 0; i-- {
		res += arr[i] + " "
	}

	fmt.Println(res)
}

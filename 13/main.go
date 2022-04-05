package main

import "fmt"

func main() {
	a := 99
	b := 19
	a, b = b, a

	fmt.Println(a, b)
	/*
		a += b
		b =int(math.Abs(float64(a - b)))
		a -= b
	*/
}

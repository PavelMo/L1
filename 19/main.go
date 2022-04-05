package main

import "fmt"

func main() {
	st := ("á Ö  Ԫ  ಊ  ವ главрыба ")
	//Что бы корректно обрабатывать unicode символы используем руны
	r := []rune(st)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	fmt.Println(string(r))
}

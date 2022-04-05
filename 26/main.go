package main

import (
	"fmt"
	"strings"
)

func main() {

	st := "abcd"
	fmt.Printf("is this string contains only unique characters: \"%s\"  ? %t \n", st, uniqueChars(st))
	st = "abCdefAaf"
	fmt.Printf("is this string contains only unique characters: \"%s\" ? %t \n", st, uniqueChars(st))
	st = "aabcd"
	fmt.Printf("is this string contains only unique characters: \"%s\"  ? %t \n", st, uniqueChars(st))
}
func uniqueChars(st string) bool {
	b := true
	mp := make(map[string]int)
	st = strings.ToLower(st)
	//Проходимся по каждому символу в строке и пишем в мапу
	for _, val := range st {
		if _, ok := mp[string(val)]; ok {
			return false
		} else {
			mp[string(val)]++
		}
	}
	return b
}

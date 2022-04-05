package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	res := []string{}
	mp := make(map[string]bool)

	for _, val := range arr {
		//Если не находим значение в мапе, то записываем его туда и в собственное множество
		if mp[val] != true {
			mp[val] = true
			res = append(res, val)
		}
	}
	fmt.Println(res)
}

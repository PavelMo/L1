package main

import "fmt"

func main() {
	arr1 := []int{3, 4, 6, 9, 2, 1}
	arr2 := []int{1, 3, 3, 5, 7, 9}
	res := []int{}
	//Создадим хэш таблицу и запишем в него значения,например, из первого массива
	mp := make(map[int]bool)

	for _, val := range arr1 {
		mp[val] = true
	}
	//Пройдём по элементам второго массива и если находим по ключу значение, то удалим это значение и запишем в пересечение
	for _, val := range arr2 {
		if _, ok := mp[val]; ok {
			delete(mp, val)
			res = append(res, val)
		}
	}

	fmt.Println(res)
}

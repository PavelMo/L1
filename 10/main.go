package main

import (
	"fmt"
)

func main() {
	input := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	//Создадим хэш таблицу для записи подмножеств с шагом 10 градусов
	resMap := make(map[int][]float64)
	for _, val := range input {
		//Округлим знаение до десятков
		key := int(val) / 10 * 10
		//По этому ключу добавим число в подмножество
		resMap[key] = append(resMap[key], val)
	}
	fmt.Println(resMap)
}

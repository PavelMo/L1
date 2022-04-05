package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	//Используем мьютекс, чтобы не было гонки потоков
	var mu sync.Mutex

	var sum int
	arr := []int{2, 4, 6, 8, 10}

	for _, val := range arr {
		wg.Add(1)
		//В функции работаем с указателем в памяти на сумму
		go SumOfSquares(val, &sum, &mu, &wg)
	}
	wg.Wait()
	fmt.Println(sum)
}

func SumOfSquares(d int, sum *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	*sum = *sum + d*d
	mu.Unlock()

}

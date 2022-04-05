package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var (
		//С помощью мьютекса только одная гоурутина сможет записывать данные в map
		//Что бы разрешить многим горутинам читать из мапы , нужно использовать RWMutex
		mu sync.Mutex
		wg sync.WaitGroup
	)

	rand.Seed(time.Now().UnixNano())
	dataMp := make(map[int]int)

	for i := 0; i < 100; i++ {
		//Запускаем гоурутины
		wg.Add(1)
		go write(i, i+1, dataMp, &mu, &wg)
	}

	wg.Wait()
	fmt.Println(dataMp)
}
func write(data1, data2 int, dataMp map[int]int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	dataMp[data1] = data2
	mu.Unlock()
}

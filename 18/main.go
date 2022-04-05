package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

//Что бы не было гонки потоков используем мьютексы
var mu sync.Mutex

type someStruct struct {
	//..........
	counter int
	//..........
}

func main() {
	st := new(someStruct)

	for i := 0; i < 873; i++ {
		wg.Add(1)
		go Increment(st)
	}

	wg.Wait()
	fmt.Println(st.counter)
}

func Increment(st *someStruct) {
	defer wg.Done()
	mu.Lock()
	st.counter++
	mu.Unlock()
}

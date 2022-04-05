package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var (
		N       int
		timeout bool
		wg      sync.WaitGroup
	)
	N = 1

	data := make([]interface{}, 0)
	data = append(data, 1, nil, 4, 33, "ok", "string", 99.213, true)

	write := make(chan interface{})

	//Запускаем горутину, которая будет читать входящие данные
	wg.Add(1)
	go read(write, &wg)

	//Создаём таймер
	timer := time.NewTimer(time.Duration(N) * time.Second)
	//Ожидаем получения тайм аута в канал C
	for !timeout {
		select {
		case <-timer.C:
			fmt.Println("Time out...")
			close(write)
			wg.Wait()
			timeout = true
		default:
			send := data[rand.Intn(len(data))]
			fmt.Println("Sending data:", send)
			write <- send

		}

	}

}
func read(write <-chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range write {
		fmt.Printf("get data type: %T data: %v \n\n", data, data)
	}
	fmt.Println("Stop reading...\n")
}

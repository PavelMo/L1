package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
)

func main() {
	var (
		//Используем контекст для остановки гоурутин
		ctx, cancel = context.WithCancel(context.Background())
		n           int
		wg          sync.WaitGroup
	)

	n = 10

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	//Какие-то данные
	data := make([]interface{}, 0)
	data = append(data, 1, nil, 4, 33, "ok", "string", 99.213, true)

	mainCh := make(chan interface{}, n)
	defer close(mainCh)

	//Запускаем n воркеров
	for id := 1; id <= n; id++ {
		wg.Add(1)
		go work(id, mainCh, ctx, &wg)
	}
	//Останавливаем воркеров, когда получаем сигнал от системы
	for {
		select {
		case <-signalChan:
			fmt.Println("Stopping program...\n")
			cancel()
			wg.Wait()
			fmt.Println("\nAll workers stopped......Shutting down")
			return
		default:
			mainCh <- data[rand.Intn(len(data))]
		}
	}

}

//читаем данные из основного канала
func work(id int, mainCh <-chan interface{}, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case data := <-mainCh:
			fmt.Printf("Worker %d get data type: %T, data: %v \n\n", id, data, data)
		case <-ctx.Done():
			fmt.Printf("Worker %d finished working \n", id)
			return
		}

	}
}

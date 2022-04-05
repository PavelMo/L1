package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//С помощью контекста
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Received a signal from context to stop goruite...")
				return
			default:
				//Какая-то работа
			}
		}
	}()
	cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("Stopped gourutine 1.... \n")

	//С помощью канала
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Received a signal from channel to stop goruite...")
				close(done)
				return
			default:
				//Какая-то работа
			}
		}
	}()

	done <- struct{}{}
	time.Sleep(1 * time.Second)
	fmt.Println("Stopped gourutine 2.... \n")
	//Так же есть такие способы остановки горутины с помощью канала:
	/*
			//С помощью закрытия канала
		ch := make(chan struct{})
		go func() {
			for val := range ch {
				_ = val
			}
			fmt.Println("Channel is closed, stopping gorutine...")
		}()
		time.Sleep(1 * time.Second)
		close(ch)
		time.Sleep(1 * time.Second)

			//С помощью таймера
		var timeout bool
		N :=3
		timer := time.NewTimer(time.Duration(N) * time.Second)
		go func (){
			for !timeout {
				select {
				case <-timer.C:
					fmt.Println("Time out...")
					timeout = true
				default:
					//Какая-то работа

			}
		}
		}()

	*/
}

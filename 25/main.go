package main

import (
	"fmt"
	"time"
)

func main() {
	sleep(2)
	fmt.Println("starting work after sleep...")
	//Какая-то работа...
}
func sleep(duration int) {
	//time.After ждёт указаное время, а потом отправляет текущее время в возвращаемый канал
	ch := time.After(time.Second * time.Duration(duration))
	for {
		select {
		case <-ch:
			fmt.Println("end sleeping..\n")
			return
		default:
			fmt.Println("sleeping\n")
		}
	}
}

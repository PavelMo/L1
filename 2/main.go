package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}

	c := make(chan int)
	done := make(chan int)

	//В отдельной горутине отправляем элементы массива в канал
	go func() {
		for _, val := range arr {
			c <- val
		}

		done <- 0

	}()
	//функция squares будет работать пока горутина не отправит 0 в канал, сигнализирующий о заверщении работы функции
	squares(c, done)
}

//Вычисление квадратов элементов массива из канала
func squares(c, done chan int) {
	var d int
	for {
		select {
		case d = <-c:
			fmt.Println(d * d)
		case <-done:
			fmt.Println("\nDone......")
			close(c)
			return
		}
	}
	// Можно просто использовать range, который автоматически проверяет закрыт канал или нет
	//	И закрывать канал в гоурутине после записи всех чисел
	/*
		for val :=range c{
		fmt.Println(val*val)
		}
	*/
}

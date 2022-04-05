package main

import "fmt"

func main() {
	ch := make(chan struct{})

	data := make([]interface{}, 0)
	data = append(data, 1, "some string", true, ch)


	for i := 0; i < len(data); i++ {
		//Можно просто использовать Printf: fmt.Printf("value: %v Type: %T \n", data[i], data[i])
		switch data[i].(type) {
		case int:
			fmt.Printf("Get an int! value: %v \n", data[i])
		case string:
			fmt.Printf("Get a string! value: %v \n", data[i])
		case bool:
			fmt.Printf("Get a bool! value: %v \n", data[i])
		default:
			fmt.Printf("value: %v Type: %T \n", data[i], data[i])
		}
	}
}

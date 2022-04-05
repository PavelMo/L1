package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	/*Из-за возможных символов юникода, которые могут занимать больше одного байта, строку нужно конвертировать в руну
	//К примеру, рассмотрим код:
		some := "波音公司(Boeing)在查尔斯顿附近的新厂破土动工时，曾宣扬这里是最先进的制造中心"

		fmt.Println(len(some))         // Выводит: 110
		fmt.Println(len([]rune(some))) // Выводит: 42
	*/
	justString = string([]rune(v)[:100])

	fmt.Println(justString)
	//fmt.Println(len([]rune(v)))
}

func main() {
	someFunc()
}
func createHugeString(size int) string {
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))

	hugeString := strings.Builder{}

	for i := 0; i < size; i++ {
		hugeString.WriteRune('丁' + rune(randomizer.Intn('俿'-'丁'+1)))
	}

	return hugeString.String()
}

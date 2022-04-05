package main

import "fmt"

type Human struct {
	//Какие-то поля......
}

//Тип Action содержит в себе экземпляр типа Human и имеет доступ к его полям и методам
type Action struct {
	some string
	Human
}

func (h *Human) Do(some string) {
	fmt.Println("Human do:", some)
}
//Конструктор
func Constructor(action string) *Action {

	return &Action{
		some:  action,
		Human: Human{},
	}
}

func main() {
	a := Constructor("something")
	a.Do(a.some)
}

package main

/*
穿衣程序第一版
人和衣服绑定在一个类上，如果要新增装扮，需要修改 Person 类，不符合开放-封闭原则
*/

import "fmt"

func main() {
	person := NewPerson("TT")

	fmt.Println("第一种装扮: ")
	person.WearTShirts()
	person.WearBigTrouser()
	person.WearSneakers()
	person.Show()

	fmt.Println()

	fmt.Println("第二种装扮: ")
	person.WearSuit()
	person.WearTie()
	person.WearLeatherShoes()
	person.Show()

}

type Person struct {
	name string
}

func NewPerson(name string) *Person {
	return &Person{name}
}

func (p *Person) WearTShirts() {
	fmt.Println("大T恤")
}

func (p *Person) WearBigTrouser() {
	fmt.Println("垮裤")
}

func (p *Person) WearSneakers() {
	fmt.Println("破球鞋")
}

func (p *Person) WearSuit() {
	fmt.Println("西装")
}

func (p *Person) WearTie() {
	fmt.Println("领带")
}

func (p *Person) WearLeatherShoes() {
	fmt.Println("皮鞋")
}

func (p *Person) Show() {
	fmt.Printf("装扮的 %v\n", p.name)
}

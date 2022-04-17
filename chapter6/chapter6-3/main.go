package main

/*
穿衣程序第二版
使用抽象将人与服饰分离，如果要增加超人装扮，只要增加类即可
但使用了抽象不等于用好了面向对象
*/

import "fmt"

func main() {
	person := NewPerson("TT")
	fmt.Println("第一种装扮：")
	var (
		tshirts    Finery = new(TShirts)
		bigTrouser Finery = new(BigTrouser)
		sneakers   Finery = new(Sneakers)
	)

	tshirts.Show()
	bigTrouser.Show()
	sneakers.Show()
	person.Show()

	fmt.Println()

	fmt.Println("第二种装扮：")
	var (
		suit         Finery = new(Suit)
		tie          Finery = new(Tie)
		leatherShoes        = new(LeatherShoes)
	)
	suit.Show()
	tie.Show()
	leatherShoes.Show()
	person.Show()
}

type Person struct {
	name string
}

func NewPerson(name string) *Person {
	return &Person{name}
}

func (p *Person) Show() {
	fmt.Printf("装扮的 %v\n", p.name)
}

type Finery interface {
	Show()
}

type TShirts struct{}

func (ts *TShirts) Show() { fmt.Println("大T恤") }

type BigTrouser struct{}

func (ts *BigTrouser) Show() { fmt.Println("垮裤") }

type Sneakers struct{}

func (ts *Sneakers) Show() { fmt.Println("破球鞋") }

type Suit struct{}

func (ts *Suit) Show() { fmt.Println("西装") }

type Tie struct{}

func (ts *Tie) Show() { fmt.Println("领带") }

type LeatherShoes struct{}

func (ts *LeatherShoes) Show() { fmt.Println("皮鞋") }

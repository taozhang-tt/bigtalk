package main

import "fmt"

/*
穿衣程序装饰模式版本
*/
func main() {
	person := NewPerson("TT")
	fmt.Println("第一种装扮：")
	var (
		tshirts    = new(TShirts)
		bigTrouser = new(BigTrouser)
		sneakers   = new(Sneakers)
	)
	tshirts.Decorate(person)
	bigTrouser.Decorate(tshirts)
	sneakers.Decorate(bigTrouser)
	sneakers.Show()

	fmt.Printf("\n\n")

	fmt.Println("第二种装扮：")
	leatherShoes := new(LeatherShoes)
	tie := new(Tie)
	suit := new(Suit)
	leatherShoes.Decorate(person)
	tie.Decorate(leatherShoes)
	suit.Decorate(tie)
	suit.Show()

}

type People interface {
	Show()
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

// 服饰类(Decorator)
type Finery struct {
	component People
}

func (f *Finery) Decorate(component People) {
	f.component = component
}

func (f *Finery) Show() {
	if f.component != nil {
		f.component.Show()
	}
}

// 具体服饰类(ContreteDecorator)
type TShirts struct {
	Finery
}

func (ts *TShirts) Show() {
	ts.Finery.Show()
	fmt.Print("大T恤 ")
}

type BigTrouser struct {
	Finery
}

func (bt *BigTrouser) Show() {
	bt.Finery.Show()
	fmt.Print("垮裤 ")
}

type Sneakers struct {
	Finery
}

func (s *Sneakers) Show() {
	s.Finery.Show()
	fmt.Print("破球鞋")
}

type Suit struct {
	Finery
}

func (s *Suit) Show() {
	s.Finery.Show()
	fmt.Print("西装 ")
}

type Tie struct {
	Finery
}

func (t *Tie) Show() {
	t.Finery.Show()
	fmt.Print("领带 ")
}

type LeatherShoes struct {
	Finery
}

func (ls *LeatherShoes) Show() {
	ls.Finery.Show()
	fmt.Print("皮鞋 ")
}

package main

import "fmt"

func main() {
	jiaojiao := &SchoolGirl{"李娇娇"}
	zhuojiayi := NewPursuit(jiaojiao)

	zhuojiayi.GiveDolls()
	zhuojiayi.GiveFlowers()
	zhuojiayi.GiveChocolate()
}

// 被追求者类
type SchoolGirl struct {
	Name string
}

// 追求者类
type Pursuit struct {
	mm *SchoolGirl
}

func NewPursuit(mm *SchoolGirl) *Pursuit {
	return &Pursuit{mm}
}

func (p *Pursuit) GiveDolls() {
	fmt.Printf("%v 送你洋娃娃\n", p.mm.Name)
}

func (p *Pursuit) GiveFlowers() {
	fmt.Printf("%v 送你鲜花\n", p.mm.Name)
}

func (p *Pursuit) GiveChocolate() {
	fmt.Printf("%v 送你巧克力\n", p.mm.Name)
}

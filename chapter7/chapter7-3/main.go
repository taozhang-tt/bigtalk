package main

import "fmt"

func main() {
	jiaojiao := &SchoolGirl{"李娇娇"}
	zhuojiayi := NewProxy(jiaojiao)

	zhuojiayi.GiveDolls()
	zhuojiayi.GiveFlowers()
	zhuojiayi.GiveChocolate()
}

// 被追求者类
type SchoolGirl struct {
	Name string
}

// 代理类
type Proxy struct {
	mm *SchoolGirl
}

func NewProxy(mm *SchoolGirl) *Proxy {
	return &Proxy{mm}
}

func (p *Proxy) GiveDolls() {
	fmt.Printf("%v 送你洋娃娃\n", p.mm.Name)
}

func (p *Proxy) GiveFlowers() {
	fmt.Printf("%v 送你鲜花\n", p.mm.Name)
}

func (p *Proxy) GiveChocolate() {
	fmt.Printf("%v 送你巧克力\n", p.mm.Name)
}

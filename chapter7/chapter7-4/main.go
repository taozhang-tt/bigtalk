package main

import "fmt"

func main() {
	jiaojiao := &SchoolGirl{"娇娇"}

	proxy := NewProxy(jiaojiao)
	proxy.GiveDolls()
	proxy.GiveFlowers()
	proxy.GiveChocolate()
}

type IGiveGift interface {
	GiveDolls()
	GiveFlowers()
	GiveChocolate()
}

type SchoolGirl struct {
	Name string
}

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

type Proxy struct {
	gg *Pursuit
}

func NewProxy(mm *SchoolGirl) *Proxy {
	return &Proxy{NewPursuit(mm)}
}

func (p *Proxy) GiveDolls() {
	p.gg.GiveDolls()
}

func (p *Proxy) GiveFlowers() {
	p.gg.GiveFlowers()
}

func (p *Proxy) GiveChocolate() {
	p.gg.GiveChocolate()
}

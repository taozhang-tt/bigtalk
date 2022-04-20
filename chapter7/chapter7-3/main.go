package main

import "fmt"

func main() {
	main2()
}

func main1() {
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

// 菜鸟教程的另一个例子
type Image interface {
	display()
}

type RealImage struct {
	fileName string
}

func (ri *RealImage) display() {
	fmt.Printf("Displaying: %v\n", ri.fileName)
}

func (ri *RealImage) loadFromDisk(fileName string) {
	fmt.Printf("Loading: %v\n", fileName)
}

func NewRealImage(fileName string) *RealImage {
	img := &RealImage{fileName}
	img.loadFromDisk(fileName)
	return img
}

type ProxyImage struct {
	realImage *RealImage
	fileName  string
}

func (pi *ProxyImage) display() {
	if pi.realImage == nil {
		pi.realImage = NewRealImage(pi.fileName)
	}
	pi.realImage.display()
}

func NewProxyImage(fileName string) *ProxyImage {
	return &ProxyImage{fileName: fileName}
}

func main2() {
	img := NewProxyImage("picture1.png")
	// 图像将从磁盘加载
	img.display()

	// 图像将会直接加载
	img.display()
}

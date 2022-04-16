package main

// 一个简易的收银系统

import "fmt"

func main() {
	cash := new(Cash)
	fmt.Println("购买99的袜子一双")
	cash.Add(99, 1)
	fmt.Println("购买199的裤子三条")
	cash.Add(199, 3)
	fmt.Printf("一共 %v 元\n", cash.Total())
}

// 收银类
type Cash struct {
	total float64
}

// 商品单价、购买数量
func (c *Cash) Add(price float64, num int) {
	c.total += price * float64(num)
}

func (c *Cash) Total() float64 {
	return c.total
}

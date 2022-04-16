package main

// 收银系统增加打折功能
// 提供打折下拉框选择商品折扣

import "fmt"

func main() {
	cash := new(Cash)
	fmt.Println("购买99的袜子一双，八折")
	cash.Add(99, 1, 1)
	fmt.Println("购买199的裤子三条, 六折")
	cash.Add(199, 3, 2)
	fmt.Printf("一共 %v 元\n", cash.Total())
}

// 收银类
type Cash struct {
	total float64
}

// 商品单价、购买数量
func (c *Cash) Add(price float64, num int, discount int) {
	switch discount {
	case 0:
		c.total += price * float64(num)
	case 1:
		c.total += price * float64(num) * 0.8 // 八折
	case 2:
		c.total += price * float64(num) * 0.6 // 六折
	case 3:
		c.total += price * float64(num) * 0.5 // 五折
	}
}

func (c *Cash) Total() float64 {
	return c.total
}

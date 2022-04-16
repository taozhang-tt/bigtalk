package main

import (
	"fmt"
	"math"
)

// 新增需求
// 需要有满300返100的促销算法
// 满200送50
// ...

// 客户端代码
func main() {
	fmt.Println("请输入收费模式：")
	var tp string
	fmt.Scanln(&tp)
	factory := new(CashFactory)
	cash := factory.CreateCashAccepter(tp)
	fmt.Println("输入优惠前金额：")
	var money float64
	fmt.Scanln(&money)
	fmt.Printf("优惠后金额：%v", cash.AcceptCash(money))

}

// 现金收费抽象接口
type CashSuper interface {
	AcceptCash(float64) float64
}

// 正常收费
type CashNormal struct{}

func (cn *CashNormal) AcceptCash(money float64) float64 {
	return money
}

// 折扣收费
type CashDiscount struct {
	Discount float64 // 具体折扣
}

func (cd *CashDiscount) AcceptCash(money float64) float64 {
	return cd.Discount * money
}

// 满减收费
type CashReturn struct {
	MoneyCondition float64 // 满额
	MoneyReturn    float64 // 返额
}

func (cr *CashReturn) AcceptCash(money float64) float64 {
	return money - math.Floor(money/cr.MoneyCondition)*cr.MoneyReturn
}

type CashFactory struct{}

func (cf *CashFactory) CreateCashAccepter(tp string) CashSuper {
	var cash CashSuper
	switch tp {
	case "正常收费":
		cash = &CashNormal{}
	case "满300减100":
		cash = &CashReturn{300, 100}
	case "打八折":
		cash = &CashDiscount{0.8}
	}
	return cash
}

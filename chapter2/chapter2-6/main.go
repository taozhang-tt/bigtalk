package main

// 策略+简单工厂

import (
	"fmt"
	"math"
)

// 客户端代码
func main() {
	fmt.Println("请输入收费模式：")
	var tp string
	fmt.Scanln(&tp)
	context := NewCashContext(tp)
	fmt.Println("输入优惠前金额：")
	var money float64
	fmt.Scanln(&money)
	fmt.Printf("优惠后金额：%v", context.GetResult(money))
}

type CashContext struct {
	cash CashSuper
}

func NewCashContext(tp string) *CashContext {
	context := new(CashContext)
	switch tp {
	case "正常收费":
		context.cash = &CashNormal{}
	case "满300减100":
		context.cash = &CashReturn{300, 100}
	case "打八折":
		context.cash = &CashDiscount{0.8}
	}
	return context
}

func (cc *CashContext) GetResult(money float64) float64 {
	return cc.cash.AcceptCash(money)
}

// 现金收费抽象策略
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

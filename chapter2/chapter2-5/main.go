package main

// 策略模式
// 客户端需要知道 CashContext，需要知道具体的策略 CashDiscount、CashReturn 等

import (
	"fmt"
	"math"
)

// 客户端代码
func main() {
	var context *CashContext
	fmt.Println("请输入收费模式：")
	var tp string
	fmt.Scanln(&tp)
	switch tp {
	case "正常收费":
		context = NewCashContext(new(CashNormal))
	case "满300减100":
		context = NewCashContext(&CashReturn{300, 100})
	case "打八折":
		context = NewCashContext(&CashDiscount{0.8})
	}
	fmt.Println("输入优惠前金额：")
	var money float64
	fmt.Scanln(&money)
	fmt.Printf("优惠后金额：%v", context.GetResult(money))
}

type CashContext struct {
	cash CashSuper
}

func NewCashContext(cash CashSuper) *CashContext {
	return &CashContext{cash}
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

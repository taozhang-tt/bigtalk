package main

// 简单工厂

import "fmt"

// 客户端
func main() {
	fmt.Println("请输入数字A：")
	var num1 float64
	fmt.Scanln(&num1)

	fmt.Println("请输入运算符号：")
	var operate string
	fmt.Scanln(&operate)

	fmt.Println("请输入数字B：")
	var num2 float64
	fmt.Scanln(&num2)

	factory := new(OperationFactory)
	operator := factory.CreateOperate(operate)
	fmt.Printf("结果是：%v\n", operator.GetResult(num1, num2))
}

type OperationFactory struct{}

func (o *OperationFactory) CreateOperate(operate string) Operation {
	var ret Operation
	switch operate {
	case "+":
		ret = new(OperationAdd)
	case "-":
		ret = new(OperationSub)
	case "*":
		ret = new(OperationMul)
	case "/":
		ret = new(OperationDiv)
	}
	return ret
}

type Operation interface {
	GetResult(num1, num2 float64) float64
}

// 加法类
type OperationAdd struct{}

func (o *OperationAdd) GetResult(num1, num2 float64) float64 {
	return num1 + num2
}

// 减法类
type OperationSub struct{}

func (o *OperationSub) GetResult(num1, num2 float64) float64 {
	return num1 - num2
}

// 乘法类
type OperationMul struct{}

func (o *OperationMul) GetResult(num1, num2 float64) float64 {
	return num1 * num2
}

// 除法类
type OperationDiv struct{}

func (o *OperationDiv) GetResult(num1, num2 float64) float64 {
	return num1 / num2
}
